package runtime

type Runtime struct {
	sourceText string
}

func New(sourceText string) *Runtime {
	return &Runtime{
		sourceText: sourceText,
	}
}

func (r *Runtime) Evaluate() error {
	panic("TODO")
}
