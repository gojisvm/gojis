package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseScript(i *isolate, p param) *ast.Script {
	return &ast.Script{
		ScriptBody: parseScriptBody(i, p),
	}
}

func parseScriptBody(i *isolate, p param) *ast.ScriptBody {
	return &ast.ScriptBody{
		StatementList: parseStatementList(i, p),
	}
}

func parseArguments(i *isolate, p param) *ast.Arguments {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.ParOpen) {
		i.restore(chck)
		return nil
	}

	argList := parseArgumentList(i, p.only(pYield|pAwait))
	if argList == nil {
		i.restore(chck)
		return nil
	}

	comma := i.acceptOneOfTypes(token.Comma) // the comma is effectively optional, we just need to get the return value to set it in the ast node

	if !i.acceptOneOfTypes(token.ParClose) {
		i.restore(chck)
		return nil
	}

	return &ast.Arguments{
		ArgumentList: argList,
		Comma:        comma,
	}
}
