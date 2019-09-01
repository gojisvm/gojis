package lang

import (
	"reflect"
	"strconv"

	"github.com/gojisvm/gojis/internal/runtime/errors"
)

var _ Value = (*Object)(nil) // ensure that Object implements Value

// FunctionMode represents the ThisMode of a function. The ThisMode defines how
// this references are interpreted within the formal parameters and code body of
// the function. lexical means that this refers to the this value of a lexically
// enclosing function. strict means that the this value is used exactly as
// provided by an invocation of the function. global means that a this value of
// undefined is interpreted as a reference to the global object.
type FunctionMode uint8

// Function modes as defined in 9.2, Table 27.
const (
	FunctionModeUnknown FunctionMode = iota
	FunctionModeLexical
	FunctionModeStrict
	FunctionModeGlobal
)

// FunctionKind is the kind of a function, as specified in 9.2, Table 27.
type FunctionKind uint8

// Function kinds as defined in 9.2, Table 27.
const (
	FunctionKindUnknown FunctionKind = iota
	FunctionKindNormal
	FunctionKindClassGenerator
	FunctionKindGenerator
	FunctionKindAsync
)

// ConstructorKind is the kind of a constructor function, as specified in 9.2,
// Table 27.
type ConstructorKind uint8

// Constructor kinds as defined in 9.2, Table 27.
const (
	ConstructorKindUnknown ConstructorKind = iota
	ConstructorKindBase
	ConstructorKindDerived
)

// Object is a language type as specified by the language spec.
type Object struct {
	fields map[StringOrSymbol]*Property
	slots  map[StringOrSymbol]Value

	Prototype  Value // *Object or Null
	Extensible bool

	// Function Object
	Environment      InternalValue
	FormalParameters interface{}
	FunctionKind     FunctionKind
	ConstructorKind  ConstructorKind
	Realm            InternalValue
	ScriptOrModule   interface{}
	ThisMode         FunctionMode
	Strict           bool
	HomeObject       *Object

	// Call is the Call function of an object as (kind of) specified by the
	// language spec. This is only not nil for function and constructor function
	// objects.
	Call func(Value, ...Value) (Value, errors.Error)

	// Constructor Function Object

	// Construct is the Construct function of an obejct as (kind of) specified
	// by the language spec. This is only not nil for constructor function
	// objects.
	Construct func(*Object, ...Value) (*Object, errors.Error)
}

// ObjectCreate creates a new ordinary object at runtime, where proto is the
// given prototype (must be an Object or Null), and internalSlotsList is a list
// of the names of additional internal slots that must be defined as part of the
// object. If none are provided, an empty list is used. ObjectCreate is
// specified in 9.1.12.
func ObjectCreate(proto Value, internalSlotsList ...StringOrSymbol) *Object {
	if internalSlotsList == nil {
		internalSlotsList = []StringOrSymbol{}
	}

	obj := new(Object)
	obj.fields = make(map[StringOrSymbol]*Property)
	obj.slots = make(map[StringOrSymbol]Value)
	for _, slot := range internalSlotsList {
		obj.slots[slot] = Undefined
	}
	EnsureTypeOneOf(proto, TypeObject, TypeNull) // panic if proto is not TypeObject or TypeNull
	obj.Prototype = proto
	obj.Extensible = true

	return obj
}

// Value returns the object itself.
//
//  var o *Object
//  ...
//  o == o.Value() // true
func (o *Object) Value() interface{} { return o }

// Type returns lang.TypeObject.
func (o *Object) Type() Type { return TypeObject }

/* -- 9.1, ordinary object internal methods and internal slots -- */

// GetPrototypeOf delegates to OrdinaryGetPrototypeOf. GetPrototypeOf is
// specified in 9.1.1.
func (o *Object) GetPrototypeOf() Value {
	return o.OrdinaryGetPrototypeOf()
}

// OrdinaryGetPrototypeOf returns the prototype of the object, either an Object
// or Null. OrdinaryGetPrototypeOf is specified in 9.1.1.1.
func (o *Object) OrdinaryGetPrototypeOf() Value {
	return o.Prototype
}

// SetPrototypeOf delegates to OrdinarySetPrototypeOf. SetPrototypeOf is
// specified in 9.1.2.
func (o *Object) SetPrototypeOf(v Value) Boolean {
	return o.OrdinarySetPrototypeOf(v)
}

