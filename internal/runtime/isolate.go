package runtime

import (
	"strings"

	"github.com/gojisvm/gojis/internal/parser"
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/runtime/agent"
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

type isolate struct {
	parser *parser.Parser
	agent  *agent.Agent // the surrounding agent
}

func newIsolate() *isolate {
	return &isolate{
		parser: parser.New(),
		agent:  agent.New(),
	}
}

// parseScript creates a Script Record based upon the result of parsing
// sourceText as a Script. parseScript is specified in 15.1.9.
func (i *isolate) parseScript(sourceText string, realm *realm.Realm, hostDefined interface{}) (*ScriptRecord, error) {
	script, err := i.parser.Parse("source text", strings.NewReader(sourceText))
	if err != nil {
		return nil, err
	}

	return &ScriptRecord{
		Realm:          realm,
		Environment:    nil,
		ECMAScriptCode: script,
		HostDefined:    hostDefined,
	}, nil
}

// scriptEvaluation evaluates a given script record. scriptEvaluation is
// specified in 15.1.10.
func (i *isolate) scriptEvaluation(scriptRecord *ScriptRecord) {
	globalEnv, ok := scriptRecord.Realm.GlobalEnv.(binding.Environment)
	if !ok {
		panic("Global environment of script record's realm is not an environment")
	}

	scriptCtx := &agent.ExecutionContext{
		Function:            lang.Null,
		Realm:               scriptRecord.Realm,
		ScriptOrModule:      scriptRecord,
		VariableEnvironment: globalEnv,
		LexicalEnvironment:  globalEnv,
	}

	// TODO: suspend the currently running execution context

	i.agent.ExecutionContextStack.Push(scriptCtx)
	scriptBody := scriptRecord.ECMAScriptCode
	i.globalDeclarationInstantiation(scriptBody, globalEnv)
	panic("TODO")
}

func (i *isolate) globalDeclarationInstantiation(script *ast.Script, env binding.Environment) {

}
