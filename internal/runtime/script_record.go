package runtime

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

// ScriptRecord encapsulates information about a script being evaluated.
// ScriptRecord is specified in 15.1.8.
type ScriptRecord struct {
	Realm          *realm.Realm
	Environment    binding.Environment
	ECMAScriptCode *ast.Script
	HostDefined    interface{}
}

// Type returns lang.TypeInternal.
func (r *ScriptRecord) Type() lang.Type { return lang.TypeInternal }

// Value returns the ScriptRecord itself.
func (r *ScriptRecord) Value() interface{} { return r }
