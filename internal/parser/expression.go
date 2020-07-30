//lint:file-ignore U1000 "work in progress"

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

	if asyncArrowFunction := parseAsyncArrowFunction(i, p.only(pIn|pYield|pAwait)); asyncArrowFunction != nil {
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

	if !i.acceptOneOfTypes(token.NewT) {
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
	// see #90
	return nil
}

func _parseCallExpression(i *isolate, p param) *ast.CallExpression {
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
	if i.acceptOneOfTypes(token.NewT) &&
		i.acceptOneOfTypes(token.Dot) &&
		i.acceptOneOfTypes(token.Target) {
		return &ast.NewTarget{}
	}
	return nil
}

func parseMemberExpression(i *isolate, p param) *ast.MemberExpression {
	/*
	   Because the MemberExpression rule is left-recursive, this has not been
	   implemented yet. Please find a modified, working, non left-recursive rule
	   directly below this paragraph. The difficulty is creating a model that
	   supports the modified rule, but can be converted to the original,
	   expected model.

	   S → S 'p' | 'q'

	   S  → 'q' ps
	   ps → 'p' ps | <nothing>

	   ---

	   MemX:
	   PrimX     MemXAux
	   SuperProp MemXAux
	   MetaProp  MemXAux
	   "new"     MemX    Args

	   MemXAux:
	   "[" X "]"       MemXAux
	   "." IdentName   MemXAux
	   TemplateLiteral MemXAux
	   <nothing>
	*/

	chck := i.checkpoint()

	if primExpr := parsePrimaryExpression(i, p.only(pYield|pAwait)); primExpr != nil {
		if aux := parseMemberExpressionAuxiliary(i, p); aux != nil {
			aux.PrimaryExpression = primExpr
			return aux.ToActual()
		}
	} else if superProp := parseSuperProperty(i, p.only(pYield|pAwait)); superProp != nil {
		if aux := parseMemberExpressionAuxiliary(i, p); aux != nil {
			aux.SuperProperty = superProp
			return aux.ToActual()
		}
	} else if metaProp := parseMetaProperty(i, p.only(pYield|pAwait)); metaProp != nil {
		if aux := parseMemberExpressionAuxiliary(i, p); aux != nil {
			aux.MetaProperty = metaProp
			return aux.ToActual()
		}
	} else if i.acceptOneOfTypes(token.NewT) { // token is parsed, so this throws errors as of here
		if memberExpr := parseMemberExpression(i, p.only(pYield|pAwait)); memberExpr != nil {
			if args := parseArguments(i, p.only(pYield|pAwait)); args != nil {
				return &ast.MemberExpression{
					MemberExpression: memberExpr,
					Arguments:        args,
				}
			}
			i.fatal(msgExpectingArguments)
		} else {
			i.fatal(msgExpectingMemberExpression)
		}
	}

	i.restore(chck)
	return nil
}

/*
   MemXAux:
   "[" X "]"       MemXAux
   "." IdentName   MemXAux
   TemplateLiteral MemXAux
   <nothing>
*/
type memberExpressionAuxiliary struct {
	Brackets        bool
	Expression      *ast.Expression
	Dot             bool
	IdentifierName  string
	TemplateLiteral *ast.TemplateLiteral
	Aux             *memberExpressionAuxiliary

	PrimaryExpression *ast.PrimaryExpression
	SuperProperty     *ast.SuperProperty
	MetaProperty      *ast.MetaProperty

	_Nothing bool
}

func (a *memberExpressionAuxiliary) ToActual() *ast.MemberExpression {
	memX := new(ast.MemberExpression)
	a.toActualRecursive(memX)
	return memX
}

func (a *memberExpressionAuxiliary) toActualRecursive(aggregator *ast.MemberExpression) *ast.MemberExpression {
	aggregator.PrimaryExpression = a.PrimaryExpression
	aggregator.MetaProperty = a.MetaProperty
	aggregator.SuperProperty = a.SuperProperty
	aggregator.TemplateLiteral = a.TemplateLiteral
	aggregator.IdentifierName = a.IdentifierName
	aggregator.Expression = a.Expression

	if a._Nothing {
		return aggregator
	}

	// no arguments, since that is not handled with auxiliary objects
	aggregator.MemberExpression = a.Aux.ToActual()

	return aggregator
}

func parseMemberExpressionAuxiliary(i *isolate, p param) *memberExpressionAuxiliary {
	if i.acceptOneOfTypes(token.BracketOpen) {
		if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
			if i.acceptOneOfTypes(token.BracketClose) {
				return &memberExpressionAuxiliary{
					Brackets:   true,
					Expression: expr,
					Aux:        parseMemberExpressionAuxiliary(i, p),
				}
			}
			i.fatal(msgExpectingBracketClose)
		} else {
			i.fatal(msgExpectingExpression)
		}
	} else if i.acceptOneOfTypes(token.Dot) {
		if t, ok := i.accept(token.IdentifierName); ok {
			return &memberExpressionAuxiliary{
				Dot:            true,
				IdentifierName: t.Value,
				Aux:            parseMemberExpressionAuxiliary(i, p),
			}
		}
		i.fatal(msgExpectingIdentifierName)
	} else if templateLiteral := parseTemplateLiteral(i, p.only(pYield|pAwait).add(pTagged)); templateLiteral != nil {
		return &memberExpressionAuxiliary{
			TemplateLiteral: templateLiteral,
			Aux:             parseMemberExpressionAuxiliary(i, p),
		}
	}

	return &memberExpressionAuxiliary{
		_Nothing: true,
	}
}

