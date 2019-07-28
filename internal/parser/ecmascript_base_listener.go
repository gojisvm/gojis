// Code generated from ECMAScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ECMAScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseECMAScriptListener is a complete listener for a parse tree produced by ECMAScriptParser.
type BaseECMAScriptListener struct{}

var _ ECMAScriptListener = &BaseECMAScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseECMAScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseECMAScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseECMAScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseECMAScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInputElementDiv is called when production inputElementDiv is entered.
func (s *BaseECMAScriptListener) EnterInputElementDiv(ctx *InputElementDivContext) {}

// ExitInputElementDiv is called when production inputElementDiv is exited.
func (s *BaseECMAScriptListener) ExitInputElementDiv(ctx *InputElementDivContext) {}

// EnterInputElementRegExp is called when production inputElementRegExp is entered.
func (s *BaseECMAScriptListener) EnterInputElementRegExp(ctx *InputElementRegExpContext) {}

// ExitInputElementRegExp is called when production inputElementRegExp is exited.
func (s *BaseECMAScriptListener) ExitInputElementRegExp(ctx *InputElementRegExpContext) {}

// EnterInputElementRegExpOrTemplateTail is called when production inputElementRegExpOrTemplateTail is entered.
func (s *BaseECMAScriptListener) EnterInputElementRegExpOrTemplateTail(ctx *InputElementRegExpOrTemplateTailContext) {
}

// ExitInputElementRegExpOrTemplateTail is called when production inputElementRegExpOrTemplateTail is exited.
func (s *BaseECMAScriptListener) ExitInputElementRegExpOrTemplateTail(ctx *InputElementRegExpOrTemplateTailContext) {
}

// EnterInputElementTemplateTail is called when production inputElementTemplateTail is entered.
func (s *BaseECMAScriptListener) EnterInputElementTemplateTail(ctx *InputElementTemplateTailContext) {}

// ExitInputElementTemplateTail is called when production inputElementTemplateTail is exited.
func (s *BaseECMAScriptListener) ExitInputElementTemplateTail(ctx *InputElementTemplateTailContext) {}

// EnterRegularExpressionLiteral is called when production regularExpressionLiteral is entered.
func (s *BaseECMAScriptListener) EnterRegularExpressionLiteral(ctx *RegularExpressionLiteralContext) {}

// ExitRegularExpressionLiteral is called when production regularExpressionLiteral is exited.
func (s *BaseECMAScriptListener) ExitRegularExpressionLiteral(ctx *RegularExpressionLiteralContext) {}

// EnterIdentifierReference is called when production identifierReference is entered.
func (s *BaseECMAScriptListener) EnterIdentifierReference(ctx *IdentifierReferenceContext) {}

// ExitIdentifierReference is called when production identifierReference is exited.
func (s *BaseECMAScriptListener) ExitIdentifierReference(ctx *IdentifierReferenceContext) {}

// EnterIdentifierReference_Yield is called when production identifierReference_Yield is entered.
func (s *BaseECMAScriptListener) EnterIdentifierReference_Yield(ctx *IdentifierReference_YieldContext) {
}

// ExitIdentifierReference_Yield is called when production identifierReference_Yield is exited.
func (s *BaseECMAScriptListener) ExitIdentifierReference_Yield(ctx *IdentifierReference_YieldContext) {
}

// EnterIdentifierReference_Await is called when production identifierReference_Await is entered.
func (s *BaseECMAScriptListener) EnterIdentifierReference_Await(ctx *IdentifierReference_AwaitContext) {
}

// ExitIdentifierReference_Await is called when production identifierReference_Await is exited.
func (s *BaseECMAScriptListener) ExitIdentifierReference_Await(ctx *IdentifierReference_AwaitContext) {
}

// EnterIdentifierReference_Yield_Await is called when production identifierReference_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterIdentifierReference_Yield_Await(ctx *IdentifierReference_Yield_AwaitContext) {
}

// ExitIdentifierReference_Yield_Await is called when production identifierReference_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitIdentifierReference_Yield_Await(ctx *IdentifierReference_Yield_AwaitContext) {
}

// EnterBindingIdentifier is called when production bindingIdentifier is entered.
func (s *BaseECMAScriptListener) EnterBindingIdentifier(ctx *BindingIdentifierContext) {}

// ExitBindingIdentifier is called when production bindingIdentifier is exited.
func (s *BaseECMAScriptListener) ExitBindingIdentifier(ctx *BindingIdentifierContext) {}

// EnterBindingIdentifier_Yield is called when production bindingIdentifier_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingIdentifier_Yield(ctx *BindingIdentifier_YieldContext) {}

// ExitBindingIdentifier_Yield is called when production bindingIdentifier_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingIdentifier_Yield(ctx *BindingIdentifier_YieldContext) {}

// EnterBindingIdentifier_Await is called when production bindingIdentifier_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingIdentifier_Await(ctx *BindingIdentifier_AwaitContext) {}

// ExitBindingIdentifier_Await is called when production bindingIdentifier_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingIdentifier_Await(ctx *BindingIdentifier_AwaitContext) {}

// EnterBindingIdentifier_Yield_Await is called when production bindingIdentifier_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingIdentifier_Yield_Await(ctx *BindingIdentifier_Yield_AwaitContext) {
}

// ExitBindingIdentifier_Yield_Await is called when production bindingIdentifier_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingIdentifier_Yield_Await(ctx *BindingIdentifier_Yield_AwaitContext) {
}

// EnterLabelIdentifier is called when production labelIdentifier is entered.
func (s *BaseECMAScriptListener) EnterLabelIdentifier(ctx *LabelIdentifierContext) {}

// ExitLabelIdentifier is called when production labelIdentifier is exited.
func (s *BaseECMAScriptListener) ExitLabelIdentifier(ctx *LabelIdentifierContext) {}

// EnterLabelIdentifier_Yield is called when production labelIdentifier_Yield is entered.
func (s *BaseECMAScriptListener) EnterLabelIdentifier_Yield(ctx *LabelIdentifier_YieldContext) {}

// ExitLabelIdentifier_Yield is called when production labelIdentifier_Yield is exited.
func (s *BaseECMAScriptListener) ExitLabelIdentifier_Yield(ctx *LabelIdentifier_YieldContext) {}

// EnterLabelIdentifier_Await is called when production labelIdentifier_Await is entered.
func (s *BaseECMAScriptListener) EnterLabelIdentifier_Await(ctx *LabelIdentifier_AwaitContext) {}

// ExitLabelIdentifier_Await is called when production labelIdentifier_Await is exited.
func (s *BaseECMAScriptListener) ExitLabelIdentifier_Await(ctx *LabelIdentifier_AwaitContext) {}

// EnterLabelIdentifier_Yield_Await is called when production labelIdentifier_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLabelIdentifier_Yield_Await(ctx *LabelIdentifier_Yield_AwaitContext) {
}

// ExitLabelIdentifier_Yield_Await is called when production labelIdentifier_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLabelIdentifier_Yield_Await(ctx *LabelIdentifier_Yield_AwaitContext) {
}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseECMAScriptListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseECMAScriptListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterPrimaryExpression_Yield is called when production primaryExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterPrimaryExpression_Yield(ctx *PrimaryExpression_YieldContext) {}

// ExitPrimaryExpression_Yield is called when production primaryExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitPrimaryExpression_Yield(ctx *PrimaryExpression_YieldContext) {}

// EnterPrimaryExpression_Await is called when production primaryExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterPrimaryExpression_Await(ctx *PrimaryExpression_AwaitContext) {}

// ExitPrimaryExpression_Await is called when production primaryExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitPrimaryExpression_Await(ctx *PrimaryExpression_AwaitContext) {}

// EnterPrimaryExpression_Yield_Await is called when production primaryExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterPrimaryExpression_Yield_Await(ctx *PrimaryExpression_Yield_AwaitContext) {
}

// ExitPrimaryExpression_Yield_Await is called when production primaryExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitPrimaryExpression_Yield_Await(ctx *PrimaryExpression_Yield_AwaitContext) {
}

// EnterCoverParenthesizedExpressionAndArrowParameterList is called when production coverParenthesizedExpressionAndArrowParameterList is entered.
func (s *BaseECMAScriptListener) EnterCoverParenthesizedExpressionAndArrowParameterList(ctx *CoverParenthesizedExpressionAndArrowParameterListContext) {
}

// ExitCoverParenthesizedExpressionAndArrowParameterList is called when production coverParenthesizedExpressionAndArrowParameterList is exited.
func (s *BaseECMAScriptListener) ExitCoverParenthesizedExpressionAndArrowParameterList(ctx *CoverParenthesizedExpressionAndArrowParameterListContext) {
}

// EnterCoverParenthesizedExpressionAndArrowParameterList_Yield is called when production coverParenthesizedExpressionAndArrowParameterList_Yield is entered.
func (s *BaseECMAScriptListener) EnterCoverParenthesizedExpressionAndArrowParameterList_Yield(ctx *CoverParenthesizedExpressionAndArrowParameterList_YieldContext) {
}

// ExitCoverParenthesizedExpressionAndArrowParameterList_Yield is called when production coverParenthesizedExpressionAndArrowParameterList_Yield is exited.
func (s *BaseECMAScriptListener) ExitCoverParenthesizedExpressionAndArrowParameterList_Yield(ctx *CoverParenthesizedExpressionAndArrowParameterList_YieldContext) {
}

// EnterCoverParenthesizedExpressionAndArrowParameterList_Await is called when production coverParenthesizedExpressionAndArrowParameterList_Await is entered.
func (s *BaseECMAScriptListener) EnterCoverParenthesizedExpressionAndArrowParameterList_Await(ctx *CoverParenthesizedExpressionAndArrowParameterList_AwaitContext) {
}

// ExitCoverParenthesizedExpressionAndArrowParameterList_Await is called when production coverParenthesizedExpressionAndArrowParameterList_Await is exited.
func (s *BaseECMAScriptListener) ExitCoverParenthesizedExpressionAndArrowParameterList_Await(ctx *CoverParenthesizedExpressionAndArrowParameterList_AwaitContext) {
}

// EnterCoverParenthesizedExpressionAndArrowParameterList_Yield_Await is called when production coverParenthesizedExpressionAndArrowParameterList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCoverParenthesizedExpressionAndArrowParameterList_Yield_Await(ctx *CoverParenthesizedExpressionAndArrowParameterList_Yield_AwaitContext) {
}

// ExitCoverParenthesizedExpressionAndArrowParameterList_Yield_Await is called when production coverParenthesizedExpressionAndArrowParameterList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCoverParenthesizedExpressionAndArrowParameterList_Yield_Await(ctx *CoverParenthesizedExpressionAndArrowParameterList_Yield_AwaitContext) {
}

// EnterLiteral is called when production literal is entered.
func (s *BaseECMAScriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseECMAScriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseECMAScriptListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseECMAScriptListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterArrayLiteral_Yield is called when production arrayLiteral_Yield is entered.
func (s *BaseECMAScriptListener) EnterArrayLiteral_Yield(ctx *ArrayLiteral_YieldContext) {}

// ExitArrayLiteral_Yield is called when production arrayLiteral_Yield is exited.
func (s *BaseECMAScriptListener) ExitArrayLiteral_Yield(ctx *ArrayLiteral_YieldContext) {}

// EnterArrayLiteral_Await is called when production arrayLiteral_Await is entered.
func (s *BaseECMAScriptListener) EnterArrayLiteral_Await(ctx *ArrayLiteral_AwaitContext) {}

// ExitArrayLiteral_Await is called when production arrayLiteral_Await is exited.
func (s *BaseECMAScriptListener) ExitArrayLiteral_Await(ctx *ArrayLiteral_AwaitContext) {}

// EnterArrayLiteral_Yield_Await is called when production arrayLiteral_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArrayLiteral_Yield_Await(ctx *ArrayLiteral_Yield_AwaitContext) {}

// ExitArrayLiteral_Yield_Await is called when production arrayLiteral_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArrayLiteral_Yield_Await(ctx *ArrayLiteral_Yield_AwaitContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseECMAScriptListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseECMAScriptListener) ExitElementList(ctx *ElementListContext) {}

// EnterElementList_Yield is called when production elementList_Yield is entered.
func (s *BaseECMAScriptListener) EnterElementList_Yield(ctx *ElementList_YieldContext) {}

// ExitElementList_Yield is called when production elementList_Yield is exited.
func (s *BaseECMAScriptListener) ExitElementList_Yield(ctx *ElementList_YieldContext) {}

// EnterElementList_Await is called when production elementList_Await is entered.
func (s *BaseECMAScriptListener) EnterElementList_Await(ctx *ElementList_AwaitContext) {}

// ExitElementList_Await is called when production elementList_Await is exited.
func (s *BaseECMAScriptListener) ExitElementList_Await(ctx *ElementList_AwaitContext) {}

// EnterElementList_Yield_Await is called when production elementList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterElementList_Yield_Await(ctx *ElementList_Yield_AwaitContext) {}

// ExitElementList_Yield_Await is called when production elementList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitElementList_Yield_Await(ctx *ElementList_Yield_AwaitContext) {}

// EnterElision is called when production elision is entered.
func (s *BaseECMAScriptListener) EnterElision(ctx *ElisionContext) {}

// ExitElision is called when production elision is exited.
func (s *BaseECMAScriptListener) ExitElision(ctx *ElisionContext) {}

// EnterSpreadElement is called when production spreadElement is entered.
func (s *BaseECMAScriptListener) EnterSpreadElement(ctx *SpreadElementContext) {}

// ExitSpreadElement is called when production spreadElement is exited.
func (s *BaseECMAScriptListener) ExitSpreadElement(ctx *SpreadElementContext) {}

// EnterSpreadElement_Yield is called when production spreadElement_Yield is entered.
func (s *BaseECMAScriptListener) EnterSpreadElement_Yield(ctx *SpreadElement_YieldContext) {}

// ExitSpreadElement_Yield is called when production spreadElement_Yield is exited.
func (s *BaseECMAScriptListener) ExitSpreadElement_Yield(ctx *SpreadElement_YieldContext) {}

// EnterSpreadElement_Await is called when production spreadElement_Await is entered.
func (s *BaseECMAScriptListener) EnterSpreadElement_Await(ctx *SpreadElement_AwaitContext) {}

// ExitSpreadElement_Await is called when production spreadElement_Await is exited.
func (s *BaseECMAScriptListener) ExitSpreadElement_Await(ctx *SpreadElement_AwaitContext) {}

// EnterSpreadElement_Yield_Await is called when production spreadElement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterSpreadElement_Yield_Await(ctx *SpreadElement_Yield_AwaitContext) {
}

// ExitSpreadElement_Yield_Await is called when production spreadElement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitSpreadElement_Yield_Await(ctx *SpreadElement_Yield_AwaitContext) {
}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseECMAScriptListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseECMAScriptListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterObjectLiteral_Yield is called when production objectLiteral_Yield is entered.
func (s *BaseECMAScriptListener) EnterObjectLiteral_Yield(ctx *ObjectLiteral_YieldContext) {}

// ExitObjectLiteral_Yield is called when production objectLiteral_Yield is exited.
func (s *BaseECMAScriptListener) ExitObjectLiteral_Yield(ctx *ObjectLiteral_YieldContext) {}

// EnterObjectLiteral_Await is called when production objectLiteral_Await is entered.
func (s *BaseECMAScriptListener) EnterObjectLiteral_Await(ctx *ObjectLiteral_AwaitContext) {}

// ExitObjectLiteral_Await is called when production objectLiteral_Await is exited.
func (s *BaseECMAScriptListener) ExitObjectLiteral_Await(ctx *ObjectLiteral_AwaitContext) {}

// EnterObjectLiteral_Yield_Await is called when production objectLiteral_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterObjectLiteral_Yield_Await(ctx *ObjectLiteral_Yield_AwaitContext) {
}

// ExitObjectLiteral_Yield_Await is called when production objectLiteral_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitObjectLiteral_Yield_Await(ctx *ObjectLiteral_Yield_AwaitContext) {
}

// EnterPropertyDefinitionList is called when production propertyDefinitionList is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinitionList(ctx *PropertyDefinitionListContext) {}

// ExitPropertyDefinitionList is called when production propertyDefinitionList is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinitionList(ctx *PropertyDefinitionListContext) {}

// EnterPropertyDefinitionList_Yield is called when production propertyDefinitionList_Yield is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinitionList_Yield(ctx *PropertyDefinitionList_YieldContext) {
}

// ExitPropertyDefinitionList_Yield is called when production propertyDefinitionList_Yield is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinitionList_Yield(ctx *PropertyDefinitionList_YieldContext) {
}

// EnterPropertyDefinitionList_Await is called when production propertyDefinitionList_Await is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinitionList_Await(ctx *PropertyDefinitionList_AwaitContext) {
}

// ExitPropertyDefinitionList_Await is called when production propertyDefinitionList_Await is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinitionList_Await(ctx *PropertyDefinitionList_AwaitContext) {
}

// EnterPropertyDefinitionList_Yield_Await is called when production propertyDefinitionList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinitionList_Yield_Await(ctx *PropertyDefinitionList_Yield_AwaitContext) {
}

// ExitPropertyDefinitionList_Yield_Await is called when production propertyDefinitionList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinitionList_Yield_Await(ctx *PropertyDefinitionList_Yield_AwaitContext) {
}

// EnterPropertyDefinition is called when production propertyDefinition is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinition(ctx *PropertyDefinitionContext) {}

// ExitPropertyDefinition is called when production propertyDefinition is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinition(ctx *PropertyDefinitionContext) {}

// EnterPropertyDefinition_Yield is called when production propertyDefinition_Yield is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinition_Yield(ctx *PropertyDefinition_YieldContext) {}

// ExitPropertyDefinition_Yield is called when production propertyDefinition_Yield is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinition_Yield(ctx *PropertyDefinition_YieldContext) {}

// EnterPropertyDefinition_Await is called when production propertyDefinition_Await is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinition_Await(ctx *PropertyDefinition_AwaitContext) {}

// ExitPropertyDefinition_Await is called when production propertyDefinition_Await is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinition_Await(ctx *PropertyDefinition_AwaitContext) {}

// EnterPropertyDefinition_Yield_Await is called when production propertyDefinition_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterPropertyDefinition_Yield_Await(ctx *PropertyDefinition_Yield_AwaitContext) {
}

// ExitPropertyDefinition_Yield_Await is called when production propertyDefinition_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitPropertyDefinition_Yield_Await(ctx *PropertyDefinition_Yield_AwaitContext) {
}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseECMAScriptListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseECMAScriptListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterPropertyName_Yield is called when production propertyName_Yield is entered.
func (s *BaseECMAScriptListener) EnterPropertyName_Yield(ctx *PropertyName_YieldContext) {}

// ExitPropertyName_Yield is called when production propertyName_Yield is exited.
func (s *BaseECMAScriptListener) ExitPropertyName_Yield(ctx *PropertyName_YieldContext) {}

// EnterPropertyName_Await is called when production propertyName_Await is entered.
func (s *BaseECMAScriptListener) EnterPropertyName_Await(ctx *PropertyName_AwaitContext) {}

// ExitPropertyName_Await is called when production propertyName_Await is exited.
func (s *BaseECMAScriptListener) ExitPropertyName_Await(ctx *PropertyName_AwaitContext) {}

// EnterPropertyName_Yield_Await is called when production propertyName_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterPropertyName_Yield_Await(ctx *PropertyName_Yield_AwaitContext) {}

// ExitPropertyName_Yield_Await is called when production propertyName_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitPropertyName_Yield_Await(ctx *PropertyName_Yield_AwaitContext) {}

// EnterLiteralPropertyName is called when production literalPropertyName is entered.
func (s *BaseECMAScriptListener) EnterLiteralPropertyName(ctx *LiteralPropertyNameContext) {}

// ExitLiteralPropertyName is called when production literalPropertyName is exited.
func (s *BaseECMAScriptListener) ExitLiteralPropertyName(ctx *LiteralPropertyNameContext) {}

// EnterComputedPropertyName is called when production computedPropertyName is entered.
func (s *BaseECMAScriptListener) EnterComputedPropertyName(ctx *ComputedPropertyNameContext) {}

// ExitComputedPropertyName is called when production computedPropertyName is exited.
func (s *BaseECMAScriptListener) ExitComputedPropertyName(ctx *ComputedPropertyNameContext) {}

// EnterComputedPropertyName_Yield is called when production computedPropertyName_Yield is entered.
func (s *BaseECMAScriptListener) EnterComputedPropertyName_Yield(ctx *ComputedPropertyName_YieldContext) {
}

// ExitComputedPropertyName_Yield is called when production computedPropertyName_Yield is exited.
func (s *BaseECMAScriptListener) ExitComputedPropertyName_Yield(ctx *ComputedPropertyName_YieldContext) {
}

// EnterComputedPropertyName_Await is called when production computedPropertyName_Await is entered.
func (s *BaseECMAScriptListener) EnterComputedPropertyName_Await(ctx *ComputedPropertyName_AwaitContext) {
}

// ExitComputedPropertyName_Await is called when production computedPropertyName_Await is exited.
func (s *BaseECMAScriptListener) ExitComputedPropertyName_Await(ctx *ComputedPropertyName_AwaitContext) {
}

// EnterComputedPropertyName_Yield_Await is called when production computedPropertyName_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterComputedPropertyName_Yield_Await(ctx *ComputedPropertyName_Yield_AwaitContext) {
}

// ExitComputedPropertyName_Yield_Await is called when production computedPropertyName_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitComputedPropertyName_Yield_Await(ctx *ComputedPropertyName_Yield_AwaitContext) {
}

// EnterCoverInitializedName is called when production coverInitializedName is entered.
func (s *BaseECMAScriptListener) EnterCoverInitializedName(ctx *CoverInitializedNameContext) {}

// ExitCoverInitializedName is called when production coverInitializedName is exited.
func (s *BaseECMAScriptListener) ExitCoverInitializedName(ctx *CoverInitializedNameContext) {}

// EnterCoverInitializedName_Yield is called when production coverInitializedName_Yield is entered.
func (s *BaseECMAScriptListener) EnterCoverInitializedName_Yield(ctx *CoverInitializedName_YieldContext) {
}

// ExitCoverInitializedName_Yield is called when production coverInitializedName_Yield is exited.
func (s *BaseECMAScriptListener) ExitCoverInitializedName_Yield(ctx *CoverInitializedName_YieldContext) {
}

// EnterCoverInitializedName_Await is called when production coverInitializedName_Await is entered.
func (s *BaseECMAScriptListener) EnterCoverInitializedName_Await(ctx *CoverInitializedName_AwaitContext) {
}

// ExitCoverInitializedName_Await is called when production coverInitializedName_Await is exited.
func (s *BaseECMAScriptListener) ExitCoverInitializedName_Await(ctx *CoverInitializedName_AwaitContext) {
}

// EnterCoverInitializedName_Yield_Await is called when production coverInitializedName_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCoverInitializedName_Yield_Await(ctx *CoverInitializedName_Yield_AwaitContext) {
}

// ExitCoverInitializedName_Yield_Await is called when production coverInitializedName_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCoverInitializedName_Yield_Await(ctx *CoverInitializedName_Yield_AwaitContext) {
}

// EnterInitializer is called when production initializer is entered.
func (s *BaseECMAScriptListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseECMAScriptListener) ExitInitializer(ctx *InitializerContext) {}

// EnterInitializer_In is called when production initializer_In is entered.
func (s *BaseECMAScriptListener) EnterInitializer_In(ctx *Initializer_InContext) {}

// ExitInitializer_In is called when production initializer_In is exited.
func (s *BaseECMAScriptListener) ExitInitializer_In(ctx *Initializer_InContext) {}

// EnterInitializer_Yield is called when production initializer_Yield is entered.
func (s *BaseECMAScriptListener) EnterInitializer_Yield(ctx *Initializer_YieldContext) {}

// ExitInitializer_Yield is called when production initializer_Yield is exited.
func (s *BaseECMAScriptListener) ExitInitializer_Yield(ctx *Initializer_YieldContext) {}

// EnterInitializer_In_Yield is called when production initializer_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterInitializer_In_Yield(ctx *Initializer_In_YieldContext) {}

// ExitInitializer_In_Yield is called when production initializer_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitInitializer_In_Yield(ctx *Initializer_In_YieldContext) {}

// EnterInitializer_Await is called when production initializer_Await is entered.
func (s *BaseECMAScriptListener) EnterInitializer_Await(ctx *Initializer_AwaitContext) {}

// ExitInitializer_Await is called when production initializer_Await is exited.
func (s *BaseECMAScriptListener) ExitInitializer_Await(ctx *Initializer_AwaitContext) {}

// EnterInitializer_In_Await is called when production initializer_In_Await is entered.
func (s *BaseECMAScriptListener) EnterInitializer_In_Await(ctx *Initializer_In_AwaitContext) {}

// ExitInitializer_In_Await is called when production initializer_In_Await is exited.
func (s *BaseECMAScriptListener) ExitInitializer_In_Await(ctx *Initializer_In_AwaitContext) {}

// EnterInitializer_Yield_Await is called when production initializer_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterInitializer_Yield_Await(ctx *Initializer_Yield_AwaitContext) {}

// ExitInitializer_Yield_Await is called when production initializer_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitInitializer_Yield_Await(ctx *Initializer_Yield_AwaitContext) {}

// EnterInitializer_In_Yield_Await is called when production initializer_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterInitializer_In_Yield_Await(ctx *Initializer_In_Yield_AwaitContext) {
}

// ExitInitializer_In_Yield_Await is called when production initializer_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitInitializer_In_Yield_Await(ctx *Initializer_In_Yield_AwaitContext) {
}

// EnterTemplateLiteral is called when production templateLiteral is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral(ctx *TemplateLiteralContext) {}

// ExitTemplateLiteral is called when production templateLiteral is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral(ctx *TemplateLiteralContext) {}

// EnterTemplateLiteral_Yield is called when production templateLiteral_Yield is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Yield(ctx *TemplateLiteral_YieldContext) {}

// ExitTemplateLiteral_Yield is called when production templateLiteral_Yield is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Yield(ctx *TemplateLiteral_YieldContext) {}

