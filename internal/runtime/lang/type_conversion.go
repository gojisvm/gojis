package lang

import (
	"fmt"
	"math"

	"github.com/gojisvm/gojis/internal/runtime/errors"
)

// ToPrimitive attempts to convert the given input to a primitive value, or, if
// preferredType is not nil, to that preferred type.
// ToPrimitive is specified in 7.1.1.
func ToPrimitive(input Value, preferredType interface{}) (Value, errors.Error) {
	if input.Type() == TypeObject {
		var hint string

		t, ok := preferredType.(Type)
		if preferredType != nil && !ok {
			panic("preferredType is not a Type")
		}

		if preferredType == nil {
			hint = "default"
		} else if t == TypeString {
			hint = "string"
		} else if t == TypeNumber {
			hint = "number"
		}

		o, ok := input.(*Object)
		if !ok {
			panic("input is TypeObject, but not *Object")
		}

		exoticToPrim, err := GetMethod(o, NewStringOrSymbol(SymbolToPrimitive))
		if err != nil {
			return nil, err
		}

		if exoticToPrim != Undefined {
			result, callErr := Call(exoticToPrim.(*Object), input, NewString(hint))
			if callErr != nil {
				return nil, callErr
			}

			if result.Type() != TypeObject {
				return result, nil
			}

			return nil, errors.NewTypeError("Call of internal primitive conversion returned non-primitive object")
		}

		if hint == "default" {
			hint = "number"
		}

		val, err := OrdinaryToPrimitive(o, hint)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	return input, nil
}

// OrdinaryToPrimitive is specified in 7.1.1.1.
func OrdinaryToPrimitive(o *Object, hint string) (Value, errors.Error) {
	methodNames := []string{"valueOf", "toString"}
	if hint == "string" {
		methodNames = []string{"toString", "valueOf"}
	}

	for _, name := range methodNames {
		method, err := Get(o, NewStringOrSymbol(NewString(name)))
		if err != nil {
			return nil, err
		}

		if IsCallable(method) {
			result, err := Call(method.(*Object), o)
			if err != nil {
				return nil, err
			}

			if result.Type() != TypeObject {
				return result, nil
			}
		}
	}

	return nil, errors.NewTypeError("Cannot convert ordinary object to primitive")
}

// ToBoolean converts the argument to a Boolean value.
// ToBoolean is specified in 7.1.2.
func ToBoolean(arg Value) Boolean {
	switch arg.Type() {
	case TypeUndefined,
		TypeNull:
		return False
	case TypeSymbol,
		TypeObject:
		return True
	case TypeBoolean:
		return arg.(Boolean)
	case TypeNumber:
		if val := arg.Value(); val == PosZero.Value() ||
			val == NegZero.Value() ||
			arg == NaN {
			return False
		}
		return True
	case TypeString:
		if arg.Value() == "" {
			return False
		}
		return True
	}

	panic(unhandledType(arg))
}

// ToNumber converts the argument to a Number.
// ToNumber is specified in 7.1.3.
func ToNumber(arg Value) (Number, errors.Error) {
	switch arg.Type() {
	case TypeUndefined:
		return NaN, nil
	case TypeNull:
		return PosZero, nil
	case TypeBoolean:
		if arg.(Boolean) {
			return NewNumber(1), nil
		}
		return PosZero, nil
	case TypeNumber:
		return arg.(Number), nil
	case TypeString:
		panic("TODO: 7.1.3.1")
	case TypeSymbol:
		return Zero, errors.NewTypeError("Cannot convert from Symbol to Number")
	case TypeObject:
		primValue, err := ToPrimitive(arg, TypeNumber)
		if err != nil {
			return Zero, err
		}
		return ToNumber(primValue)
	}

	panic(unhandledType(arg))
}

// ToInteger converts the argument to an integer Number value.
// ToInteger is specified in 7.1.4.
func ToInteger(arg Value) (Number, errors.Error) {
	number, err := ToNumber(arg)
	if err != nil {
		return Zero, err
	}

	if number == NaN {
		return PosZero, nil
	}

	val := arg.Value()

	if val == PosZero.Value() ||
		val == NegZero.Value() ||
		arg == PosInfinity ||
		arg == NegInfinity {
		return arg.(Number), nil
	}

	return NewNumber(math.Floor(val.(float64))), nil
}

func toInt(arg Value, bits uint) (Number, errors.Error) {
	uintval, err := toUint(arg, bits)
	if err != nil {
		return Zero, err
	}

	float64val := uintval.Value().(float64)
	if float64val >= float64(int64(1)<<(bits-1)) {
		return NewNumber(float64val - float64(int64(1)<<bits)), nil
	}
	return uintval, nil
}

func toUint(arg Value, bits uint) (Number, errors.Error) {
	number, err := ToNumber(arg)
	if err != nil {
		return Zero, err
	}

	if number == NaN ||
		number == PosZero || number == NegZero ||
		number == PosInfinity || number == NegInfinity {
		return PosZero, nil
	}

	floatval := number.Value().(float64)
	intval := int64(floatval)
	intXXbit := intval % 1 << bits
	return NewNumber(float64(intXXbit)), nil
}

// ToInt32 converts the argument to an int32 Number value.
// ToInt32 is specified in 7.1.5.
func ToInt32(arg Value) (Number, errors.Error) {
	return toInt(arg, 32)
}

// ToUint32 converts the argument to an uint32 Number value.
// ToUint32 is specified in 7.1.6.
func ToUint32(arg Value) (Number, errors.Error) {
	return toUint(arg, 32)
}

// ToInt16 converts the argument to an int16 Number value.
// ToInt16 is specified in 7.1.7.
func ToInt16(arg Value) (Number, errors.Error) {
	return toInt(arg, 16)
}

// ToUint16 converts the argument to an uint16 Number value.
// ToUint16 is specified in 7.1.8.
func ToUint16(arg Value) (Number, errors.Error) {
	return toUint(arg, 16)
}

// ToInt8 converts the argument to an int8 Number value.
// ToInt8 is specified in 7.1.9.
func ToInt8(arg Value) (Number, errors.Error) {
	return toInt(arg, 8)
}

// ToUint8 converts the argument to an uint8 Number value.
// ToUint8 is specified in 7.1.10.
func ToUint8(arg Value) (Number, errors.Error) {
	return toUint(arg, 8)
}

// ToUint8Clamp converts the argument to one of 28 integer values in the range 0
// through 255, inclusive.
// ToUint8Clamp is specified in 7.1.11.
func ToUint8Clamp(arg Value) (Number, errors.Error) {
	number, err := ToNumber(arg)
	if err != nil {
		return Zero, nil
	}

	if number == NaN {
		return PosZero, nil
	}

	floatval := number.Value().(float64)
	if floatval <= 0 {
		return PosZero, nil
	}
	if floatval >= 255 {
		return NewNumber(255), nil
	}

	f := math.Floor(floatval)
	if f+0.5 < floatval {
		return NewNumber(floatval + 1), nil
	}
	if floatval < f+0.5 {
		return NewNumber(f), nil
	}

	// f == floatval, also f is an integer because of math.Floor
	if int64(f)%2 != 0 {
		return NewNumber(f + 1), nil
	}
	return NewNumber(f), nil
}

// ToString converts the argument to a String.
// ToString is specified in 7.1.12.
func ToString(arg Value) String {
	panic("TODO")
}

// NumberToString converts the given Number to a String.
// NumberToString is specified in 7.1.12.1.
func NumberToString(n Number) String {
	panic("TODO")
}

// ToObject converts the given argument to an Object.
// ToObject is specified in 7.1.13.
func ToObject(arg Value) *Object {
	panic("TODO")
}

// ToPropertyKey converts the given argument to a StringOrSymbol.
// ToPropertyKey is specified in 7.1.14.
func ToPropertyKey(arg Value) StringOrSymbol {
	panic("TODO")
}

// ToLength converts argument to an integer suitable for use as the length of an
// array-like object.
// ToLength is specified in 7.1.15.
func ToLength(arg Value) Number {
	panic("TODO")
}

// CanonicalNumericIndexString returns argument converted to a numeric value if
// it is a String representation of a Number that would be produced by ToString,
// or the string "-0". Otherwise, it returns Undefined.
// CanonicalNumericIndexString is specified in 7.1.16.
func CanonicalNumericIndexString(arg Value) Number {
	panic("TODO")
}

// ToIndex returns value argument converted to a numeric value if it is a valid
// integer index value.
// ToIndex is specified in 7.1.17.
func ToIndex(arg Value) Number {
	panic("TODO")
}

func unhandledType(arg Value) error {
	return fmt.Errorf("Unhandled type in type conversion: '%v'", arg.Type())
}