func _parseMemberExpression(i *isolate, p param) *ast.MemberExpression {
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
	} else if i.acceptOneOfTypes(token.NewT) {
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
		// ensure that no line terminator ahead
		if !i.negativeLookahead(token.LineTerminator) {
			return &ast.YieldExpression{}
		}

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

func parseIdentifierReference(i *isolate, p param) *ast.IdentifierReference {
	if !p.is(pYield) && i.acceptOneOfTypes(token.Yield) {
		return &ast.IdentifierReference{
			Yield: true,
		}
	}
	if !p.is(pAwait) && i.acceptOneOfTypes(token.Await) {
		return &ast.IdentifierReference{
			Await: true,
		}
	}
	if ident := parseIdentifier(i, 0); ident != nil {
		return &ast.IdentifierReference{
			Identifier: ident,
		}
	}
	return nil
}

func parseAwaitExpression(i *isolate, p param) *ast.AwaitExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Await) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield).add(pAwait)); unaryExpr != nil {
			return &ast.AwaitExpression{
				UnaryExpression: unaryExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseUnaryExpression(i *isolate, p param) *ast.UnaryExpression {
	chck := i.checkpoint()

	if p.is(pAwait) {
		if awaitExpr := parseAwaitExpression(i, p.only(pYield)); awaitExpr != nil {
			return &ast.UnaryExpression{
				AwaitExpression: awaitExpr,
			}
		}
	}

	if updateExpr := parseUpdateExpression(i, p.only(pYield|pAwait)); updateExpr != nil {
		return &ast.UnaryExpression{
			UpdateExpression: updateExpr,
		}
	} else if i.acceptOneOfTypes(token.Delete) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				Delete:          true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.Void) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				Void:            true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.Typeof) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				Typeof:          true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.Plus) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				Plus:            true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.Minus) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				Minus:           true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.Tilde) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				Tilde:           true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.ExclamationMark) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UnaryExpression{
				ExclamationMark: true,
				UnaryExpression: unaryExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseUpdateExpression(i *isolate, p param) *ast.UpdateExpression {
	chck := i.checkpoint()

	if leftHandSideExpression := parseLeftHandSideExpression(i, p.only(pYield|pAwait)); leftHandSideExpression != nil {
		var plusPlus, minusMinus bool
		if i.negativeLookahead(token.LineTerminator) {
			if i.acceptOneOfTypes(token.UpdatePlus) {
				plusPlus = true
			} else if i.acceptOneOfTypes(token.UpdateMinus) {
				minusMinus = true
			}
		}
		return &ast.UpdateExpression{
			LeftHandSideExpression: leftHandSideExpression,
			PlusPlus:               plusPlus,
			MinusMinus:             minusMinus,
		}
	} else if i.acceptOneOfTypes(token.UpdatePlus) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UpdateExpression{
				PlusPlus:        true,
				UnaryExpression: unaryExpr,
			}
		}
	} else if i.acceptOneOfTypes(token.UpdateMinus) {
		if unaryExpr := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpr != nil {
			return &ast.UpdateExpression{
				MinusMinus:      true,
				UnaryExpression: unaryExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncFunctionExpression(i *isolate, p param) *ast.AsyncFunctionExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if i.acceptOneOfTypes(token.Function) {
				bindingIdent := parseBindingIdentifier(i, pAwait) // binding identifier is effectively optional
				if i.acceptOneOfTypes(token.ParOpen) {
					if formalParameters := parseFormalParameters(i, pAwait); formalParameters != nil {
						if i.acceptOneOfTypes(token.ParClose) {
							if i.acceptOneOfTypes(token.BraceOpen) {
								if asyncFunctionBody := parseAsyncFunctionBody(i, 0); asyncFunctionBody != nil {
									if i.acceptOneOfTypes(token.BraceClose) {
										return &ast.AsyncFunctionExpression{
											BindingIdentifier: bindingIdent,
											FormalParameters:  formalParameters,
											AsyncFunctionBody: asyncFunctionBody,
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncGeneratorExpression(i *isolate, p param) *ast.AsyncGeneratorExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if i.acceptOneOfTypes(token.Function) {
				if i.acceptOneOfTypes(token.Asterisk) {
					bindingIdent := parseBindingIdentifier(i, pYield|pAwait) // optional
					if i.acceptOneOfTypes(token.ParOpen) {
						if formalParameters := parseFormalParameters(i, pYield|pAwait); formalParameters != nil {
							if i.acceptOneOfTypes(token.ParClose) {
								if i.acceptOneOfTypes(token.BraceOpen) {
									if asyncGeneratorBody := parseAsyncGeneratorBody(i, 0); asyncGeneratorBody != nil {
										if i.acceptOneOfTypes(token.BraceClose) {
											return &ast.AsyncGeneratorExpression{
												BindingIdentifier:  bindingIdent,
												FormalParameters:   formalParameters,
												AsyncGeneratorBody: asyncGeneratorBody,
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseSuperCall(i *isolate, p param) *ast.SuperCall {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Super) {
		if args := parseArguments(i, p.only(pYield|pAwait)); args != nil {
			return &ast.SuperCall{
				Arguments: args,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseClassExpression(i *isolate, p param) *ast.ClassExpression {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Class) {
		bindingIdent := parseBindingIdentifier(i, p.only(pYield|pAwait))
		if classTail := parseClassTail(i, p.only(pYield|pAwait)); classTail != nil {
			return &ast.ClassExpression{
				BindingIdentifier: bindingIdent,
				ClassTail:         classTail,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseArrowFunction(i *isolate, p param) *ast.ArrowFunction {
	chck := i.checkpoint()

	if arrowParameters := parseArrowParameters(i, p.only(pYield|pAwait)); arrowParameters != nil {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if i.acceptOneOfTypes(token.Arrow) {
				if conciseBody := parseConciseBody(i, p.only(pIn)); conciseBody != nil {
					return &ast.ArrowFunction{
						ArrowParameters: arrowParameters,
						ConciseBody:     conciseBody,
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseArrowParameters(i *isolate, p param) *ast.ArrowParameters {
	if bindingIdent := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdent != nil {
		return &ast.ArrowParameters{
			BindingIdentifier: bindingIdent,
		}
	} else if coverParenthesizedExpressionAndArrowParameterList := parseCoverParenthesizedExpressionAndArrowParameterList(i, p.only(pYield|pAwait)); coverParenthesizedExpressionAndArrowParameterList != nil {
		return &ast.ArrowParameters{
			CoverParenthesizedExpressionAndArrowParameterList: coverParenthesizedExpressionAndArrowParameterList,
		}
	}

	return nil
}

func parseConciseBody(i *isolate, p param) *ast.ConciseBody {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.BraceOpen) {
		if assignmentExpr := parseAssignmentExpression(i, p.only(pIn)); assignmentExpr != nil {
			return &ast.ConciseBody{
				AssignmentExpression: assignmentExpr,
			}
		}
	}

	// lookahead: no token.BraceOpen
	if functionBody := parseFunctionBody(i, 0); functionBody != nil {
		if i.acceptOneOfTypes(token.BraceClose) {
			return &ast.ConciseBody{
				FunctionBody: functionBody,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseCoverCallExpressionAndAsyncArrowHead(i *isolate, p param) *ast.CoverCallExpressionAndAsyncArrowHead {
	chck := i.checkpoint()

	if memberExpression := parseMemberExpression(i, p.only(pYield|pAwait)); memberExpression != nil {
		if args := parseArguments(i, p.only(pYield|pAwait)); args != nil {
			return &ast.CoverCallExpressionAndAsyncArrowHead{
				MemberExpression: memberExpression,
				Arguments:        args,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncArrowFunction(i *isolate, p param) *ast.AsyncArrowFunction {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) {
			if asyncArrowBindingIdent := parseAsyncArrowBindingIdentifier(i, p.only(pYield)); asyncArrowBindingIdent != nil {
				if i.negativeLookahead(token.LineTerminator) {
					if i.acceptOneOfTypes(token.Arrow) {
						if asyncConciseBode := parseAsyncConciseBody(i, p.only(pIn)); asyncConciseBode != nil {
							return &ast.AsyncArrowFunction{
								AsyncArrowBindingIdentifier: asyncArrowBindingIdent,
								AsyncConciseBody:            asyncConciseBode,
							}
						}
					}
				}
			}
		}
	} else if coverCallExpressionAndAsyncArrowHead := parseCoverCallExpressionAndAsyncArrowHead(i, p.only(pYield|pAwait)); coverCallExpressionAndAsyncArrowHead != nil {
		if i.negativeLookahead(token.LineTerminator) {
			if i.acceptOneOfTypes(token.Arrow) {
				if asyncConciseBode := parseAsyncConciseBody(i, p.only(pIn)); asyncConciseBode != nil {
					return &ast.AsyncArrowFunction{
						CoverCallExpressionAndAsyncArrowHead: coverCallExpressionAndAsyncArrowHead,
						AsyncConciseBody:                     asyncConciseBode,
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncConciseBody(i *isolate, p param) *ast.AsyncConciseBody {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.BraceOpen) {
		if assignmentExpr := parseAssignmentExpression(i, p.only(pIn)); assignmentExpr != nil {
			return &ast.AsyncConciseBody{
				AssignmentExpression: assignmentExpr,
			}
		}
	}

	// lookahead: no token.BraceOpen
	if asyncFunctionBody := parseAsyncFunctionBody(i, 0); asyncFunctionBody != nil {
		if i.acceptOneOfTypes(token.BraceClose) {
			return &ast.AsyncConciseBody{
				AsyncFunctionBody: asyncFunctionBody,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncArrowBindingIdentifier(i *isolate, p param) *ast.AsyncArrowBindingIdentifier {
	if bindingIdent := parseBindingIdentifier(i, p.only(pYield).add(pAwait)); bindingIdent != nil {
		return &ast.AsyncArrowBindingIdentifier{
			BindingIdentifier: bindingIdent,
		}
	}

	return nil
}

func parseConditionalExpression(i *isolate, p param) *ast.ConditionalExpression {
	// see #90
	return nil
}

func _parseConditionalExpression(i *isolate, p param) *ast.ConditionalExpression {
	chck := i.checkpoint()

	if orExpr := parseLogicalORExpression(i, p.only(pIn|pYield|pAwait)); orExpr != nil {
		if i.acceptOneOfTypes(token.QuestionMark) {
			if assignmentExpr1 := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr1 != nil {
				if i.acceptOneOfTypes(token.Colon) {
					if assignmentExpr2 := parseAssignmentExpression(i, p.only(pIn|pYield|pAwait)); assignmentExpr2 != nil {
						return &ast.ConditionalExpression{
							LogicalORExpression:   orExpr,
							AssignmentExpression1: assignmentExpr1,
							AssignmentExpression2: assignmentExpr2,
						}
					}
				}
			}
		}

		return &ast.ConditionalExpression{
			LogicalORExpression: orExpr,
		}
	}

	i.restore(chck)
	return nil
}

func parseLogicalORExpression(i *isolate, p param) *ast.LogicalORExpression {
	// see #90
	return nil
}

func _parseLogicalORExpression(i *isolate, p param) *ast.LogicalORExpression {
	chck := i.checkpoint()

	if andExpr := parseLogicalANDExpression(i, p.only(pIn|pYield|pAwait)); andExpr != nil {
		return &ast.LogicalORExpression{
			LogicalANDExpression: andExpr,
		}
	} else if orExpr := parseLogicalORExpression(i, p.only(pIn|pYield|pAwait)); orExpr != nil {
		if i.acceptOneOfTypes(token.LogicalOr) {
			if andExpr := parseLogicalANDExpression(i, p.only(pIn|pYield|pAwait)); andExpr != nil {
				return &ast.LogicalORExpression{
					LogicalORExpression:  orExpr,
					LogicalANDExpression: andExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseLogicalANDExpression(i *isolate, p param) *ast.LogicalANDExpression {
	// see #90
	return nil
}

func _parseLogicalANDExpression(i *isolate, p param) *ast.LogicalANDExpression {
	chck := i.checkpoint()

	if orExpr := parseBitwiseORExpression(i, p.only(pIn|pYield|pAwait)); orExpr != nil {
		return &ast.LogicalANDExpression{
			BitwiseORExpression: orExpr,
		}
	} else if andExpr := parseLogicalANDExpression(i, p.only(pIn|pYield|pAwait)); andExpr != nil {
		if i.acceptOneOfTypes(token.LogicalAnd) {
			if orExpr := parseBitwiseORExpression(i, p.only(pIn|pYield|pAwait)); orExpr != nil {
				return &ast.LogicalANDExpression{
					BitwiseORExpression:  orExpr,
					LogicalANDExpression: andExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBitwiseANDExpression(i *isolate, p param) *ast.BitwiseANDExpression {
	// see #90
	return nil
}

func _parseBitwiseANDExpression(i *isolate, p param) *ast.BitwiseANDExpression {
	chck := i.checkpoint()

	if eqExpr := parseEqualityExpression(i, p.only(pIn|pYield|pAwait)); eqExpr != nil {
		return &ast.BitwiseANDExpression{
			EqualityExpression: eqExpr,
		}
	} else if andExpr := parseBitwiseANDExpression(i, p.only(pIn|pYield|pAwait)); andExpr != nil {
		if i.acceptOneOfTypes(token.BitwiseAnd) {
			if eqExpr := parseEqualityExpression(i, p.only(pIn|pYield|pAwait)); eqExpr != nil {
				return &ast.BitwiseANDExpression{
					BitwiseANDExpression: andExpr,
					EqualityExpression:   eqExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBitwiseXORExpression(i *isolate, p param) *ast.BitwiseXORExpression {
	// see #90
	return nil
}

func _parseBitwiseXORExpression(i *isolate, p param) *ast.BitwiseXORExpression {
	chck := i.checkpoint()

	if andExpr := parseBitwiseANDExpression(i, p.only(pIn|pYield|pAwait)); andExpr != nil {
		return &ast.BitwiseXORExpression{
			BitwiseANDExpression: andExpr,
		}
	} else if xorExpr := parseBitwiseXORExpression(i, p.only(pIn|pYield|pAwait)); xorExpr != nil {
		if i.acceptOneOfTypes(token.BitwiseXor) {
			if andExpr := parseBitwiseANDExpression(i, p.only(pIn|pYield|pAwait)); andExpr != nil {
				return &ast.BitwiseXORExpression{
					BitwiseXORExpression: xorExpr,
					BitwiseANDExpression: andExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBitwiseORExpression(i *isolate, p param) *ast.BitwiseORExpression {
	// see #90
	return nil
}

func _parseBitwiseORExpression(i *isolate, p param) *ast.BitwiseORExpression {
	chck := i.checkpoint()

	if xorExpr := parseBitwiseXORExpression(i, p.only(pIn|pYield|pAwait)); xorExpr != nil {
		return &ast.BitwiseORExpression{
			BitwiseXORExpression: xorExpr,
		}
	} else if orExpr := parseBitwiseORExpression(i, p.only(pIn|pYield|pAwait)); orExpr != nil {
		if i.acceptOneOfTypes(token.BitwiseXor) {
			if xorExpr := parseBitwiseXORExpression(i, p.only(pIn|pYield|pAwait)); xorExpr != nil {
				return &ast.BitwiseORExpression{
					BitwiseORExpression:  orExpr,
					BitwiseXORExpression: xorExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseEqualityExpression(i *isolate, p param) *ast.EqualityExpression {
	// see #90
	return nil
}

func _parseEqualityExpression(i *isolate, p param) *ast.EqualityExpression {
	chck := i.checkpoint()

	if relationalExpr := parseRelationalExpression(i, p.only(pIn|pYield|pAwait)); relationalExpr != nil {
		return &ast.EqualityExpression{
			RelationalExpression: relationalExpr,
		}
	} else if eqExpr := parseEqualityExpression(i, p.only(pIn|pYield|pAwait)); eqExpr != nil {
		if t, ok := i.acceptOneOf(token.Equals, token.NotEquals, token.StrictEquals, token.StrictNotEquals); ok {
			if relationalExpr := parseRelationalExpression(i, p.only(pIn|pYield|pAwait)); relationalExpr != nil {
				return &ast.EqualityExpression{
					RelationalExpression: relationalExpr,
					EqualityExpression:   eqExpr,
					Equals:               t.Type == token.Equals,
					NotEquals:            t.Type == token.NotEquals,
					StrictEquals:         t.Type == token.StrictEquals,
					StrictNotEquals:      t.Type == token.StrictNotEquals,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseRelationalExpression(i *isolate, p param) *ast.RelationalExpression {
	// see #90
	return nil
}

func _parseRelationalExpression(i *isolate, p param) *ast.RelationalExpression {
	chck := i.checkpoint()

	if shiftExpr := parseShiftExpression(i, p.only(pYield|pAwait)); shiftExpr != nil {
		return &ast.RelationalExpression{
			ShiftExpression: shiftExpr,
		}
	}

	if p.is(pIn) {
		if relationalExpr := parseRelationalExpression(i, p.only(pYield|pAwait).add(pIn)); relationalExpr != nil {
			if i.acceptOneOfTypes(token.In) {
				if shiftExpr := parseShiftExpression(i, p.only(pYield|pAwait)); shiftExpr != nil {
					return &ast.RelationalExpression{
						RelationalExpression: relationalExpr,
						In:                   true,
						ShiftExpression:      shiftExpr,
					}
				}
			}
		}
	} else if relationalExpr := parseRelationalExpression(i, p.only(pIn|pYield|pAwait)); relationalExpr != nil {
		if t, ok := i.acceptOneOf(token.LessThan, token.GreaterThan, token.LessThanOrEqualTo, token.GreaterThanOrEqualTo, token.Instanceof); ok {
			if shiftExpr := parseShiftExpression(i, p.only(pYield|pAwait)); shiftExpr != nil {
				return &ast.RelationalExpression{
					RelationalExpression: relationalExpr,
					ShiftExpression:      shiftExpr,
					LessThan:             t.Type == token.LessThan,
					GreaterThan:          t.Type == token.GreaterThan,
					LessThanOrEqualTo:    t.Type == token.LessThanOrEqualTo,
					GreaterThanOrEqualTo: t.Type == token.GreaterThanOrEqualTo,
					Instanceof:           t.Type == token.Instanceof,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseShiftExpression(i *isolate, p param) *ast.ShiftExpression {
	// see #90
	return nil
}

func _parseShiftExpression(i *isolate, p param) *ast.ShiftExpression {
	chck := i.checkpoint()

	if additiveExpr := parseAdditiveExpression(i, p.only(pYield|pAwait)); additiveExpr != nil {
		return &ast.ShiftExpression{
			AdditiveExpression: additiveExpr,
		}
	} else if shiftExpr := parseShiftExpression(i, p.only(pYield|pAwait)); shiftExpr != nil {
		if t, ok := i.acceptOneOf(token.LeftShift, token.RightShift, token.UnsignedRightShift); ok {
			if additiveExpr := parseAdditiveExpression(i, p.only(pYield|pAwait)); additiveExpr != nil {
				return &ast.ShiftExpression{
					ShiftExpression:    shiftExpr,
					AdditiveExpression: additiveExpr,
					LeftShift:          t.Type == token.LeftShift,
					RightShift:         t.Type == token.RightShift,
					UnsignedRightShift: t.Type == token.UnsignedRightShift,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAdditiveExpression(i *isolate, p param) *ast.AdditiveExpression {
	// see #90
	return nil
}

func _parseAdditiveExpression(i *isolate, p param) *ast.AdditiveExpression {
	chck := i.checkpoint()

	if multiplicativeExpression := parseMultiplicativeExpression(i, p.only(pYield|pAwait)); multiplicativeExpression != nil {
		return &ast.AdditiveExpression{
			MultiplicativeExpression: multiplicativeExpression,
		}
	} else if additiveExpr := parseAdditiveExpression(i, p.only(pYield|pAwait)); additiveExpr != nil {
		if t, ok := i.acceptOneOf(token.Plus, token.Minus); ok {
			if multiplicativeExpr := parseMultiplicativeExpression(i, p.only(pYield|pAwait)); multiplicativeExpr != nil {
				return &ast.AdditiveExpression{
					AdditiveExpression:       additiveExpr,
					MultiplicativeExpression: multiplicativeExpr,
					Plus:                     t.Type == token.Plus,
					Minus:                    t.Type == token.Minus,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseMultiplicativeExpression(i *isolate, p param) *ast.MultiplicativeExpression {
	// se #90
	return nil
}

func _parseMultiplicativeExpression(i *isolate, p param) *ast.MultiplicativeExpression {
	chck := i.checkpoint()

	if exponentiationExpression := parseExponentiationExpression(i, p.only(pYield|pAwait)); exponentiationExpression != nil {
		return &ast.MultiplicativeExpression{
			ExponentiationExpression: exponentiationExpression,
		}
	} else if multiplicativeExpr := parseMultiplicativeExpression(i, p.only(pYield|pAwait)); multiplicativeExpr != nil {
		if t, ok := i.acceptOneOf(token.Asterisk, token.Slash, token.Modulo); ok {
			if exponentiationExpression := parseExponentiationExpression(i, p.only(pYield|pAwait)); exponentiationExpression != nil {
				return &ast.MultiplicativeExpression{
					MultiplicativeExpression: multiplicativeExpr,
					ExponentiationExpression: exponentiationExpression,
					Asterisk:                 t.Type == token.Asterisk,
					Slash:                    t.Type == token.Slash,
					Modulo:                   t.Type == token.Modulo,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseExponentiationExpression(i *isolate, p param) *ast.ExponentiationExpression {
	chck := i.checkpoint()

	if unaryExpression := parseUnaryExpression(i, p.only(pYield|pAwait)); unaryExpression != nil {
		return &ast.ExponentiationExpression{
			UnaryExpression: unaryExpression,
		}
	} else if updateExpr := parseUpdateExpression(i, p.only(pYield|pAwait)); updateExpr != nil {
		if i.acceptOneOfTypes(token.Power) {
			if exponentiationExpr := parseExponentiationExpression(i, p.only(pYield|pAwait)); exponentiationExpr != nil {
				return &ast.ExponentiationExpression{
					UpdateExpression:         updateExpr,
					ExponentiationExpression: exponentiationExpr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}
