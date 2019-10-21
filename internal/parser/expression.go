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

func parseLeftHandSideExpression(i *isolate, p param) *ast.LeftHandSideExpression {
	if newExpression := parseNewExpression(i, p.only(pYield|pAwait)); newExpression != nil {
		return &ast.LeftHandSideExpression{
			NewExpression: newExpression,
		}
	}

	if callExpression := parseCallExpression(i, p.only(pYield|pAwait)); callExpression != nil {
		return &ast.LeftHandSideExpression{
			CallExpression: callExpression,
		}
	}

	return nil
}

func parseNewExpression(i *isolate, p param) *ast.NewExpression {
	chck := i.checkpoint()

	if memberExpression := parseMemberExpression(i, p.only(pYield|pAwait)); memberExpression != nil {
		return &ast.NewExpression{
			MemberExpression: memberExpression,
		}
	}

	if !i.acceptOneOfTypes(token.New_) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if newExpression := parseNewExpression(i, p.only(pYield|pAwait)); newExpression != nil {
		return &ast.NewExpression{
			NewExpression: newExpression,
		}
	}

	i.restore(chck)
	return nil
}

func parseCallExpression(i *isolate, p param) *ast.CallExpression {
	chck := i.checkpoint()

	if coverCallExpressionAndAsyncArrowHead := parseCoverCallExpressionAndAsyncArrowHead(i, p.only(pYield|pAwait)); coverCallExpressionAndAsyncArrowHead != nil {
		return &ast.CallExpression{
			CoverCallExpressionAndAsyncArrowHead: coverCallExpressionAndAsyncArrowHead,
		}
	}

	if superCall := parseSuperCall(i, p.only(pYield|pAwait)); superCall != nil {
		return &ast.CallExpression{
			SuperCall: superCall,
		}
	}

	if callExpression := parseCallExpression(i, p.only(pYield|pAwait)); callExpression != nil {
		if arguments := parseArguments(i, p.only(pYield|pAwait)); arguments != nil {
			return &ast.CallExpression{
				CallExpression: callExpression,
				Arguments:      arguments,
			}
		} else if i.acceptOneOfTypes(token.BracketOpen) {
			if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
				if i.acceptOneOfTypes(token.BracketClose) {
					return &ast.CallExpression{
						CallExpression: callExpression,
						Expression:     expr,
					}
				}
			}
		} else if i.acceptOneOfTypes(token.Dot) {
			if ident, ok := i.accept(token.IdentifierName); ok {
				return &ast.CallExpression{
					CallExpression: callExpression,
					IdentifierName: ident.Value,
				}
			}
		} else if templateLiteral := parseTemplateLiteral(i, p.only(pYield|pAwait).add(pTagged)); templateLiteral != nil {
			return &ast.CallExpression{
				CallExpression:  callExpression,
				TemplateLiteral: templateLiteral,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseMemberExpression(i *isolate, p param) *ast.MemberExpression {
	chck := i.checkpoint()

	if primaryExpression := parsePrimaryExpression(i, p.only(pYield|pAwait)); primaryExpression != nil {
		return &ast.MemberExpression{
			PrimaryExpression: primaryExpression,
		}
	}

	if memberExpression := parseMemberExpression(i, p.only(pYield|pAwait)); memberExpression != nil {
		if i.acceptOneOfTypes(token.BracketOpen) {
			if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
				if i.acceptOneOfTypes(token.BracketClose) {
					return &ast.MemberExpression{
						MemberExpression: memberExpression,
						Expression:       expr,
					}
				}
			}
		} else if i.acceptOneOfTypes(token.Dot) {
			if t, ok := i.accept(token.IdentifierName); ok {
				return &ast.MemberExpression{
					MemberExpression: memberExpression,
					IdentifierName:   t.Value,
				}
			}
		} else {
			if templateLiteral := parseTemplateLiteral(i, p.only(pYield|pAwait).add(pTagged)); templateLiteral != nil {
				return &ast.MemberExpression{
					MemberExpression: memberExpression,
					TemplateLiteral:  templateLiteral,
				}
			}
		}

		panic("TODO")

		i.restore(chck)
		return nil
	}

	i.restore(chck)
	return nil
}
