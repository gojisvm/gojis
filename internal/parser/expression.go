package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseExpression(i *isolate, p param) *ast.Expression {
	chck := i.checkpoint()

	var decls []*ast.AssignmentExpression

	first := parseAssignmentExpression(i, p.only(pYield|pAwait|pReturn)) // expression consists of at least one entry
	if first == nil {
		return nil
	}
	decls = append(decls, first)

	for { // parse until there are no more assignment expressions parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(chck)
			return nil
		}

		next := parseAssignmentExpression(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no assignment expression, so reset to before comma
			break
		}
		decls = append(decls, next)
	}

	return &ast.Expression{
		AssignmentExpressions: decls,
	}
}

func parseAssignmentExpression(i *isolate, p param) *ast.AssignmentExpression {
	chck := i.checkpoint()

	if conditionalExpression := parseConditionalExpression(i, p.only(pIn|pYield|pAwait)); conditionalExpression != nil {
		return &ast.AssignmentExpression{
			ConditionalExpression: conditionalExpression,
		}
	}

	if p.is(pYield) {
		if yieldExpression := parseYieldExpression(i, p.only(pIn|pAwait)); yieldExpression != nil {
			return &ast.AssignmentExpression{
				YieldExpression: yieldExpression,
			}
		}
	}

	if arrowFunction := parseArrowFunction(i, p.only(pIn|pYield|pAwait)); arrowFunction != nil {
		return &ast.AssignmentExpression{
			ArrowFunction: arrowFunction,
		}
	}

	if asyncArrowFunction := parseAsyncparseArrowFunction(i, p.only(pIn|pYield|pAwait)); asyncArrowFunction != nil {
		return &ast.AssignmentExpression{
			AsyncArrowFunction: asyncArrowFunction,
		}
	}

	if leftHandSideExpression := parseLeftHandSideExpression(i, p.only(pYield|pAwait)); leftHandSideExpression != nil {
		if t, ok := i.acceptOneOf(token.Assign,
			token.MultiplyAssign,
			token.DivAssign,
			token.ModuloAssign,
			token.PlusAssign,
			token.MinusAssign,
			token.LeftShiftAssign,
			token.RightShiftAssign,
			token.UnsignedRightShiftAssign,
			token.AndAssign,
			token.XorAssign,
			token.OrAssign,
			token.PowerAssign); ok {
			if assignmentExpression := parseAssignmentExpression(i, p.only(pIn|pYield|pAwait)); assignmentExpression != nil {
				return &ast.AssignmentExpression{
					LeftHandSideExpression: leftHandSideExpression,
					Assign:                 t.Type == token.Assign,
					AssignmentOperator:     t.Value,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}
