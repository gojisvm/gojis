package lang

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/runtime/errors"
)

// Get retrieves the value of a specific property of an object.
// Get is specified in 7.3.1.
func Get(o *Object, p StringOrSymbol) (InternalValue, errors.Error) {
	return o.Get(p, o)
}

// GetV retrieves the value of a specific property of an object,
// assuming that the value is an ECMAScript language value.
// If the value is not an object, the property lookup is performed
// using a wrapper object appropriate for the type of the value.
// GetV is specified in 7.3.2.
func GetV(v Value, p StringOrSymbol) (Value, errors.Error) {
	o := ToObject(v)
	return o.Get(p, v)
}

// Set is used to set the value of a specific property of an object.
// If throw == true and the property cannot be set, a TypeError will be returned
// which then must be thrown.
// Set is specified in 7.3.3.
func Set(o *Object, p StringOrSymbol, v Value, throw bool) (Boolean, errors.Error) {
	success, err := o.Set(p, v, o)
	if err != nil {
		return False, err
	}

	if !success.Value().(bool) && throw {
		return False, errors.NewTypeError(fmt.Sprintf("Cannot set '%v' of object", p.Value()))
	}

	return success, nil
}

// CreateDataProperty creates a new own property of an object.
// CreateDataProperty is specified in 7.3.4.
func CreateDataProperty(o *Object, p StringOrSymbol, v Value) Boolean {
	desc := NewDataProperty(v, True, True, True)
	return o.DefineOwnProperty(p, desc)
}

// CreateMethodProperty creates a new own property of an object.
// CreateMethodProperty is specified in 7.3.5.
func CreateMethodProperty(o *Object, p StringOrSymbol, v Value) Boolean {
	desc := NewDataProperty(v, True, False, True)
	return o.DefineOwnProperty(p, desc)
}

// CreateDataPropertyOrThrow creates a new own property of an object.
// It returnes a TypeError exception if the requested property update cannot
// be performed.
// CreateDataPropertyOrThrow is specified in 7.3.6.
func CreateDataPropertyOrThrow(o *Object, p StringOrSymbol, v Value) (Boolean, errors.Error) {
	success := CreateDataProperty(o, p, v)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to create data property '%v'", p.Value()))
	}
	return success, nil
}

// DefinePropertyOrThrow is used to call the DefineOwnProperty internal method of an object.
// If the requested property update cannot be performed, a TypeError is returned.
// DefinePropertyOrThrow is specified in 7.3.7.
func DefinePropertyOrThrow(o *Object, p StringOrSymbol, desc *Property) (Boolean, errors.Error) {
	success := o.DefineOwnProperty(p, desc)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to define property '%v'", p.Value()))
	}

	return success, nil
}

// DeletePropertyOrThrow removes a specific own property of an object.
// If the property is not configurable, an exception is returned.
// DeletePropertyOrThrow is specified in 7.3.8.
func DeletePropertyOrThrow(o *Object, p StringOrSymbol) (Boolean, errors.Error) {
	success := o.Delete(p)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to delete property '%v'", p.Value()))
	}

	return success, nil
}

// GetMethod retrieves a callable object of a property of an ECMAScript language value.
// If no error is returned, the object is guaranteed to be callable.
// GetMethod is specified in 7.3.9.
func GetMethod(v Value, p StringOrSymbol) (Value, errors.Error) {
	f, err := GetV(v, p)
	if err != nil {
		return nil, err
	}

	if f == Undefined || f == Null {
		return Undefined, nil
	}

	if !InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f, nil
}

// HasProperty is used to determine whether an object has a property with the specified property
// key. The property can either be inherited or an own property of the object.
// HasProperty is specified in 7.3.10.
func HasProperty(o *Object, p StringOrSymbol) Boolean {
	return o.HasProperty(p)
}

// HasOwnProperty is used to determine whether an object has a property with the specified property
// key. The property can only be an own property of the object.
// HasProperty is specified in 7.3.11.
func HasOwnProperty(o *Object, p StringOrSymbol) Boolean {
	return o.GetOwnProperty(p) != nil
}

// Call is used to call the Call internal method of a function object.
// Call is specified in 7.3.12.
func Call(f *Object, thisValue Value, args ...Value) (Value, errors.Error) {
	if args == nil {
		args = []Value{}
	}

	if !InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f.Call(thisValue, args...)
}