// EnterTemplateLiteral_Await is called when production templateLiteral_Await is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Await(ctx *TemplateLiteral_AwaitContext) {}

// ExitTemplateLiteral_Await is called when production templateLiteral_Await is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Await(ctx *TemplateLiteral_AwaitContext) {}

// EnterTemplateLiteral_Yield_Await is called when production templateLiteral_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Yield_Await(ctx *TemplateLiteral_Yield_AwaitContext) {
}

// ExitTemplateLiteral_Yield_Await is called when production templateLiteral_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Yield_Await(ctx *TemplateLiteral_Yield_AwaitContext) {
}

// EnterTemplateLiteral_Tagged is called when production templateLiteral_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Tagged(ctx *TemplateLiteral_TaggedContext) {}

// ExitTemplateLiteral_Tagged is called when production templateLiteral_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Tagged(ctx *TemplateLiteral_TaggedContext) {}

// EnterTemplateLiteral_Yield_Tagged is called when production templateLiteral_Yield_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Yield_Tagged(ctx *TemplateLiteral_Yield_TaggedContext) {
}

// ExitTemplateLiteral_Yield_Tagged is called when production templateLiteral_Yield_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Yield_Tagged(ctx *TemplateLiteral_Yield_TaggedContext) {
}

// EnterTemplateLiteral_Await_Tagged is called when production templateLiteral_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Await_Tagged(ctx *TemplateLiteral_Await_TaggedContext) {
}

// ExitTemplateLiteral_Await_Tagged is called when production templateLiteral_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Await_Tagged(ctx *TemplateLiteral_Await_TaggedContext) {
}

// EnterTemplateLiteral_Yield_Await_Tagged is called when production templateLiteral_Yield_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateLiteral_Yield_Await_Tagged(ctx *TemplateLiteral_Yield_Await_TaggedContext) {
}

// ExitTemplateLiteral_Yield_Await_Tagged is called when production templateLiteral_Yield_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateLiteral_Yield_Await_Tagged(ctx *TemplateLiteral_Yield_Await_TaggedContext) {
}

// EnterSubstitutionTemplate is called when production substitutionTemplate is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate(ctx *SubstitutionTemplateContext) {}

// ExitSubstitutionTemplate is called when production substitutionTemplate is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate(ctx *SubstitutionTemplateContext) {}

// EnterSubstitutionTemplate_Yield is called when production substitutionTemplate_Yield is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Yield(ctx *SubstitutionTemplate_YieldContext) {
}

// ExitSubstitutionTemplate_Yield is called when production substitutionTemplate_Yield is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Yield(ctx *SubstitutionTemplate_YieldContext) {
}

// EnterSubstitutionTemplate_Await is called when production substitutionTemplate_Await is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Await(ctx *SubstitutionTemplate_AwaitContext) {
}

// ExitSubstitutionTemplate_Await is called when production substitutionTemplate_Await is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Await(ctx *SubstitutionTemplate_AwaitContext) {
}

// EnterSubstitutionTemplate_Yield_Await is called when production substitutionTemplate_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Yield_Await(ctx *SubstitutionTemplate_Yield_AwaitContext) {
}

// ExitSubstitutionTemplate_Yield_Await is called when production substitutionTemplate_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Yield_Await(ctx *SubstitutionTemplate_Yield_AwaitContext) {
}

// EnterSubstitutionTemplate_Tagged is called when production substitutionTemplate_Tagged is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Tagged(ctx *SubstitutionTemplate_TaggedContext) {
}

// ExitSubstitutionTemplate_Tagged is called when production substitutionTemplate_Tagged is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Tagged(ctx *SubstitutionTemplate_TaggedContext) {
}

// EnterSubstitutionTemplate_Yield_Tagged is called when production substitutionTemplate_Yield_Tagged is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Yield_Tagged(ctx *SubstitutionTemplate_Yield_TaggedContext) {
}

// ExitSubstitutionTemplate_Yield_Tagged is called when production substitutionTemplate_Yield_Tagged is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Yield_Tagged(ctx *SubstitutionTemplate_Yield_TaggedContext) {
}

// EnterSubstitutionTemplate_Await_Tagged is called when production substitutionTemplate_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Await_Tagged(ctx *SubstitutionTemplate_Await_TaggedContext) {
}

// ExitSubstitutionTemplate_Await_Tagged is called when production substitutionTemplate_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Await_Tagged(ctx *SubstitutionTemplate_Await_TaggedContext) {
}

// EnterSubstitutionTemplate_Yield_Await_Tagged is called when production substitutionTemplate_Yield_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterSubstitutionTemplate_Yield_Await_Tagged(ctx *SubstitutionTemplate_Yield_Await_TaggedContext) {
}

// ExitSubstitutionTemplate_Yield_Await_Tagged is called when production substitutionTemplate_Yield_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitSubstitutionTemplate_Yield_Await_Tagged(ctx *SubstitutionTemplate_Yield_Await_TaggedContext) {
}

// EnterTemplateSpans is called when production templateSpans is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans(ctx *TemplateSpansContext) {}

// ExitTemplateSpans is called when production templateSpans is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans(ctx *TemplateSpansContext) {}

// EnterTemplateSpans_Yield is called when production templateSpans_Yield is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Yield(ctx *TemplateSpans_YieldContext) {}

// ExitTemplateSpans_Yield is called when production templateSpans_Yield is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Yield(ctx *TemplateSpans_YieldContext) {}

// EnterTemplateSpans_Await is called when production templateSpans_Await is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Await(ctx *TemplateSpans_AwaitContext) {}

// ExitTemplateSpans_Await is called when production templateSpans_Await is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Await(ctx *TemplateSpans_AwaitContext) {}

// EnterTemplateSpans_Yield_Await is called when production templateSpans_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Yield_Await(ctx *TemplateSpans_Yield_AwaitContext) {
}

// ExitTemplateSpans_Yield_Await is called when production templateSpans_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Yield_Await(ctx *TemplateSpans_Yield_AwaitContext) {
}

// EnterTemplateSpans_Tagged is called when production templateSpans_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Tagged(ctx *TemplateSpans_TaggedContext) {}

// ExitTemplateSpans_Tagged is called when production templateSpans_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Tagged(ctx *TemplateSpans_TaggedContext) {}

// EnterTemplateSpans_Yield_Tagged is called when production templateSpans_Yield_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Yield_Tagged(ctx *TemplateSpans_Yield_TaggedContext) {
}

// ExitTemplateSpans_Yield_Tagged is called when production templateSpans_Yield_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Yield_Tagged(ctx *TemplateSpans_Yield_TaggedContext) {
}

// EnterTemplateSpans_Await_Tagged is called when production templateSpans_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Await_Tagged(ctx *TemplateSpans_Await_TaggedContext) {
}

// ExitTemplateSpans_Await_Tagged is called when production templateSpans_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Await_Tagged(ctx *TemplateSpans_Await_TaggedContext) {
}

// EnterTemplateSpans_Yield_Await_Tagged is called when production templateSpans_Yield_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateSpans_Yield_Await_Tagged(ctx *TemplateSpans_Yield_Await_TaggedContext) {
}

// ExitTemplateSpans_Yield_Await_Tagged is called when production templateSpans_Yield_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateSpans_Yield_Await_Tagged(ctx *TemplateSpans_Yield_Await_TaggedContext) {
}

// EnterTemplateMiddleList is called when production templateMiddleList is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList(ctx *TemplateMiddleListContext) {}

// ExitTemplateMiddleList is called when production templateMiddleList is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList(ctx *TemplateMiddleListContext) {}

// EnterTemplateMiddleList_Yield is called when production templateMiddleList_Yield is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Yield(ctx *TemplateMiddleList_YieldContext) {}

// ExitTemplateMiddleList_Yield is called when production templateMiddleList_Yield is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Yield(ctx *TemplateMiddleList_YieldContext) {}

// EnterTemplateMiddleList_Await is called when production templateMiddleList_Await is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Await(ctx *TemplateMiddleList_AwaitContext) {}

// ExitTemplateMiddleList_Await is called when production templateMiddleList_Await is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Await(ctx *TemplateMiddleList_AwaitContext) {}

// EnterTemplateMiddleList_Yield_Await is called when production templateMiddleList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Yield_Await(ctx *TemplateMiddleList_Yield_AwaitContext) {
}

// ExitTemplateMiddleList_Yield_Await is called when production templateMiddleList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Yield_Await(ctx *TemplateMiddleList_Yield_AwaitContext) {
}

// EnterTemplateMiddleList_Tagged is called when production templateMiddleList_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Tagged(ctx *TemplateMiddleList_TaggedContext) {
}

// ExitTemplateMiddleList_Tagged is called when production templateMiddleList_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Tagged(ctx *TemplateMiddleList_TaggedContext) {
}

// EnterTemplateMiddleList_Yield_Tagged is called when production templateMiddleList_Yield_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Yield_Tagged(ctx *TemplateMiddleList_Yield_TaggedContext) {
}

// ExitTemplateMiddleList_Yield_Tagged is called when production templateMiddleList_Yield_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Yield_Tagged(ctx *TemplateMiddleList_Yield_TaggedContext) {
}

// EnterTemplateMiddleList_Await_Tagged is called when production templateMiddleList_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Await_Tagged(ctx *TemplateMiddleList_Await_TaggedContext) {
}

// ExitTemplateMiddleList_Await_Tagged is called when production templateMiddleList_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Await_Tagged(ctx *TemplateMiddleList_Await_TaggedContext) {
}

// EnterTemplateMiddleList_Yield_Await_Tagged is called when production templateMiddleList_Yield_Await_Tagged is entered.
func (s *BaseECMAScriptListener) EnterTemplateMiddleList_Yield_Await_Tagged(ctx *TemplateMiddleList_Yield_Await_TaggedContext) {
}

// ExitTemplateMiddleList_Yield_Await_Tagged is called when production templateMiddleList_Yield_Await_Tagged is exited.
func (s *BaseECMAScriptListener) ExitTemplateMiddleList_Yield_Await_Tagged(ctx *TemplateMiddleList_Yield_Await_TaggedContext) {
}

// EnterMemberExpression is called when production memberExpression is entered.
func (s *BaseECMAScriptListener) EnterMemberExpression(ctx *MemberExpressionContext) {}

// ExitMemberExpression is called when production memberExpression is exited.
func (s *BaseECMAScriptListener) ExitMemberExpression(ctx *MemberExpressionContext) {}

// EnterMemberExpression_Yield is called when production memberExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterMemberExpression_Yield(ctx *MemberExpression_YieldContext) {}

// ExitMemberExpression_Yield is called when production memberExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitMemberExpression_Yield(ctx *MemberExpression_YieldContext) {}

// EnterMemberExpression_Await is called when production memberExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterMemberExpression_Await(ctx *MemberExpression_AwaitContext) {}

// ExitMemberExpression_Await is called when production memberExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitMemberExpression_Await(ctx *MemberExpression_AwaitContext) {}

// EnterMemberExpression_Yield_Await is called when production memberExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterMemberExpression_Yield_Await(ctx *MemberExpression_Yield_AwaitContext) {
}

// ExitMemberExpression_Yield_Await is called when production memberExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitMemberExpression_Yield_Await(ctx *MemberExpression_Yield_AwaitContext) {
}

// EnterSuperProperty is called when production superProperty is entered.
func (s *BaseECMAScriptListener) EnterSuperProperty(ctx *SuperPropertyContext) {}

// ExitSuperProperty is called when production superProperty is exited.
func (s *BaseECMAScriptListener) ExitSuperProperty(ctx *SuperPropertyContext) {}

// EnterSuperProperty_Yield is called when production superProperty_Yield is entered.
func (s *BaseECMAScriptListener) EnterSuperProperty_Yield(ctx *SuperProperty_YieldContext) {}

// ExitSuperProperty_Yield is called when production superProperty_Yield is exited.
func (s *BaseECMAScriptListener) ExitSuperProperty_Yield(ctx *SuperProperty_YieldContext) {}

// EnterSuperProperty_Await is called when production superProperty_Await is entered.
func (s *BaseECMAScriptListener) EnterSuperProperty_Await(ctx *SuperProperty_AwaitContext) {}

// ExitSuperProperty_Await is called when production superProperty_Await is exited.
func (s *BaseECMAScriptListener) ExitSuperProperty_Await(ctx *SuperProperty_AwaitContext) {}

// EnterSuperProperty_Yield_Await is called when production superProperty_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterSuperProperty_Yield_Await(ctx *SuperProperty_Yield_AwaitContext) {
}

// ExitSuperProperty_Yield_Await is called when production superProperty_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitSuperProperty_Yield_Await(ctx *SuperProperty_Yield_AwaitContext) {
}

// EnterMetaProperty is called when production metaProperty is entered.
func (s *BaseECMAScriptListener) EnterMetaProperty(ctx *MetaPropertyContext) {}

// ExitMetaProperty is called when production metaProperty is exited.
func (s *BaseECMAScriptListener) ExitMetaProperty(ctx *MetaPropertyContext) {}

// EnterNewTarget is called when production newTarget is entered.
func (s *BaseECMAScriptListener) EnterNewTarget(ctx *NewTargetContext) {}

// ExitNewTarget is called when production newTarget is exited.
func (s *BaseECMAScriptListener) ExitNewTarget(ctx *NewTargetContext) {}

// EnterTheNewExpression is called when production theNewExpression is entered.
func (s *BaseECMAScriptListener) EnterTheNewExpression(ctx *TheNewExpressionContext) {}

// ExitTheNewExpression is called when production theNewExpression is exited.
func (s *BaseECMAScriptListener) ExitTheNewExpression(ctx *TheNewExpressionContext) {}

// EnterTheNewExpression_Yield is called when production theNewExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterTheNewExpression_Yield(ctx *TheNewExpression_YieldContext) {}

// ExitTheNewExpression_Yield is called when production theNewExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitTheNewExpression_Yield(ctx *TheNewExpression_YieldContext) {}

// EnterTheNewExpression_Await is called when production theNewExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterTheNewExpression_Await(ctx *TheNewExpression_AwaitContext) {}

// ExitTheNewExpression_Await is called when production theNewExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitTheNewExpression_Await(ctx *TheNewExpression_AwaitContext) {}

// EnterTheNewExpression_Yield_Await is called when production theNewExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterTheNewExpression_Yield_Await(ctx *TheNewExpression_Yield_AwaitContext) {
}

// ExitTheNewExpression_Yield_Await is called when production theNewExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitTheNewExpression_Yield_Await(ctx *TheNewExpression_Yield_AwaitContext) {
}

// EnterCallExpression is called when production callExpression is entered.
func (s *BaseECMAScriptListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production callExpression is exited.
func (s *BaseECMAScriptListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterCallExpression_Yield is called when production callExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterCallExpression_Yield(ctx *CallExpression_YieldContext) {}

// ExitCallExpression_Yield is called when production callExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitCallExpression_Yield(ctx *CallExpression_YieldContext) {}

// EnterCallExpression_Await is called when production callExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterCallExpression_Await(ctx *CallExpression_AwaitContext) {}

// ExitCallExpression_Await is called when production callExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitCallExpression_Await(ctx *CallExpression_AwaitContext) {}

// EnterCallExpression_Yield_Await is called when production callExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCallExpression_Yield_Await(ctx *CallExpression_Yield_AwaitContext) {
}

// ExitCallExpression_Yield_Await is called when production callExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCallExpression_Yield_Await(ctx *CallExpression_Yield_AwaitContext) {
}

// EnterSuperCall is called when production superCall is entered.
func (s *BaseECMAScriptListener) EnterSuperCall(ctx *SuperCallContext) {}

// ExitSuperCall is called when production superCall is exited.
func (s *BaseECMAScriptListener) ExitSuperCall(ctx *SuperCallContext) {}

// EnterSuperCall_Yield is called when production superCall_Yield is entered.
func (s *BaseECMAScriptListener) EnterSuperCall_Yield(ctx *SuperCall_YieldContext) {}

// ExitSuperCall_Yield is called when production superCall_Yield is exited.
func (s *BaseECMAScriptListener) ExitSuperCall_Yield(ctx *SuperCall_YieldContext) {}

// EnterSuperCall_Await is called when production superCall_Await is entered.
func (s *BaseECMAScriptListener) EnterSuperCall_Await(ctx *SuperCall_AwaitContext) {}

// ExitSuperCall_Await is called when production superCall_Await is exited.
func (s *BaseECMAScriptListener) ExitSuperCall_Await(ctx *SuperCall_AwaitContext) {}

// EnterSuperCall_Yield_Await is called when production superCall_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterSuperCall_Yield_Await(ctx *SuperCall_Yield_AwaitContext) {}

// ExitSuperCall_Yield_Await is called when production superCall_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitSuperCall_Yield_Await(ctx *SuperCall_Yield_AwaitContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseECMAScriptListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseECMAScriptListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArguments_Yield is called when production arguments_Yield is entered.
func (s *BaseECMAScriptListener) EnterArguments_Yield(ctx *Arguments_YieldContext) {}

// ExitArguments_Yield is called when production arguments_Yield is exited.
func (s *BaseECMAScriptListener) ExitArguments_Yield(ctx *Arguments_YieldContext) {}

// EnterArguments_Await is called when production arguments_Await is entered.
func (s *BaseECMAScriptListener) EnterArguments_Await(ctx *Arguments_AwaitContext) {}

// ExitArguments_Await is called when production arguments_Await is exited.
func (s *BaseECMAScriptListener) ExitArguments_Await(ctx *Arguments_AwaitContext) {}

// EnterArguments_Yield_Await is called when production arguments_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArguments_Yield_Await(ctx *Arguments_Yield_AwaitContext) {}

// ExitArguments_Yield_Await is called when production arguments_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArguments_Yield_Await(ctx *Arguments_Yield_AwaitContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseECMAScriptListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseECMAScriptListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArgumentList_Yield is called when production argumentList_Yield is entered.
func (s *BaseECMAScriptListener) EnterArgumentList_Yield(ctx *ArgumentList_YieldContext) {}

// ExitArgumentList_Yield is called when production argumentList_Yield is exited.
func (s *BaseECMAScriptListener) ExitArgumentList_Yield(ctx *ArgumentList_YieldContext) {}

// EnterArgumentList_Await is called when production argumentList_Await is entered.
func (s *BaseECMAScriptListener) EnterArgumentList_Await(ctx *ArgumentList_AwaitContext) {}

// ExitArgumentList_Await is called when production argumentList_Await is exited.
func (s *BaseECMAScriptListener) ExitArgumentList_Await(ctx *ArgumentList_AwaitContext) {}

// EnterArgumentList_Yield_Await is called when production argumentList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArgumentList_Yield_Await(ctx *ArgumentList_Yield_AwaitContext) {}

// ExitArgumentList_Yield_Await is called when production argumentList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArgumentList_Yield_Await(ctx *ArgumentList_Yield_AwaitContext) {}

// EnterLeftHandSideExpression is called when production leftHandSideExpression is entered.
func (s *BaseECMAScriptListener) EnterLeftHandSideExpression(ctx *LeftHandSideExpressionContext) {}

// ExitLeftHandSideExpression is called when production leftHandSideExpression is exited.
func (s *BaseECMAScriptListener) ExitLeftHandSideExpression(ctx *LeftHandSideExpressionContext) {}

// EnterLeftHandSideExpression_Yield is called when production leftHandSideExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterLeftHandSideExpression_Yield(ctx *LeftHandSideExpression_YieldContext) {
}

// ExitLeftHandSideExpression_Yield is called when production leftHandSideExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitLeftHandSideExpression_Yield(ctx *LeftHandSideExpression_YieldContext) {
}

// EnterLeftHandSideExpression_Await is called when production leftHandSideExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterLeftHandSideExpression_Await(ctx *LeftHandSideExpression_AwaitContext) {
}

// ExitLeftHandSideExpression_Await is called when production leftHandSideExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitLeftHandSideExpression_Await(ctx *LeftHandSideExpression_AwaitContext) {
}

// EnterLeftHandSideExpression_Yield_Await is called when production leftHandSideExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLeftHandSideExpression_Yield_Await(ctx *LeftHandSideExpression_Yield_AwaitContext) {
}

// ExitLeftHandSideExpression_Yield_Await is called when production leftHandSideExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLeftHandSideExpression_Yield_Await(ctx *LeftHandSideExpression_Yield_AwaitContext) {
}

// EnterUpdateExpression is called when production updateExpression is entered.
func (s *BaseECMAScriptListener) EnterUpdateExpression(ctx *UpdateExpressionContext) {}

// ExitUpdateExpression is called when production updateExpression is exited.
func (s *BaseECMAScriptListener) ExitUpdateExpression(ctx *UpdateExpressionContext) {}

// EnterUpdateExpression_Yield is called when production updateExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterUpdateExpression_Yield(ctx *UpdateExpression_YieldContext) {}

// ExitUpdateExpression_Yield is called when production updateExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitUpdateExpression_Yield(ctx *UpdateExpression_YieldContext) {}

// EnterUpdateExpression_Await is called when production updateExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterUpdateExpression_Await(ctx *UpdateExpression_AwaitContext) {}

// ExitUpdateExpression_Await is called when production updateExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitUpdateExpression_Await(ctx *UpdateExpression_AwaitContext) {}

// EnterUpdateExpression_Yield_Await is called when production updateExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterUpdateExpression_Yield_Await(ctx *UpdateExpression_Yield_AwaitContext) {
}

// ExitUpdateExpression_Yield_Await is called when production updateExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitUpdateExpression_Yield_Await(ctx *UpdateExpression_Yield_AwaitContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseECMAScriptListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseECMAScriptListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryExpression_Yield is called when production unaryExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterUnaryExpression_Yield(ctx *UnaryExpression_YieldContext) {}

// ExitUnaryExpression_Yield is called when production unaryExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitUnaryExpression_Yield(ctx *UnaryExpression_YieldContext) {}

// EnterUnaryExpression_Await is called when production unaryExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterUnaryExpression_Await(ctx *UnaryExpression_AwaitContext) {}

// ExitUnaryExpression_Await is called when production unaryExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitUnaryExpression_Await(ctx *UnaryExpression_AwaitContext) {}

// EnterUnaryExpression_Yield_Await is called when production unaryExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterUnaryExpression_Yield_Await(ctx *UnaryExpression_Yield_AwaitContext) {
}

// ExitUnaryExpression_Yield_Await is called when production unaryExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitUnaryExpression_Yield_Await(ctx *UnaryExpression_Yield_AwaitContext) {
}

// EnterExponentationExpression is called when production exponentationExpression is entered.
func (s *BaseECMAScriptListener) EnterExponentationExpression(ctx *ExponentationExpressionContext) {}

// ExitExponentationExpression is called when production exponentationExpression is exited.
func (s *BaseECMAScriptListener) ExitExponentationExpression(ctx *ExponentationExpressionContext) {}

// EnterExponentationExpression_Yield is called when production exponentationExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterExponentationExpression_Yield(ctx *ExponentationExpression_YieldContext) {
}

// ExitExponentationExpression_Yield is called when production exponentationExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitExponentationExpression_Yield(ctx *ExponentationExpression_YieldContext) {
}

// EnterExponentationExpression_Await is called when production exponentationExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterExponentationExpression_Await(ctx *ExponentationExpression_AwaitContext) {
}

// ExitExponentationExpression_Await is called when production exponentationExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitExponentationExpression_Await(ctx *ExponentationExpression_AwaitContext) {
}

// EnterExponentationExpression_Yield_Await is called when production exponentationExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterExponentationExpression_Yield_Await(ctx *ExponentationExpression_Yield_AwaitContext) {
}

// ExitExponentationExpression_Yield_Await is called when production exponentationExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitExponentationExpression_Yield_Await(ctx *ExponentationExpression_Yield_AwaitContext) {
}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseECMAScriptListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseECMAScriptListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterMultiplicativeExpression_Yield is called when production multiplicativeExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterMultiplicativeExpression_Yield(ctx *MultiplicativeExpression_YieldContext) {
}

// ExitMultiplicativeExpression_Yield is called when production multiplicativeExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitMultiplicativeExpression_Yield(ctx *MultiplicativeExpression_YieldContext) {
}

// EnterMultiplicativeExpression_Await is called when production multiplicativeExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterMultiplicativeExpression_Await(ctx *MultiplicativeExpression_AwaitContext) {
}

// ExitMultiplicativeExpression_Await is called when production multiplicativeExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitMultiplicativeExpression_Await(ctx *MultiplicativeExpression_AwaitContext) {
}

// EnterMultiplicativeExpression_Yield_Await is called when production multiplicativeExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterMultiplicativeExpression_Yield_Await(ctx *MultiplicativeExpression_Yield_AwaitContext) {
}

// ExitMultiplicativeExpression_Yield_Await is called when production multiplicativeExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitMultiplicativeExpression_Yield_Await(ctx *MultiplicativeExpression_Yield_AwaitContext) {
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseECMAScriptListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseECMAScriptListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterAdditiveExpression_Yield is called when production additiveExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterAdditiveExpression_Yield(ctx *AdditiveExpression_YieldContext) {}

// ExitAdditiveExpression_Yield is called when production additiveExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitAdditiveExpression_Yield(ctx *AdditiveExpression_YieldContext) {}

// EnterAdditiveExpression_Await is called when production additiveExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterAdditiveExpression_Await(ctx *AdditiveExpression_AwaitContext) {}

// ExitAdditiveExpression_Await is called when production additiveExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitAdditiveExpression_Await(ctx *AdditiveExpression_AwaitContext) {}

// EnterAdditiveExpression_Yield_Await is called when production additiveExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAdditiveExpression_Yield_Await(ctx *AdditiveExpression_Yield_AwaitContext) {
}

// ExitAdditiveExpression_Yield_Await is called when production additiveExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAdditiveExpression_Yield_Await(ctx *AdditiveExpression_Yield_AwaitContext) {
}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseECMAScriptListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseECMAScriptListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterShiftExpression_Yield is called when production shiftExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterShiftExpression_Yield(ctx *ShiftExpression_YieldContext) {}

