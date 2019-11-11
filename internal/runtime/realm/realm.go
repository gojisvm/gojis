package realm

import (
	"sync/atomic"

	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

// Intrinsic names used by the specification. The specification denotes the
// usage of these names as e.g. %ThrowTypeError% (enclosed in '%').
const (
	IntrinsicNameObjectPrototype   = "ObjectPrototype"
	IntrinsicNameFunctionPrototype = "FunctionPrototype"
	IntrinsicNameThrowTypeError    = "ThrowTypeError"
)

var (
	currentRealm atomic.Value // holds *realm.Realm
)

// CurrentRealm returns the current realm as used in the specification.
func CurrentRealm() *Realm { return currentRealm.Load().(*Realm) }

// Realm is a struct that contains fields specified in 8.2.
type Realm struct {
	Intrinsics  *lang.Record
	GlobalObj   lang.Value                   // Object or Undefined
	GlobalEnv   lang.Value                   // Object or Undefined
	TemplateMap map[interface{}]*lang.Object // Parse Node -> Object
	HostDefined lang.InternalValue
}

// Type returns lang.TypeInternal.
func (*Realm) Type() lang.Type { return lang.TypeInternal }

// Value returns the Realm itself.
func (r *Realm) Value() interface{} { return r }

// CreateRealm creates a realm with Undefined GlobalObj and Undefined GlobalEnv,
// an empty TemplateMap, Undefined as HostDefined and the Intrinsics specified
// in 8.2.2. CreateRealm itself is specified in 8.2.1.
func CreateRealm() *Realm {
	r := new(Realm)
	CreateIntrinsics(r)
	r.GlobalObj = lang.Undefined
	r.GlobalEnv = lang.Undefined
	r.TemplateMap = make(map[interface{}]*lang.Object)
	return r
}

// CreateIntrinsics sets intrinsic objects of a record as specified in 8.2.2.
func CreateIntrinsics(r *Realm) lang.InternalValue {
	r.Intrinsics = lang.NewRecord()
	objProto := lang.ObjectCreate(lang.Null)
	r.Intrinsics.SetField(IntrinsicNameObjectPrototype, objProto)

	// %ThrowTypeError% as in 8.2.2 and 9.2.9.1 is not implemented here, since
	// "throwing" is represented as returning an error in this implementation.
	// If the decision is made, that we will use an actual function to throw an
	// error, its prototype must be set to the funcProto that is created
	// hereafter (8.2.2, Step 11).

	funcProto := CreateBuiltinFunction(func(lang.Value, ...lang.Value) (lang.Value, errors.Error) { return nil, nil }, r, lang.Null)
	r.Intrinsics.SetField(IntrinsicNameFunctionPrototype, funcProto)
	r.AddRestrictedFunctionProperties(funcProto)

	// #64: 8.2.2, Step 13

	return r.Intrinsics
}

// AddRestrictedFunctionProperties adds restricted function properties to the
// given object within this realm as specified in 9.2.9.
func (r *Realm) AddRestrictedFunctionProperties(f *lang.Object) {
	lang.DefinePropertyOrThrow(f, lang.NewStringOrSymbol(lang.NewString("caller")), lang.NewAccessorProperty(nil, nil, false, true))
	lang.DefinePropertyOrThrow(f, lang.NewStringOrSymbol(lang.NewString("arguments")), lang.NewAccessorProperty(nil, nil, false, true))
}

// GetIntrinsicObject returns the intrinsic object of the realm specified by the
// given name, or Undefined if no intrinsic object with that name could be
// found.
func (r *Realm) GetIntrinsicObject(n string) lang.Value {
	val, ok := r.Intrinsics.GetField(n)
	if !ok {
		return lang.Undefined
	}
	return val.(lang.Value)
}

// SetRealmGlobalObject sets the global object of a realm, as specified in
// 8.2.3. If the passed global object is Undefined, a new object created from
// the intrinsic %ObjectPrototype% object of this realm will be used instead.
func (r *Realm) SetRealmGlobalObject(globalObj, thisValue lang.Value) *Realm {
	if globalObj == lang.Undefined {
		panic("TODO: 8.2.3")
	}

	if thisValue == lang.Undefined {
		thisValue = globalObj
	}

	r.GlobalObj = globalObj

	GlobalEnv := binding.NewGlobalEnvironment(globalObj.(*lang.Object), thisValue.(*lang.Object))
	r.GlobalEnv = GlobalEnv

	return r
}

// SetDefaultGlobalBindings is specified in 8.2.4. TODO: improve godoc
func (r *Realm) SetDefaultGlobalBindings() lang.Value {
	global := r.GlobalObj.(*lang.Object)
	_ = global
	panic("TODO: for every property\n" + `2. For each property of the Global Object specified in clause 18, do
	a. Let name be the String value of the property name.
	b. Let desc be the fully populated data property descriptor for the property containing the specified
	attributes for the property. For properties listed in 18.2, 18.3, or 18.4 the value of the [[Value]] attribute is
	the corresponding intrinsic object from realmRec.
	c. Perform ? DefinePropertyOrThrow(global, name, desc).`)
}

// GetFunctionRealm returns the realm that the object is belonging to. That is,
// if the object has a Realm internal slot, its value will be returned. If it is
// a proxy object, the proxy target's function realm will be returned.
// Otherwise, this function returns the current realm record. GetFunctionRealm
// is specified in 7.3.22.
func GetFunctionRealm(obj *lang.Object) *Realm {
	panic("TODO")
}

// OrdinaryCreateFromConstructor creates an object, whose prototype will be the
// prototype of the passed constructor object. If that constructor's property is
// not set, the intrinsic default prototype will be used instead.
// OrdinaryCreateFromConstructor is specified in 9.1.13.
func OrdinaryCreateFromConstructor(constructor *lang.Object, intrinsicDefaultProto lang.String, internalSlotsList ...lang.StringOrSymbol) (*lang.Object, errors.Error) {
	proto, err := GetPrototypeFromConstructor(constructor, intrinsicDefaultProto)
	if err != nil {
		return nil, err
	}

	return lang.ObjectCreate(proto, internalSlotsList...), nil
}

// GetPrototypeFromConstructor determines the [[Prototype]] value that should be
// used to create an object corresponding to a specific constructor.
// GetPrototypeFromConstructor is specified in 9.1.14.
func GetPrototypeFromConstructor(constructor *lang.Object, intrinsicDefaultProto lang.String) (*lang.Object, errors.Error) {
	proto, err := lang.Get(constructor, lang.NewStringOrSymbol(lang.NewString("prototype")))
	if err != nil {
		return nil, err
	}

	if proto.Type() != lang.TypeObject {
		realm := GetFunctionRealm(constructor)
		proto = realm.GetIntrinsicObject(intrinsicDefaultProto.Value().(string))
	}
	return proto.(*lang.Object), nil
}
