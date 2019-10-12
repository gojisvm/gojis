package runtime

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

type ScriptRecord struct {
	Realm          *realm.Realm
	Environment    binding.Environment
	ECMAScriptCode *ast.Script
	HostDefined    interface{}
}

func (r *ScriptRecord) Type() lang.Type    { return lang.TypeInternal }
func (r *ScriptRecord) Value() interface{} { return r }
