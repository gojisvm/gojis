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

func parseSuperProperty(i *isolate, p param) *ast.SuperProperty {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Super) {
		i.restore(chck)
		return nil
	}

	if i.acceptOneOfTypes(token.BracketOpen) {
		if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
			if i.acceptOneOfTypes(token.BracketClose) {
				return &ast.SuperProperty{
					Expression: expr,
				}
			}
		}
	} else if i.acceptOneOfTypes(token.Dot) {
		if t, ok := i.accept(token.IdentifierName); ok {
			return &ast.SuperProperty{
				IdentifierName: t.Value,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseMetaProperty(i *isolate, p param) *ast.MetaProperty {
	if newTarget := parseNewTarget(i, 0); newTarget != nil {
		return &ast.MetaProperty{
			NewTarget: newTarget,
		}
	}

	return nil
}

func parseNewTarget(i *isolate, p param) *ast.NewTarget {
	// new . target
	if i.acceptOneOfTypes(token.New_) &&
		i.acceptOneOfTypes(token.Dot) &&
		i.acceptOneOfTypes(token.Target) {
		return &ast.NewTarget{}
	}
	return nil
}

func parseMemberExpression(i *isolate, p param) *ast.MemberExpression {
	chck := i.checkpoint()

	if primaryExpression := parsePrimaryExpression(i, p.only(pYield|pAwait)); primaryExpression != nil {
		return &ast.MemberExpression{
			PrimaryExpression: primaryExpression,
		}
	} else if memberExpression := parseMemberExpression(i, p.only(pYield|pAwait)); memberExpression != nil {
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
	} else if superProperty := parseSuperProperty(i, p.only(pYield|pAwait)); superProperty != nil {
		return &ast.MemberExpression{
			SuperProperty: superProperty,
		}
	} else if metaProperty := parseMetaProperty(i, 0); metaProperty != nil {
		return &ast.MemberExpression{
			MetaProperty: metaProperty,
		}
	} else if i.acceptOneOfTypes(token.New_) {
		if memberExpression := parseMemberExpression(i, p.only(pYield|pAwait)); memberExpression != nil {
			if arguments := parseArguments(i, p.only(pYield|pAwait)); arguments != nil {
				return &ast.MemberExpression{
					MemberExpression: memberExpression,
					Arguments:        arguments,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parsePrimaryExpression(i *isolate, p param) *ast.PrimaryExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.This) {
		return &ast.PrimaryExpression{
			This: true,
		}
	} else if identRef := parseIdentifierReference(i, p.only(pYield|pAwait)); identRef != nil {
		return &ast.PrimaryExpression{
			IdentifierReference: identRef,
		}
	} else if literal := parseLiteral(i, 0); literal != nil {
		return &ast.PrimaryExpression{
			Literal: literal,
		}
	} else if arrayLiteral := parseArrayLiteral(i, p.only(pYield|pAwait)); arrayLiteral != nil {
		return &ast.PrimaryExpression{
			ArrayLiteral: arrayLiteral,
		}
	} else if objectLiteral := parseObjectLiteral(i, p.only(pYield|pAwait)); objectLiteral != nil {
		return &ast.PrimaryExpression{
			ObjectLiteral: objectLiteral,
		}
	} else if funcExpr := parseFunctionExpression(i, 0); funcExpr != nil {
		return &ast.PrimaryExpression{
			FunctionExpression: funcExpr,
		}
	} else if classExpression := parseClassExpression(i, p.only(pYield|pAwait)); classExpression != nil {
		return &ast.PrimaryExpression{
			ClassExpression: classExpression,
		}
	} else if genExpr := parseGeneratorExpression(i, 0); genExpr != nil {
		return &ast.PrimaryExpression{
			GeneratorExpression: genExpr,
		}
	} else if asyncFunctionExpr := parseAsyncFunctionExpression(i, 0); asyncFunctionExpr != nil {
		return &ast.PrimaryExpression{
			AsyncFunctionExpression: asyncFunctionExpr,
		}
	} else if asyncGeneratorExpr := parseAsyncGeneratorExpression(i, 0); asyncGeneratorExpr != nil {
		return &ast.PrimaryExpression{
			AsyncGeneratorExpression: asyncGeneratorExpr,
		}
	} else if regularExpressionLiteral := parseRegularExpressionLiteral(i, 0); regularExpressionLiteral != nil {
		return &ast.PrimaryExpression{
			RegularExpressionLiteral: regularExpressionLiteral,
		}
	} else if templateLiteral := parseTemplateLiteral(i, p.only(pYield|pAwait)); templateLiteral != nil {
		return &ast.PrimaryExpression{
			TemplateLiteral: templateLiteral,
		}
	} else if coverParenthesizedExpressionAndArrowParameterList := parseCoverParenthesizedExpressionAndArrowParameterList(i, p.only(pYield|pAwait)); coverParenthesizedExpressionAndArrowParameterList != nil {
		return &ast.PrimaryExpression{
			CoverParenthesizedExpressionAndArrowParameterList: coverParenthesizedExpressionAndArrowParameterList,
		}
	}

	i.restore(chck)
	return nil
}

func parseCoverParenthesizedExpressionAndArrowParameterList(i *isolate, p param) *ast.CoverParenthesizedExpressionAndArrowParameterList {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.ParOpen) {
		i.restore(chck)
		return nil
	}

	if i.acceptOneOfTypes(token.ParClose) {

	} else if i.acceptOneOfTypes(token.Ellipsis) {
		if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
			if i.acceptOneOfTypes(token.ParClose) {
				return &ast.CoverParenthesizedExpressionAndArrowParameterList{
					BindingIdentifier: bindingIdentifier,
					Ellipsis:          true,
				}
			}
		} else if bindingPattern := parseBindingPattern(i, p.only(pYield|pAwait)); bindingPattern != nil {
			if i.acceptOneOfTypes(token.ParClose) {
				return &ast.CoverParenthesizedExpressionAndArrowParameterList{
					BindingPattern: bindingPattern,
					Ellipsis:       true,
				}
			}
		}
	} else if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
		if i.acceptOneOfTypes(token.ParClose) {
			return &ast.CoverParenthesizedExpressionAndArrowParameterList{
				Expression: expr,
			}
		} else if i.acceptOneOfTypes(token.Comma) {
			if i.acceptOneOfTypes(token.ParClose) {
				return &ast.CoverParenthesizedExpressionAndArrowParameterList{
					Expression: expr,
					Comma:      true,
				}
			} else if i.acceptOneOfTypes(token.Ellipsis) {
				if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
					if i.acceptOneOfTypes(token.ParClose) {
						return &ast.CoverParenthesizedExpressionAndArrowParameterList{
							BindingIdentifier: bindingIdentifier,
							Ellipsis:          true,
							Comma:             true,
						}
					}
				} else if bindingPattern := parseBindingPattern(i, p.only(pYield|pAwait)); bindingPattern != nil {
					if i.acceptOneOfTypes(token.ParClose) {
						return &ast.CoverParenthesizedExpressionAndArrowParameterList{
							BindingPattern: bindingPattern,
							Ellipsis:       true,
							Comma:          true,
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseFunctionExpression(i *isolate, p param) *ast.FunctionExpression {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Function) {
		i.restore(chck)
		return nil
	}

	bindingIdentifier := parseBindingIdentifier(i, 0) // bindingIdentifier is optional

	if i.acceptOneOfTypes(token.ParOpen) {
		if formalParameters := parseFormalParameters(i, 0); formalParameters != nil {
			if i.acceptOneOfTypes(token.ParClose) &&
				i.acceptOneOfTypes(token.BraceOpen) {
				if functionBody := parseFunctionBody(i, 0); functionBody != nil {
					if i.acceptOneOfTypes(token.BraceClose) {
						return &ast.FunctionExpression{
							BindingIdentifier: bindingIdentifier,
							FormalParameters:  formalParameters,
							FunctionBody:      functionBody,
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseFunctionBody(i *isolate, p param) *ast.FunctionBody {
	if functionStmtList := parseFunctionStatementList(i, p.only(pYield|pAwait)); functionStmtList != nil {
		return &ast.FunctionBody{
			FunctionStatementList: functionStmtList,
		}
	}
	return nil
}

func parseFunctionStatementList(i *isolate, p param) *ast.FunctionStatementList {
	return &ast.FunctionStatementList{
		StatementList: parseStatementList(i, p.only(pYield|pAwait).add(pReturn)),
	}
}

func parseGeneratorExpression(i *isolate, p param) *ast.GeneratorExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Function) &&
		i.acceptOneOfTypes(token.Asterisk) {
		bindingIdent := parseBindingIdentifier(i, pYield) // binding identifier is optional
		if i.acceptOneOfTypes(token.ParOpen) {
			if formalParameters := parseFormalParameters(i, pYield); formalParameters != nil {
				if i.acceptOneOfTypes(token.ParClose) &&
					i.acceptOneOfTypes(token.BraceOpen) {
					if generatorBody := parseGeneratorBody(i, 0); generatorBody != nil {
						return &ast.GeneratorExpression{
							BindingIdentifier: bindingIdent,
							FormalParameters:  formalParameters,
							GeneratorBody:     generatorBody,
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseYieldExpression(i *isolate, p param) *ast.YieldExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Yield) {
		if i.acceptOneOfTypes(token.LineTerminator) {
			i.unread()
			return &ast.YieldExpression{}
		}

		// no line terminator ahead
		_, asterisk := i.accept(token.Asterisk)
		if assignmentExpr := parseAssignmentExpression(i, p.only(pIn|pAwait).add(pYield)); assignmentExpr != nil {
			return &ast.YieldExpression{
				Asterisk:             asterisk,
				AssignmentExpression: assignmentExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}