// ExitShiftExpression_Yield is called when production shiftExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitShiftExpression_Yield(ctx *ShiftExpression_YieldContext) {}

// EnterShiftExpression_Await is called when production shiftExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterShiftExpression_Await(ctx *ShiftExpression_AwaitContext) {}

// ExitShiftExpression_Await is called when production shiftExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitShiftExpression_Await(ctx *ShiftExpression_AwaitContext) {}

// EnterShiftExpression_Yield_Await is called when production shiftExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterShiftExpression_Yield_Await(ctx *ShiftExpression_Yield_AwaitContext) {
}

// ExitShiftExpression_Yield_Await is called when production shiftExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitShiftExpression_Yield_Await(ctx *ShiftExpression_Yield_AwaitContext) {
}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterRelationalExpression_In is called when production relationalExpression_In is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_In(ctx *RelationalExpression_InContext) {}

// ExitRelationalExpression_In is called when production relationalExpression_In is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_In(ctx *RelationalExpression_InContext) {}

// EnterRelationalExpression_Yield is called when production relationalExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_Yield(ctx *RelationalExpression_YieldContext) {
}

// ExitRelationalExpression_Yield is called when production relationalExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_Yield(ctx *RelationalExpression_YieldContext) {
}

// EnterRelationalExpression_In_Yield is called when production relationalExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_In_Yield(ctx *RelationalExpression_In_YieldContext) {
}

// ExitRelationalExpression_In_Yield is called when production relationalExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_In_Yield(ctx *RelationalExpression_In_YieldContext) {
}

// EnterRelationalExpression_Await is called when production relationalExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_Await(ctx *RelationalExpression_AwaitContext) {
}

// ExitRelationalExpression_Await is called when production relationalExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_Await(ctx *RelationalExpression_AwaitContext) {
}

// EnterRelationalExpression_In_Await is called when production relationalExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_In_Await(ctx *RelationalExpression_In_AwaitContext) {
}

// ExitRelationalExpression_In_Await is called when production relationalExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_In_Await(ctx *RelationalExpression_In_AwaitContext) {
}

// EnterRelationalExpression_Yield_Await is called when production relationalExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_Yield_Await(ctx *RelationalExpression_Yield_AwaitContext) {
}

// ExitRelationalExpression_Yield_Await is called when production relationalExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_Yield_Await(ctx *RelationalExpression_Yield_AwaitContext) {
}

// EnterRelationalExpression_In_Yield_Await is called when production relationalExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression_In_Yield_Await(ctx *RelationalExpression_In_Yield_AwaitContext) {
}

// ExitRelationalExpression_In_Yield_Await is called when production relationalExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression_In_Yield_Await(ctx *RelationalExpression_In_Yield_AwaitContext) {
}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterEqualityExpression_In is called when production equalityExpression_In is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_In(ctx *EqualityExpression_InContext) {}

// ExitEqualityExpression_In is called when production equalityExpression_In is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_In(ctx *EqualityExpression_InContext) {}

// EnterEqualityExpression_Yield is called when production equalityExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_Yield(ctx *EqualityExpression_YieldContext) {}

// ExitEqualityExpression_Yield is called when production equalityExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_Yield(ctx *EqualityExpression_YieldContext) {}

// EnterEqualityExpression_In_Yield is called when production equalityExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_In_Yield(ctx *EqualityExpression_In_YieldContext) {
}

// ExitEqualityExpression_In_Yield is called when production equalityExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_In_Yield(ctx *EqualityExpression_In_YieldContext) {
}

// EnterEqualityExpression_Await is called when production equalityExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_Await(ctx *EqualityExpression_AwaitContext) {}

// ExitEqualityExpression_Await is called when production equalityExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_Await(ctx *EqualityExpression_AwaitContext) {}

// EnterEqualityExpression_In_Await is called when production equalityExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_In_Await(ctx *EqualityExpression_In_AwaitContext) {
}

// ExitEqualityExpression_In_Await is called when production equalityExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_In_Await(ctx *EqualityExpression_In_AwaitContext) {
}

// EnterEqualityExpression_Yield_Await is called when production equalityExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_Yield_Await(ctx *EqualityExpression_Yield_AwaitContext) {
}

// ExitEqualityExpression_Yield_Await is called when production equalityExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_Yield_Await(ctx *EqualityExpression_Yield_AwaitContext) {
}

// EnterEqualityExpression_In_Yield_Await is called when production equalityExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression_In_Yield_Await(ctx *EqualityExpression_In_Yield_AwaitContext) {
}

// ExitEqualityExpression_In_Yield_Await is called when production equalityExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression_In_Yield_Await(ctx *EqualityExpression_In_Yield_AwaitContext) {
}

// EnterBitwiseANDExpression is called when production bitwiseANDExpression is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression(ctx *BitwiseANDExpressionContext) {}

// ExitBitwiseANDExpression is called when production bitwiseANDExpression is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression(ctx *BitwiseANDExpressionContext) {}

// EnterBitwiseANDExpression_In is called when production bitwiseANDExpression_In is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_In(ctx *BitwiseANDExpression_InContext) {}

// ExitBitwiseANDExpression_In is called when production bitwiseANDExpression_In is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_In(ctx *BitwiseANDExpression_InContext) {}

// EnterBitwiseANDExpression_Yield is called when production bitwiseANDExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_Yield(ctx *BitwiseANDExpression_YieldContext) {
}

// ExitBitwiseANDExpression_Yield is called when production bitwiseANDExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_Yield(ctx *BitwiseANDExpression_YieldContext) {
}

// EnterBitwiseANDExpression_In_Yield is called when production bitwiseANDExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_In_Yield(ctx *BitwiseANDExpression_In_YieldContext) {
}

// ExitBitwiseANDExpression_In_Yield is called when production bitwiseANDExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_In_Yield(ctx *BitwiseANDExpression_In_YieldContext) {
}

// EnterBitwiseANDExpression_Await is called when production bitwiseANDExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_Await(ctx *BitwiseANDExpression_AwaitContext) {
}

// ExitBitwiseANDExpression_Await is called when production bitwiseANDExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_Await(ctx *BitwiseANDExpression_AwaitContext) {
}

// EnterBitwiseANDExpression_In_Await is called when production bitwiseANDExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_In_Await(ctx *BitwiseANDExpression_In_AwaitContext) {
}

// ExitBitwiseANDExpression_In_Await is called when production bitwiseANDExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_In_Await(ctx *BitwiseANDExpression_In_AwaitContext) {
}

// EnterBitwiseANDExpression_Yield_Await is called when production bitwiseANDExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_Yield_Await(ctx *BitwiseANDExpression_Yield_AwaitContext) {
}

// ExitBitwiseANDExpression_Yield_Await is called when production bitwiseANDExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_Yield_Await(ctx *BitwiseANDExpression_Yield_AwaitContext) {
}

// EnterBitwiseANDExpression_In_Yield_Await is called when production bitwiseANDExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseANDExpression_In_Yield_Await(ctx *BitwiseANDExpression_In_Yield_AwaitContext) {
}

// ExitBitwiseANDExpression_In_Yield_Await is called when production bitwiseANDExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseANDExpression_In_Yield_Await(ctx *BitwiseANDExpression_In_Yield_AwaitContext) {
}

// EnterBitwiseXORExpression is called when production bitwiseXORExpression is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression(ctx *BitwiseXORExpressionContext) {}

// ExitBitwiseXORExpression is called when production bitwiseXORExpression is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression(ctx *BitwiseXORExpressionContext) {}

// EnterBitwiseXORExpression_In is called when production bitwiseXORExpression_In is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_In(ctx *BitwiseXORExpression_InContext) {}

// ExitBitwiseXORExpression_In is called when production bitwiseXORExpression_In is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_In(ctx *BitwiseXORExpression_InContext) {}

// EnterBitwiseXORExpression_Yield is called when production bitwiseXORExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_Yield(ctx *BitwiseXORExpression_YieldContext) {
}

// ExitBitwiseXORExpression_Yield is called when production bitwiseXORExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_Yield(ctx *BitwiseXORExpression_YieldContext) {
}

// EnterBitwiseXORExpression_In_Yield is called when production bitwiseXORExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_In_Yield(ctx *BitwiseXORExpression_In_YieldContext) {
}

// ExitBitwiseXORExpression_In_Yield is called when production bitwiseXORExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_In_Yield(ctx *BitwiseXORExpression_In_YieldContext) {
}

// EnterBitwiseXORExpression_Await is called when production bitwiseXORExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_Await(ctx *BitwiseXORExpression_AwaitContext) {
}

// ExitBitwiseXORExpression_Await is called when production bitwiseXORExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_Await(ctx *BitwiseXORExpression_AwaitContext) {
}

// EnterBitwiseXORExpression_In_Await is called when production bitwiseXORExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_In_Await(ctx *BitwiseXORExpression_In_AwaitContext) {
}

// ExitBitwiseXORExpression_In_Await is called when production bitwiseXORExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_In_Await(ctx *BitwiseXORExpression_In_AwaitContext) {
}

// EnterBitwiseXORExpression_Yield_Await is called when production bitwiseXORExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_Yield_Await(ctx *BitwiseXORExpression_Yield_AwaitContext) {
}

// ExitBitwiseXORExpression_Yield_Await is called when production bitwiseXORExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_Yield_Await(ctx *BitwiseXORExpression_Yield_AwaitContext) {
}

// EnterBitwiseXORExpression_In_Yield_Await is called when production bitwiseXORExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseXORExpression_In_Yield_Await(ctx *BitwiseXORExpression_In_Yield_AwaitContext) {
}

// ExitBitwiseXORExpression_In_Yield_Await is called when production bitwiseXORExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseXORExpression_In_Yield_Await(ctx *BitwiseXORExpression_In_Yield_AwaitContext) {
}

// EnterBitwiseORExpression is called when production bitwiseORExpression is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression(ctx *BitwiseORExpressionContext) {}

// ExitBitwiseORExpression is called when production bitwiseORExpression is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression(ctx *BitwiseORExpressionContext) {}

// EnterBitwiseORExpression_In is called when production bitwiseORExpression_In is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_In(ctx *BitwiseORExpression_InContext) {}

// ExitBitwiseORExpression_In is called when production bitwiseORExpression_In is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_In(ctx *BitwiseORExpression_InContext) {}

// EnterBitwiseORExpression_Yield is called when production bitwiseORExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_Yield(ctx *BitwiseORExpression_YieldContext) {
}

// ExitBitwiseORExpression_Yield is called when production bitwiseORExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_Yield(ctx *BitwiseORExpression_YieldContext) {
}

// EnterBitwiseORExpression_In_Yield is called when production bitwiseORExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_In_Yield(ctx *BitwiseORExpression_In_YieldContext) {
}

// ExitBitwiseORExpression_In_Yield is called when production bitwiseORExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_In_Yield(ctx *BitwiseORExpression_In_YieldContext) {
}

// EnterBitwiseORExpression_Await is called when production bitwiseORExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_Await(ctx *BitwiseORExpression_AwaitContext) {
}

// ExitBitwiseORExpression_Await is called when production bitwiseORExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_Await(ctx *BitwiseORExpression_AwaitContext) {
}

// EnterBitwiseORExpression_In_Await is called when production bitwiseORExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_In_Await(ctx *BitwiseORExpression_In_AwaitContext) {
}

// ExitBitwiseORExpression_In_Await is called when production bitwiseORExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_In_Await(ctx *BitwiseORExpression_In_AwaitContext) {
}

// EnterBitwiseORExpression_Yield_Await is called when production bitwiseORExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_Yield_Await(ctx *BitwiseORExpression_Yield_AwaitContext) {
}

// ExitBitwiseORExpression_Yield_Await is called when production bitwiseORExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_Yield_Await(ctx *BitwiseORExpression_Yield_AwaitContext) {
}

// EnterBitwiseORExpression_In_Yield_Await is called when production bitwiseORExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBitwiseORExpression_In_Yield_Await(ctx *BitwiseORExpression_In_Yield_AwaitContext) {
}

// ExitBitwiseORExpression_In_Yield_Await is called when production bitwiseORExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBitwiseORExpression_In_Yield_Await(ctx *BitwiseORExpression_In_Yield_AwaitContext) {
}

// EnterLogicalANDExpression is called when production logicalANDExpression is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression(ctx *LogicalANDExpressionContext) {}

// ExitLogicalANDExpression is called when production logicalANDExpression is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression(ctx *LogicalANDExpressionContext) {}

// EnterLogicalANDExpression_In is called when production logicalANDExpression_In is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_In(ctx *LogicalANDExpression_InContext) {}

// ExitLogicalANDExpression_In is called when production logicalANDExpression_In is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_In(ctx *LogicalANDExpression_InContext) {}

// EnterLogicalANDExpression_Yield is called when production logicalANDExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_Yield(ctx *LogicalANDExpression_YieldContext) {
}

// ExitLogicalANDExpression_Yield is called when production logicalANDExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_Yield(ctx *LogicalANDExpression_YieldContext) {
}

// EnterLogicalANDExpression_In_Yield is called when production logicalANDExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_In_Yield(ctx *LogicalANDExpression_In_YieldContext) {
}

// ExitLogicalANDExpression_In_Yield is called when production logicalANDExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_In_Yield(ctx *LogicalANDExpression_In_YieldContext) {
}

// EnterLogicalANDExpression_Await is called when production logicalANDExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_Await(ctx *LogicalANDExpression_AwaitContext) {
}

// ExitLogicalANDExpression_Await is called when production logicalANDExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_Await(ctx *LogicalANDExpression_AwaitContext) {
}

// EnterLogicalANDExpression_In_Await is called when production logicalANDExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_In_Await(ctx *LogicalANDExpression_In_AwaitContext) {
}

// ExitLogicalANDExpression_In_Await is called when production logicalANDExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_In_Await(ctx *LogicalANDExpression_In_AwaitContext) {
}

// EnterLogicalANDExpression_Yield_Await is called when production logicalANDExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_Yield_Await(ctx *LogicalANDExpression_Yield_AwaitContext) {
}

// ExitLogicalANDExpression_Yield_Await is called when production logicalANDExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_Yield_Await(ctx *LogicalANDExpression_Yield_AwaitContext) {
}

// EnterLogicalANDExpression_In_Yield_Await is called when production logicalANDExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalANDExpression_In_Yield_Await(ctx *LogicalANDExpression_In_Yield_AwaitContext) {
}

// ExitLogicalANDExpression_In_Yield_Await is called when production logicalANDExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalANDExpression_In_Yield_Await(ctx *LogicalANDExpression_In_Yield_AwaitContext) {
}

// EnterLogicalORExpression is called when production logicalORExpression is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression(ctx *LogicalORExpressionContext) {}

// ExitLogicalORExpression is called when production logicalORExpression is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression(ctx *LogicalORExpressionContext) {}

// EnterLogicalORExpression_In is called when production logicalORExpression_In is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_In(ctx *LogicalORExpression_InContext) {}

// ExitLogicalORExpression_In is called when production logicalORExpression_In is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_In(ctx *LogicalORExpression_InContext) {}

// EnterLogicalORExpression_Yield is called when production logicalORExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_Yield(ctx *LogicalORExpression_YieldContext) {
}

// ExitLogicalORExpression_Yield is called when production logicalORExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_Yield(ctx *LogicalORExpression_YieldContext) {
}

// EnterLogicalORExpression_In_Yield is called when production logicalORExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_In_Yield(ctx *LogicalORExpression_In_YieldContext) {
}

// ExitLogicalORExpression_In_Yield is called when production logicalORExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_In_Yield(ctx *LogicalORExpression_In_YieldContext) {
}

// EnterLogicalORExpression_Await is called when production logicalORExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_Await(ctx *LogicalORExpression_AwaitContext) {
}

// ExitLogicalORExpression_Await is called when production logicalORExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_Await(ctx *LogicalORExpression_AwaitContext) {
}

// EnterLogicalORExpression_In_Await is called when production logicalORExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_In_Await(ctx *LogicalORExpression_In_AwaitContext) {
}

// ExitLogicalORExpression_In_Await is called when production logicalORExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_In_Await(ctx *LogicalORExpression_In_AwaitContext) {
}

// EnterLogicalORExpression_Yield_Await is called when production logicalORExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_Yield_Await(ctx *LogicalORExpression_Yield_AwaitContext) {
}

// ExitLogicalORExpression_Yield_Await is called when production logicalORExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_Yield_Await(ctx *LogicalORExpression_Yield_AwaitContext) {
}

// EnterLogicalORExpression_In_Yield_Await is called when production logicalORExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLogicalORExpression_In_Yield_Await(ctx *LogicalORExpression_In_Yield_AwaitContext) {
}

// ExitLogicalORExpression_In_Yield_Await is called when production logicalORExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLogicalORExpression_In_Yield_Await(ctx *LogicalORExpression_In_Yield_AwaitContext) {
}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterConditionalExpression_In is called when production conditionalExpression_In is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_In(ctx *ConditionalExpression_InContext) {}

// ExitConditionalExpression_In is called when production conditionalExpression_In is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_In(ctx *ConditionalExpression_InContext) {}

// EnterConditionalExpression_Yield is called when production conditionalExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_Yield(ctx *ConditionalExpression_YieldContext) {
}

// ExitConditionalExpression_Yield is called when production conditionalExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_Yield(ctx *ConditionalExpression_YieldContext) {
}

// EnterConditionalExpression_In_Yield is called when production conditionalExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_In_Yield(ctx *ConditionalExpression_In_YieldContext) {
}

// ExitConditionalExpression_In_Yield is called when production conditionalExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_In_Yield(ctx *ConditionalExpression_In_YieldContext) {
}

// EnterConditionalExpression_Await is called when production conditionalExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_Await(ctx *ConditionalExpression_AwaitContext) {
}

// ExitConditionalExpression_Await is called when production conditionalExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_Await(ctx *ConditionalExpression_AwaitContext) {
}

// EnterConditionalExpression_In_Await is called when production conditionalExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_In_Await(ctx *ConditionalExpression_In_AwaitContext) {
}

// ExitConditionalExpression_In_Await is called when production conditionalExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_In_Await(ctx *ConditionalExpression_In_AwaitContext) {
}

// EnterConditionalExpression_Yield_Await is called when production conditionalExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_Yield_Await(ctx *ConditionalExpression_Yield_AwaitContext) {
}

// ExitConditionalExpression_Yield_Await is called when production conditionalExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_Yield_Await(ctx *ConditionalExpression_Yield_AwaitContext) {
}

// EnterConditionalExpression_In_Yield_Await is called when production conditionalExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterConditionalExpression_In_Yield_Await(ctx *ConditionalExpression_In_Yield_AwaitContext) {
}

// ExitConditionalExpression_In_Yield_Await is called when production conditionalExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitConditionalExpression_In_Yield_Await(ctx *ConditionalExpression_In_Yield_AwaitContext) {
}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseECMAScriptListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseECMAScriptListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentExpression_In is called when production assignmentExpression_In is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_In(ctx *AssignmentExpression_InContext) {}

// ExitAssignmentExpression_In is called when production assignmentExpression_In is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_In(ctx *AssignmentExpression_InContext) {}

// EnterAssignmentExpression_Yield is called when production assignmentExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_Yield(ctx *AssignmentExpression_YieldContext) {
}

// ExitAssignmentExpression_Yield is called when production assignmentExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_Yield(ctx *AssignmentExpression_YieldContext) {
}

// EnterAssignmentExpression_In_Yield is called when production assignmentExpression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_In_Yield(ctx *AssignmentExpression_In_YieldContext) {
}

// ExitAssignmentExpression_In_Yield is called when production assignmentExpression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_In_Yield(ctx *AssignmentExpression_In_YieldContext) {
}

// EnterAssignmentExpression_Await is called when production assignmentExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_Await(ctx *AssignmentExpression_AwaitContext) {
}

// ExitAssignmentExpression_Await is called when production assignmentExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_Await(ctx *AssignmentExpression_AwaitContext) {
}

// EnterAssignmentExpression_In_Await is called when production assignmentExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_In_Await(ctx *AssignmentExpression_In_AwaitContext) {
}

// ExitAssignmentExpression_In_Await is called when production assignmentExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_In_Await(ctx *AssignmentExpression_In_AwaitContext) {
}

// EnterAssignmentExpression_Yield_Await is called when production assignmentExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_Yield_Await(ctx *AssignmentExpression_Yield_AwaitContext) {
}

// ExitAssignmentExpression_Yield_Await is called when production assignmentExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_Yield_Await(ctx *AssignmentExpression_Yield_AwaitContext) {
}

// EnterAssignmentExpression_In_Yield_Await is called when production assignmentExpression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression_In_Yield_Await(ctx *AssignmentExpression_In_Yield_AwaitContext) {
}

// ExitAssignmentExpression_In_Yield_Await is called when production assignmentExpression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression_In_Yield_Await(ctx *AssignmentExpression_In_Yield_AwaitContext) {
}

// EnterExpression is called when production expression is entered.
func (s *BaseECMAScriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseECMAScriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpression_In is called when production expression_In is entered.
func (s *BaseECMAScriptListener) EnterExpression_In(ctx *Expression_InContext) {}

// ExitExpression_In is called when production expression_In is exited.
func (s *BaseECMAScriptListener) ExitExpression_In(ctx *Expression_InContext) {}

// EnterExpression_Yield is called when production expression_Yield is entered.
func (s *BaseECMAScriptListener) EnterExpression_Yield(ctx *Expression_YieldContext) {}

// ExitExpression_Yield is called when production expression_Yield is exited.
func (s *BaseECMAScriptListener) ExitExpression_Yield(ctx *Expression_YieldContext) {}

// EnterExpression_In_Yield is called when production expression_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterExpression_In_Yield(ctx *Expression_In_YieldContext) {}

// ExitExpression_In_Yield is called when production expression_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitExpression_In_Yield(ctx *Expression_In_YieldContext) {}

// EnterExpression_Await is called when production expression_Await is entered.
func (s *BaseECMAScriptListener) EnterExpression_Await(ctx *Expression_AwaitContext) {}

// ExitExpression_Await is called when production expression_Await is exited.
func (s *BaseECMAScriptListener) ExitExpression_Await(ctx *Expression_AwaitContext) {}

// EnterExpression_In_Await is called when production expression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterExpression_In_Await(ctx *Expression_In_AwaitContext) {}

// ExitExpression_In_Await is called when production expression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitExpression_In_Await(ctx *Expression_In_AwaitContext) {}

// EnterExpression_Yield_Await is called when production expression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterExpression_Yield_Await(ctx *Expression_Yield_AwaitContext) {}

// ExitExpression_Yield_Await is called when production expression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitExpression_Yield_Await(ctx *Expression_Yield_AwaitContext) {}

// EnterExpression_In_Yield_Await is called when production expression_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterExpression_In_Yield_Await(ctx *Expression_In_Yield_AwaitContext) {
}

// ExitExpression_In_Yield_Await is called when production expression_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitExpression_In_Yield_Await(ctx *Expression_In_Yield_AwaitContext) {
}

// EnterStatement is called when production statement is entered.
func (s *BaseECMAScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseECMAScriptListener) ExitStatement(ctx *StatementContext) {}

// EnterStatement_Yield is called when production statement_Yield is entered.
func (s *BaseECMAScriptListener) EnterStatement_Yield(ctx *Statement_YieldContext) {}

// ExitStatement_Yield is called when production statement_Yield is exited.
func (s *BaseECMAScriptListener) ExitStatement_Yield(ctx *Statement_YieldContext) {}

// EnterStatement_Await is called when production statement_Await is entered.
func (s *BaseECMAScriptListener) EnterStatement_Await(ctx *Statement_AwaitContext) {}

// ExitStatement_Await is called when production statement_Await is exited.
func (s *BaseECMAScriptListener) ExitStatement_Await(ctx *Statement_AwaitContext) {}

// EnterStatement_Yield_Await is called when production statement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterStatement_Yield_Await(ctx *Statement_Yield_AwaitContext) {}

// ExitStatement_Yield_Await is called when production statement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitStatement_Yield_Await(ctx *Statement_Yield_AwaitContext) {}

// EnterStatement_Return is called when production statement_Return is entered.
func (s *BaseECMAScriptListener) EnterStatement_Return(ctx *Statement_ReturnContext) {}

// ExitStatement_Return is called when production statement_Return is exited.
func (s *BaseECMAScriptListener) ExitStatement_Return(ctx *Statement_ReturnContext) {}

// EnterStatement_Yield_Return is called when production statement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterStatement_Yield_Return(ctx *Statement_Yield_ReturnContext) {}

// ExitStatement_Yield_Return is called when production statement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitStatement_Yield_Return(ctx *Statement_Yield_ReturnContext) {}

// EnterStatement_Await_Return is called when production statement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterStatement_Await_Return(ctx *Statement_Await_ReturnContext) {}

// ExitStatement_Await_Return is called when production statement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitStatement_Await_Return(ctx *Statement_Await_ReturnContext) {}

// EnterStatement_Yield_Await_Return is called when production statement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterStatement_Yield_Await_Return(ctx *Statement_Yield_Await_ReturnContext) {
}

// ExitStatement_Yield_Await_Return is called when production statement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitStatement_Yield_Await_Return(ctx *Statement_Yield_Await_ReturnContext) {
}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseECMAScriptListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseECMAScriptListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclaration_Yield is called when production declaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterDeclaration_Yield(ctx *Declaration_YieldContext) {}

// ExitDeclaration_Yield is called when production declaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitDeclaration_Yield(ctx *Declaration_YieldContext) {}

// EnterDeclaration_Await is called when production declaration_Await is entered.
func (s *BaseECMAScriptListener) EnterDeclaration_Await(ctx *Declaration_AwaitContext) {}

// ExitDeclaration_Await is called when production declaration_Await is exited.
func (s *BaseECMAScriptListener) ExitDeclaration_Await(ctx *Declaration_AwaitContext) {}

