package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseTemplateLiteral(i *isolate, p param) *ast.TemplateLiteral {
	chck := i.checkpoint()

	if t, ok := i.accept(token.NoSubstitutionTemplate); ok {
		return &ast.TemplateLiteral{
			NoSubstitutionTemplate: t.Value,
		}
	}

	if substitutionTemplate := parseSubstitutionTemplate(i, p.only(pYield|pAwait|pTagged)); substitutionTemplate != nil {
		return &ast.TemplateLiteral{
			SubstitutionTemplate: substitutionTemplate,
		}
	}

	i.restore(chck)
	return nil
}

func parseSubstitutionTemplate(i *isolate, p param) *ast.SubstitutionTemplate {
	chck := i.checkpoint()

	var templateHead string
	var expr *ast.Expression
	var templateSpans *ast.TemplateSpans

	t, ok := i.accept(token.TemplateHead)
	if !ok {
		i.restore(chck)
		return nil
	}
	templateHead = t.Value

	if expr = parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr == nil {
		i.restore(chck)
		return nil
	}

	if templateSpans = parseTemplateSpans(i, p.only(pYield|pAwait|pTagged)); templateSpans == nil {
		i.restore(chck)
		return nil
	}

	return &ast.SubstitutionTemplate{
		TemplateHead:  templateHead,
		Expression:    expr,
		TemplateSpans: templateSpans,
	}
}