// OrdinarySetPrototypeOf sets the prototype of the object. The given prototype
// value can either be an Object or Null. OrdinarySetPrototypeOf is specified in
// 9.1.2.1.
func (o *Object) OrdinarySetPrototypeOf(v Value) Boolean {
	EnsureTypeOneOf(v, TypeObject, TypeNull)

	extensible := o.Extensible
	current := o.Prototype

	if InternalSameValue(v, current) {
		return True
	}

	if !extensible {
		return False
	}

	/*
	   This loop guarantees that there will be no circularities in any prototype
	   chain that only includes objects that use the ordinary object definitions
	   for [[GetPrototypeOf]] and [[SetPrototypeOf]].
	*/
	p := v
	done := false
	for !done {
		if p == Null {
			done = true
		} else if InternalSameValue(p, o) {
			return False
		} else {
			// FIXME: if p.GetPrototypeOf is not the ordinary object internal
			// method defined in 9.1.1, set done to true. else {
			p = p.(*Object).Prototype // this type assertion cannot fail, since p is checked to be Object or Null, and Null is handled above
			// }
		}
	}

	o.Prototype = v
	return True
}

// IsExtensible delegates to OrdinaryIsExtensible. IsExtensible is specified in
// 9.1.3.
func (o *Object) IsExtensible() Boolean {
	return o.OrdinaryIsExtensible()
}

// OrdinaryIsExtensible is used to determine whether the object is extensible.
// OrdinaryIsExtensible is specified in 9.1.3.1.
func (o *Object) OrdinaryIsExtensible() Boolean {
	return Boolean(o.Extensible)
}

// PreventExtensions delegates to OrdinaryPreventExtensions. PreventExtensions
// is specified in 9.1.4.
func (o *Object) PreventExtensions() Boolean {
	return o.OrdinaryPreventExtensions()
}

// OrdinaryPreventExtensions makes the object non-extensible. After a call to
// this method, IsExtensible and OrdinaryIsExtensible calls to the object will
// always return false. OrdinaryPreventExtensions is specified in 9.1.4.1.
func (o *Object) OrdinaryPreventExtensions() Boolean {
	o.Extensible = false
	return True
}

// GetOwnProperty delegates to OrdinaryGetOwnProperty. GetOwnProperty is
// specified in 9.1.5.
func (o *Object) GetOwnProperty(p StringOrSymbol) *Property {
	return o.OrdinaryGetOwnProperty(p)
}

// OrdinaryGetOwnProperty returns a property with the given key of this object.
// Where Undefined in the specification is used to indicate no result, here, a
// nil pointer is used. OrdinaryGetOwnProperty is specified in 9.1.5.1.
func (o *Object) OrdinaryGetOwnProperty(p StringOrSymbol) *Property {
	x, ok := o.fields[p]
	if !ok {
		return nil // actually Undefined
	}

	// TODO: maybe introduce a Property#copy function?
	d := NewProperty()
	if x.IsDataDescriptor() {
		d.SetField(FieldNameValue, x.Value())
		d.SetField(FieldNameWritable, x.Writable())
	} else if x.IsAccessorDescriptor() {
		d.SetField(FieldNameGet, x.Get())
		d.SetField(FieldNameSet, x.Set())
	}
	d.SetField(FieldNameEnumerable, x.Enumerable())
	d.SetField(FieldNameConfigurable, x.Configurable())

	return d
}

// DefineOwnProperty delegates to OrdinaryDefineOwnProperty. DefineOwnProperty
// is specified in 9.1.6.
func (o *Object) DefineOwnProperty(p StringOrSymbol, desc *Property) Boolean {
	return o.OrdinaryDefineOwnProperty(p, desc)
}

// OrdinaryDefineOwnProperty is used to define an own property of the object.
// OrdinaryDefineOwnProperty is specified in 9.1.6.1.
func (o *Object) OrdinaryDefineOwnProperty(p StringOrSymbol, desc *Property) Boolean {
	current := o.GetOwnProperty(p)
	extensible := o.Extensible
	return o.ValidateAndApplyPropertyDescriptor(p, extensible, desc, current)
}

// IsCompatiblePropertyDescriptor delegates to
// ValidateAndApplyPropertyDescriptor on a nil Object and a zero StringOrSymbol.
// This resembles the behaviour of the object and the property key being
// Undefined. IsCompatiblePropertyDescriptor is specified in 9.1.6.2.
func (o *Object) IsCompatiblePropertyDescriptor(extensible bool, desc, current *Property) Boolean {
	// TODO: make code more beautiful
	return ((*Object)(nil)).ValidateAndApplyPropertyDescriptor(StringOrSymbol{}, extensible, desc, current)
}

