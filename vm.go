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

	runtime *runtime.Runtime

	// configuration

	// debug indicates that the VM should output diagnostic messages
	debug bool
}

// Option represents configuration for a VM instance.
type Option func(*VM)

// NewVM creates a new, initialized VM that is ready to use.
func NewVM(opts ...Option) *VM {
	vm := new(VM)
	for _, opt := range opts {
		opt(vm)
	}

	return vm
}

// Opts represents functional options.
var Opts = optStruct{}

type optStruct struct{}

func (optStruct) Debug(vm *VM) {
	vm.debug = true
}

// ------------------

// Eval delegates to the builtin eval function of this VM's global object.
func (vm *VM) Eval(script string) (Object, error) {
	return vm.Object.Lookup("eval").CallWithArgs(script)
}

// SetConsole is used to change the console of the VM. Calls like 'console.log'
// will be written to the given Console.
func (vm *VM) SetConsole(console Console) {
	vm.Object.SetObject("console", console)
}
