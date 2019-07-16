package runtime

import (
	"github.com/rs/zerolog"
)

// Runtime is an object that will evaluate
// an AST.
type Runtime struct {
	log zerolog.Logger

	ast interface{}
}

// New creates a new runtime using the given logger and
// evaluating the given AST.
func New(log zerolog.Logger, ast interface{}) *Runtime {
	r := new(Runtime)
	r.log = log
	r.ast = ast
	return r
}