// ValidateAndApplyPropertyDescriptor is pretty complicated and should be
// simplified and/or refactored into multiple sub-functions (TODO:).
// ValidateAndApplyPropertyDescriptor is specified in 9.1.6.3.
func (o *Object) ValidateAndApplyPropertyDescriptor(p StringOrSymbol, extensible bool, desc, current *Property) Boolean {
	// if o != nil, p is not zero value
	if current == nil {
		if !extensible {
			return False
		}

		if desc.IsGenericDescriptor() || desc.IsDataDescriptor() {
			if o != nil {
				/*
					If O is not undefined, create an own accessor property named
					P of object O whose [[Get]], [[Set]], [[Enumerable]] and
					[[Configurable]] attribute values are described by Desc. If
					the value of an attribute field of Desc is absent, the
					attribute of the newly created property is set to its
					default value.
				*/
				panic("TODO")
			}
		} else {
			// desc.IsAccessorDescriptor() is true
			/*
			   If O is not undefined, create an own accessor property named P of
			   object O whose [[Get]], [[Set]], [[Enumerable]] and
			   [[Configurable]] attribute values are described by Desc. If the
			   value of an attribute field of Desc is absent, the attribute of
			   the newly created property is set to its default value.
			*/
			panic("TODO")
		}
	}

	if len(desc.Record.fields) == 0 {
		return True
	}

	if !current.Configurable() {
		if desc.Configurable() {
			return False
		}

		if desc.Enumerable() && (current.Enumerable() == !desc.Enumerable()) {
			return False
		}
	}

	if desc.IsGenericDescriptor() {
		// no further valudation is required
	} else if current.IsDataDescriptor() != desc.IsDataDescriptor() {
		if !current.Configurable() {
			return False
		}

		if current.IsDataDescriptor() {
			/*
			   If O is not undefined, convert the property named P of object O
			   from a data property to an accessor property. Preserve the
			   existing values of the converted property's [[Configurable]] and
			   [[Enumerable]] attributes and set the rest of the property's
			   attributes to their default values.
			*/
			panic("TODO")
		} else {
			/*
			   If O is not undefined, convert the property named P of object O
			   from an accessor property to a data property. Preserve the
			   existing values of the converted property's [[Configurable]] and
			   [[Enumerable]] attributes and set the rest of the property's
			   attributes to their default values.
			*/
			panic("TODO")
		}
	} else if current.IsDataDescriptor() && desc.IsDataDescriptor() {
		if !current.Configurable() && !current.Writable() {
			if desc.Writable() {
				return False
			}

			if _, ok := desc.GetField(FieldNameValue); ok && !InternalSameValue(desc.Value(), current.Value()) {
				return False
			}

			return True
		}
	} else if current.IsAccessorDescriptor() && desc.IsAccessorDescriptor() {
		if !current.Configurable() {
			if _, ok := desc.GetField(FieldNameSet); ok && reflect.ValueOf(desc.Set()).Pointer() != reflect.ValueOf(current.Set()).Pointer() {
				return False
			}

			return True
		}
	}

	if o != nil {
		prop := o.fields[p]
		for k, v := range desc.Record.fields {
			prop.fields[k] = v
		}
	}

	return True
}

// HasProperty delegates to OrdinaryHasProperty. HasProperty is specified in
// 9.1.7.
func (o *Object) HasProperty(p StringOrSymbol) Boolean {
	return o.OrdinaryHasProperty(p)
}

// OrdinaryHasProperty is used to determine whether an object or any object in
// its prototype chain has a property with the given name. OrdinaryHasProperty
// is specified in 9.1.7.1.
func (o *Object) OrdinaryHasProperty(p StringOrSymbol) Boolean {
	if o.GetOwnProperty(p) != nil {
		return True
	}

	parent := o.GetPrototypeOf()
	if parent != Null {
		return parent.(*Object).HasProperty(p)
	}

	return False
}

// Get delegates to OrdinaryGet. Get is specified in 9.1.8.
func (o *Object) Get(p StringOrSymbol, receiver Value) (Value, errors.Error) {
	return o.OrdinaryGet(p, receiver)
}

// OrdinaryGet returns the value of the property with the given name. If the
// property is an own property of this object, its value is returned. If this
// object does not have a property with the given name, its prototype chain will
// be checked.
//
// If the found property is a data property descriptor, its value is returned.
//
// If the found property is an accessor property descriptor, the value returned
// by its [[Get]] method is returned.
//
// If no property with the given name could be found, Undefined is returned.
func (o *Object) OrdinaryGet(p StringOrSymbol, receiver Value) (Value, errors.Error) {
	desc := o.GetOwnProperty(p)
	if desc == nil {
		parent := o.GetPrototypeOf()
		if parent == Null {
			return Undefined, nil
		}
		return parent.(*Object).Get(p, receiver)
	}

	if desc.IsDataDescriptor() {
		return desc.Value(), nil
	}

	if getter := desc.Get(); getter != Undefined {
		return Call(getter.(*Object), receiver)
	}
	return Undefined, nil
}