// EnterDeclaration_Yield_Await is called when production declaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterDeclaration_Yield_Await(ctx *Declaration_Yield_AwaitContext) {}

// ExitDeclaration_Yield_Await is called when production declaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitDeclaration_Yield_Await(ctx *Declaration_Yield_AwaitContext) {}

// EnterHoistableDeclaration is called when production hoistableDeclaration is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration(ctx *HoistableDeclarationContext) {}

// ExitHoistableDeclaration is called when production hoistableDeclaration is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration(ctx *HoistableDeclarationContext) {}

// EnterHoistableDeclaration_Yield is called when production hoistableDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Yield(ctx *HoistableDeclaration_YieldContext) {
}

// ExitHoistableDeclaration_Yield is called when production hoistableDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Yield(ctx *HoistableDeclaration_YieldContext) {
}

// EnterHoistableDeclaration_Await is called when production hoistableDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Await(ctx *HoistableDeclaration_AwaitContext) {
}

// ExitHoistableDeclaration_Await is called when production hoistableDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Await(ctx *HoistableDeclaration_AwaitContext) {
}

// EnterHoistableDeclaration_Yield_Await is called when production hoistableDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Yield_Await(ctx *HoistableDeclaration_Yield_AwaitContext) {
}

// ExitHoistableDeclaration_Yield_Await is called when production hoistableDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Yield_Await(ctx *HoistableDeclaration_Yield_AwaitContext) {
}

// EnterHoistableDeclaration_Default is called when production hoistableDeclaration_Default is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Default(ctx *HoistableDeclaration_DefaultContext) {
}

// ExitHoistableDeclaration_Default is called when production hoistableDeclaration_Default is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Default(ctx *HoistableDeclaration_DefaultContext) {
}

// EnterHoistableDeclaration_Yield_Default is called when production hoistableDeclaration_Yield_Default is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Yield_Default(ctx *HoistableDeclaration_Yield_DefaultContext) {
}

// ExitHoistableDeclaration_Yield_Default is called when production hoistableDeclaration_Yield_Default is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Yield_Default(ctx *HoistableDeclaration_Yield_DefaultContext) {
}

// EnterHoistableDeclaration_Await_Default is called when production hoistableDeclaration_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Await_Default(ctx *HoistableDeclaration_Await_DefaultContext) {
}

// ExitHoistableDeclaration_Await_Default is called when production hoistableDeclaration_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Await_Default(ctx *HoistableDeclaration_Await_DefaultContext) {
}

// EnterHoistableDeclaration_Yield_Await_Default is called when production hoistableDeclaration_Yield_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterHoistableDeclaration_Yield_Await_Default(ctx *HoistableDeclaration_Yield_Await_DefaultContext) {
}

// ExitHoistableDeclaration_Yield_Await_Default is called when production hoistableDeclaration_Yield_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitHoistableDeclaration_Yield_Await_Default(ctx *HoistableDeclaration_Yield_Await_DefaultContext) {
}

// EnterBreakableStatement is called when production breakableStatement is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement(ctx *BreakableStatementContext) {}

// ExitBreakableStatement is called when production breakableStatement is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement(ctx *BreakableStatementContext) {}

// EnterBreakableStatement_Yield is called when production breakableStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Yield(ctx *BreakableStatement_YieldContext) {}

// ExitBreakableStatement_Yield is called when production breakableStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Yield(ctx *BreakableStatement_YieldContext) {}

// EnterBreakableStatement_Await is called when production breakableStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Await(ctx *BreakableStatement_AwaitContext) {}

// ExitBreakableStatement_Await is called when production breakableStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Await(ctx *BreakableStatement_AwaitContext) {}

// EnterBreakableStatement_Yield_Await is called when production breakableStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Yield_Await(ctx *BreakableStatement_Yield_AwaitContext) {
}

// ExitBreakableStatement_Yield_Await is called when production breakableStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Yield_Await(ctx *BreakableStatement_Yield_AwaitContext) {
}

// EnterBreakableStatement_Return is called when production breakableStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Return(ctx *BreakableStatement_ReturnContext) {
}

// ExitBreakableStatement_Return is called when production breakableStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Return(ctx *BreakableStatement_ReturnContext) {
}

// EnterBreakableStatement_Yield_Return is called when production breakableStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Yield_Return(ctx *BreakableStatement_Yield_ReturnContext) {
}

// ExitBreakableStatement_Yield_Return is called when production breakableStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Yield_Return(ctx *BreakableStatement_Yield_ReturnContext) {
}

// EnterBreakableStatement_Await_Return is called when production breakableStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Await_Return(ctx *BreakableStatement_Await_ReturnContext) {
}

// ExitBreakableStatement_Await_Return is called when production breakableStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Await_Return(ctx *BreakableStatement_Await_ReturnContext) {
}

// EnterBreakableStatement_Yield_Await_Return is called when production breakableStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterBreakableStatement_Yield_Await_Return(ctx *BreakableStatement_Yield_Await_ReturnContext) {
}

// ExitBreakableStatement_Yield_Await_Return is called when production breakableStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitBreakableStatement_Yield_Await_Return(ctx *BreakableStatement_Yield_Await_ReturnContext) {
}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterBlockStatement_Yield is called when production blockStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Yield(ctx *BlockStatement_YieldContext) {}

// ExitBlockStatement_Yield is called when production blockStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Yield(ctx *BlockStatement_YieldContext) {}

// EnterBlockStatement_Await is called when production blockStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Await(ctx *BlockStatement_AwaitContext) {}

// ExitBlockStatement_Await is called when production blockStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Await(ctx *BlockStatement_AwaitContext) {}

// EnterBlockStatement_Yield_Await is called when production blockStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Yield_Await(ctx *BlockStatement_Yield_AwaitContext) {
}

// ExitBlockStatement_Yield_Await is called when production blockStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Yield_Await(ctx *BlockStatement_Yield_AwaitContext) {
}

// EnterBlockStatement_Return is called when production blockStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Return(ctx *BlockStatement_ReturnContext) {}

// ExitBlockStatement_Return is called when production blockStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Return(ctx *BlockStatement_ReturnContext) {}

// EnterBlockStatement_Yield_Return is called when production blockStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Yield_Return(ctx *BlockStatement_Yield_ReturnContext) {
}

// ExitBlockStatement_Yield_Return is called when production blockStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Yield_Return(ctx *BlockStatement_Yield_ReturnContext) {
}

// EnterBlockStatement_Await_Return is called when production blockStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Await_Return(ctx *BlockStatement_Await_ReturnContext) {
}

// ExitBlockStatement_Await_Return is called when production blockStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Await_Return(ctx *BlockStatement_Await_ReturnContext) {
}

// EnterBlockStatement_Yield_Await_Return is called when production blockStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterBlockStatement_Yield_Await_Return(ctx *BlockStatement_Yield_Await_ReturnContext) {
}

// ExitBlockStatement_Yield_Await_Return is called when production blockStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitBlockStatement_Yield_Await_Return(ctx *BlockStatement_Yield_Await_ReturnContext) {
}

// EnterBlock is called when production block is entered.
func (s *BaseECMAScriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseECMAScriptListener) ExitBlock(ctx *BlockContext) {}

// EnterBlock_Yield is called when production block_Yield is entered.
func (s *BaseECMAScriptListener) EnterBlock_Yield(ctx *Block_YieldContext) {}

// ExitBlock_Yield is called when production block_Yield is exited.
func (s *BaseECMAScriptListener) ExitBlock_Yield(ctx *Block_YieldContext) {}

// EnterBlock_Await is called when production block_Await is entered.
func (s *BaseECMAScriptListener) EnterBlock_Await(ctx *Block_AwaitContext) {}

// ExitBlock_Await is called when production block_Await is exited.
func (s *BaseECMAScriptListener) ExitBlock_Await(ctx *Block_AwaitContext) {}

// EnterBlock_Yield_Await is called when production block_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBlock_Yield_Await(ctx *Block_Yield_AwaitContext) {}

// ExitBlock_Yield_Await is called when production block_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBlock_Yield_Await(ctx *Block_Yield_AwaitContext) {}

// EnterBlock_Return is called when production block_Return is entered.
func (s *BaseECMAScriptListener) EnterBlock_Return(ctx *Block_ReturnContext) {}

// ExitBlock_Return is called when production block_Return is exited.
func (s *BaseECMAScriptListener) ExitBlock_Return(ctx *Block_ReturnContext) {}

// EnterBlock_Yield_Return is called when production block_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterBlock_Yield_Return(ctx *Block_Yield_ReturnContext) {}

// ExitBlock_Yield_Return is called when production block_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitBlock_Yield_Return(ctx *Block_Yield_ReturnContext) {}

// EnterBlock_Await_Return is called when production block_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterBlock_Await_Return(ctx *Block_Await_ReturnContext) {}

// ExitBlock_Await_Return is called when production block_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitBlock_Await_Return(ctx *Block_Await_ReturnContext) {}

// EnterBlock_Yield_Await_Return is called when production block_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterBlock_Yield_Await_Return(ctx *Block_Yield_Await_ReturnContext) {}

// ExitBlock_Yield_Await_Return is called when production block_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitBlock_Yield_Await_Return(ctx *Block_Yield_Await_ReturnContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseECMAScriptListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseECMAScriptListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatementList_Yield is called when production statementList_Yield is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Yield(ctx *StatementList_YieldContext) {}

// ExitStatementList_Yield is called when production statementList_Yield is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Yield(ctx *StatementList_YieldContext) {}

// EnterStatementList_Await is called when production statementList_Await is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Await(ctx *StatementList_AwaitContext) {}

// ExitStatementList_Await is called when production statementList_Await is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Await(ctx *StatementList_AwaitContext) {}

// EnterStatementList_Yield_Await is called when production statementList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Yield_Await(ctx *StatementList_Yield_AwaitContext) {
}

// ExitStatementList_Yield_Await is called when production statementList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Yield_Await(ctx *StatementList_Yield_AwaitContext) {
}

// EnterStatementList_Return is called when production statementList_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Return(ctx *StatementList_ReturnContext) {}

// ExitStatementList_Return is called when production statementList_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Return(ctx *StatementList_ReturnContext) {}

// EnterStatementList_Yield_Return is called when production statementList_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Yield_Return(ctx *StatementList_Yield_ReturnContext) {
}

// ExitStatementList_Yield_Return is called when production statementList_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Yield_Return(ctx *StatementList_Yield_ReturnContext) {
}

// EnterStatementList_Await_Return is called when production statementList_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Await_Return(ctx *StatementList_Await_ReturnContext) {
}

// ExitStatementList_Await_Return is called when production statementList_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Await_Return(ctx *StatementList_Await_ReturnContext) {
}

// EnterStatementList_Yield_Await_Return is called when production statementList_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementList_Yield_Await_Return(ctx *StatementList_Yield_Await_ReturnContext) {
}

// ExitStatementList_Yield_Await_Return is called when production statementList_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementList_Yield_Await_Return(ctx *StatementList_Yield_Await_ReturnContext) {
}

// EnterStatementListItem is called when production statementListItem is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem(ctx *StatementListItemContext) {}

// ExitStatementListItem is called when production statementListItem is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem(ctx *StatementListItemContext) {}

// EnterStatementListItem_Yield is called when production statementListItem_Yield is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Yield(ctx *StatementListItem_YieldContext) {}

// ExitStatementListItem_Yield is called when production statementListItem_Yield is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Yield(ctx *StatementListItem_YieldContext) {}

// EnterStatementListItem_Await is called when production statementListItem_Await is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Await(ctx *StatementListItem_AwaitContext) {}

// ExitStatementListItem_Await is called when production statementListItem_Await is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Await(ctx *StatementListItem_AwaitContext) {}

// EnterStatementListItem_Yield_Await is called when production statementListItem_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Yield_Await(ctx *StatementListItem_Yield_AwaitContext) {
}

// ExitStatementListItem_Yield_Await is called when production statementListItem_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Yield_Await(ctx *StatementListItem_Yield_AwaitContext) {
}

// EnterStatementListItem_Return is called when production statementListItem_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Return(ctx *StatementListItem_ReturnContext) {}

// ExitStatementListItem_Return is called when production statementListItem_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Return(ctx *StatementListItem_ReturnContext) {}

// EnterStatementListItem_Yield_Return is called when production statementListItem_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Yield_Return(ctx *StatementListItem_Yield_ReturnContext) {
}

// ExitStatementListItem_Yield_Return is called when production statementListItem_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Yield_Return(ctx *StatementListItem_Yield_ReturnContext) {
}

// EnterStatementListItem_Await_Return is called when production statementListItem_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Await_Return(ctx *StatementListItem_Await_ReturnContext) {
}

// ExitStatementListItem_Await_Return is called when production statementListItem_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Await_Return(ctx *StatementListItem_Await_ReturnContext) {
}

// EnterStatementListItem_Yield_Await_Return is called when production statementListItem_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterStatementListItem_Yield_Await_Return(ctx *StatementListItem_Yield_Await_ReturnContext) {
}

// ExitStatementListItem_Yield_Await_Return is called when production statementListItem_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitStatementListItem_Yield_Await_Return(ctx *StatementListItem_Yield_Await_ReturnContext) {
}

// EnterLexicalDeclaration is called when production lexicalDeclaration is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration(ctx *LexicalDeclarationContext) {}

// ExitLexicalDeclaration is called when production lexicalDeclaration is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration(ctx *LexicalDeclarationContext) {}

// EnterLexicalDeclaration_In is called when production lexicalDeclaration_In is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_In(ctx *LexicalDeclaration_InContext) {}

// ExitLexicalDeclaration_In is called when production lexicalDeclaration_In is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_In(ctx *LexicalDeclaration_InContext) {}

// EnterLexicalDeclaration_Yield is called when production lexicalDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_Yield(ctx *LexicalDeclaration_YieldContext) {}

// ExitLexicalDeclaration_Yield is called when production lexicalDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_Yield(ctx *LexicalDeclaration_YieldContext) {}

// EnterLexicalDeclaration_In_Yield is called when production lexicalDeclaration_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_In_Yield(ctx *LexicalDeclaration_In_YieldContext) {
}

// ExitLexicalDeclaration_In_Yield is called when production lexicalDeclaration_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_In_Yield(ctx *LexicalDeclaration_In_YieldContext) {
}

// EnterLexicalDeclaration_Await is called when production lexicalDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_Await(ctx *LexicalDeclaration_AwaitContext) {}

// ExitLexicalDeclaration_Await is called when production lexicalDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_Await(ctx *LexicalDeclaration_AwaitContext) {}

// EnterLexicalDeclaration_In_Await is called when production lexicalDeclaration_In_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_In_Await(ctx *LexicalDeclaration_In_AwaitContext) {
}

// ExitLexicalDeclaration_In_Await is called when production lexicalDeclaration_In_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_In_Await(ctx *LexicalDeclaration_In_AwaitContext) {
}

// EnterLexicalDeclaration_Yield_Await is called when production lexicalDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_Yield_Await(ctx *LexicalDeclaration_Yield_AwaitContext) {
}

// ExitLexicalDeclaration_Yield_Await is called when production lexicalDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_Yield_Await(ctx *LexicalDeclaration_Yield_AwaitContext) {
}

// EnterLexicalDeclaration_In_Yield_Await is called when production lexicalDeclaration_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalDeclaration_In_Yield_Await(ctx *LexicalDeclaration_In_Yield_AwaitContext) {
}

// ExitLexicalDeclaration_In_Yield_Await is called when production lexicalDeclaration_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalDeclaration_In_Yield_Await(ctx *LexicalDeclaration_In_Yield_AwaitContext) {
}

// EnterLetOrConst is called when production letOrConst is entered.
func (s *BaseECMAScriptListener) EnterLetOrConst(ctx *LetOrConstContext) {}

// ExitLetOrConst is called when production letOrConst is exited.
func (s *BaseECMAScriptListener) ExitLetOrConst(ctx *LetOrConstContext) {}

// EnterBindingList is called when production bindingList is entered.
func (s *BaseECMAScriptListener) EnterBindingList(ctx *BindingListContext) {}

// ExitBindingList is called when production bindingList is exited.
func (s *BaseECMAScriptListener) ExitBindingList(ctx *BindingListContext) {}

// EnterBindingList_In is called when production bindingList_In is entered.
func (s *BaseECMAScriptListener) EnterBindingList_In(ctx *BindingList_InContext) {}

// ExitBindingList_In is called when production bindingList_In is exited.
func (s *BaseECMAScriptListener) ExitBindingList_In(ctx *BindingList_InContext) {}

// EnterBindingList_Yield is called when production bindingList_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingList_Yield(ctx *BindingList_YieldContext) {}

// ExitBindingList_Yield is called when production bindingList_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingList_Yield(ctx *BindingList_YieldContext) {}

// EnterBindingList_In_Yield is called when production bindingList_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingList_In_Yield(ctx *BindingList_In_YieldContext) {}

// ExitBindingList_In_Yield is called when production bindingList_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingList_In_Yield(ctx *BindingList_In_YieldContext) {}

// EnterBindingList_Await is called when production bindingList_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingList_Await(ctx *BindingList_AwaitContext) {}

// ExitBindingList_Await is called when production bindingList_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingList_Await(ctx *BindingList_AwaitContext) {}

// EnterBindingList_In_Await is called when production bindingList_In_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingList_In_Await(ctx *BindingList_In_AwaitContext) {}

// ExitBindingList_In_Await is called when production bindingList_In_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingList_In_Await(ctx *BindingList_In_AwaitContext) {}

// EnterBindingList_Yield_Await is called when production bindingList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingList_Yield_Await(ctx *BindingList_Yield_AwaitContext) {}

// ExitBindingList_Yield_Await is called when production bindingList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingList_Yield_Await(ctx *BindingList_Yield_AwaitContext) {}

// EnterBindingList_In_Yield_Await is called when production bindingList_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingList_In_Yield_Await(ctx *BindingList_In_Yield_AwaitContext) {
}

// ExitBindingList_In_Yield_Await is called when production bindingList_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingList_In_Yield_Await(ctx *BindingList_In_Yield_AwaitContext) {
}

// EnterLexicalBinding is called when production lexicalBinding is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding(ctx *LexicalBindingContext) {}

// ExitLexicalBinding is called when production lexicalBinding is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding(ctx *LexicalBindingContext) {}

// EnterLexicalBinding_In is called when production lexicalBinding_In is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_In(ctx *LexicalBinding_InContext) {}

// ExitLexicalBinding_In is called when production lexicalBinding_In is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_In(ctx *LexicalBinding_InContext) {}

// EnterLexicalBinding_Yield is called when production lexicalBinding_Yield is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_Yield(ctx *LexicalBinding_YieldContext) {}

// ExitLexicalBinding_Yield is called when production lexicalBinding_Yield is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_Yield(ctx *LexicalBinding_YieldContext) {}

// EnterLexicalBinding_In_Yield is called when production lexicalBinding_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_In_Yield(ctx *LexicalBinding_In_YieldContext) {}

// ExitLexicalBinding_In_Yield is called when production lexicalBinding_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_In_Yield(ctx *LexicalBinding_In_YieldContext) {}

// EnterLexicalBinding_Await is called when production lexicalBinding_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_Await(ctx *LexicalBinding_AwaitContext) {}

// ExitLexicalBinding_Await is called when production lexicalBinding_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_Await(ctx *LexicalBinding_AwaitContext) {}

// EnterLexicalBinding_In_Await is called when production lexicalBinding_In_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_In_Await(ctx *LexicalBinding_In_AwaitContext) {}

// ExitLexicalBinding_In_Await is called when production lexicalBinding_In_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_In_Await(ctx *LexicalBinding_In_AwaitContext) {}

// EnterLexicalBinding_Yield_Await is called when production lexicalBinding_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_Yield_Await(ctx *LexicalBinding_Yield_AwaitContext) {
}

// ExitLexicalBinding_Yield_Await is called when production lexicalBinding_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_Yield_Await(ctx *LexicalBinding_Yield_AwaitContext) {
}

// EnterLexicalBinding_In_Yield_Await is called when production lexicalBinding_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLexicalBinding_In_Yield_Await(ctx *LexicalBinding_In_Yield_AwaitContext) {
}

// ExitLexicalBinding_In_Yield_Await is called when production lexicalBinding_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLexicalBinding_In_Yield_Await(ctx *LexicalBinding_In_Yield_AwaitContext) {
}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseECMAScriptListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseECMAScriptListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterVariableStatement_Yield is called when production variableStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterVariableStatement_Yield(ctx *VariableStatement_YieldContext) {}

// ExitVariableStatement_Yield is called when production variableStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitVariableStatement_Yield(ctx *VariableStatement_YieldContext) {}

// EnterVariableStatement_Await is called when production variableStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableStatement_Await(ctx *VariableStatement_AwaitContext) {}

// ExitVariableStatement_Await is called when production variableStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableStatement_Await(ctx *VariableStatement_AwaitContext) {}

// EnterVariableStatement_Yield_Await is called when production variableStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableStatement_Yield_Await(ctx *VariableStatement_Yield_AwaitContext) {
}

// ExitVariableStatement_Yield_Await is called when production variableStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableStatement_Yield_Await(ctx *VariableStatement_Yield_AwaitContext) {
}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// EnterVariableDeclarationList_In is called when production variableDeclarationList_In is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_In(ctx *VariableDeclarationList_InContext) {
}

// ExitVariableDeclarationList_In is called when production variableDeclarationList_In is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_In(ctx *VariableDeclarationList_InContext) {
}

// EnterVariableDeclarationList_Yield is called when production variableDeclarationList_Yield is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_Yield(ctx *VariableDeclarationList_YieldContext) {
}

// ExitVariableDeclarationList_Yield is called when production variableDeclarationList_Yield is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_Yield(ctx *VariableDeclarationList_YieldContext) {
}

// EnterVariableDeclarationList_In_Yield is called when production variableDeclarationList_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_In_Yield(ctx *VariableDeclarationList_In_YieldContext) {
}

// ExitVariableDeclarationList_In_Yield is called when production variableDeclarationList_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_In_Yield(ctx *VariableDeclarationList_In_YieldContext) {
}

// EnterVariableDeclarationList_Await is called when production variableDeclarationList_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_Await(ctx *VariableDeclarationList_AwaitContext) {
}

// ExitVariableDeclarationList_Await is called when production variableDeclarationList_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_Await(ctx *VariableDeclarationList_AwaitContext) {
}

// EnterVariableDeclarationList_In_Await is called when production variableDeclarationList_In_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_In_Await(ctx *VariableDeclarationList_In_AwaitContext) {
}

// ExitVariableDeclarationList_In_Await is called when production variableDeclarationList_In_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_In_Await(ctx *VariableDeclarationList_In_AwaitContext) {
}

// EnterVariableDeclarationList_Yield_Await is called when production variableDeclarationList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_Yield_Await(ctx *VariableDeclarationList_Yield_AwaitContext) {
}

// ExitVariableDeclarationList_Yield_Await is called when production variableDeclarationList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_Yield_Await(ctx *VariableDeclarationList_Yield_AwaitContext) {
}

// EnterVariableDeclarationList_In_Yield_Await is called when production variableDeclarationList_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList_In_Yield_Await(ctx *VariableDeclarationList_In_Yield_AwaitContext) {
}

// ExitVariableDeclarationList_In_Yield_Await is called when production variableDeclarationList_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList_In_Yield_Await(ctx *VariableDeclarationList_In_Yield_AwaitContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableDeclaration_In is called when production variableDeclaration_In is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_In(ctx *VariableDeclaration_InContext) {}

// ExitVariableDeclaration_In is called when production variableDeclaration_In is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_In(ctx *VariableDeclaration_InContext) {}

// EnterVariableDeclaration_Yield is called when production variableDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_Yield(ctx *VariableDeclaration_YieldContext) {
}

// ExitVariableDeclaration_Yield is called when production variableDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_Yield(ctx *VariableDeclaration_YieldContext) {
}

// EnterVariableDeclaration_In_Yield is called when production variableDeclaration_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_In_Yield(ctx *VariableDeclaration_In_YieldContext) {
}

// ExitVariableDeclaration_In_Yield is called when production variableDeclaration_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_In_Yield(ctx *VariableDeclaration_In_YieldContext) {
}

// EnterVariableDeclaration_Await is called when production variableDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_Await(ctx *VariableDeclaration_AwaitContext) {
}

// ExitVariableDeclaration_Await is called when production variableDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_Await(ctx *VariableDeclaration_AwaitContext) {
}

// EnterVariableDeclaration_In_Await is called when production variableDeclaration_In_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_In_Await(ctx *VariableDeclaration_In_AwaitContext) {
}

// ExitVariableDeclaration_In_Await is called when production variableDeclaration_In_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_In_Await(ctx *VariableDeclaration_In_AwaitContext) {
}

// EnterVariableDeclaration_Yield_Await is called when production variableDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_Yield_Await(ctx *VariableDeclaration_Yield_AwaitContext) {
}

// ExitVariableDeclaration_Yield_Await is called when production variableDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_Yield_Await(ctx *VariableDeclaration_Yield_AwaitContext) {
}

// EnterVariableDeclaration_In_Yield_Await is called when production variableDeclaration_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration_In_Yield_Await(ctx *VariableDeclaration_In_Yield_AwaitContext) {
}

// ExitVariableDeclaration_In_Yield_Await is called when production variableDeclaration_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration_In_Yield_Await(ctx *VariableDeclaration_In_Yield_AwaitContext) {
}

// EnterBindingPattern is called when production bindingPattern is entered.
func (s *BaseECMAScriptListener) EnterBindingPattern(ctx *BindingPatternContext) {}

// ExitBindingPattern is called when production bindingPattern is exited.
func (s *BaseECMAScriptListener) ExitBindingPattern(ctx *BindingPatternContext) {}

// EnterBindingPattern_Yield is called when production bindingPattern_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingPattern_Yield(ctx *BindingPattern_YieldContext) {}

// ExitBindingPattern_Yield is called when production bindingPattern_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingPattern_Yield(ctx *BindingPattern_YieldContext) {}