func parseNumericLiteral(i *isolate, p param) *ast.NumericLiteral {
	chck := i.checkpoint()

	if t, ok := i.accept(token.DecimalLiteral); ok {
		return &ast.NumericLiteral{
			DecimalLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.BinaryIntegerLiteral); ok {
		return &ast.NumericLiteral{
			BinaryIntegerLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.OctalIntegerLiteral); ok {
		return &ast.NumericLiteral{
			OctalIntegerLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.HexIntegerLiteral); ok {
		return &ast.NumericLiteral{
			HexIntegerLiteral: t.Value,
		}
	}

	i.restore(chck)
	return nil
}

func parseTemplateSpans(i *isolate, p param) *ast.TemplateSpans {
	chck := i.checkpoint()

	templateMiddleList := parseTemplateMiddleList(i, p.only(pYield|pAwait|pTagged)) // templateMiddleList is optional

	if t, ok := i.accept(token.TemplateTail); ok {
		return &ast.TemplateSpans{
			TemplateMiddleList: templateMiddleList,
			TemplateTail:       t.Value,
		}
	}

	i.restore(chck)
	return nil
}

func parseTemplateMiddleList(i *isolate, p param) *ast.TemplateMiddleList {
	chck := i.checkpoint()

	if t, ok := i.accept(token.TemplateMiddle); ok {
		if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
			return &ast.TemplateMiddleList{
				TemplateMiddle: t.Value,
				Expression:     expr,
			}
		}
	} else {
		templateMiddleList := parseTemplateMiddleList(i, p.only(pYield|pAwait|pTagged)) // templateMiddleList is optional
		if t, ok := i.accept(token.TemplateMiddle); ok {
			if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
				return &ast.TemplateMiddleList{
					TemplateMiddleList: templateMiddleList,
					TemplateMiddle:     t.Value,
					Expression:         expr,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseLiteral(i *isolate, p param) *ast.Literal {
	if t, ok := i.accept(token.Null); ok {
		return &ast.Literal{
			NullLiteral: t.Value,
		}
	} else if t, ok := i.accept(token.Boolean); ok {
		return &ast.Literal{
			BooleanLiteral: t.Value,
		}
	} else if numericLiteral := parseNumericLiteral(i, 0); numericLiteral != nil {
		return &ast.Literal{
			NumericLiteral: numericLiteral,
		}
	} else if t, ok := i.accept(token.StringLiteral); ok {
		return &ast.Literal{
			StringLiteral: t.Value,
		}
	}

	return nil
}

func parseArrayLiteral(i *isolate, p param) *ast.ArrayLiteral {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.BracketOpen) {
		i.restore(chck)
		return nil
	}

	if elision := parseElision(i, 0); elision != nil {
		if i.acceptOneOfTypes(token.BracketClose) {
			return &ast.ArrayLiteral{
				Elision: elision,
			}
		}
	} else if elementList := parseElementList(i, p.only(pYield|pAwait)); elementList != nil {
		if i.acceptOneOfTypes(token.Comma) {
			elision := parseElision(i, 0)
			if i.acceptOneOfTypes(token.BracketClose) {
				return &ast.ArrayLiteral{
					ElementList: elementList,
					Comma:       true,
					Elision:     elision,
				}
			}
		}
		if i.acceptOneOfTypes(token.BracketClose) {
			return &ast.ArrayLiteral{
				ElementList: elementList,
			}
		}
	} else if i.acceptOneOfTypes(token.BracketClose) {
		return &ast.ArrayLiteral{} // elision is optional
	}

	i.restore(chck)
	return nil
}

func parseElementList(i *isolate, p param) *ast.ElementList {
	chck := i.checkpoint()

	if elementList := parseElementList(i, p.only(pYield|pAwait)); elementList != nil {
		if i.acceptOneOfTypes(token.Comma) {
			elision := parseElision(i, 0)
			if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
				return &ast.ElementList{
					ElementList:          elementList,
					Comma:                true,
					Elision:              elision,
					AssignmentExpression: assignmentExpr,
				}
			} else if spreadElement := parseSpreadElement(i, p.only(pYield|pAwait)); spreadElement != nil {
				return &ast.ElementList{
					ElementList:   elementList,
					Comma:         true,
					Elision:       elision,
					SpreadElement: spreadElement,
				}
			}
		}
	}

	elision := parseElision(i, 0)
	if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
		return &ast.ElementList{
			Comma:                true,
			Elision:              elision,
			AssignmentExpression: assignmentExpr,
		}
	} else if spreadElement := parseSpreadElement(i, p.only(pYield|pAwait)); spreadElement != nil {
		return &ast.ElementList{
			Comma:         true,
			Elision:       elision,
			SpreadElement: spreadElement,
		}
	}

	i.restore(chck)
	return nil
}

func parseSpreadElement(i *isolate, p param) *ast.SpreadElement {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Ellipsis) {
		if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
			return &ast.SpreadElement{
				AssignmentExpression: assignmentExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseObjectLiteral(i *isolate, p param) *ast.ObjectLiteral {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.BraceOpen) {
		i.restore(chck)
		return nil
	}

	if i.acceptOneOfTypes(token.BraceClose) {
		return &ast.ObjectLiteral{}
	} else if propertyDefinitionList := parsePropertyDefinitionList(i, p.only(pYield|pAwait)); propertyDefinitionList != nil {
		comma := i.acceptOneOfTypes(token.Comma)
		if i.acceptOneOfTypes(token.BraceClose) {
			return &ast.ObjectLiteral{
				PropertyDefinitionList: propertyDefinitionList,
				Comma:                  comma,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parsePropertyDefinitionList(i *isolate, p param) *ast.PropertyDefinitionList {
	chck := i.checkpoint()

	var defs []*ast.PropertyDefinition

	first := parsePropertyDefinition(i, p.only(pYield|pAwait|pReturn)) // definition list consists of at least one entry
	if first == nil {
		return nil
	}
	defs = append(defs, first)

	for { // parse until there are no more assignment definitions parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(chck)
			return nil
		}

		next := parsePropertyDefinition(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no definition, so reset to before comma
			break
		}
		defs = append(defs, next)
	}

	return &ast.PropertyDefinitionList{
		PropertyDefinitions: defs,
	}
}

func parsePropertyDefinition(i *isolate, p param) *ast.PropertyDefinition {
	chck := i.checkpoint()

	if identRef := parseIdentifierReference(i, p.only(pYield|pAwait)); identRef != nil {
		return &ast.PropertyDefinition{
			IdentifierReference: identRef,
		}
	} else if coverInitializedName := parseCoverInitializedName(i, p.only(pYield|pAwait)); coverInitializedName != nil {
		return &ast.PropertyDefinition{
			CoverInitializedName: coverInitializedName,
		}
	} else if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
		if i.acceptOneOfTypes(token.Colon) {
			if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
				return &ast.PropertyDefinition{
					Colon:                true,
					PropertyName:         propertyName,
					AssignmentExpression: assignmentExpr,
				}
			}
		}
	} else if methodDefinition := parseMethodDefinition(i, p.only(pYield|pAwait)); methodDefinition != nil {
		return &ast.PropertyDefinition{
			MethodDefinition: methodDefinition,
		}
	} else if i.acceptOneOfTypes(token.Ellipsis) {
		if assignmentExpr := parseAssignmentExpression(i, p.only(pYield|pAwait).add(pIn)); assignmentExpr != nil {
			return &ast.PropertyDefinition{
				Ellipsis:             true,
				AssignmentExpression: assignmentExpr,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseCoverInitializedName(i *isolate, p param) *ast.CoverInitializedName {
	chck := i.checkpoint()

	if identRef := parseIdentifierReference(i, p.only(pYield|pAwait)); identRef != nil {
		if initializer := parseInitializer(i, p.only(pYield|pAwait).add(pIn)); initializer != nil {
			return &ast.CoverInitializedName{
				IdentifierReference: identRef,
				Initializer:         initializer,
			}
		}
	}

	i.restore(chck)
	return nil
}
