package runtime

import (
	"strings"

	"github.com/gojisvm/gojis/internal/parser"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

type isolate struct {
	parser *parser.Parser
}

func newIsolate() *isolate {
	return &isolate{
		parser: parser.New(),
	}
}

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

func (i *isolate) scriptEvaluation(scriptRecord *ScriptRecord) {
	panic("TODO")
}