// EnterBindingPattern_Await is called when production bindingPattern_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingPattern_Await(ctx *BindingPattern_AwaitContext) {}

// ExitBindingPattern_Await is called when production bindingPattern_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingPattern_Await(ctx *BindingPattern_AwaitContext) {}

// EnterBindingPattern_Yield_Await is called when production bindingPattern_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingPattern_Yield_Await(ctx *BindingPattern_Yield_AwaitContext) {
}

// ExitBindingPattern_Yield_Await is called when production bindingPattern_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingPattern_Yield_Await(ctx *BindingPattern_Yield_AwaitContext) {
}

// EnterObjectBindingPattern is called when production objectBindingPattern is entered.
func (s *BaseECMAScriptListener) EnterObjectBindingPattern(ctx *ObjectBindingPatternContext) {}

// ExitObjectBindingPattern is called when production objectBindingPattern is exited.
func (s *BaseECMAScriptListener) ExitObjectBindingPattern(ctx *ObjectBindingPatternContext) {}

// EnterObjectBindingPattern_Yield is called when production objectBindingPattern_Yield is entered.
func (s *BaseECMAScriptListener) EnterObjectBindingPattern_Yield(ctx *ObjectBindingPattern_YieldContext) {
}

// ExitObjectBindingPattern_Yield is called when production objectBindingPattern_Yield is exited.
func (s *BaseECMAScriptListener) ExitObjectBindingPattern_Yield(ctx *ObjectBindingPattern_YieldContext) {
}

// EnterObjectBindingPattern_Await is called when production objectBindingPattern_Await is entered.
func (s *BaseECMAScriptListener) EnterObjectBindingPattern_Await(ctx *ObjectBindingPattern_AwaitContext) {
}

// ExitObjectBindingPattern_Await is called when production objectBindingPattern_Await is exited.
func (s *BaseECMAScriptListener) ExitObjectBindingPattern_Await(ctx *ObjectBindingPattern_AwaitContext) {
}

// EnterObjectBindingPattern_Yield_Await is called when production objectBindingPattern_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterObjectBindingPattern_Yield_Await(ctx *ObjectBindingPattern_Yield_AwaitContext) {
}

// ExitObjectBindingPattern_Yield_Await is called when production objectBindingPattern_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitObjectBindingPattern_Yield_Await(ctx *ObjectBindingPattern_Yield_AwaitContext) {
}

// EnterArrayBindingPattern is called when production arrayBindingPattern is entered.
func (s *BaseECMAScriptListener) EnterArrayBindingPattern(ctx *ArrayBindingPatternContext) {}

// ExitArrayBindingPattern is called when production arrayBindingPattern is exited.
func (s *BaseECMAScriptListener) ExitArrayBindingPattern(ctx *ArrayBindingPatternContext) {}

// EnterArrayBindingPattern_Yield is called when production arrayBindingPattern_Yield is entered.
func (s *BaseECMAScriptListener) EnterArrayBindingPattern_Yield(ctx *ArrayBindingPattern_YieldContext) {
}

// ExitArrayBindingPattern_Yield is called when production arrayBindingPattern_Yield is exited.
func (s *BaseECMAScriptListener) ExitArrayBindingPattern_Yield(ctx *ArrayBindingPattern_YieldContext) {
}

// EnterArrayBindingPattern_Await is called when production arrayBindingPattern_Await is entered.
func (s *BaseECMAScriptListener) EnterArrayBindingPattern_Await(ctx *ArrayBindingPattern_AwaitContext) {
}

// ExitArrayBindingPattern_Await is called when production arrayBindingPattern_Await is exited.
func (s *BaseECMAScriptListener) ExitArrayBindingPattern_Await(ctx *ArrayBindingPattern_AwaitContext) {
}

// EnterArrayBindingPattern_Yield_Await is called when production arrayBindingPattern_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArrayBindingPattern_Yield_Await(ctx *ArrayBindingPattern_Yield_AwaitContext) {
}

// ExitArrayBindingPattern_Yield_Await is called when production arrayBindingPattern_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArrayBindingPattern_Yield_Await(ctx *ArrayBindingPattern_Yield_AwaitContext) {
}

// EnterBindingRestProperty is called when production bindingRestProperty is entered.
func (s *BaseECMAScriptListener) EnterBindingRestProperty(ctx *BindingRestPropertyContext) {}

// ExitBindingRestProperty is called when production bindingRestProperty is exited.
func (s *BaseECMAScriptListener) ExitBindingRestProperty(ctx *BindingRestPropertyContext) {}

// EnterBindingRestProperty_Yield is called when production bindingRestProperty_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingRestProperty_Yield(ctx *BindingRestProperty_YieldContext) {
}

// ExitBindingRestProperty_Yield is called when production bindingRestProperty_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingRestProperty_Yield(ctx *BindingRestProperty_YieldContext) {
}

// EnterBindingRestProperty_Await is called when production bindingRestProperty_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingRestProperty_Await(ctx *BindingRestProperty_AwaitContext) {
}

// ExitBindingRestProperty_Await is called when production bindingRestProperty_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingRestProperty_Await(ctx *BindingRestProperty_AwaitContext) {
}

// EnterBindingRestProperty_Yield_Await is called when production bindingRestProperty_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingRestProperty_Yield_Await(ctx *BindingRestProperty_Yield_AwaitContext) {
}

// ExitBindingRestProperty_Yield_Await is called when production bindingRestProperty_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingRestProperty_Yield_Await(ctx *BindingRestProperty_Yield_AwaitContext) {
}

// EnterBindingPropertyList is called when production bindingPropertyList is entered.
func (s *BaseECMAScriptListener) EnterBindingPropertyList(ctx *BindingPropertyListContext) {}

// ExitBindingPropertyList is called when production bindingPropertyList is exited.
func (s *BaseECMAScriptListener) ExitBindingPropertyList(ctx *BindingPropertyListContext) {}

// EnterBindingPropertyList_Yield is called when production bindingPropertyList_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingPropertyList_Yield(ctx *BindingPropertyList_YieldContext) {
}

// ExitBindingPropertyList_Yield is called when production bindingPropertyList_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingPropertyList_Yield(ctx *BindingPropertyList_YieldContext) {
}

// EnterBindingPropertyList_Await is called when production bindingPropertyList_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingPropertyList_Await(ctx *BindingPropertyList_AwaitContext) {
}

// ExitBindingPropertyList_Await is called when production bindingPropertyList_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingPropertyList_Await(ctx *BindingPropertyList_AwaitContext) {
}

// EnterBindingPropertyList_Yield_Await is called when production bindingPropertyList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingPropertyList_Yield_Await(ctx *BindingPropertyList_Yield_AwaitContext) {
}

// ExitBindingPropertyList_Yield_Await is called when production bindingPropertyList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingPropertyList_Yield_Await(ctx *BindingPropertyList_Yield_AwaitContext) {
}

// EnterBindingElementList is called when production bindingElementList is entered.
func (s *BaseECMAScriptListener) EnterBindingElementList(ctx *BindingElementListContext) {}

// ExitBindingElementList is called when production bindingElementList is exited.
func (s *BaseECMAScriptListener) ExitBindingElementList(ctx *BindingElementListContext) {}

// EnterBindingElementList_Yield is called when production bindingElementList_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingElementList_Yield(ctx *BindingElementList_YieldContext) {}

// ExitBindingElementList_Yield is called when production bindingElementList_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingElementList_Yield(ctx *BindingElementList_YieldContext) {}

// EnterBindingElementList_Await is called when production bindingElementList_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingElementList_Await(ctx *BindingElementList_AwaitContext) {}

// ExitBindingElementList_Await is called when production bindingElementList_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingElementList_Await(ctx *BindingElementList_AwaitContext) {}

// EnterBindingElementList_Yield_Await is called when production bindingElementList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingElementList_Yield_Await(ctx *BindingElementList_Yield_AwaitContext) {
}

// ExitBindingElementList_Yield_Await is called when production bindingElementList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingElementList_Yield_Await(ctx *BindingElementList_Yield_AwaitContext) {
}

// EnterBindingElisionElement is called when production bindingElisionElement is entered.
func (s *BaseECMAScriptListener) EnterBindingElisionElement(ctx *BindingElisionElementContext) {}

// ExitBindingElisionElement is called when production bindingElisionElement is exited.
func (s *BaseECMAScriptListener) ExitBindingElisionElement(ctx *BindingElisionElementContext) {}

// EnterBindingElisionElement_Yield is called when production bindingElisionElement_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingElisionElement_Yield(ctx *BindingElisionElement_YieldContext) {
}

// ExitBindingElisionElement_Yield is called when production bindingElisionElement_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingElisionElement_Yield(ctx *BindingElisionElement_YieldContext) {
}

// EnterBindingElisionElement_Await is called when production bindingElisionElement_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingElisionElement_Await(ctx *BindingElisionElement_AwaitContext) {
}

// ExitBindingElisionElement_Await is called when production bindingElisionElement_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingElisionElement_Await(ctx *BindingElisionElement_AwaitContext) {
}

// EnterBindingElisionElement_Yield_Await is called when production bindingElisionElement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingElisionElement_Yield_Await(ctx *BindingElisionElement_Yield_AwaitContext) {
}

// ExitBindingElisionElement_Yield_Await is called when production bindingElisionElement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingElisionElement_Yield_Await(ctx *BindingElisionElement_Yield_AwaitContext) {
}

// EnterBindingProperty is called when production bindingProperty is entered.
func (s *BaseECMAScriptListener) EnterBindingProperty(ctx *BindingPropertyContext) {}

// ExitBindingProperty is called when production bindingProperty is exited.
func (s *BaseECMAScriptListener) ExitBindingProperty(ctx *BindingPropertyContext) {}

// EnterBindingProperty_Yield is called when production bindingProperty_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingProperty_Yield(ctx *BindingProperty_YieldContext) {}

// ExitBindingProperty_Yield is called when production bindingProperty_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingProperty_Yield(ctx *BindingProperty_YieldContext) {}

// EnterBindingProperty_Await is called when production bindingProperty_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingProperty_Await(ctx *BindingProperty_AwaitContext) {}

// ExitBindingProperty_Await is called when production bindingProperty_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingProperty_Await(ctx *BindingProperty_AwaitContext) {}

// EnterBindingProperty_Yield_Await is called when production bindingProperty_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingProperty_Yield_Await(ctx *BindingProperty_Yield_AwaitContext) {
}

// ExitBindingProperty_Yield_Await is called when production bindingProperty_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingProperty_Yield_Await(ctx *BindingProperty_Yield_AwaitContext) {
}

// EnterBindingElement is called when production bindingElement is entered.
func (s *BaseECMAScriptListener) EnterBindingElement(ctx *BindingElementContext) {}

// ExitBindingElement is called when production bindingElement is exited.
func (s *BaseECMAScriptListener) ExitBindingElement(ctx *BindingElementContext) {}

// EnterBindingElement_Yield is called when production bindingElement_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingElement_Yield(ctx *BindingElement_YieldContext) {}

// ExitBindingElement_Yield is called when production bindingElement_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingElement_Yield(ctx *BindingElement_YieldContext) {}

// EnterBindingElement_Await is called when production bindingElement_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingElement_Await(ctx *BindingElement_AwaitContext) {}

// ExitBindingElement_Await is called when production bindingElement_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingElement_Await(ctx *BindingElement_AwaitContext) {}

// EnterBindingElement_Yield_Await is called when production bindingElement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingElement_Yield_Await(ctx *BindingElement_Yield_AwaitContext) {
}

// ExitBindingElement_Yield_Await is called when production bindingElement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingElement_Yield_Await(ctx *BindingElement_Yield_AwaitContext) {
}

// EnterSingleNameBinding is called when production singleNameBinding is entered.
func (s *BaseECMAScriptListener) EnterSingleNameBinding(ctx *SingleNameBindingContext) {}

// ExitSingleNameBinding is called when production singleNameBinding is exited.
func (s *BaseECMAScriptListener) ExitSingleNameBinding(ctx *SingleNameBindingContext) {}

// EnterSingleNameBinding_Yield is called when production singleNameBinding_Yield is entered.
func (s *BaseECMAScriptListener) EnterSingleNameBinding_Yield(ctx *SingleNameBinding_YieldContext) {}

// ExitSingleNameBinding_Yield is called when production singleNameBinding_Yield is exited.
func (s *BaseECMAScriptListener) ExitSingleNameBinding_Yield(ctx *SingleNameBinding_YieldContext) {}

// EnterSingleNameBinding_Await is called when production singleNameBinding_Await is entered.
func (s *BaseECMAScriptListener) EnterSingleNameBinding_Await(ctx *SingleNameBinding_AwaitContext) {}

// ExitSingleNameBinding_Await is called when production singleNameBinding_Await is exited.
func (s *BaseECMAScriptListener) ExitSingleNameBinding_Await(ctx *SingleNameBinding_AwaitContext) {}

// EnterSingleNameBinding_Yield_Await is called when production singleNameBinding_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterSingleNameBinding_Yield_Await(ctx *SingleNameBinding_Yield_AwaitContext) {
}

// ExitSingleNameBinding_Yield_Await is called when production singleNameBinding_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitSingleNameBinding_Yield_Await(ctx *SingleNameBinding_Yield_AwaitContext) {
}

// EnterBindingRestElement is called when production bindingRestElement is entered.
func (s *BaseECMAScriptListener) EnterBindingRestElement(ctx *BindingRestElementContext) {}

// ExitBindingRestElement is called when production bindingRestElement is exited.
func (s *BaseECMAScriptListener) ExitBindingRestElement(ctx *BindingRestElementContext) {}

// EnterBindingRestElement_Yield is called when production bindingRestElement_Yield is entered.
func (s *BaseECMAScriptListener) EnterBindingRestElement_Yield(ctx *BindingRestElement_YieldContext) {}

// ExitBindingRestElement_Yield is called when production bindingRestElement_Yield is exited.
func (s *BaseECMAScriptListener) ExitBindingRestElement_Yield(ctx *BindingRestElement_YieldContext) {}

// EnterBindingRestElement_Await is called when production bindingRestElement_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingRestElement_Await(ctx *BindingRestElement_AwaitContext) {}

// ExitBindingRestElement_Await is called when production bindingRestElement_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingRestElement_Await(ctx *BindingRestElement_AwaitContext) {}

// EnterBindingRestElement_Yield_Await is called when production bindingRestElement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBindingRestElement_Yield_Await(ctx *BindingRestElement_Yield_AwaitContext) {
}

// ExitBindingRestElement_Yield_Await is called when production bindingRestElement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBindingRestElement_Yield_Await(ctx *BindingRestElement_Yield_AwaitContext) {
}

// EnterTheEmptyStatement is called when production theEmptyStatement is entered.
func (s *BaseECMAScriptListener) EnterTheEmptyStatement(ctx *TheEmptyStatementContext) {}

// ExitTheEmptyStatement is called when production theEmptyStatement is exited.
func (s *BaseECMAScriptListener) ExitTheEmptyStatement(ctx *TheEmptyStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseECMAScriptListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseECMAScriptListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterExpressionStatement_Yield is called when production expressionStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterExpressionStatement_Yield(ctx *ExpressionStatement_YieldContext) {
}

// ExitExpressionStatement_Yield is called when production expressionStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitExpressionStatement_Yield(ctx *ExpressionStatement_YieldContext) {
}

// EnterExpressionStatement_Await is called when production expressionStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterExpressionStatement_Await(ctx *ExpressionStatement_AwaitContext) {
}

// ExitExpressionStatement_Await is called when production expressionStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitExpressionStatement_Await(ctx *ExpressionStatement_AwaitContext) {
}

// EnterExpressionStatement_Yield_Await is called when production expressionStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterExpressionStatement_Yield_Await(ctx *ExpressionStatement_Yield_AwaitContext) {
}

// ExitExpressionStatement_Yield_Await is called when production expressionStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitExpressionStatement_Yield_Await(ctx *ExpressionStatement_Yield_AwaitContext) {
}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseECMAScriptListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseECMAScriptListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfStatement_Yield is called when production ifStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Yield(ctx *IfStatement_YieldContext) {}

// ExitIfStatement_Yield is called when production ifStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Yield(ctx *IfStatement_YieldContext) {}

// EnterIfStatement_Await is called when production ifStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Await(ctx *IfStatement_AwaitContext) {}

// ExitIfStatement_Await is called when production ifStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Await(ctx *IfStatement_AwaitContext) {}

// EnterIfStatement_Yield_Await is called when production ifStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Yield_Await(ctx *IfStatement_Yield_AwaitContext) {}

// ExitIfStatement_Yield_Await is called when production ifStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Yield_Await(ctx *IfStatement_Yield_AwaitContext) {}

// EnterIfStatement_Return is called when production ifStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Return(ctx *IfStatement_ReturnContext) {}

// ExitIfStatement_Return is called when production ifStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Return(ctx *IfStatement_ReturnContext) {}

// EnterIfStatement_Yield_Return is called when production ifStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Yield_Return(ctx *IfStatement_Yield_ReturnContext) {}

// ExitIfStatement_Yield_Return is called when production ifStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Yield_Return(ctx *IfStatement_Yield_ReturnContext) {}

// EnterIfStatement_Await_Return is called when production ifStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Await_Return(ctx *IfStatement_Await_ReturnContext) {}

// ExitIfStatement_Await_Return is called when production ifStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Await_Return(ctx *IfStatement_Await_ReturnContext) {}

// EnterIfStatement_Yield_Await_Return is called when production ifStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterIfStatement_Yield_Await_Return(ctx *IfStatement_Yield_Await_ReturnContext) {
}

// ExitIfStatement_Yield_Await_Return is called when production ifStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitIfStatement_Yield_Await_Return(ctx *IfStatement_Yield_Await_ReturnContext) {
}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterIterationStatement_Yield is called when production iterationStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Yield(ctx *IterationStatement_YieldContext) {}

// ExitIterationStatement_Yield is called when production iterationStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Yield(ctx *IterationStatement_YieldContext) {}

// EnterIterationStatement_Await is called when production iterationStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Await(ctx *IterationStatement_AwaitContext) {}

// ExitIterationStatement_Await is called when production iterationStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Await(ctx *IterationStatement_AwaitContext) {}

// EnterIterationStatement_Yield_Await is called when production iterationStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Yield_Await(ctx *IterationStatement_Yield_AwaitContext) {
}

// ExitIterationStatement_Yield_Await is called when production iterationStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Yield_Await(ctx *IterationStatement_Yield_AwaitContext) {
}

// EnterIterationStatement_Return is called when production iterationStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Return(ctx *IterationStatement_ReturnContext) {
}

// ExitIterationStatement_Return is called when production iterationStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Return(ctx *IterationStatement_ReturnContext) {
}

// EnterIterationStatement_Yield_Return is called when production iterationStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Yield_Return(ctx *IterationStatement_Yield_ReturnContext) {
}

// ExitIterationStatement_Yield_Return is called when production iterationStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Yield_Return(ctx *IterationStatement_Yield_ReturnContext) {
}

// EnterIterationStatement_Await_Return is called when production iterationStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Await_Return(ctx *IterationStatement_Await_ReturnContext) {
}

// ExitIterationStatement_Await_Return is called when production iterationStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Await_Return(ctx *IterationStatement_Await_ReturnContext) {
}

// EnterIterationStatement_Yield_Await_Return is called when production iterationStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterIterationStatement_Yield_Await_Return(ctx *IterationStatement_Yield_Await_ReturnContext) {
}

// ExitIterationStatement_Yield_Await_Return is called when production iterationStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitIterationStatement_Yield_Await_Return(ctx *IterationStatement_Yield_Await_ReturnContext) {
}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *BaseECMAScriptListener) EnterForDeclaration(ctx *ForDeclarationContext) {}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *BaseECMAScriptListener) ExitForDeclaration(ctx *ForDeclarationContext) {}

// EnterForDeclaration_Yield is called when production forDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterForDeclaration_Yield(ctx *ForDeclaration_YieldContext) {}

// ExitForDeclaration_Yield is called when production forDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitForDeclaration_Yield(ctx *ForDeclaration_YieldContext) {}

// EnterForDeclaration_Await is called when production forDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterForDeclaration_Await(ctx *ForDeclaration_AwaitContext) {}

// ExitForDeclaration_Await is called when production forDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitForDeclaration_Await(ctx *ForDeclaration_AwaitContext) {}

// EnterForDeclaration_Yield_Await is called when production forDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterForDeclaration_Yield_Await(ctx *ForDeclaration_Yield_AwaitContext) {
}

// ExitForDeclaration_Yield_Await is called when production forDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitForDeclaration_Yield_Await(ctx *ForDeclaration_Yield_AwaitContext) {
}

// EnterForBinding is called when production forBinding is entered.
func (s *BaseECMAScriptListener) EnterForBinding(ctx *ForBindingContext) {}

// ExitForBinding is called when production forBinding is exited.
func (s *BaseECMAScriptListener) ExitForBinding(ctx *ForBindingContext) {}

// EnterForBinding_Yield is called when production forBinding_Yield is entered.
func (s *BaseECMAScriptListener) EnterForBinding_Yield(ctx *ForBinding_YieldContext) {}

// ExitForBinding_Yield is called when production forBinding_Yield is exited.
func (s *BaseECMAScriptListener) ExitForBinding_Yield(ctx *ForBinding_YieldContext) {}

// EnterForBinding_Await is called when production forBinding_Await is entered.
func (s *BaseECMAScriptListener) EnterForBinding_Await(ctx *ForBinding_AwaitContext) {}

// ExitForBinding_Await is called when production forBinding_Await is exited.
func (s *BaseECMAScriptListener) ExitForBinding_Await(ctx *ForBinding_AwaitContext) {}

// EnterForBinding_Yield_Await is called when production forBinding_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterForBinding_Yield_Await(ctx *ForBinding_Yield_AwaitContext) {}

// ExitForBinding_Yield_Await is called when production forBinding_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitForBinding_Yield_Await(ctx *ForBinding_Yield_AwaitContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseECMAScriptListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseECMAScriptListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterContinueStatement_Yield is called when production continueStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterContinueStatement_Yield(ctx *ContinueStatement_YieldContext) {}

// ExitContinueStatement_Yield is called when production continueStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitContinueStatement_Yield(ctx *ContinueStatement_YieldContext) {}

// EnterContinueStatement_Await is called when production continueStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterContinueStatement_Await(ctx *ContinueStatement_AwaitContext) {}

// ExitContinueStatement_Await is called when production continueStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitContinueStatement_Await(ctx *ContinueStatement_AwaitContext) {}

// EnterContinueStatement_Yield_Await is called when production continueStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterContinueStatement_Yield_Await(ctx *ContinueStatement_Yield_AwaitContext) {
}

// ExitContinueStatement_Yield_Await is called when production continueStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitContinueStatement_Yield_Await(ctx *ContinueStatement_Yield_AwaitContext) {
}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseECMAScriptListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseECMAScriptListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterBreakStatement_Yield is called when production breakStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterBreakStatement_Yield(ctx *BreakStatement_YieldContext) {}

// ExitBreakStatement_Yield is called when production breakStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitBreakStatement_Yield(ctx *BreakStatement_YieldContext) {}

// EnterBreakStatement_Await is called when production breakStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterBreakStatement_Await(ctx *BreakStatement_AwaitContext) {}

// ExitBreakStatement_Await is called when production breakStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitBreakStatement_Await(ctx *BreakStatement_AwaitContext) {}

// EnterBreakStatement_Yield_Await is called when production breakStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterBreakStatement_Yield_Await(ctx *BreakStatement_Yield_AwaitContext) {
}

// ExitBreakStatement_Yield_Await is called when production breakStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitBreakStatement_Yield_Await(ctx *BreakStatement_Yield_AwaitContext) {
}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseECMAScriptListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseECMAScriptListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterReturnStatement_Yield is called when production returnStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterReturnStatement_Yield(ctx *ReturnStatement_YieldContext) {}

// ExitReturnStatement_Yield is called when production returnStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitReturnStatement_Yield(ctx *ReturnStatement_YieldContext) {}

// EnterReturnStatement_Await is called when production returnStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterReturnStatement_Await(ctx *ReturnStatement_AwaitContext) {}

// ExitReturnStatement_Await is called when production returnStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitReturnStatement_Await(ctx *ReturnStatement_AwaitContext) {}

// EnterReturnStatement_Yield_Await is called when production returnStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterReturnStatement_Yield_Await(ctx *ReturnStatement_Yield_AwaitContext) {
}

// ExitReturnStatement_Yield_Await is called when production returnStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitReturnStatement_Yield_Await(ctx *ReturnStatement_Yield_AwaitContext) {
}

// EnterWithStatement is called when production withStatement is entered.
func (s *BaseECMAScriptListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BaseECMAScriptListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterWithStatement_Yield is called when production withStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Yield(ctx *WithStatement_YieldContext) {}

// ExitWithStatement_Yield is called when production withStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Yield(ctx *WithStatement_YieldContext) {}

// EnterWithStatement_Await is called when production withStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Await(ctx *WithStatement_AwaitContext) {}

// ExitWithStatement_Await is called when production withStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Await(ctx *WithStatement_AwaitContext) {}

// EnterWithStatement_Yield_Await is called when production withStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Yield_Await(ctx *WithStatement_Yield_AwaitContext) {
}

// ExitWithStatement_Yield_Await is called when production withStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Yield_Await(ctx *WithStatement_Yield_AwaitContext) {
}

// EnterWithStatement_Return is called when production withStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Return(ctx *WithStatement_ReturnContext) {}

// ExitWithStatement_Return is called when production withStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Return(ctx *WithStatement_ReturnContext) {}

// EnterWithStatement_Yield_Return is called when production withStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Yield_Return(ctx *WithStatement_Yield_ReturnContext) {
}

// ExitWithStatement_Yield_Return is called when production withStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Yield_Return(ctx *WithStatement_Yield_ReturnContext) {
}

