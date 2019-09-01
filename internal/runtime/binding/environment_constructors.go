package binding

import "github.com/gojisvm/gojis/internal/runtime/lang"

// GetIdentifierReference returns a named reference with n as the name. If
// strict is true, the reference will be a strict reference. The given
// environment must be of type Environment or Null. If the environment does not
// have a binding with the given name, its outer environment is used.
// GetIdentifierReference is specified in 8.1.2.1.
func GetIdentifierReference(env lang.InternalValue, n lang.String, strict bool) *Reference {
	if env == lang.Null {
		return NewReference(lang.NewStringOrSymbol(n), lang.Undefined, strict)
	}

	lex := env.(Environment)
	if lex.HasBinding(n) {
		return NewReference(lang.NewStringOrSymbol(n), lex, strict)
	}

	return GetIdentifierReference(lex.Outer(), n, strict)
}

// NewDeclarativeEnvironment creates a new declarative environment with the
// given outer environment. NewDeclarativeEnvironment is specified in 8.1.2.2.
func NewDeclarativeEnvironment(outer Environment) *DeclarativeEnvironment {
	env := new(DeclarativeEnvironment)
	env.outer = outer
	env.bindings = make(map[string]*Binding)
	return env
}

// NewObjectEnvironment creates a new object environment with the given outer
// environment for the given object. NewObjectEnvironment is specified in
// 8.1.2.3.
func NewObjectEnvironment(outer Environment, obj *lang.Object) *ObjectEnvironment {
	env := new(ObjectEnvironment)
	env.outer = outer
	env.bindingObject = obj
	return env
}

// NewFunctionEnvironment creates a new function environment for the given
// function object. The passed function object can either be Undefined or an
// Object. NewFunctionEnvironment is specified in 8.1.2.4.
func NewFunctionEnvironment(f *lang.Object, newTarget lang.Value) *FunctionEnvironment {
	lang.EnsureTypeOneOf(f, lang.TypeObject, lang.TypeUndefined)

	env := new(FunctionEnvironment)
	env.FunctionObject = f
	env.ThisBindingStatus = StatusUninitialized
	if f.ThisMode == lang.FunctionModeLexical {
		env.ThisBindingStatus = StatusLexical
	}

	env.HomeObject = f.HomeObject
	env.NewTarget = newTarget
	env.outer = f.Environment.(Environment)
	return env
}

// NewGlobalEnvironment creates a new global environment with the given global
// object and the given thisValue as this value. NewGlobalEnvironment is
// specified in 8.1.2.5.
func NewGlobalEnvironment(globalObj, thisValue *lang.Object) *GlobalEnvironment {
	env := new(GlobalEnvironment)
	objRec := NewObjectEnvironment(env, globalObj)
	dclRec := NewDeclarativeEnvironment(env)
	env.ObjectRecord = objRec
	env.GlobalThisValue = thisValue
	env.DeclarativeRecord = dclRec
	env.VarNames = make([]string, 0)
	return env
}

// NewModuleEnvironment creates a new module environment with the given
// environment as outer environment.
func NewModuleEnvironment(outer Environment) *ModuleEnvironment {
	env := new(ModuleEnvironment)
	env.outer = outer
	return env
}