// Construct is used to call the Construct internal method of a constructor object.
// Construct is specified in 7.3.13.
func Construct(f, newTarget *Object, args ...Value) (*Object, errors.Error) {
	if newTarget == nil {
		newTarget = f
	}

	if args == nil {
		args = []Value{}
	}

	return f.Construct(newTarget, args...)
}

// Integrity levels as specified in 7.3.14 in the specification.
const (
	IntegrityLevelSealed = "sealed"
	IntegrityLevelFrozen = "frozen"
)

// SetIntegrityLevel is used to fix the set of own properties of an object.
// SetIntegrityLevel is specified in 7.3.14.
func SetIntegrityLevel(o *Object, level string) (Boolean, errors.Error) {
	status := o.PreventExtensions()
	if !status {
		return False, nil
	}

	keys := o.OwnPropertyKeys()
	switch level {
	case IntegrityLevelSealed:
		return setIntegrityLevelSealed(o, keys)
	case IntegrityLevelFrozen:
		return setIntegrityLevelFrozen(o, keys)
	default:
		panic(fmt.Errorf("unknown integrity level '%v'", level))
	}
}

func setIntegrityLevelSealed(o *Object, keys []StringOrSymbol) (Boolean, errors.Error) {
	for _, k := range keys {
		p := NewProperty()
		p.SetField(FieldNameConfigurable, False)
		_, err := DefinePropertyOrThrow(o, k, p)
		if err != nil {
			return False, err
		}
	}
	return True, nil
}

func setIntegrityLevelFrozen(o *Object, keys []StringOrSymbol) (Boolean, errors.Error) {
	for _, k := range keys {
		p := o.GetOwnProperty(k)
		if p != nil {
			desc := NewProperty()
			if p.IsAccessorDescriptor() {
				desc.SetField(FieldNameConfigurable, False)
			} else {
				desc.SetField(FieldNameConfigurable, False)
				desc.SetField(FieldNameWritable, False)
			}
			_, err := DefinePropertyOrThrow(o, k, desc)
			if err != nil {
				return False, err
			}
		}
	}
	return True, nil
}

// TestIntegrityLevel is used to determine if the set of own properties of an object
// are fixed.
// TestIntegrityLevel is specified in 7.3.15.
func TestIntegrityLevel(o *Object, level string) Boolean {
	if o.IsExtensible() {
		return False
	}

	keys := o.OwnPropertyKeys()
	for _, k := range keys {
		desc := o.GetOwnProperty(k)
		if desc != nil {
			if desc.Configurable() {
				return False
			}
			if level == IntegrityLevelFrozen && desc.IsDataDescriptor() {
				if desc.Writable() {
					return false
				}
			}
		}
	}
	return true
}

// CreateArrayFromList creates an array whose elements are provided by a List.
// CreateArrayFromList is specified in 7.3.16.
func CreateArrayFromList(elements []Value) {
	panic("TODO: Arrays")
}

// CreateListFromArrayLike creates a List value whose elements are provided by the indexed
// properties of an array-like object.
// CreateListFromArrayLike is specified in 7.3.17.
func CreateListFromArrayLike(o *Object, elementTypes []Type) {
	panic("TODO: Arrays")
}

// Invoke is used to call a method property of an ECMAScript language value.
// Invoke is specified in 7.3.18.
func Invoke(v Value, p StringOrSymbol, args ...Value) (Value, errors.Error) {
	if len(args) == 0 {
		args = []Value{}
		_ = args
	}

	f, err := GetV(v, p)
	if err != nil {
		return nil, err
	}

	return Call(f.(*Object), v, args...)
}

// OrdinaryHasInstance implements the default algorithm for determining if an object o inherits
// from the instance object inheritance path provided by constructor c.
// OrdinaryHasInstance is specified in 7.3.19.
func OrdinaryHasInstance(c, o *Object) Boolean {
	panic("TODO")
}

// SpeciesConstructor retrieves the constructor that should be used to create new objects
// that are derived from the argument object o. The defaultConstructor argument is the
// constructor to use if a constructor's @@species property cannot be found starting from o.
// SpeciesConstructor is specified in 7.3.20.
func SpeciesConstructor() {
	panic("TODO")
}

// EnumerableOwnPropertyNames is specified in 7.3.21.
func EnumerableOwnPropertyNames() {
	panic("TODO")
}

// 7.3.22 GetFunctionRealm is not implemented here. You will find it in
// internal/runtime/realm/relam.go. The reason for that is, that otherwise,
// an import cycle lang -> realm -> lang rises.

// CopyDataProperties is specified in 7.3.23.
func CopyDataProperties() {
	panic("TODO")
}