// EnterWithStatement_Await_Return is called when production withStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Await_Return(ctx *WithStatement_Await_ReturnContext) {
}

// ExitWithStatement_Await_Return is called when production withStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Await_Return(ctx *WithStatement_Await_ReturnContext) {
}

// EnterWithStatement_Yield_Await_Return is called when production withStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterWithStatement_Yield_Await_Return(ctx *WithStatement_Yield_Await_ReturnContext) {
}

// ExitWithStatement_Yield_Await_Return is called when production withStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitWithStatement_Yield_Await_Return(ctx *WithStatement_Yield_Await_ReturnContext) {
}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchStatement_Yield is called when production switchStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Yield(ctx *SwitchStatement_YieldContext) {}

// ExitSwitchStatement_Yield is called when production switchStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Yield(ctx *SwitchStatement_YieldContext) {}

// EnterSwitchStatement_Await is called when production switchStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Await(ctx *SwitchStatement_AwaitContext) {}

// ExitSwitchStatement_Await is called when production switchStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Await(ctx *SwitchStatement_AwaitContext) {}

// EnterSwitchStatement_Yield_Await is called when production switchStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Yield_Await(ctx *SwitchStatement_Yield_AwaitContext) {
}

// ExitSwitchStatement_Yield_Await is called when production switchStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Yield_Await(ctx *SwitchStatement_Yield_AwaitContext) {
}

// EnterSwitchStatement_Return is called when production switchStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Return(ctx *SwitchStatement_ReturnContext) {}

// ExitSwitchStatement_Return is called when production switchStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Return(ctx *SwitchStatement_ReturnContext) {}

// EnterSwitchStatement_Yield_Return is called when production switchStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Yield_Return(ctx *SwitchStatement_Yield_ReturnContext) {
}

// ExitSwitchStatement_Yield_Return is called when production switchStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Yield_Return(ctx *SwitchStatement_Yield_ReturnContext) {
}

// EnterSwitchStatement_Await_Return is called when production switchStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Await_Return(ctx *SwitchStatement_Await_ReturnContext) {
}

// ExitSwitchStatement_Await_Return is called when production switchStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Await_Return(ctx *SwitchStatement_Await_ReturnContext) {
}

// EnterSwitchStatement_Yield_Await_Return is called when production switchStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement_Yield_Await_Return(ctx *SwitchStatement_Yield_Await_ReturnContext) {
}

// ExitSwitchStatement_Yield_Await_Return is called when production switchStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement_Yield_Await_Return(ctx *SwitchStatement_Yield_Await_ReturnContext) {
}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterCaseBlock_Yield is called when production caseBlock_Yield is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Yield(ctx *CaseBlock_YieldContext) {}

// ExitCaseBlock_Yield is called when production caseBlock_Yield is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Yield(ctx *CaseBlock_YieldContext) {}

// EnterCaseBlock_Await is called when production caseBlock_Await is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Await(ctx *CaseBlock_AwaitContext) {}

// ExitCaseBlock_Await is called when production caseBlock_Await is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Await(ctx *CaseBlock_AwaitContext) {}

// EnterCaseBlock_Yield_Await is called when production caseBlock_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Yield_Await(ctx *CaseBlock_Yield_AwaitContext) {}

// ExitCaseBlock_Yield_Await is called when production caseBlock_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Yield_Await(ctx *CaseBlock_Yield_AwaitContext) {}

// EnterCaseBlock_Return is called when production caseBlock_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Return(ctx *CaseBlock_ReturnContext) {}

// ExitCaseBlock_Return is called when production caseBlock_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Return(ctx *CaseBlock_ReturnContext) {}

// EnterCaseBlock_Yield_Return is called when production caseBlock_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Yield_Return(ctx *CaseBlock_Yield_ReturnContext) {}

// ExitCaseBlock_Yield_Return is called when production caseBlock_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Yield_Return(ctx *CaseBlock_Yield_ReturnContext) {}

// EnterCaseBlock_Await_Return is called when production caseBlock_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Await_Return(ctx *CaseBlock_Await_ReturnContext) {}

// ExitCaseBlock_Await_Return is called when production caseBlock_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Await_Return(ctx *CaseBlock_Await_ReturnContext) {}

// EnterCaseBlock_Yield_Await_Return is called when production caseBlock_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock_Yield_Await_Return(ctx *CaseBlock_Yield_Await_ReturnContext) {
}

// ExitCaseBlock_Yield_Await_Return is called when production caseBlock_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock_Yield_Await_Return(ctx *CaseBlock_Yield_Await_ReturnContext) {
}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseECMAScriptListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseECMAScriptListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterCaseClause_Yield is called when production caseClause_Yield is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Yield(ctx *CaseClause_YieldContext) {}

// ExitCaseClause_Yield is called when production caseClause_Yield is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Yield(ctx *CaseClause_YieldContext) {}

// EnterCaseClause_Await is called when production caseClause_Await is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Await(ctx *CaseClause_AwaitContext) {}

// ExitCaseClause_Await is called when production caseClause_Await is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Await(ctx *CaseClause_AwaitContext) {}

// EnterCaseClause_Yield_Await is called when production caseClause_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Yield_Await(ctx *CaseClause_Yield_AwaitContext) {}

// ExitCaseClause_Yield_Await is called when production caseClause_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Yield_Await(ctx *CaseClause_Yield_AwaitContext) {}

// EnterCaseClause_Return is called when production caseClause_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Return(ctx *CaseClause_ReturnContext) {}

// ExitCaseClause_Return is called when production caseClause_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Return(ctx *CaseClause_ReturnContext) {}

// EnterCaseClause_Yield_Return is called when production caseClause_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Yield_Return(ctx *CaseClause_Yield_ReturnContext) {}

// ExitCaseClause_Yield_Return is called when production caseClause_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Yield_Return(ctx *CaseClause_Yield_ReturnContext) {}

// EnterCaseClause_Await_Return is called when production caseClause_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Await_Return(ctx *CaseClause_Await_ReturnContext) {}

// ExitCaseClause_Await_Return is called when production caseClause_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Await_Return(ctx *CaseClause_Await_ReturnContext) {}

// EnterCaseClause_Yield_Await_Return is called when production caseClause_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterCaseClause_Yield_Await_Return(ctx *CaseClause_Yield_Await_ReturnContext) {
}

// ExitCaseClause_Yield_Await_Return is called when production caseClause_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitCaseClause_Yield_Await_Return(ctx *CaseClause_Yield_Await_ReturnContext) {
}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterDefaultClause_Yield is called when production defaultClause_Yield is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Yield(ctx *DefaultClause_YieldContext) {}

// ExitDefaultClause_Yield is called when production defaultClause_Yield is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Yield(ctx *DefaultClause_YieldContext) {}

// EnterDefaultClause_Await is called when production defaultClause_Await is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Await(ctx *DefaultClause_AwaitContext) {}

// ExitDefaultClause_Await is called when production defaultClause_Await is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Await(ctx *DefaultClause_AwaitContext) {}

// EnterDefaultClause_Yield_Await is called when production defaultClause_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Yield_Await(ctx *DefaultClause_Yield_AwaitContext) {
}

// ExitDefaultClause_Yield_Await is called when production defaultClause_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Yield_Await(ctx *DefaultClause_Yield_AwaitContext) {
}

// EnterDefaultClause_Return is called when production defaultClause_Return is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Return(ctx *DefaultClause_ReturnContext) {}

// ExitDefaultClause_Return is called when production defaultClause_Return is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Return(ctx *DefaultClause_ReturnContext) {}

// EnterDefaultClause_Yield_Return is called when production defaultClause_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Yield_Return(ctx *DefaultClause_Yield_ReturnContext) {
}

// ExitDefaultClause_Yield_Return is called when production defaultClause_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Yield_Return(ctx *DefaultClause_Yield_ReturnContext) {
}

// EnterDefaultClause_Await_Return is called when production defaultClause_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Await_Return(ctx *DefaultClause_Await_ReturnContext) {
}

// ExitDefaultClause_Await_Return is called when production defaultClause_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Await_Return(ctx *DefaultClause_Await_ReturnContext) {
}

// EnterDefaultClause_Yield_Await_Return is called when production defaultClause_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause_Yield_Await_Return(ctx *DefaultClause_Yield_Await_ReturnContext) {
}

// ExitDefaultClause_Yield_Await_Return is called when production defaultClause_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause_Yield_Await_Return(ctx *DefaultClause_Yield_Await_ReturnContext) {
}

// EnterLabelledStatement is called when production labelledStatement is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement(ctx *LabelledStatementContext) {}

// ExitLabelledStatement is called when production labelledStatement is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement(ctx *LabelledStatementContext) {}

// EnterLabelledStatement_Yield is called when production labelledStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Yield(ctx *LabelledStatement_YieldContext) {}

// ExitLabelledStatement_Yield is called when production labelledStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Yield(ctx *LabelledStatement_YieldContext) {}

// EnterLabelledStatement_Await is called when production labelledStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Await(ctx *LabelledStatement_AwaitContext) {}

// ExitLabelledStatement_Await is called when production labelledStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Await(ctx *LabelledStatement_AwaitContext) {}

// EnterLabelledStatement_Yield_Await is called when production labelledStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Yield_Await(ctx *LabelledStatement_Yield_AwaitContext) {
}

// ExitLabelledStatement_Yield_Await is called when production labelledStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Yield_Await(ctx *LabelledStatement_Yield_AwaitContext) {
}

// EnterLabelledStatement_Return is called when production labelledStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Return(ctx *LabelledStatement_ReturnContext) {}

// ExitLabelledStatement_Return is called when production labelledStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Return(ctx *LabelledStatement_ReturnContext) {}

// EnterLabelledStatement_Yield_Return is called when production labelledStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Yield_Return(ctx *LabelledStatement_Yield_ReturnContext) {
}

// ExitLabelledStatement_Yield_Return is called when production labelledStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Yield_Return(ctx *LabelledStatement_Yield_ReturnContext) {
}

// EnterLabelledStatement_Await_Return is called when production labelledStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Await_Return(ctx *LabelledStatement_Await_ReturnContext) {
}

// ExitLabelledStatement_Await_Return is called when production labelledStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Await_Return(ctx *LabelledStatement_Await_ReturnContext) {
}

// EnterLabelledStatement_Yield_Await_Return is called when production labelledStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement_Yield_Await_Return(ctx *LabelledStatement_Yield_Await_ReturnContext) {
}

// ExitLabelledStatement_Yield_Await_Return is called when production labelledStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement_Yield_Await_Return(ctx *LabelledStatement_Yield_Await_ReturnContext) {
}

// EnterLabelledItem is called when production labelledItem is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem(ctx *LabelledItemContext) {}

// ExitLabelledItem is called when production labelledItem is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem(ctx *LabelledItemContext) {}

// EnterLabelledItem_Yield is called when production labelledItem_Yield is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Yield(ctx *LabelledItem_YieldContext) {}

// ExitLabelledItem_Yield is called when production labelledItem_Yield is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Yield(ctx *LabelledItem_YieldContext) {}

// EnterLabelledItem_Await is called when production labelledItem_Await is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Await(ctx *LabelledItem_AwaitContext) {}

// ExitLabelledItem_Await is called when production labelledItem_Await is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Await(ctx *LabelledItem_AwaitContext) {}

// EnterLabelledItem_Yield_Await is called when production labelledItem_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Yield_Await(ctx *LabelledItem_Yield_AwaitContext) {}

// ExitLabelledItem_Yield_Await is called when production labelledItem_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Yield_Await(ctx *LabelledItem_Yield_AwaitContext) {}

// EnterLabelledItem_Return is called when production labelledItem_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Return(ctx *LabelledItem_ReturnContext) {}

// ExitLabelledItem_Return is called when production labelledItem_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Return(ctx *LabelledItem_ReturnContext) {}

// EnterLabelledItem_Yield_Return is called when production labelledItem_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Yield_Return(ctx *LabelledItem_Yield_ReturnContext) {
}

// ExitLabelledItem_Yield_Return is called when production labelledItem_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Yield_Return(ctx *LabelledItem_Yield_ReturnContext) {
}

// EnterLabelledItem_Await_Return is called when production labelledItem_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Await_Return(ctx *LabelledItem_Await_ReturnContext) {
}

// ExitLabelledItem_Await_Return is called when production labelledItem_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Await_Return(ctx *LabelledItem_Await_ReturnContext) {
}

// EnterLabelledItem_Yield_Await_Return is called when production labelledItem_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterLabelledItem_Yield_Await_Return(ctx *LabelledItem_Yield_Await_ReturnContext) {
}

// ExitLabelledItem_Yield_Await_Return is called when production labelledItem_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitLabelledItem_Yield_Await_Return(ctx *LabelledItem_Yield_Await_ReturnContext) {
}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseECMAScriptListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseECMAScriptListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterThrowStatement_Yield is called when production throwStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterThrowStatement_Yield(ctx *ThrowStatement_YieldContext) {}

// ExitThrowStatement_Yield is called when production throwStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitThrowStatement_Yield(ctx *ThrowStatement_YieldContext) {}

// EnterThrowStatement_Await is called when production throwStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterThrowStatement_Await(ctx *ThrowStatement_AwaitContext) {}

// ExitThrowStatement_Await is called when production throwStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitThrowStatement_Await(ctx *ThrowStatement_AwaitContext) {}

// EnterThrowStatement_Yield_Await is called when production throwStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterThrowStatement_Yield_Await(ctx *ThrowStatement_Yield_AwaitContext) {
}

// ExitThrowStatement_Yield_Await is called when production throwStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitThrowStatement_Yield_Await(ctx *ThrowStatement_Yield_AwaitContext) {
}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseECMAScriptListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseECMAScriptListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterTryStatement_Yield is called when production tryStatement_Yield is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Yield(ctx *TryStatement_YieldContext) {}

// ExitTryStatement_Yield is called when production tryStatement_Yield is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Yield(ctx *TryStatement_YieldContext) {}

// EnterTryStatement_Await is called when production tryStatement_Await is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Await(ctx *TryStatement_AwaitContext) {}

// ExitTryStatement_Await is called when production tryStatement_Await is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Await(ctx *TryStatement_AwaitContext) {}

// EnterTryStatement_Yield_Await is called when production tryStatement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Yield_Await(ctx *TryStatement_Yield_AwaitContext) {}

// ExitTryStatement_Yield_Await is called when production tryStatement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Yield_Await(ctx *TryStatement_Yield_AwaitContext) {}

// EnterTryStatement_Return is called when production tryStatement_Return is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Return(ctx *TryStatement_ReturnContext) {}

// ExitTryStatement_Return is called when production tryStatement_Return is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Return(ctx *TryStatement_ReturnContext) {}

// EnterTryStatement_Yield_Return is called when production tryStatement_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Yield_Return(ctx *TryStatement_Yield_ReturnContext) {
}

// ExitTryStatement_Yield_Return is called when production tryStatement_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Yield_Return(ctx *TryStatement_Yield_ReturnContext) {
}

// EnterTryStatement_Await_Return is called when production tryStatement_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Await_Return(ctx *TryStatement_Await_ReturnContext) {
}

// ExitTryStatement_Await_Return is called when production tryStatement_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Await_Return(ctx *TryStatement_Await_ReturnContext) {
}

// EnterTryStatement_Yield_Await_Return is called when production tryStatement_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterTryStatement_Yield_Await_Return(ctx *TryStatement_Yield_Await_ReturnContext) {
}

// ExitTryStatement_Yield_Await_Return is called when production tryStatement_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitTryStatement_Yield_Await_Return(ctx *TryStatement_Yield_Await_ReturnContext) {
}

// EnterCatch_ is called when production catch_ is entered.
func (s *BaseECMAScriptListener) EnterCatch_(ctx *Catch_Context) {}

// ExitCatch_ is called when production catch_ is exited.
func (s *BaseECMAScriptListener) ExitCatch_(ctx *Catch_Context) {}

// EnterCatch_Yield is called when production catch_Yield is entered.
func (s *BaseECMAScriptListener) EnterCatch_Yield(ctx *Catch_YieldContext) {}

// ExitCatch_Yield is called when production catch_Yield is exited.
func (s *BaseECMAScriptListener) ExitCatch_Yield(ctx *Catch_YieldContext) {}

// EnterCatch_Await is called when production catch_Await is entered.
func (s *BaseECMAScriptListener) EnterCatch_Await(ctx *Catch_AwaitContext) {}

// ExitCatch_Await is called when production catch_Await is exited.
func (s *BaseECMAScriptListener) ExitCatch_Await(ctx *Catch_AwaitContext) {}

// EnterCatch_Yield_Await is called when production catch_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCatch_Yield_Await(ctx *Catch_Yield_AwaitContext) {}

// ExitCatch_Yield_Await is called when production catch_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCatch_Yield_Await(ctx *Catch_Yield_AwaitContext) {}

// EnterCatch_Return is called when production catch_Return is entered.
func (s *BaseECMAScriptListener) EnterCatch_Return(ctx *Catch_ReturnContext) {}

// ExitCatch_Return is called when production catch_Return is exited.
func (s *BaseECMAScriptListener) ExitCatch_Return(ctx *Catch_ReturnContext) {}

// EnterCatch_Yield_Return is called when production catch_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterCatch_Yield_Return(ctx *Catch_Yield_ReturnContext) {}

// ExitCatch_Yield_Return is called when production catch_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitCatch_Yield_Return(ctx *Catch_Yield_ReturnContext) {}

// EnterCatch_Await_Return is called when production catch_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterCatch_Await_Return(ctx *Catch_Await_ReturnContext) {}

// ExitCatch_Await_Return is called when production catch_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitCatch_Await_Return(ctx *Catch_Await_ReturnContext) {}

// EnterCatch_Yield_Await_Return is called when production catch_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterCatch_Yield_Await_Return(ctx *Catch_Yield_Await_ReturnContext) {}

// ExitCatch_Yield_Await_Return is called when production catch_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitCatch_Yield_Await_Return(ctx *Catch_Yield_Await_ReturnContext) {}

// EnterFinally_ is called when production finally_ is entered.
func (s *BaseECMAScriptListener) EnterFinally_(ctx *Finally_Context) {}

// ExitFinally_ is called when production finally_ is exited.
func (s *BaseECMAScriptListener) ExitFinally_(ctx *Finally_Context) {}

// EnterFinally_Yield is called when production finally_Yield is entered.
func (s *BaseECMAScriptListener) EnterFinally_Yield(ctx *Finally_YieldContext) {}

// ExitFinally_Yield is called when production finally_Yield is exited.
func (s *BaseECMAScriptListener) ExitFinally_Yield(ctx *Finally_YieldContext) {}

// EnterFinally_Await is called when production finally_Await is entered.
func (s *BaseECMAScriptListener) EnterFinally_Await(ctx *Finally_AwaitContext) {}

// ExitFinally_Await is called when production finally_Await is exited.
func (s *BaseECMAScriptListener) ExitFinally_Await(ctx *Finally_AwaitContext) {}

// EnterFinally_Yield_Await is called when production finally_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFinally_Yield_Await(ctx *Finally_Yield_AwaitContext) {}

// ExitFinally_Yield_Await is called when production finally_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFinally_Yield_Await(ctx *Finally_Yield_AwaitContext) {}

// EnterFinally_Return is called when production finally_Return is entered.
func (s *BaseECMAScriptListener) EnterFinally_Return(ctx *Finally_ReturnContext) {}

// ExitFinally_Return is called when production finally_Return is exited.
func (s *BaseECMAScriptListener) ExitFinally_Return(ctx *Finally_ReturnContext) {}

// EnterFinally_Yield_Return is called when production finally_Yield_Return is entered.
func (s *BaseECMAScriptListener) EnterFinally_Yield_Return(ctx *Finally_Yield_ReturnContext) {}

// ExitFinally_Yield_Return is called when production finally_Yield_Return is exited.
func (s *BaseECMAScriptListener) ExitFinally_Yield_Return(ctx *Finally_Yield_ReturnContext) {}

// EnterFinally_Await_Return is called when production finally_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterFinally_Await_Return(ctx *Finally_Await_ReturnContext) {}

// ExitFinally_Await_Return is called when production finally_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitFinally_Await_Return(ctx *Finally_Await_ReturnContext) {}

// EnterFinally_Yield_Await_Return is called when production finally_Yield_Await_Return is entered.
func (s *BaseECMAScriptListener) EnterFinally_Yield_Await_Return(ctx *Finally_Yield_Await_ReturnContext) {
}

// ExitFinally_Yield_Await_Return is called when production finally_Yield_Await_Return is exited.
func (s *BaseECMAScriptListener) ExitFinally_Yield_Await_Return(ctx *Finally_Yield_Await_ReturnContext) {
}

// EnterCatchParameter is called when production catchParameter is entered.
func (s *BaseECMAScriptListener) EnterCatchParameter(ctx *CatchParameterContext) {}

// ExitCatchParameter is called when production catchParameter is exited.
func (s *BaseECMAScriptListener) ExitCatchParameter(ctx *CatchParameterContext) {}

// EnterCatchParameter_Yield is called when production catchParameter_Yield is entered.
func (s *BaseECMAScriptListener) EnterCatchParameter_Yield(ctx *CatchParameter_YieldContext) {}

// ExitCatchParameter_Yield is called when production catchParameter_Yield is exited.
func (s *BaseECMAScriptListener) ExitCatchParameter_Yield(ctx *CatchParameter_YieldContext) {}

// EnterCatchParameter_Await is called when production catchParameter_Await is entered.
func (s *BaseECMAScriptListener) EnterCatchParameter_Await(ctx *CatchParameter_AwaitContext) {}

// ExitCatchParameter_Await is called when production catchParameter_Await is exited.
func (s *BaseECMAScriptListener) ExitCatchParameter_Await(ctx *CatchParameter_AwaitContext) {}

// EnterCatchParameter_Yield_Await is called when production catchParameter_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCatchParameter_Yield_Await(ctx *CatchParameter_Yield_AwaitContext) {
}

// ExitCatchParameter_Yield_Await is called when production catchParameter_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCatchParameter_Yield_Await(ctx *CatchParameter_Yield_AwaitContext) {
}

// EnterDebuggerStatement is called when production debuggerStatement is entered.
func (s *BaseECMAScriptListener) EnterDebuggerStatement(ctx *DebuggerStatementContext) {}

// ExitDebuggerStatement is called when production debuggerStatement is exited.
func (s *BaseECMAScriptListener) ExitDebuggerStatement(ctx *DebuggerStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionDeclaration_Yield is called when production functionDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Yield(ctx *FunctionDeclaration_YieldContext) {
}

// ExitFunctionDeclaration_Yield is called when production functionDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Yield(ctx *FunctionDeclaration_YieldContext) {
}

// EnterFunctionDeclaration_Await is called when production functionDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Await(ctx *FunctionDeclaration_AwaitContext) {
}

// ExitFunctionDeclaration_Await is called when production functionDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Await(ctx *FunctionDeclaration_AwaitContext) {
}

// EnterFunctionDeclaration_Yield_Await is called when production functionDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Yield_Await(ctx *FunctionDeclaration_Yield_AwaitContext) {
}

// ExitFunctionDeclaration_Yield_Await is called when production functionDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Yield_Await(ctx *FunctionDeclaration_Yield_AwaitContext) {
}

// EnterFunctionDeclaration_Default is called when production functionDeclaration_Default is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Default(ctx *FunctionDeclaration_DefaultContext) {
}

// ExitFunctionDeclaration_Default is called when production functionDeclaration_Default is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Default(ctx *FunctionDeclaration_DefaultContext) {
}

// EnterFunctionDeclaration_Yield_Default is called when production functionDeclaration_Yield_Default is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Yield_Default(ctx *FunctionDeclaration_Yield_DefaultContext) {
}

// ExitFunctionDeclaration_Yield_Default is called when production functionDeclaration_Yield_Default is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Yield_Default(ctx *FunctionDeclaration_Yield_DefaultContext) {
}

// EnterFunctionDeclaration_Await_Default is called when production functionDeclaration_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Await_Default(ctx *FunctionDeclaration_Await_DefaultContext) {
}

// ExitFunctionDeclaration_Await_Default is called when production functionDeclaration_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Await_Default(ctx *FunctionDeclaration_Await_DefaultContext) {
}

// EnterFunctionDeclaration_Yield_Await_Default is called when production functionDeclaration_Yield_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration_Yield_Await_Default(ctx *FunctionDeclaration_Yield_Await_DefaultContext) {
}

// ExitFunctionDeclaration_Yield_Await_Default is called when production functionDeclaration_Yield_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration_Yield_Await_Default(ctx *FunctionDeclaration_Yield_Await_DefaultContext) {
}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseECMAScriptListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseECMAScriptListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterUniqueFormalParameters is called when production uniqueFormalParameters is entered.
func (s *BaseECMAScriptListener) EnterUniqueFormalParameters(ctx *UniqueFormalParametersContext) {}

// ExitUniqueFormalParameters is called when production uniqueFormalParameters is exited.
func (s *BaseECMAScriptListener) ExitUniqueFormalParameters(ctx *UniqueFormalParametersContext) {}

// EnterUniqueFormalParameters_Yield is called when production uniqueFormalParameters_Yield is entered.
func (s *BaseECMAScriptListener) EnterUniqueFormalParameters_Yield(ctx *UniqueFormalParameters_YieldContext) {
}

// ExitUniqueFormalParameters_Yield is called when production uniqueFormalParameters_Yield is exited.
func (s *BaseECMAScriptListener) ExitUniqueFormalParameters_Yield(ctx *UniqueFormalParameters_YieldContext) {
}

// EnterUniqueFormalParameters_Await is called when production uniqueFormalParameters_Await is entered.
func (s *BaseECMAScriptListener) EnterUniqueFormalParameters_Await(ctx *UniqueFormalParameters_AwaitContext) {
}

// ExitUniqueFormalParameters_Await is called when production uniqueFormalParameters_Await is exited.
func (s *BaseECMAScriptListener) ExitUniqueFormalParameters_Await(ctx *UniqueFormalParameters_AwaitContext) {
}