// Set delegates to OrdinarySet. Set is specified in 9.1.9.
func (o *Object) Set(p StringOrSymbol, v, receiver Value) (Boolean, errors.Error) {
	return o.OrdinarySet(p, v, receiver)
}

// OrdinarySet is a wrapper around OrdinarySetWithOwnDescriptor, which will
// resolve p to an own property of the object. OrdinarySet is specified in
// 9.1.9.1.
func (o *Object) OrdinarySet(p StringOrSymbol, v, receiver Value) (Boolean, errors.Error) {
	return o.OrdinarySetWithOwnDescriptor(p, v, receiver, o.GetOwnProperty(p))
}

// OrdinarySetWithOwnDescriptor is pretty complicated and should be simplified
// and/or refactored into multiple sub-functions (TODO:).
// OrdinarySetWithOwnDescriptor is specified in 9.1.9.2.
func (o *Object) OrdinarySetWithOwnDescriptor(p StringOrSymbol, v, receiver Value, ownDesc *Property) (Boolean, errors.Error) {
	if ownDesc == nil {
		parent := o.GetPrototypeOf()
		if parent != Null {
			return parent.(*Object).Set(p, v, receiver)
		}

		ownDesc = NewDataProperty(Undefined, True, True, True)
	}

	if ownDesc.IsDataDescriptor() {
		if !ownDesc.Writable() {
			return False, nil
		}

		if receiver.Type() != TypeObject {
			return False, nil
		}

		receiverObj := receiver.(*Object)
		existingDescriptor := receiverObj.GetOwnProperty(p)
		if existingDescriptor != nil {
			if existingDescriptor.IsAccessorDescriptor() {
				return False, nil
			}

			if !existingDescriptor.Writable() {
				return False, nil
			}

			valueDesc := NewProperty()
			valueDesc.SetField(FieldNameValue, v)
			return receiverObj.DefineOwnProperty(p, valueDesc), nil
		}

		return CreateDataProperty(receiverObj, p, v), nil
	}

	// assert: ownDesc.IsAccessorDescriptor is true

	if setter := ownDesc.Set(); setter != Undefined {
		_, err := Call(setter.(*Object), receiver, v)
		if err != nil {
			return False, err
		}
		return True, nil
	}
	return False, nil
}

// Delete delegates to OrdinaryDelete. Delete is specified in 9.1.10.
func (o *Object) Delete(p StringOrSymbol) Boolean {
	return o.OrdinaryDelete(p)
}

// OrdinaryDelete removes a property with the name p from the object.
// OrdinaryDelete is specified in 9.1.10.1.
func (o *Object) OrdinaryDelete(p StringOrSymbol) Boolean {
	desc := o.GetOwnProperty(p)
	if desc == nil {
		return True
	}

	if desc.Configurable() {
		delete(o.fields, p)
		return True
	}

	return False
}

// OwnPropertyKeys delegates to OrdinaryOwnPropertyKeys. OwnPropertyKeys is
// specified in 9.1.11.
func (o *Object) OwnPropertyKeys() []StringOrSymbol {
	return o.OrdinaryOwnPropertyKeys()
}

// OrdinaryOwnPropertyKeys returns a list of names of the properties of this
// object. That is, given an object with the properties 'A', 'B', and 'C',
// OrdinaryOwnPropertyKeys will return ['A', 'B', 'C'].
//
// TODO: add comment on ordering and grouping of property names
//
// OrdinaryOwnPropertyKeys is specified in 9.1.11.1.
func (o *Object) OrdinaryOwnPropertyKeys() []StringOrSymbol {
	keys := []StringOrSymbol{}
	secs := make(map[int][]StringOrSymbol)
	const (
		integerIndex int = iota
		stringIndex
		symbolIndex
	)

	for k := range o.fields {
		if n, err := strconv.ParseInt(k.String().Value().(string), 10, 64); err == nil && n >= 0 {
			// k is an integer index
			secs[integerIndex] = append(secs[integerIndex], k)
			continue
		}

		if k.Type() == TypeString {
			secs[stringIndex] = append(secs[stringIndex], k)
		}

		if k.Type() == TypeSymbol {
			secs[symbolIndex] = append(secs[symbolIndex], k)
		}
	}

	keys = append(keys, secs[integerIndex]...)
	keys = append(keys, secs[stringIndex]...)
	keys = append(keys, secs[symbolIndex]...)
	return keys
}
