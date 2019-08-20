package gojis

// Object represents any ECMAScript language value. This can be a String or a
// Number as well as Null or Undefined. To check if the object represents Null,
// use Object#IsNull. To check if the object represents Undefined, use
// Object#IsUndefined.
type Object interface {
	// Lookup returns a property of this object with the given name. If no such
	// property exists, Undefined will be returned. This method never returns nil.
	Lookup(string) Object
	// SetFunction adds a property (more specific, a function object) to this
	// object. When the function object is invoked, the given function will be
	// executed, with all parameters wrapped into the Args object.
	SetFunction(string, func(Args) Object)
	// CallWithArgs attempts to invoke this objects 'Call' property. Note that this
	// property is only set if this object is a function or constructor object. If
	// this object does not have a 'Call' property that is callable, an error will
	// be returned.
	CallWithArgs(...interface{}) (Object, error)
	// SetObbject adds a property with the given name to this object. The property's
	// value will be the given object.
	SetObject(string, Object)

	// IsUndefined is used to determine whether this object represents the Undefined
	// value.
	IsUndefined() bool
	// IsNull is used to determine whether this object represents the Null value.
	IsNull() bool
	// IsFunction is used to determine whether this object can be invoked. If this
	// returns true, CallWithArgs will not return an error.
	IsFunction() bool

	// Type returns the ECMAScript language type of this object.
	Type() Type
	// Value returns the Go value corresponding to the ECMAScript language value of
	// this object.
	Value() interface{}
}
