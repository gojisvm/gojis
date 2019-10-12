package gojis

import (
	"github.com/gojisvm/gojis/internal/runtime"
)

// Type represents the ECMAScript language types.
type Type uint8

// Available types. TypeUnknown must not be used. If it appears, this indicates,
// that somewhere the type was not set (it is the default value for Type).
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

// VM represents an instance of the GojisVM. It can be used to evaluate
// ECMAScript code.
type VM struct {
	Object // the global object

	runtime runtime.Runtime
}

// NewVM creates a new, initialized VM that is ready to use.
func NewVM() *VM {
	return &VM{
		Object: nil,
	}
}

func (vm *VM) Eval(script string) error {
	panic("TODO")
}

func (vm *VM) EvalFile(path string) error {
	panic("TODO")
}

// SetConsole is used to change the console of the VM. Calls like 'console.log'
// will be written to the given Console.
func (vm *VM) SetConsole(console Console) {
	vm.Object.SetObject("console", console)
}
