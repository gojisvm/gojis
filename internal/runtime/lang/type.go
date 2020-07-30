package lang

import "fmt"

// Type represents a language type as specified by the
// ECMAScript Language Types.
type Type uint8

// Language Type as specified in 6.1
const (
	TypeInternal Type = iota

	TypeUndefined
	TypeNull
	TypeBoolean
	TypeString
	TypeSymbol
	TypeNumber
	TypeObject
)

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "Undefined"
	case TypeNull:
		return "Null"
	case TypeBoolean:
		return "Boolean"
	case TypeString:
		return "String"
	case TypeSymbol:
		return "Symbol"
	case TypeNumber:
		return "Number"
	case TypeObject:
		return "Object"
	default:
		return "Unknown"
	}
}

// EnsureTypeOneOf panics, if the type of the given value is not one of
// the given ECMALanguage data types.
func EnsureTypeOneOf(arg Value, ts ...Type) {
	if !TypeIsOneOf(arg, ts...) {
		panic(fmt.Errorf("value's type must be one of %v, but was %v", ts, arg.Type()))
	}
}

// TypeIsOneOf is used to determine whether the type of the given value
// is one of the given ECMALanguage data types.
func TypeIsOneOf(arg Value, ts ...Type) bool {
	got := arg.Type()
	for _, t := range ts {
		if t == got {
			return true
		}
	}
	return false
}
