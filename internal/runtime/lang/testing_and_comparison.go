package lang

import (
	"fmt"
	"strings"

	"github.com/gojisvm/gojis/internal/runtime/errors"
)

// RequireObjectCoercible returns the value and no error
// if the Value is coercible, meaning that the type of the
// value is neither Null nor Undefined.
// RequireObjectCoercible is specified in 7.2.1.
func RequireObjectCoercible(arg Value) (Value, errors.Error) {
	switch arg.Type() {
	case TypeNull,
		TypeUndefined:
		return nil, errors.NewTypeError("Object is not coercible")
	case TypeBoolean,
		TypeNumber,
		TypeString,
		TypeSymbol,
		TypeObject:
		return arg, nil
	}

	panic(fmt.Errorf("unhandled argument type: %v", arg.Type()))
}

// IsArray is used to determine whether the given value is an array.
// To obtain a bool value, use InternalIsArray instead.
// IsArray is specified in 7.2.2.
func IsArray(arg Value) Boolean { return Boolean(InternalIsArray(arg)) }

// InternalIsArray is used to determine whether the given value is an array.
func InternalIsArray(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	panic("TODO: Arrays")
}

// IsCallable is used to determine whether the value has a Call internal method.
// If this function returns True, the value can be used as an argument to
// the function Call.
// To obtain a bool value, use InternalIsCallable instead.
// IsCallable is specified in 7.2.3.
func IsCallable(arg Value) Boolean { return Boolean(InternalIsCallable(arg)) }

// InternalIsCallable is used to determine whether the value has a Call internal method.
// If this function returns true, the value can be used as an argument to
// the function Call.
func InternalIsCallable(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	if fObj, ok := arg.(*Object); ok {
		return fObj.Call != nil
	}
	return false
}

// IsConstructor is used to determine whether the value has a Construct internal method.
// To obtain a bool value, use InternalIsConstructor instead.
// IsConstructor is specified in 7.2.3.
func IsConstructor(arg Value) Boolean { return Boolean(InternalIsConstructor(arg)) }

// InternalIsConstructor is used to determine whether the value has a Construct internal method.
func InternalIsConstructor(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	if cObj, ok := arg.(*Object); ok {
		return cObj.Construct != nil
	}
	return false
}

// IsExtensible is used to determine whether the object is extensible,
// meaning if new properties can be added.
// IsExtensible is specified in 7.2.5.
func IsExtensible(o *Object) Boolean { return o.IsExtensible() }

// InternalIsExtensible is used to determine whether the object is extensible,
// meaning if new properties can be added.
func InternalIsExtensible(o *Object) bool {
	return IsExtensible(o).Value().(bool)
}

// IsInteger is used to determine whether the value is an integer.
// To obtain a bool value, use InternalIsInteger instead.
// IsInteger is specified in 7.2.6.
func IsInteger(arg Value) Boolean { return Boolean(InternalIsInteger(arg)) }

// InternalIsInteger is used to determine whether the value is an integer.
func InternalIsInteger(arg Value) bool {
	if arg.Type() != TypeNumber {
		return false
	}

	if arg == NaN ||
		arg == PosInfinity ||
		arg == NegInfinity {
		return false
	}

	val := arg.Value().(float64)
	return val == float64(int(val))
}

// IsPropertyKey is used to determine whether the type of the value is String or Symbol.
// To obtain a bool value, use InternalIsPropertyKey instead.
// IsPropertyKey is specified in 7.2.7.
func IsPropertyKey(arg Value) Boolean { return Boolean(InternalIsPropertyKey(arg)) }

// InternalIsPropertyKey is used to determine whether the type of the value is String or Symbol.
func InternalIsPropertyKey(arg Value) bool {
	return arg.Type() == TypeString ||
		arg.Type() == TypeSymbol
}

// IsRegExp is used to determine whether the value has a @@match property, or, if not,
// if it has a RegExpMatcher internal slot.
// To obtain a bool value, use InternalIsRegExp instead.
// IsRegExp is specified in 7.2.8.
func IsRegExp(arg Value) Boolean { return Boolean(InternalIsRegExp(arg)) }

// InternalIsRegExp is used to determine whether the value has a @@match property, or, if not,
// if it has a RegExpMatcher internal slot.
func InternalIsRegExp(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	panic("TODO: 7.2.8 InternalIsRegExp")
}

// IsStringPrefix is used to determine whether p is a prefix of q or not.
// To obtain a bool value, use InternalIsStringPrefix instead.
// IsStringPrefix is specified in 7.2.9.
func IsStringPrefix(p, q String) Boolean { return Boolean(InternalIsStringPrefix(p, q)) }

// InternalIsStringPrefix is used to determine whether p is a prefix of q or not.
func InternalIsStringPrefix(p, q String) bool {
	pVal, qVal := p.Value().(string), q.Value().(string)
	return strings.HasPrefix(qVal, pVal)
}

// SameValue is used to determine, whether x and y have the same value.
// To obtain a bool value, use InternalSameValue instead.
// SameValue is specified in 7.2.10.
func SameValue(x, y Value) Boolean { return Boolean(InternalSameValue(x, y)) }

// InternalSameValue is used to determine, whether x and y have the same value.
func InternalSameValue(x, y Value) bool {
	if x.Type() != y.Type() {
		return false
	}

	if x.Type() == TypeNumber {
		if x == NaN && y == NaN {
			return true
		}

		if x == PosZero && y == NegZero {
			return false
		}

		if x == NegZero && y == PosZero {
			return false
		}

		return x.Value() == y.Value()
	}

	return InternalSameValueNonNumber(x, y)
}

// SameValueZero is used to determine, whether x and y have the same value.
// The difference to SameValue is, that here, +0 == -0.
// To obtain a bool value, use InternalSameValueZero instead.
// SameValueZero is specified in 7.2.11.
func SameValueZero(x, y Value) Boolean { return Boolean(InternalSameValueZero(x, y)) }

// InternalSameValueZero is used to determine, whether x and y have the same value.
// The difference to InternalSameValue is, that here, +0 == -0.
func InternalSameValueZero(x, y Value) bool {
	if x.Type() != y.Type() {
		return false
	}
	if x.Type() != y.Type() {
		return false
	}

	if x.Type() == TypeNumber {
		if x == NaN && y == NaN {
			return true
		}

		if x == PosZero && y == NegZero {
			return true
		}

		if x == NegZero && y == PosZero {
			return true
		}

		return x.Value() == y.Value()
	}

	return InternalSameValueNonNumber(x, y)
}

// SameValueNonNumber is used to determine, whether x and y have the same value,
// assuming their type is not Number.
// To obtain a bool value, use InternalSameValueNonNumber instead.
// SameValueNonNumber is specified in 7.2.12.
func SameValueNonNumber(x, y Value) Boolean {
	return Boolean(InternalSameValueNonNumber(x, y))
}

// InternalSameValueNonNumber is used to determine, whether x and y have the same value,
// assuming their type is not Number.
func InternalSameValueNonNumber(x, y Value) bool {
	if x.Type() == TypeUndefined {
		return true
	}

	if x.Type() == TypeNull {
		return true
	}

	if x.Type() == TypeString {
		xVal, yVal := x.Value().(String), y.Value().(String)
		return StringsEqual(xVal, yVal)
	}

	if x.Type() == TypeBoolean {
		return x.Value() == y.Value()
	}

	if x.Type() == TypeSymbol {
		return x.Value() == y.Value()
	}

	return x.Value() == y.Value()
}
