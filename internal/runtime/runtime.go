package runtime

// Runtime is a struct that is used to evaluate an AST.
type Runtime struct {
	sourceText string
}

// New creates a new ready to evaluate runtime that will evaluate the given
// source text.
func New(sourceText string) *Runtime {
	return &Runtime{
		sourceText: sourceText,
	}
}

// Evaluate evaluates the runtime's AST.
func (r *Runtime) Evaluate() error {
	panic("TODO")
}