// EnterUniqueFormalParameters_Yield_Await is called when production uniqueFormalParameters_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterUniqueFormalParameters_Yield_Await(ctx *UniqueFormalParameters_Yield_AwaitContext) {
}

// ExitUniqueFormalParameters_Yield_Await is called when production uniqueFormalParameters_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitUniqueFormalParameters_Yield_Await(ctx *UniqueFormalParameters_Yield_AwaitContext) {
}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseECMAScriptListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseECMAScriptListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameters_Yield is called when production formalParameters_Yield is entered.
func (s *BaseECMAScriptListener) EnterFormalParameters_Yield(ctx *FormalParameters_YieldContext) {}

// ExitFormalParameters_Yield is called when production formalParameters_Yield is exited.
func (s *BaseECMAScriptListener) ExitFormalParameters_Yield(ctx *FormalParameters_YieldContext) {}

// EnterFormalParameters_Await is called when production formalParameters_Await is entered.
func (s *BaseECMAScriptListener) EnterFormalParameters_Await(ctx *FormalParameters_AwaitContext) {}

// ExitFormalParameters_Await is called when production formalParameters_Await is exited.
func (s *BaseECMAScriptListener) ExitFormalParameters_Await(ctx *FormalParameters_AwaitContext) {}

// EnterFormalParameters_Yield_Await is called when production formalParameters_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFormalParameters_Yield_Await(ctx *FormalParameters_Yield_AwaitContext) {
}

// ExitFormalParameters_Yield_Await is called when production formalParameters_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFormalParameters_Yield_Await(ctx *FormalParameters_Yield_AwaitContext) {
}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseECMAScriptListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseECMAScriptListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameterList_Yield is called when production formalParameterList_Yield is entered.
func (s *BaseECMAScriptListener) EnterFormalParameterList_Yield(ctx *FormalParameterList_YieldContext) {
}

// ExitFormalParameterList_Yield is called when production formalParameterList_Yield is exited.
func (s *BaseECMAScriptListener) ExitFormalParameterList_Yield(ctx *FormalParameterList_YieldContext) {
}

// EnterFormalParameterList_Await is called when production formalParameterList_Await is entered.
func (s *BaseECMAScriptListener) EnterFormalParameterList_Await(ctx *FormalParameterList_AwaitContext) {
}

// ExitFormalParameterList_Await is called when production formalParameterList_Await is exited.
func (s *BaseECMAScriptListener) ExitFormalParameterList_Await(ctx *FormalParameterList_AwaitContext) {
}

// EnterFormalParameterList_Yield_Await is called when production formalParameterList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFormalParameterList_Yield_Await(ctx *FormalParameterList_Yield_AwaitContext) {
}

// ExitFormalParameterList_Yield_Await is called when production formalParameterList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFormalParameterList_Yield_Await(ctx *FormalParameterList_Yield_AwaitContext) {
}

// EnterFunctionRestParameter is called when production functionRestParameter is entered.
func (s *BaseECMAScriptListener) EnterFunctionRestParameter(ctx *FunctionRestParameterContext) {}

// ExitFunctionRestParameter is called when production functionRestParameter is exited.
func (s *BaseECMAScriptListener) ExitFunctionRestParameter(ctx *FunctionRestParameterContext) {}

// EnterFunctionRestParameter_Yield is called when production functionRestParameter_Yield is entered.
func (s *BaseECMAScriptListener) EnterFunctionRestParameter_Yield(ctx *FunctionRestParameter_YieldContext) {
}

// ExitFunctionRestParameter_Yield is called when production functionRestParameter_Yield is exited.
func (s *BaseECMAScriptListener) ExitFunctionRestParameter_Yield(ctx *FunctionRestParameter_YieldContext) {
}

// EnterFunctionRestParameter_Await is called when production functionRestParameter_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionRestParameter_Await(ctx *FunctionRestParameter_AwaitContext) {
}

// ExitFunctionRestParameter_Await is called when production functionRestParameter_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionRestParameter_Await(ctx *FunctionRestParameter_AwaitContext) {
}

// EnterFunctionRestParameter_Yield_Await is called when production functionRestParameter_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionRestParameter_Yield_Await(ctx *FunctionRestParameter_Yield_AwaitContext) {
}

// ExitFunctionRestParameter_Yield_Await is called when production functionRestParameter_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionRestParameter_Yield_Await(ctx *FunctionRestParameter_Yield_AwaitContext) {
}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseECMAScriptListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseECMAScriptListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterFormalParameter_Yield is called when production formalParameter_Yield is entered.
func (s *BaseECMAScriptListener) EnterFormalParameter_Yield(ctx *FormalParameter_YieldContext) {}

// ExitFormalParameter_Yield is called when production formalParameter_Yield is exited.
func (s *BaseECMAScriptListener) ExitFormalParameter_Yield(ctx *FormalParameter_YieldContext) {}

// EnterFormalParameter_Await is called when production formalParameter_Await is entered.
func (s *BaseECMAScriptListener) EnterFormalParameter_Await(ctx *FormalParameter_AwaitContext) {}

// ExitFormalParameter_Await is called when production formalParameter_Await is exited.
func (s *BaseECMAScriptListener) ExitFormalParameter_Await(ctx *FormalParameter_AwaitContext) {}

// EnterFormalParameter_Yield_Await is called when production formalParameter_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFormalParameter_Yield_Await(ctx *FormalParameter_Yield_AwaitContext) {
}

// ExitFormalParameter_Yield_Await is called when production formalParameter_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFormalParameter_Yield_Await(ctx *FormalParameter_Yield_AwaitContext) {
}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseECMAScriptListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseECMAScriptListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterFunctionBody_Yield is called when production functionBody_Yield is entered.
func (s *BaseECMAScriptListener) EnterFunctionBody_Yield(ctx *FunctionBody_YieldContext) {}

// ExitFunctionBody_Yield is called when production functionBody_Yield is exited.
func (s *BaseECMAScriptListener) ExitFunctionBody_Yield(ctx *FunctionBody_YieldContext) {}

// EnterFunctionBody_Await is called when production functionBody_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionBody_Await(ctx *FunctionBody_AwaitContext) {}

// ExitFunctionBody_Await is called when production functionBody_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionBody_Await(ctx *FunctionBody_AwaitContext) {}

// EnterFunctionBody_Yield_Await is called when production functionBody_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionBody_Yield_Await(ctx *FunctionBody_Yield_AwaitContext) {}

// ExitFunctionBody_Yield_Await is called when production functionBody_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionBody_Yield_Await(ctx *FunctionBody_Yield_AwaitContext) {}

// EnterFunctionStatementList is called when production functionStatementList is entered.
func (s *BaseECMAScriptListener) EnterFunctionStatementList(ctx *FunctionStatementListContext) {}

// ExitFunctionStatementList is called when production functionStatementList is exited.
func (s *BaseECMAScriptListener) ExitFunctionStatementList(ctx *FunctionStatementListContext) {}

// EnterFunctionStatementList_Yield is called when production functionStatementList_Yield is entered.
func (s *BaseECMAScriptListener) EnterFunctionStatementList_Yield(ctx *FunctionStatementList_YieldContext) {
}

// ExitFunctionStatementList_Yield is called when production functionStatementList_Yield is exited.
func (s *BaseECMAScriptListener) ExitFunctionStatementList_Yield(ctx *FunctionStatementList_YieldContext) {
}

// EnterFunctionStatementList_Await is called when production functionStatementList_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionStatementList_Await(ctx *FunctionStatementList_AwaitContext) {
}

// ExitFunctionStatementList_Await is called when production functionStatementList_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionStatementList_Await(ctx *FunctionStatementList_AwaitContext) {
}

// EnterFunctionStatementList_Yield_Await is called when production functionStatementList_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterFunctionStatementList_Yield_Await(ctx *FunctionStatementList_Yield_AwaitContext) {
}

// ExitFunctionStatementList_Yield_Await is called when production functionStatementList_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitFunctionStatementList_Yield_Await(ctx *FunctionStatementList_Yield_AwaitContext) {
}

// EnterArrowFunction is called when production arrowFunction is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction(ctx *ArrowFunctionContext) {}

// ExitArrowFunction is called when production arrowFunction is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction(ctx *ArrowFunctionContext) {}

// EnterArrowFunction_In is called when production arrowFunction_In is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_In(ctx *ArrowFunction_InContext) {}

// ExitArrowFunction_In is called when production arrowFunction_In is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_In(ctx *ArrowFunction_InContext) {}

// EnterArrowFunction_Yield is called when production arrowFunction_Yield is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_Yield(ctx *ArrowFunction_YieldContext) {}

// ExitArrowFunction_Yield is called when production arrowFunction_Yield is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_Yield(ctx *ArrowFunction_YieldContext) {}

// EnterArrowFunction_In_Yield is called when production arrowFunction_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_In_Yield(ctx *ArrowFunction_In_YieldContext) {}

// ExitArrowFunction_In_Yield is called when production arrowFunction_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_In_Yield(ctx *ArrowFunction_In_YieldContext) {}

// EnterArrowFunction_Await is called when production arrowFunction_Await is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_Await(ctx *ArrowFunction_AwaitContext) {}

// ExitArrowFunction_Await is called when production arrowFunction_Await is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_Await(ctx *ArrowFunction_AwaitContext) {}

// EnterArrowFunction_In_Await is called when production arrowFunction_In_Await is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_In_Await(ctx *ArrowFunction_In_AwaitContext) {}

// ExitArrowFunction_In_Await is called when production arrowFunction_In_Await is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_In_Await(ctx *ArrowFunction_In_AwaitContext) {}

// EnterArrowFunction_Yield_Await is called when production arrowFunction_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_Yield_Await(ctx *ArrowFunction_Yield_AwaitContext) {
}

// ExitArrowFunction_Yield_Await is called when production arrowFunction_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_Yield_Await(ctx *ArrowFunction_Yield_AwaitContext) {
}

// EnterArrowFunction_In_Yield_Await is called when production arrowFunction_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArrowFunction_In_Yield_Await(ctx *ArrowFunction_In_Yield_AwaitContext) {
}

// ExitArrowFunction_In_Yield_Await is called when production arrowFunction_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArrowFunction_In_Yield_Await(ctx *ArrowFunction_In_Yield_AwaitContext) {
}

// EnterArrowParameters is called when production arrowParameters is entered.
func (s *BaseECMAScriptListener) EnterArrowParameters(ctx *ArrowParametersContext) {}

// ExitArrowParameters is called when production arrowParameters is exited.
func (s *BaseECMAScriptListener) ExitArrowParameters(ctx *ArrowParametersContext) {}

// EnterArrowParameters_Yield is called when production arrowParameters_Yield is entered.
func (s *BaseECMAScriptListener) EnterArrowParameters_Yield(ctx *ArrowParameters_YieldContext) {}

// ExitArrowParameters_Yield is called when production arrowParameters_Yield is exited.
func (s *BaseECMAScriptListener) ExitArrowParameters_Yield(ctx *ArrowParameters_YieldContext) {}

// EnterArrowParameters_Await is called when production arrowParameters_Await is entered.
func (s *BaseECMAScriptListener) EnterArrowParameters_Await(ctx *ArrowParameters_AwaitContext) {}

// ExitArrowParameters_Await is called when production arrowParameters_Await is exited.
func (s *BaseECMAScriptListener) ExitArrowParameters_Await(ctx *ArrowParameters_AwaitContext) {}

// EnterArrowParameters_Yield_Await is called when production arrowParameters_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterArrowParameters_Yield_Await(ctx *ArrowParameters_Yield_AwaitContext) {
}

// ExitArrowParameters_Yield_Await is called when production arrowParameters_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitArrowParameters_Yield_Await(ctx *ArrowParameters_Yield_AwaitContext) {
}

// EnterConciseBody is called when production conciseBody is entered.
func (s *BaseECMAScriptListener) EnterConciseBody(ctx *ConciseBodyContext) {}

// ExitConciseBody is called when production conciseBody is exited.
func (s *BaseECMAScriptListener) ExitConciseBody(ctx *ConciseBodyContext) {}

// EnterConciseBody_In is called when production conciseBody_In is entered.
func (s *BaseECMAScriptListener) EnterConciseBody_In(ctx *ConciseBody_InContext) {}

// ExitConciseBody_In is called when production conciseBody_In is exited.
func (s *BaseECMAScriptListener) ExitConciseBody_In(ctx *ConciseBody_InContext) {}

// EnterMethodDefinition is called when production methodDefinition is entered.
func (s *BaseECMAScriptListener) EnterMethodDefinition(ctx *MethodDefinitionContext) {}

// ExitMethodDefinition is called when production methodDefinition is exited.
func (s *BaseECMAScriptListener) ExitMethodDefinition(ctx *MethodDefinitionContext) {}

// EnterMethodDefinition_Yield is called when production methodDefinition_Yield is entered.
func (s *BaseECMAScriptListener) EnterMethodDefinition_Yield(ctx *MethodDefinition_YieldContext) {}

// ExitMethodDefinition_Yield is called when production methodDefinition_Yield is exited.
func (s *BaseECMAScriptListener) ExitMethodDefinition_Yield(ctx *MethodDefinition_YieldContext) {}

// EnterMethodDefinition_Await is called when production methodDefinition_Await is entered.
func (s *BaseECMAScriptListener) EnterMethodDefinition_Await(ctx *MethodDefinition_AwaitContext) {}

// ExitMethodDefinition_Await is called when production methodDefinition_Await is exited.
func (s *BaseECMAScriptListener) ExitMethodDefinition_Await(ctx *MethodDefinition_AwaitContext) {}

// EnterMethodDefinition_Yield_Await is called when production methodDefinition_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterMethodDefinition_Yield_Await(ctx *MethodDefinition_Yield_AwaitContext) {
}

// ExitMethodDefinition_Yield_Await is called when production methodDefinition_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitMethodDefinition_Yield_Await(ctx *MethodDefinition_Yield_AwaitContext) {
}

// EnterPropertySetParameterList is called when production propertySetParameterList is entered.
func (s *BaseECMAScriptListener) EnterPropertySetParameterList(ctx *PropertySetParameterListContext) {}

// ExitPropertySetParameterList is called when production propertySetParameterList is exited.
func (s *BaseECMAScriptListener) ExitPropertySetParameterList(ctx *PropertySetParameterListContext) {}

// EnterGeneratorMethod is called when production generatorMethod is entered.
func (s *BaseECMAScriptListener) EnterGeneratorMethod(ctx *GeneratorMethodContext) {}

// ExitGeneratorMethod is called when production generatorMethod is exited.
func (s *BaseECMAScriptListener) ExitGeneratorMethod(ctx *GeneratorMethodContext) {}

// EnterGeneratorMethod_Yield is called when production generatorMethod_Yield is entered.
func (s *BaseECMAScriptListener) EnterGeneratorMethod_Yield(ctx *GeneratorMethod_YieldContext) {}

// ExitGeneratorMethod_Yield is called when production generatorMethod_Yield is exited.
func (s *BaseECMAScriptListener) ExitGeneratorMethod_Yield(ctx *GeneratorMethod_YieldContext) {}

// EnterGeneratorMethod_Await is called when production generatorMethod_Await is entered.
func (s *BaseECMAScriptListener) EnterGeneratorMethod_Await(ctx *GeneratorMethod_AwaitContext) {}

// ExitGeneratorMethod_Await is called when production generatorMethod_Await is exited.
func (s *BaseECMAScriptListener) ExitGeneratorMethod_Await(ctx *GeneratorMethod_AwaitContext) {}

// EnterGeneratorMethod_Yield_Await is called when production generatorMethod_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterGeneratorMethod_Yield_Await(ctx *GeneratorMethod_Yield_AwaitContext) {
}

// ExitGeneratorMethod_Yield_Await is called when production generatorMethod_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitGeneratorMethod_Yield_Await(ctx *GeneratorMethod_Yield_AwaitContext) {
}

// EnterGeneratorDeclaration is called when production generatorDeclaration is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration(ctx *GeneratorDeclarationContext) {}

// ExitGeneratorDeclaration is called when production generatorDeclaration is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration(ctx *GeneratorDeclarationContext) {}

// EnterGeneratorDeclaration_Yield is called when production generatorDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Yield(ctx *GeneratorDeclaration_YieldContext) {
}

// ExitGeneratorDeclaration_Yield is called when production generatorDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Yield(ctx *GeneratorDeclaration_YieldContext) {
}

// EnterGeneratorDeclaration_Await is called when production generatorDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Await(ctx *GeneratorDeclaration_AwaitContext) {
}

// ExitGeneratorDeclaration_Await is called when production generatorDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Await(ctx *GeneratorDeclaration_AwaitContext) {
}

// EnterGeneratorDeclaration_Yield_Await is called when production generatorDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Yield_Await(ctx *GeneratorDeclaration_Yield_AwaitContext) {
}

// ExitGeneratorDeclaration_Yield_Await is called when production generatorDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Yield_Await(ctx *GeneratorDeclaration_Yield_AwaitContext) {
}

// EnterGeneratorDeclaration_Default is called when production generatorDeclaration_Default is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Default(ctx *GeneratorDeclaration_DefaultContext) {
}

// ExitGeneratorDeclaration_Default is called when production generatorDeclaration_Default is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Default(ctx *GeneratorDeclaration_DefaultContext) {
}

// EnterGeneratorDeclaration_Yield_Default is called when production generatorDeclaration_Yield_Default is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Yield_Default(ctx *GeneratorDeclaration_Yield_DefaultContext) {
}

// ExitGeneratorDeclaration_Yield_Default is called when production generatorDeclaration_Yield_Default is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Yield_Default(ctx *GeneratorDeclaration_Yield_DefaultContext) {
}

// EnterGeneratorDeclaration_Await_Default is called when production generatorDeclaration_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Await_Default(ctx *GeneratorDeclaration_Await_DefaultContext) {
}

// ExitGeneratorDeclaration_Await_Default is called when production generatorDeclaration_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Await_Default(ctx *GeneratorDeclaration_Await_DefaultContext) {
}

// EnterGeneratorDeclaration_Yield_Await_Default is called when production generatorDeclaration_Yield_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterGeneratorDeclaration_Yield_Await_Default(ctx *GeneratorDeclaration_Yield_Await_DefaultContext) {
}

// ExitGeneratorDeclaration_Yield_Await_Default is called when production generatorDeclaration_Yield_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitGeneratorDeclaration_Yield_Await_Default(ctx *GeneratorDeclaration_Yield_Await_DefaultContext) {
}

// EnterGeneratorExpression is called when production generatorExpression is entered.
func (s *BaseECMAScriptListener) EnterGeneratorExpression(ctx *GeneratorExpressionContext) {}

// ExitGeneratorExpression is called when production generatorExpression is exited.
func (s *BaseECMAScriptListener) ExitGeneratorExpression(ctx *GeneratorExpressionContext) {}

// EnterGeneratorBody is called when production generatorBody is entered.
func (s *BaseECMAScriptListener) EnterGeneratorBody(ctx *GeneratorBodyContext) {}

// ExitGeneratorBody is called when production generatorBody is exited.
func (s *BaseECMAScriptListener) ExitGeneratorBody(ctx *GeneratorBodyContext) {}

// EnterYieldExpression is called when production yieldExpression is entered.
func (s *BaseECMAScriptListener) EnterYieldExpression(ctx *YieldExpressionContext) {}

// ExitYieldExpression is called when production yieldExpression is exited.
func (s *BaseECMAScriptListener) ExitYieldExpression(ctx *YieldExpressionContext) {}

// EnterYieldExpression_In is called when production yieldExpression_In is entered.
func (s *BaseECMAScriptListener) EnterYieldExpression_In(ctx *YieldExpression_InContext) {}

// ExitYieldExpression_In is called when production yieldExpression_In is exited.
func (s *BaseECMAScriptListener) ExitYieldExpression_In(ctx *YieldExpression_InContext) {}

// EnterYieldExpression_Await is called when production yieldExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterYieldExpression_Await(ctx *YieldExpression_AwaitContext) {}

// ExitYieldExpression_Await is called when production yieldExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitYieldExpression_Await(ctx *YieldExpression_AwaitContext) {}

// EnterYieldExpression_In_Await is called when production yieldExpression_In_Await is entered.
func (s *BaseECMAScriptListener) EnterYieldExpression_In_Await(ctx *YieldExpression_In_AwaitContext) {}

// ExitYieldExpression_In_Await is called when production yieldExpression_In_Await is exited.
func (s *BaseECMAScriptListener) ExitYieldExpression_In_Await(ctx *YieldExpression_In_AwaitContext) {}

// EnterAsyncGeneratorMethod is called when production asyncGeneratorMethod is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorMethod(ctx *AsyncGeneratorMethodContext) {}

// ExitAsyncGeneratorMethod is called when production asyncGeneratorMethod is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorMethod(ctx *AsyncGeneratorMethodContext) {}

// EnterAsyncGeneratorMethod_Yield is called when production asyncGeneratorMethod_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorMethod_Yield(ctx *AsyncGeneratorMethod_YieldContext) {
}

// ExitAsyncGeneratorMethod_Yield is called when production asyncGeneratorMethod_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorMethod_Yield(ctx *AsyncGeneratorMethod_YieldContext) {
}

// EnterAsyncGeneratorMethod_Await is called when production asyncGeneratorMethod_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorMethod_Await(ctx *AsyncGeneratorMethod_AwaitContext) {
}

// ExitAsyncGeneratorMethod_Await is called when production asyncGeneratorMethod_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorMethod_Await(ctx *AsyncGeneratorMethod_AwaitContext) {
}

// EnterAsyncGeneratorMethod_Yield_Await is called when production asyncGeneratorMethod_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorMethod_Yield_Await(ctx *AsyncGeneratorMethod_Yield_AwaitContext) {
}

// ExitAsyncGeneratorMethod_Yield_Await is called when production asyncGeneratorMethod_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorMethod_Yield_Await(ctx *AsyncGeneratorMethod_Yield_AwaitContext) {
}

// EnterAsyncGeneratorDeclaration is called when production asyncGeneratorDeclaration is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration(ctx *AsyncGeneratorDeclarationContext) {
}

// ExitAsyncGeneratorDeclaration is called when production asyncGeneratorDeclaration is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration(ctx *AsyncGeneratorDeclarationContext) {
}

// EnterAsyncGeneratorDeclaration_Yield is called when production asyncGeneratorDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Yield(ctx *AsyncGeneratorDeclaration_YieldContext) {
}

// ExitAsyncGeneratorDeclaration_Yield is called when production asyncGeneratorDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Yield(ctx *AsyncGeneratorDeclaration_YieldContext) {
}

// EnterAsyncGeneratorDeclaration_Await is called when production asyncGeneratorDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Await(ctx *AsyncGeneratorDeclaration_AwaitContext) {
}

// ExitAsyncGeneratorDeclaration_Await is called when production asyncGeneratorDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Await(ctx *AsyncGeneratorDeclaration_AwaitContext) {
}

// EnterAsyncGeneratorDeclaration_Yield_Await is called when production asyncGeneratorDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Yield_Await(ctx *AsyncGeneratorDeclaration_Yield_AwaitContext) {
}

// ExitAsyncGeneratorDeclaration_Yield_Await is called when production asyncGeneratorDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Yield_Await(ctx *AsyncGeneratorDeclaration_Yield_AwaitContext) {
}

// EnterAsyncGeneratorDeclaration_Default is called when production asyncGeneratorDeclaration_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Default(ctx *AsyncGeneratorDeclaration_DefaultContext) {
}

// ExitAsyncGeneratorDeclaration_Default is called when production asyncGeneratorDeclaration_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Default(ctx *AsyncGeneratorDeclaration_DefaultContext) {
}

// EnterAsyncGeneratorDeclaration_Yield_Default is called when production asyncGeneratorDeclaration_Yield_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Yield_Default(ctx *AsyncGeneratorDeclaration_Yield_DefaultContext) {
}

// ExitAsyncGeneratorDeclaration_Yield_Default is called when production asyncGeneratorDeclaration_Yield_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Yield_Default(ctx *AsyncGeneratorDeclaration_Yield_DefaultContext) {
}

// EnterAsyncGeneratorDeclaration_Await_Default is called when production asyncGeneratorDeclaration_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Await_Default(ctx *AsyncGeneratorDeclaration_Await_DefaultContext) {
}

// ExitAsyncGeneratorDeclaration_Await_Default is called when production asyncGeneratorDeclaration_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Await_Default(ctx *AsyncGeneratorDeclaration_Await_DefaultContext) {
}

// EnterAsyncGeneratorDeclaration_Yield_Await_Default is called when production asyncGeneratorDeclaration_Yield_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorDeclaration_Yield_Await_Default(ctx *AsyncGeneratorDeclaration_Yield_Await_DefaultContext) {
}

// ExitAsyncGeneratorDeclaration_Yield_Await_Default is called when production asyncGeneratorDeclaration_Yield_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorDeclaration_Yield_Await_Default(ctx *AsyncGeneratorDeclaration_Yield_Await_DefaultContext) {
}

// EnterAsyncGeneratorExpression is called when production asyncGeneratorExpression is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorExpression(ctx *AsyncGeneratorExpressionContext) {}

// ExitAsyncGeneratorExpression is called when production asyncGeneratorExpression is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorExpression(ctx *AsyncGeneratorExpressionContext) {}

// EnterAsyncGeneratorBody is called when production asyncGeneratorBody is entered.
func (s *BaseECMAScriptListener) EnterAsyncGeneratorBody(ctx *AsyncGeneratorBodyContext) {}

// ExitAsyncGeneratorBody is called when production asyncGeneratorBody is exited.
func (s *BaseECMAScriptListener) ExitAsyncGeneratorBody(ctx *AsyncGeneratorBodyContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassDeclaration_Yield is called when production classDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Yield(ctx *ClassDeclaration_YieldContext) {}

// ExitClassDeclaration_Yield is called when production classDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Yield(ctx *ClassDeclaration_YieldContext) {}

// EnterClassDeclaration_Await is called when production classDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Await(ctx *ClassDeclaration_AwaitContext) {}

// ExitClassDeclaration_Await is called when production classDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Await(ctx *ClassDeclaration_AwaitContext) {}

// EnterClassDeclaration_Yield_Await is called when production classDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Yield_Await(ctx *ClassDeclaration_Yield_AwaitContext) {
}

