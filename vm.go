package gojis

// Type represents the ECMAScript language types.
type Type uint8

// Available types. TypeUnknown must not be used.
// If it appears, this indicates, that somewhere the type was not set
// (it is the default value for Type).
const (
	TypeUnknown = iota
	TypeUndefined
	TypeNull
	TypeBoolean
	TypeString
	TypeSymbol
	TypeNumber
	TypeObject
)

// VM represents an instance of the GojisVM.
// It can be used to evaluate ECMAScript code.
type VM struct {
	Object // the global object
}

// NewVM creates a new, initialized VM that is ready to use.
func NewVM() *VM {
	panic("TODO")
}

// Eval evaluates the given ECMAScript code, and returns an Object, representing the
// result of the evaluation. The result may be Null or Undefined.
//
// Internally, this function directly delegates to the method 'eval' that is specified
// by the ECMAScript language specification.
func (vm *VM) Eval(script string) Object {
	return vm.Lookup("eval")
}

// SetConsole is used to change the console of the VM.
// Calls like 'console.log' will be written to the given Console.
func (vm *VM) SetConsole(console Console) {
	vm.Object.SetObject("console", console)
}
