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

func parseArgumentList(i *isolate, p param) *ast.ArgumentList {
	chck := i.checkpoint()

	ellipsis := i.acceptOneOfTypes(token.Ellipsis) // ellipsis is optional
	if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
		return &ast.ArgumentList{
			AssignmentExpression: assignmentExpr,
			Ellipsis:             true,
		}
	} else if !ellipsis { // of no assignment expression could be parsed, there musn't be an ellipsis before the recursive argument list
		if argList := parseArgumentList(i, p.only(pYield|pAwait)); argList != nil {
			if i.acceptOneOfTypes(token.Comma) {
				i.acceptOneOfTypes(token.Ellipsis) // ellipsis is optional
				if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
					return &ast.ArgumentList{
						ArgumentList:         argList,
						AssignmentExpression: assignmentExpr,
						Ellipsis:             true,
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
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

func parseComputedPropertyName(i *isolate, p param) *ast.ComputedPropertyName {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.BracketOpen) {
		if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
			if i.acceptOneOfTypes(token.BracketClose) {
				return &ast.ComputedPropertyName{
					AssignmentExpression: assignmentExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseFormalParameters(i *isolate, p param) *ast.FormalParameters {
	chck := i.checkpoint()

	if functionRestParameter := parseFunctionRestParameter(i, p.only(pYield|pAwait)); functionRestParameter != nil {
		return &ast.FormalParameters{
			FunctionRestParameter: functionRestParameter,
		}
	} else if formalParameterList := parseFormalParameterList(i, p.only(pYield|pAwait)); formalParameterList != nil {
		if i.acceptOneOfTypes(token.Comma) {
			if functionRestParameter := parseFunctionRestParameter(i, p.only(pYield|pAwait)); functionRestParameter != nil {
				return &ast.FormalParameters{
					FormalParameterList:   formalParameterList,
					FunctionRestParameter: functionRestParameter,
					Comma:                 true,
				}
			}
			return &ast.FormalParameters{
				FormalParameterList: formalParameterList,
				Comma:               true,
			}
		}
		return &ast.FormalParameters{
			FormalParameterList: formalParameterList,
		}
	}

	i.restore(chck)
	return &ast.FormalParameters{}
}

func parseFunctionRestParameter(i *isolate, p param) *ast.FunctionRestParameter {
	if bindingRestElement := parseBindingRestElement(i, p.only(pYield|pAwait)); bindingRestElement != nil {
		return &ast.FunctionRestParameter{
			BindingRestElement: bindingRestElement,
		}
	}
	return nil
}

func parseFormalParameterList(i *isolate, p param) *ast.FormalParameterList {
	chck := i.checkpoint()

	var params []*ast.FormalParameter

	first := parseFormalParameter(i, p.only(pYield|pAwait)) // expression consists of at least one entry
	if first == nil {
		return nil
	}
	params = append(params, first)

	for { // parse until there are no more formal parameter parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(chck)
			return nil
		}

		next := parseFormalParameter(i, p.only(pYield|pAwait))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no formal parameter, so reset to before comma
			break
		}
		params = append(params, next)
	}

	return &ast.FormalParameterList{
		FormalParameters: params,
	}
}

func parseFormalParameter(i *isolate, p param) *ast.FormalParameter {
	if bindingElement := parseBindingElement(i, p.only(pYield|pAwait)); bindingElement != nil {
		return &ast.FormalParameter{
			BindingElement: bindingElement,
		}
	}
	return nil
}