// ExitClassDeclaration_Yield_Await is called when production classDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Yield_Await(ctx *ClassDeclaration_Yield_AwaitContext) {
}

// EnterClassDeclaration_Default is called when production classDeclaration_Default is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Default(ctx *ClassDeclaration_DefaultContext) {}

// ExitClassDeclaration_Default is called when production classDeclaration_Default is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Default(ctx *ClassDeclaration_DefaultContext) {}

// EnterClassDeclaration_Yield_Default is called when production classDeclaration_Yield_Default is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Yield_Default(ctx *ClassDeclaration_Yield_DefaultContext) {
}

// ExitClassDeclaration_Yield_Default is called when production classDeclaration_Yield_Default is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Yield_Default(ctx *ClassDeclaration_Yield_DefaultContext) {
}

// EnterClassDeclaration_Await_Default is called when production classDeclaration_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Await_Default(ctx *ClassDeclaration_Await_DefaultContext) {
}

// ExitClassDeclaration_Await_Default is called when production classDeclaration_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Await_Default(ctx *ClassDeclaration_Await_DefaultContext) {
}

// EnterClassDeclaration_Yield_Await_Default is called when production classDeclaration_Yield_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterClassDeclaration_Yield_Await_Default(ctx *ClassDeclaration_Yield_Await_DefaultContext) {
}

// ExitClassDeclaration_Yield_Await_Default is called when production classDeclaration_Yield_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitClassDeclaration_Yield_Await_Default(ctx *ClassDeclaration_Yield_Await_DefaultContext) {
}

// EnterClassExpression is called when production classExpression is entered.
func (s *BaseECMAScriptListener) EnterClassExpression(ctx *ClassExpressionContext) {}

// ExitClassExpression is called when production classExpression is exited.
func (s *BaseECMAScriptListener) ExitClassExpression(ctx *ClassExpressionContext) {}

// EnterClassExpression_Yield is called when production classExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterClassExpression_Yield(ctx *ClassExpression_YieldContext) {}

// ExitClassExpression_Yield is called when production classExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitClassExpression_Yield(ctx *ClassExpression_YieldContext) {}

// EnterClassExpression_Await is called when production classExpression_Await is entered.
func (s *BaseECMAScriptListener) EnterClassExpression_Await(ctx *ClassExpression_AwaitContext) {}

// ExitClassExpression_Await is called when production classExpression_Await is exited.
func (s *BaseECMAScriptListener) ExitClassExpression_Await(ctx *ClassExpression_AwaitContext) {}

// EnterClassExpression_Yield_Await is called when production classExpression_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterClassExpression_Yield_Await(ctx *ClassExpression_Yield_AwaitContext) {
}

// ExitClassExpression_Yield_Await is called when production classExpression_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitClassExpression_Yield_Await(ctx *ClassExpression_Yield_AwaitContext) {
}

// EnterClassTail is called when production classTail is entered.
func (s *BaseECMAScriptListener) EnterClassTail(ctx *ClassTailContext) {}

// ExitClassTail is called when production classTail is exited.
func (s *BaseECMAScriptListener) ExitClassTail(ctx *ClassTailContext) {}

// EnterClassTail_Yield is called when production classTail_Yield is entered.
func (s *BaseECMAScriptListener) EnterClassTail_Yield(ctx *ClassTail_YieldContext) {}

// ExitClassTail_Yield is called when production classTail_Yield is exited.
func (s *BaseECMAScriptListener) ExitClassTail_Yield(ctx *ClassTail_YieldContext) {}

// EnterClassTail_Await is called when production classTail_Await is entered.
func (s *BaseECMAScriptListener) EnterClassTail_Await(ctx *ClassTail_AwaitContext) {}

// ExitClassTail_Await is called when production classTail_Await is exited.
func (s *BaseECMAScriptListener) ExitClassTail_Await(ctx *ClassTail_AwaitContext) {}

// EnterClassTail_Yield_Await is called when production classTail_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterClassTail_Yield_Await(ctx *ClassTail_Yield_AwaitContext) {}

// ExitClassTail_Yield_Await is called when production classTail_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitClassTail_Yield_Await(ctx *ClassTail_Yield_AwaitContext) {}

// EnterClassHeritage is called when production classHeritage is entered.
func (s *BaseECMAScriptListener) EnterClassHeritage(ctx *ClassHeritageContext) {}

// ExitClassHeritage is called when production classHeritage is exited.
func (s *BaseECMAScriptListener) ExitClassHeritage(ctx *ClassHeritageContext) {}

// EnterClassHeritage_Yield is called when production classHeritage_Yield is entered.
func (s *BaseECMAScriptListener) EnterClassHeritage_Yield(ctx *ClassHeritage_YieldContext) {}

// ExitClassHeritage_Yield is called when production classHeritage_Yield is exited.
func (s *BaseECMAScriptListener) ExitClassHeritage_Yield(ctx *ClassHeritage_YieldContext) {}

// EnterClassHeritage_Await is called when production classHeritage_Await is entered.
func (s *BaseECMAScriptListener) EnterClassHeritage_Await(ctx *ClassHeritage_AwaitContext) {}

// ExitClassHeritage_Await is called when production classHeritage_Await is exited.
func (s *BaseECMAScriptListener) ExitClassHeritage_Await(ctx *ClassHeritage_AwaitContext) {}

// EnterClassHeritage_Yield_Await is called when production classHeritage_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterClassHeritage_Yield_Await(ctx *ClassHeritage_Yield_AwaitContext) {
}

// ExitClassHeritage_Yield_Await is called when production classHeritage_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitClassHeritage_Yield_Await(ctx *ClassHeritage_Yield_AwaitContext) {
}

// EnterClassBody is called when production classBody is entered.
func (s *BaseECMAScriptListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseECMAScriptListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBody_Yield is called when production classBody_Yield is entered.
func (s *BaseECMAScriptListener) EnterClassBody_Yield(ctx *ClassBody_YieldContext) {}

// ExitClassBody_Yield is called when production classBody_Yield is exited.
func (s *BaseECMAScriptListener) ExitClassBody_Yield(ctx *ClassBody_YieldContext) {}

// EnterClassBody_Await is called when production classBody_Await is entered.
func (s *BaseECMAScriptListener) EnterClassBody_Await(ctx *ClassBody_AwaitContext) {}

// ExitClassBody_Await is called when production classBody_Await is exited.
func (s *BaseECMAScriptListener) ExitClassBody_Await(ctx *ClassBody_AwaitContext) {}

// EnterClassBody_Yield_Await is called when production classBody_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterClassBody_Yield_Await(ctx *ClassBody_Yield_AwaitContext) {}

// ExitClassBody_Yield_Await is called when production classBody_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitClassBody_Yield_Await(ctx *ClassBody_Yield_AwaitContext) {}

// EnterClassElement is called when production classElement is entered.
func (s *BaseECMAScriptListener) EnterClassElement(ctx *ClassElementContext) {}

// ExitClassElement is called when production classElement is exited.
func (s *BaseECMAScriptListener) ExitClassElement(ctx *ClassElementContext) {}

// EnterClassElement_Yield is called when production classElement_Yield is entered.
func (s *BaseECMAScriptListener) EnterClassElement_Yield(ctx *ClassElement_YieldContext) {}

// ExitClassElement_Yield is called when production classElement_Yield is exited.
func (s *BaseECMAScriptListener) ExitClassElement_Yield(ctx *ClassElement_YieldContext) {}

// EnterClassElement_Await is called when production classElement_Await is entered.
func (s *BaseECMAScriptListener) EnterClassElement_Await(ctx *ClassElement_AwaitContext) {}

// ExitClassElement_Await is called when production classElement_Await is exited.
func (s *BaseECMAScriptListener) ExitClassElement_Await(ctx *ClassElement_AwaitContext) {}

// EnterClassElement_Yield_Await is called when production classElement_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterClassElement_Yield_Await(ctx *ClassElement_Yield_AwaitContext) {}

// ExitClassElement_Yield_Await is called when production classElement_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitClassElement_Yield_Await(ctx *ClassElement_Yield_AwaitContext) {}

// EnterAsyncFunctionDeclaration is called when production asyncFunctionDeclaration is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration(ctx *AsyncFunctionDeclarationContext) {}

// ExitAsyncFunctionDeclaration is called when production asyncFunctionDeclaration is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration(ctx *AsyncFunctionDeclarationContext) {}

// EnterAsyncFunctionDeclaration_Yield is called when production asyncFunctionDeclaration_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Yield(ctx *AsyncFunctionDeclaration_YieldContext) {
}

// ExitAsyncFunctionDeclaration_Yield is called when production asyncFunctionDeclaration_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Yield(ctx *AsyncFunctionDeclaration_YieldContext) {
}

// EnterAsyncFunctionDeclaration_Await is called when production asyncFunctionDeclaration_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Await(ctx *AsyncFunctionDeclaration_AwaitContext) {
}

// ExitAsyncFunctionDeclaration_Await is called when production asyncFunctionDeclaration_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Await(ctx *AsyncFunctionDeclaration_AwaitContext) {
}

// EnterAsyncFunctionDeclaration_Yield_Await is called when production asyncFunctionDeclaration_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Yield_Await(ctx *AsyncFunctionDeclaration_Yield_AwaitContext) {
}

// ExitAsyncFunctionDeclaration_Yield_Await is called when production asyncFunctionDeclaration_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Yield_Await(ctx *AsyncFunctionDeclaration_Yield_AwaitContext) {
}

// EnterAsyncFunctionDeclaration_Default is called when production asyncFunctionDeclaration_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Default(ctx *AsyncFunctionDeclaration_DefaultContext) {
}

// ExitAsyncFunctionDeclaration_Default is called when production asyncFunctionDeclaration_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Default(ctx *AsyncFunctionDeclaration_DefaultContext) {
}

// EnterAsyncFunctionDeclaration_Yield_Default is called when production asyncFunctionDeclaration_Yield_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Yield_Default(ctx *AsyncFunctionDeclaration_Yield_DefaultContext) {
}

// ExitAsyncFunctionDeclaration_Yield_Default is called when production asyncFunctionDeclaration_Yield_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Yield_Default(ctx *AsyncFunctionDeclaration_Yield_DefaultContext) {
}

// EnterAsyncFunctionDeclaration_Await_Default is called when production asyncFunctionDeclaration_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Await_Default(ctx *AsyncFunctionDeclaration_Await_DefaultContext) {
}

// ExitAsyncFunctionDeclaration_Await_Default is called when production asyncFunctionDeclaration_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Await_Default(ctx *AsyncFunctionDeclaration_Await_DefaultContext) {
}

// EnterAsyncFunctionDeclaration_Yield_Await_Default is called when production asyncFunctionDeclaration_Yield_Await_Default is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionDeclaration_Yield_Await_Default(ctx *AsyncFunctionDeclaration_Yield_Await_DefaultContext) {
}

// ExitAsyncFunctionDeclaration_Yield_Await_Default is called when production asyncFunctionDeclaration_Yield_Await_Default is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionDeclaration_Yield_Await_Default(ctx *AsyncFunctionDeclaration_Yield_Await_DefaultContext) {
}

// EnterAsyncFunctionExpression is called when production asyncFunctionExpression is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionExpression(ctx *AsyncFunctionExpressionContext) {}

// ExitAsyncFunctionExpression is called when production asyncFunctionExpression is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionExpression(ctx *AsyncFunctionExpressionContext) {}

// EnterAsyncMethod is called when production asyncMethod is entered.
func (s *BaseECMAScriptListener) EnterAsyncMethod(ctx *AsyncMethodContext) {}

// ExitAsyncMethod is called when production asyncMethod is exited.
func (s *BaseECMAScriptListener) ExitAsyncMethod(ctx *AsyncMethodContext) {}

// EnterAsyncMethod_Yield is called when production asyncMethod_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncMethod_Yield(ctx *AsyncMethod_YieldContext) {}

// ExitAsyncMethod_Yield is called when production asyncMethod_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncMethod_Yield(ctx *AsyncMethod_YieldContext) {}

// EnterAsyncMethod_Await is called when production asyncMethod_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncMethod_Await(ctx *AsyncMethod_AwaitContext) {}

// ExitAsyncMethod_Await is called when production asyncMethod_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncMethod_Await(ctx *AsyncMethod_AwaitContext) {}

// EnterAsyncMethod_Yield_Await is called when production asyncMethod_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncMethod_Yield_Await(ctx *AsyncMethod_Yield_AwaitContext) {}

// ExitAsyncMethod_Yield_Await is called when production asyncMethod_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncMethod_Yield_Await(ctx *AsyncMethod_Yield_AwaitContext) {}

// EnterAsyncFunctionBody is called when production asyncFunctionBody is entered.
func (s *BaseECMAScriptListener) EnterAsyncFunctionBody(ctx *AsyncFunctionBodyContext) {}

// ExitAsyncFunctionBody is called when production asyncFunctionBody is exited.
func (s *BaseECMAScriptListener) ExitAsyncFunctionBody(ctx *AsyncFunctionBodyContext) {}

// EnterAwaitExpression is called when production awaitExpression is entered.
func (s *BaseECMAScriptListener) EnterAwaitExpression(ctx *AwaitExpressionContext) {}

// ExitAwaitExpression is called when production awaitExpression is exited.
func (s *BaseECMAScriptListener) ExitAwaitExpression(ctx *AwaitExpressionContext) {}

// EnterAwaitExpression_Yield is called when production awaitExpression_Yield is entered.
func (s *BaseECMAScriptListener) EnterAwaitExpression_Yield(ctx *AwaitExpression_YieldContext) {}

// ExitAwaitExpression_Yield is called when production awaitExpression_Yield is exited.
func (s *BaseECMAScriptListener) ExitAwaitExpression_Yield(ctx *AwaitExpression_YieldContext) {}

// EnterAsyncArrowFunction is called when production asyncArrowFunction is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction(ctx *AsyncArrowFunctionContext) {}

// ExitAsyncArrowFunction is called when production asyncArrowFunction is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction(ctx *AsyncArrowFunctionContext) {}

// EnterAsyncArrowFunction_In is called when production asyncArrowFunction_In is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_In(ctx *AsyncArrowFunction_InContext) {}

// ExitAsyncArrowFunction_In is called when production asyncArrowFunction_In is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_In(ctx *AsyncArrowFunction_InContext) {}

// EnterAsyncArrowFunction_Yield is called when production asyncArrowFunction_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_Yield(ctx *AsyncArrowFunction_YieldContext) {}

// ExitAsyncArrowFunction_Yield is called when production asyncArrowFunction_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_Yield(ctx *AsyncArrowFunction_YieldContext) {}

// EnterAsyncArrowFunction_In_Yield is called when production asyncArrowFunction_In_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_In_Yield(ctx *AsyncArrowFunction_In_YieldContext) {
}

// ExitAsyncArrowFunction_In_Yield is called when production asyncArrowFunction_In_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_In_Yield(ctx *AsyncArrowFunction_In_YieldContext) {
}

// EnterAsyncArrowFunction_Await is called when production asyncArrowFunction_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_Await(ctx *AsyncArrowFunction_AwaitContext) {}

// ExitAsyncArrowFunction_Await is called when production asyncArrowFunction_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_Await(ctx *AsyncArrowFunction_AwaitContext) {}

// EnterAsyncArrowFunction_In_Await is called when production asyncArrowFunction_In_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_In_Await(ctx *AsyncArrowFunction_In_AwaitContext) {
}

// ExitAsyncArrowFunction_In_Await is called when production asyncArrowFunction_In_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_In_Await(ctx *AsyncArrowFunction_In_AwaitContext) {
}

// EnterAsyncArrowFunction_Yield_Await is called when production asyncArrowFunction_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_Yield_Await(ctx *AsyncArrowFunction_Yield_AwaitContext) {
}

// ExitAsyncArrowFunction_Yield_Await is called when production asyncArrowFunction_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_Yield_Await(ctx *AsyncArrowFunction_Yield_AwaitContext) {
}

// EnterAsyncArrowFunction_In_Yield_Await is called when production asyncArrowFunction_In_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowFunction_In_Yield_Await(ctx *AsyncArrowFunction_In_Yield_AwaitContext) {
}

// ExitAsyncArrowFunction_In_Yield_Await is called when production asyncArrowFunction_In_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowFunction_In_Yield_Await(ctx *AsyncArrowFunction_In_Yield_AwaitContext) {
}

// EnterAsyncArrowBindingIdentifier is called when production asyncArrowBindingIdentifier is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowBindingIdentifier(ctx *AsyncArrowBindingIdentifierContext) {
}

// ExitAsyncArrowBindingIdentifier is called when production asyncArrowBindingIdentifier is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowBindingIdentifier(ctx *AsyncArrowBindingIdentifierContext) {
}

// EnterAsyncArrowBindingIdentifier_Yield is called when production asyncArrowBindingIdentifier_Yield is entered.
func (s *BaseECMAScriptListener) EnterAsyncArrowBindingIdentifier_Yield(ctx *AsyncArrowBindingIdentifier_YieldContext) {
}

// ExitAsyncArrowBindingIdentifier_Yield is called when production asyncArrowBindingIdentifier_Yield is exited.
func (s *BaseECMAScriptListener) ExitAsyncArrowBindingIdentifier_Yield(ctx *AsyncArrowBindingIdentifier_YieldContext) {
}

// EnterCoverCallExpressionAndAsyncArrowHead is called when production coverCallExpressionAndAsyncArrowHead is entered.
func (s *BaseECMAScriptListener) EnterCoverCallExpressionAndAsyncArrowHead(ctx *CoverCallExpressionAndAsyncArrowHeadContext) {
}

// ExitCoverCallExpressionAndAsyncArrowHead is called when production coverCallExpressionAndAsyncArrowHead is exited.
func (s *BaseECMAScriptListener) ExitCoverCallExpressionAndAsyncArrowHead(ctx *CoverCallExpressionAndAsyncArrowHeadContext) {
}

// EnterCoverCallExpressionAndAsyncArrowHead_Yield is called when production coverCallExpressionAndAsyncArrowHead_Yield is entered.
func (s *BaseECMAScriptListener) EnterCoverCallExpressionAndAsyncArrowHead_Yield(ctx *CoverCallExpressionAndAsyncArrowHead_YieldContext) {
}

// ExitCoverCallExpressionAndAsyncArrowHead_Yield is called when production coverCallExpressionAndAsyncArrowHead_Yield is exited.
func (s *BaseECMAScriptListener) ExitCoverCallExpressionAndAsyncArrowHead_Yield(ctx *CoverCallExpressionAndAsyncArrowHead_YieldContext) {
}

// EnterCoverCallExpressionAndAsyncArrowHead_Await is called when production coverCallExpressionAndAsyncArrowHead_Await is entered.
func (s *BaseECMAScriptListener) EnterCoverCallExpressionAndAsyncArrowHead_Await(ctx *CoverCallExpressionAndAsyncArrowHead_AwaitContext) {
}

// ExitCoverCallExpressionAndAsyncArrowHead_Await is called when production coverCallExpressionAndAsyncArrowHead_Await is exited.
func (s *BaseECMAScriptListener) ExitCoverCallExpressionAndAsyncArrowHead_Await(ctx *CoverCallExpressionAndAsyncArrowHead_AwaitContext) {
}

// EnterCoverCallExpressionAndAsyncArrowHead_Yield_Await is called when production coverCallExpressionAndAsyncArrowHead_Yield_Await is entered.
func (s *BaseECMAScriptListener) EnterCoverCallExpressionAndAsyncArrowHead_Yield_Await(ctx *CoverCallExpressionAndAsyncArrowHead_Yield_AwaitContext) {
}

// ExitCoverCallExpressionAndAsyncArrowHead_Yield_Await is called when production coverCallExpressionAndAsyncArrowHead_Yield_Await is exited.
func (s *BaseECMAScriptListener) ExitCoverCallExpressionAndAsyncArrowHead_Yield_Await(ctx *CoverCallExpressionAndAsyncArrowHead_Yield_AwaitContext) {
}

// EnterScript is called when production script is entered.
func (s *BaseECMAScriptListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseECMAScriptListener) ExitScript(ctx *ScriptContext) {}

// EnterScriptBody is called when production scriptBody is entered.
func (s *BaseECMAScriptListener) EnterScriptBody(ctx *ScriptBodyContext) {}

// ExitScriptBody is called when production scriptBody is exited.
func (s *BaseECMAScriptListener) ExitScriptBody(ctx *ScriptBodyContext) {}

// EnterModule is called when production module is entered.
func (s *BaseECMAScriptListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseECMAScriptListener) ExitModule(ctx *ModuleContext) {}

// EnterModuleBody is called when production moduleBody is entered.
func (s *BaseECMAScriptListener) EnterModuleBody(ctx *ModuleBodyContext) {}

// ExitModuleBody is called when production moduleBody is exited.
func (s *BaseECMAScriptListener) ExitModuleBody(ctx *ModuleBodyContext) {}

// EnterModuleItem is called when production moduleItem is entered.
func (s *BaseECMAScriptListener) EnterModuleItem(ctx *ModuleItemContext) {}

// ExitModuleItem is called when production moduleItem is exited.
func (s *BaseECMAScriptListener) ExitModuleItem(ctx *ModuleItemContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseECMAScriptListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseECMAScriptListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterImportClause is called when production importClause is entered.
func (s *BaseECMAScriptListener) EnterImportClause(ctx *ImportClauseContext) {}

// ExitImportClause is called when production importClause is exited.
func (s *BaseECMAScriptListener) ExitImportClause(ctx *ImportClauseContext) {}

// EnterImportedDefaultBinding is called when production importedDefaultBinding is entered.
func (s *BaseECMAScriptListener) EnterImportedDefaultBinding(ctx *ImportedDefaultBindingContext) {}

// ExitImportedDefaultBinding is called when production importedDefaultBinding is exited.
func (s *BaseECMAScriptListener) ExitImportedDefaultBinding(ctx *ImportedDefaultBindingContext) {}

// EnterNameSpaceImport is called when production nameSpaceImport is entered.
func (s *BaseECMAScriptListener) EnterNameSpaceImport(ctx *NameSpaceImportContext) {}

// ExitNameSpaceImport is called when production nameSpaceImport is exited.
func (s *BaseECMAScriptListener) ExitNameSpaceImport(ctx *NameSpaceImportContext) {}

// EnterNamedImports is called when production namedImports is entered.
func (s *BaseECMAScriptListener) EnterNamedImports(ctx *NamedImportsContext) {}

// ExitNamedImports is called when production namedImports is exited.
func (s *BaseECMAScriptListener) ExitNamedImports(ctx *NamedImportsContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseECMAScriptListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseECMAScriptListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterImportsList is called when production importsList is entered.
func (s *BaseECMAScriptListener) EnterImportsList(ctx *ImportsListContext) {}

// ExitImportsList is called when production importsList is exited.
func (s *BaseECMAScriptListener) ExitImportsList(ctx *ImportsListContext) {}

// EnterImportSpecifier is called when production importSpecifier is entered.
func (s *BaseECMAScriptListener) EnterImportSpecifier(ctx *ImportSpecifierContext) {}

// ExitImportSpecifier is called when production importSpecifier is exited.
func (s *BaseECMAScriptListener) ExitImportSpecifier(ctx *ImportSpecifierContext) {}

// EnterModuleSpecifier is called when production moduleSpecifier is entered.
func (s *BaseECMAScriptListener) EnterModuleSpecifier(ctx *ModuleSpecifierContext) {}

// ExitModuleSpecifier is called when production moduleSpecifier is exited.
func (s *BaseECMAScriptListener) ExitModuleSpecifier(ctx *ModuleSpecifierContext) {}

// EnterImportedBinding is called when production importedBinding is entered.
func (s *BaseECMAScriptListener) EnterImportedBinding(ctx *ImportedBindingContext) {}

// ExitImportedBinding is called when production importedBinding is exited.
func (s *BaseECMAScriptListener) ExitImportedBinding(ctx *ImportedBindingContext) {}

// EnterExportDeclaration is called when production exportDeclaration is entered.
func (s *BaseECMAScriptListener) EnterExportDeclaration(ctx *ExportDeclarationContext) {}

// ExitExportDeclaration is called when production exportDeclaration is exited.
func (s *BaseECMAScriptListener) ExitExportDeclaration(ctx *ExportDeclarationContext) {}

// EnterExportClause is called when production exportClause is entered.
func (s *BaseECMAScriptListener) EnterExportClause(ctx *ExportClauseContext) {}

// ExitExportClause is called when production exportClause is exited.
func (s *BaseECMAScriptListener) ExitExportClause(ctx *ExportClauseContext) {}

// EnterExportsList is called when production exportsList is entered.
func (s *BaseECMAScriptListener) EnterExportsList(ctx *ExportsListContext) {}

// ExitExportsList is called when production exportsList is exited.
func (s *BaseECMAScriptListener) ExitExportsList(ctx *ExportsListContext) {}

// EnterExportSpecifier is called when production exportSpecifier is entered.
func (s *BaseECMAScriptListener) EnterExportSpecifier(ctx *ExportSpecifierContext) {}

// ExitExportSpecifier is called when production exportSpecifier is exited.
func (s *BaseECMAScriptListener) ExitExportSpecifier(ctx *ExportSpecifierContext) {}

// EnterAsyncConciseBody is called when production asyncConciseBody is entered.
func (s *BaseECMAScriptListener) EnterAsyncConciseBody(ctx *AsyncConciseBodyContext) {}

// ExitAsyncConciseBody is called when production asyncConciseBody is exited.
func (s *BaseECMAScriptListener) ExitAsyncConciseBody(ctx *AsyncConciseBodyContext) {}

// EnterAsyncConciseBody_In is called when production asyncConciseBody_In is entered.
func (s *BaseECMAScriptListener) EnterAsyncConciseBody_In(ctx *AsyncConciseBody_InContext) {}

// ExitAsyncConciseBody_In is called when production asyncConciseBody_In is exited.
func (s *BaseECMAScriptListener) ExitAsyncConciseBody_In(ctx *AsyncConciseBody_InContext) {}
