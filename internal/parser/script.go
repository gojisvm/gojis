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

func parseElision(i *isolate, p param) *ast.Elision {
	elision := new(ast.Elision)

	for i.acceptOneOfTypes(token.Comma) {
		elision.Commas++
	}

	return elision
}

func parseInitializer(i *isolate, p param) *ast.Initializer {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Assign) {
		if assignmentExpr := parseAssignmentExpression(i, p.only(pIn|pYield|pAwait)); assignmentExpr != nil {
			return &ast.Initializer{
				AssignmentExpression: assignmentExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parsePropertyName(i *isolate, p param) *ast.PropertyName {
	chck := i.checkpoint()

	if literalPropertyName := parseLiteralPropertyName(i, 0); literalPropertyName != nil {
		return &ast.PropertyName{
			LiteralPropertyName: literalPropertyName,
		}
	} else if computedPropertyName := parseComputedPropertyName(i, p.only(pYield|pAwait)); computedPropertyName != nil {
		return &ast.PropertyName{
			ComputedPropertyName: computedPropertyName,
		}
	}

	i.restore(chck)
	return nil
}

func parseLiteralPropertyName(i *isolate, p param) *ast.LiteralPropertyName {
	chck := i.checkpoint()

	if t, ok := i.accept(token.IdentifierName); ok {
		return &ast.LiteralPropertyName{
			IdentifierName: t.Value,
		}
	} else if t, ok := i.accept(token.StringLiteral); ok {
		return &ast.LiteralPropertyName{
			StringLiteral: t.Value,
		}
	} else if numericLiteral := parseNumericLiteral(i, 0); numericLiteral != nil {
		return &ast.LiteralPropertyName{
			NumericLiteral: numericLiteral,
		}
	}

	i.restore(chck)
	return nil
}
