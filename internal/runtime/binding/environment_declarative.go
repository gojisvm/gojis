package binding

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

var _ Environment = (*DeclarativeEnvironment)(nil)

// DeclarativeEnvironment represents a declarative environment record
// as described in 8.1.1.1.
type DeclarativeEnvironment struct {
	outer Environment

	bindings map[string]*Binding
}

// Outer returns the outer environment of this declarative environment.
func (e *DeclarativeEnvironment) Outer() Environment {
	return e.outer
}

func (e *DeclarativeEnvironment) setBinding(n lang.String, b *Binding) {
	e.bindings[n.Value().(string)] = b
}

func (e *DeclarativeEnvironment) getBinding(n lang.String) (*Binding, bool) {
	val, ok := e.bindings[n.Value().(string)]
	return val, ok
}

func (e *DeclarativeEnvironment) deleteBinding(n lang.String) {
	delete(e.bindings, n.Value().(string))
}

// mustGetBinding returns the binding with the given name.
// This function panics if no binding could be found.
// Only use this if you know what you are doing.
func (e *DeclarativeEnvironment) mustGetBinding(n lang.String) *Binding {
	val, ok := e.getBinding(n)
	if !ok {
		panic(fmt.Errorf("Missing binding for name '%v'", n.Value()))
	}

	return val
}

// HasBinding determines whether the given name is an identifier bound
// by the record.
// HasBinding is specified in 8.1.1.1.1.
func (e *DeclarativeEnvironment) HasBinding(n lang.String) bool {
	_, ok := e.getBinding(n)
	return ok
}

// CreateMutableBinding creates a new mutable binding for the name N that is uninitialized.
// A binding must not already exist in this Environment Record for the given name.
// If Boolean argument deletable is true the new binding can be deleted by a subsequent
// DeleteBinding call.
// CreateMutableBinding is specified in 8.1.1.1.2.
func (e *DeclarativeEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	b := NewBinding(n)
	b.deletable = deletable

	e.setBinding(n, b)

	return nil
}

// CreateImmutableBinding creates a new immutable binding for the given name that is uninitialized.
// A binding must not already exist in this Environment
// Record with that name. If strict is true the new binding is marked as a strict binding.
// CreateImmutableBinding is specified in 8.1.1.1.3.
func (e *DeclarativeEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	b := NewBinding(n)
	b.immutable = true
	b.strict = strict

	e.setBinding(n, b)

	return nil
}

// InitializeBinding is used to set the bound value of the current binding of the environment
// with the given name to the given value. An uninitialized binding for the given name must
// already exist.
// InitializeBinding is specified in 8.1.1.1.4.
func (e *DeclarativeEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	e.mustGetBinding(n).value = val

	return nil
}

// SetMutableBinding attempts to change the bound value of the binding with the given name in this
// environment to the given new value.
// If the binding is an immutable binding AND strict is true,
// a TypeError is raised.
// If no binding with the given name exists AND strict is true, a ReferenceError is raised.
// If the binding with the given name is not yet initialized, a ReferenceError is raised.
// SetMutableBinding is specified in 8.1.1.1.5.
func (e *DeclarativeEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	if !e.HasBinding(n) {
		if strict {
			return errors.NewReferenceError("Cannot set mutable binding in strict mode if binding does not exist")
		}

		_ = e.CreateMutableBinding(n, true)
		_ = e.InitializeBinding(n, val)
		return nil
	}

	b := e.mustGetBinding(n)
	if b.IsStrict() {
		strict = true
	}

	if !b.IsInitialized() {
		return errors.NewReferenceError(fmt.Sprintf("The binding for '%v' has not been initialized yet", n))
	} else if !b.IsImmutable() {
		b.Set(val)
	} else {
		if strict {
			return errors.NewTypeError(fmt.Sprintf("Attempting to change value of an immutable binding '%v'", n))
		}
	}

	return nil
}

// GetThisBinding is not supported for DeclarativeEnvironment.
// It will panic.
func (e *DeclarativeEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	panic("Not supported for DeclarativeEnvironment") // #34: better abstraction of environments?
}

// GetBindingValue returns the value of its bound identifier whose name is the given name.
// GetBindingValue is specified in 8.1.1.1.6.
func (e *DeclarativeEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	b := e.mustGetBinding(n)
	if !b.IsInitialized() {
		return nil, errors.NewReferenceError(fmt.Sprintf("Binding '%v' has not been initialized yet", n))
	}

	return b.Value(), nil
}

// DeleteBinding deletes a binding with the given name if the binding is deletable.
// If the binding is not deletable, this function returns false.
// If the binding does not exist or is deletable, this function returns true.
// DeleteBinding is specified in 8.1.1.1.7.
func (e *DeclarativeEnvironment) DeleteBinding(n lang.String) bool {
	b, ok := e.getBinding(n)
	if !ok {
		return true
	}

	if !b.IsDeletable() {
		return false
	}

	e.deleteBinding(n)
	return true
}

// HasThisBinding returns false.
// HasThisBinding is specified in 8.1.1.1.8.
func (e *DeclarativeEnvironment) HasThisBinding() bool {
	return false
}

// HasSuperBinding returns false.
// HasSuperBinding is specified in 8.1.1.1.9.
func (e *DeclarativeEnvironment) HasSuperBinding() bool {
	return false
}

// WithBaseObject returns Undefined.
// WithBaseObject is specified in 8.1.1.1.10.
func (e *DeclarativeEnvironment) WithBaseObject() lang.Value {
	return lang.Undefined
}

// Type returns lang.TypeInternal.
func (e *DeclarativeEnvironment) Type() lang.Type { return lang.TypeInternal }

// Value returns the declarative environment itself.
func (e *DeclarativeEnvironment) Value() interface{} { return e }
