package parser

import "github.com/gojisvm/gojis/internal/parser/ast"

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
