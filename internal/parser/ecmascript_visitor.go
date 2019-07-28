// Code generated from ECMAScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ECMAScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ECMAScriptParser.
type ECMAScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ECMAScriptParser#inputElementDiv.
	VisitInputElementDiv(ctx *InputElementDivContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#inputElementRegExp.
	VisitInputElementRegExp(ctx *InputElementRegExpContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#inputElementRegExpOrTemplateTail.
	VisitInputElementRegExpOrTemplateTail(ctx *InputElementRegExpOrTemplateTailContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#inputElementTemplateTail.
	VisitInputElementTemplateTail(ctx *InputElementTemplateTailContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#regularExpressionLiteral.
	VisitRegularExpressionLiteral(ctx *RegularExpressionLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#identifierReference.
	VisitIdentifierReference(ctx *IdentifierReferenceContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#identifierReference_Yield.
	VisitIdentifierReference_Yield(ctx *IdentifierReference_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#identifierReference_Await.
	VisitIdentifierReference_Await(ctx *IdentifierReference_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#identifierReference_Yield_Await.
	VisitIdentifierReference_Yield_Await(ctx *IdentifierReference_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingIdentifier.
	VisitBindingIdentifier(ctx *BindingIdentifierContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingIdentifier_Yield.
	VisitBindingIdentifier_Yield(ctx *BindingIdentifier_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingIdentifier_Await.
	VisitBindingIdentifier_Await(ctx *BindingIdentifier_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingIdentifier_Yield_Await.
	VisitBindingIdentifier_Yield_Await(ctx *BindingIdentifier_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelIdentifier.
	VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelIdentifier_Yield.
	VisitLabelIdentifier_Yield(ctx *LabelIdentifier_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelIdentifier_Await.
	VisitLabelIdentifier_Await(ctx *LabelIdentifier_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelIdentifier_Yield_Await.
	VisitLabelIdentifier_Yield_Await(ctx *LabelIdentifier_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#primaryExpression_Yield.
	VisitPrimaryExpression_Yield(ctx *PrimaryExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#primaryExpression_Await.
	VisitPrimaryExpression_Await(ctx *PrimaryExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#primaryExpression_Yield_Await.
	VisitPrimaryExpression_Yield_Await(ctx *PrimaryExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverParenthesizedExpressionAndArrowParameterList.
	VisitCoverParenthesizedExpressionAndArrowParameterList(ctx *CoverParenthesizedExpressionAndArrowParameterListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverParenthesizedExpressionAndArrowParameterList_Yield.
	VisitCoverParenthesizedExpressionAndArrowParameterList_Yield(ctx *CoverParenthesizedExpressionAndArrowParameterList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverParenthesizedExpressionAndArrowParameterList_Await.
	VisitCoverParenthesizedExpressionAndArrowParameterList_Await(ctx *CoverParenthesizedExpressionAndArrowParameterList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverParenthesizedExpressionAndArrowParameterList_Yield_Await.
	VisitCoverParenthesizedExpressionAndArrowParameterList_Yield_Await(ctx *CoverParenthesizedExpressionAndArrowParameterList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayLiteral_Yield.
	VisitArrayLiteral_Yield(ctx *ArrayLiteral_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayLiteral_Await.
	VisitArrayLiteral_Await(ctx *ArrayLiteral_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayLiteral_Yield_Await.
	VisitArrayLiteral_Yield_Await(ctx *ArrayLiteral_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elementList_Yield.
	VisitElementList_Yield(ctx *ElementList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elementList_Await.
	VisitElementList_Await(ctx *ElementList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elementList_Yield_Await.
	VisitElementList_Yield_Await(ctx *ElementList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elision.
	VisitElision(ctx *ElisionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#spreadElement.
	VisitSpreadElement(ctx *SpreadElementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#spreadElement_Yield.
	VisitSpreadElement_Yield(ctx *SpreadElement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#spreadElement_Await.
	VisitSpreadElement_Await(ctx *SpreadElement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#spreadElement_Yield_Await.
	VisitSpreadElement_Yield_Await(ctx *SpreadElement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectLiteral_Yield.
	VisitObjectLiteral_Yield(ctx *ObjectLiteral_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectLiteral_Await.
	VisitObjectLiteral_Await(ctx *ObjectLiteral_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectLiteral_Yield_Await.
	VisitObjectLiteral_Yield_Await(ctx *ObjectLiteral_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinitionList.
	VisitPropertyDefinitionList(ctx *PropertyDefinitionListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinitionList_Yield.
	VisitPropertyDefinitionList_Yield(ctx *PropertyDefinitionList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinitionList_Await.
	VisitPropertyDefinitionList_Await(ctx *PropertyDefinitionList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinitionList_Yield_Await.
	VisitPropertyDefinitionList_Yield_Await(ctx *PropertyDefinitionList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinition.
	VisitPropertyDefinition(ctx *PropertyDefinitionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinition_Yield.
	VisitPropertyDefinition_Yield(ctx *PropertyDefinition_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinition_Await.
	VisitPropertyDefinition_Await(ctx *PropertyDefinition_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyDefinition_Yield_Await.
	VisitPropertyDefinition_Yield_Await(ctx *PropertyDefinition_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyName_Yield.
	VisitPropertyName_Yield(ctx *PropertyName_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyName_Await.
	VisitPropertyName_Await(ctx *PropertyName_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyName_Yield_Await.
	VisitPropertyName_Yield_Await(ctx *PropertyName_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#literalPropertyName.
	VisitLiteralPropertyName(ctx *LiteralPropertyNameContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#computedPropertyName.
	VisitComputedPropertyName(ctx *ComputedPropertyNameContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#computedPropertyName_Yield.
	VisitComputedPropertyName_Yield(ctx *ComputedPropertyName_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#computedPropertyName_Await.
	VisitComputedPropertyName_Await(ctx *ComputedPropertyName_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#computedPropertyName_Yield_Await.
	VisitComputedPropertyName_Yield_Await(ctx *ComputedPropertyName_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverInitializedName.
	VisitCoverInitializedName(ctx *CoverInitializedNameContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverInitializedName_Yield.
	VisitCoverInitializedName_Yield(ctx *CoverInitializedName_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverInitializedName_Await.
	VisitCoverInitializedName_Await(ctx *CoverInitializedName_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverInitializedName_Yield_Await.
	VisitCoverInitializedName_Yield_Await(ctx *CoverInitializedName_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_In.
	VisitInitializer_In(ctx *Initializer_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_Yield.
	VisitInitializer_Yield(ctx *Initializer_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_In_Yield.
	VisitInitializer_In_Yield(ctx *Initializer_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_Await.
	VisitInitializer_Await(ctx *Initializer_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_In_Await.
	VisitInitializer_In_Await(ctx *Initializer_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_Yield_Await.
	VisitInitializer_Yield_Await(ctx *Initializer_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initializer_In_Yield_Await.
	VisitInitializer_In_Yield_Await(ctx *Initializer_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral.
	VisitTemplateLiteral(ctx *TemplateLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Yield.
	VisitTemplateLiteral_Yield(ctx *TemplateLiteral_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Await.
	VisitTemplateLiteral_Await(ctx *TemplateLiteral_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Yield_Await.
	VisitTemplateLiteral_Yield_Await(ctx *TemplateLiteral_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Tagged.
	VisitTemplateLiteral_Tagged(ctx *TemplateLiteral_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Yield_Tagged.
	VisitTemplateLiteral_Yield_Tagged(ctx *TemplateLiteral_Yield_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Await_Tagged.
	VisitTemplateLiteral_Await_Tagged(ctx *TemplateLiteral_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateLiteral_Yield_Await_Tagged.
	VisitTemplateLiteral_Yield_Await_Tagged(ctx *TemplateLiteral_Yield_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate.
	VisitSubstitutionTemplate(ctx *SubstitutionTemplateContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Yield.
	VisitSubstitutionTemplate_Yield(ctx *SubstitutionTemplate_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Await.
	VisitSubstitutionTemplate_Await(ctx *SubstitutionTemplate_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Yield_Await.
	VisitSubstitutionTemplate_Yield_Await(ctx *SubstitutionTemplate_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Tagged.
	VisitSubstitutionTemplate_Tagged(ctx *SubstitutionTemplate_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Yield_Tagged.
	VisitSubstitutionTemplate_Yield_Tagged(ctx *SubstitutionTemplate_Yield_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Await_Tagged.
	VisitSubstitutionTemplate_Await_Tagged(ctx *SubstitutionTemplate_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#substitutionTemplate_Yield_Await_Tagged.
	VisitSubstitutionTemplate_Yield_Await_Tagged(ctx *SubstitutionTemplate_Yield_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans.
	VisitTemplateSpans(ctx *TemplateSpansContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Yield.
	VisitTemplateSpans_Yield(ctx *TemplateSpans_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Await.
	VisitTemplateSpans_Await(ctx *TemplateSpans_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Yield_Await.
	VisitTemplateSpans_Yield_Await(ctx *TemplateSpans_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Tagged.
	VisitTemplateSpans_Tagged(ctx *TemplateSpans_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Yield_Tagged.
	VisitTemplateSpans_Yield_Tagged(ctx *TemplateSpans_Yield_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Await_Tagged.
	VisitTemplateSpans_Await_Tagged(ctx *TemplateSpans_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateSpans_Yield_Await_Tagged.
	VisitTemplateSpans_Yield_Await_Tagged(ctx *TemplateSpans_Yield_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList.
	VisitTemplateMiddleList(ctx *TemplateMiddleListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Yield.
	VisitTemplateMiddleList_Yield(ctx *TemplateMiddleList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Await.
	VisitTemplateMiddleList_Await(ctx *TemplateMiddleList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Yield_Await.
	VisitTemplateMiddleList_Yield_Await(ctx *TemplateMiddleList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Tagged.
	VisitTemplateMiddleList_Tagged(ctx *TemplateMiddleList_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Yield_Tagged.
	VisitTemplateMiddleList_Yield_Tagged(ctx *TemplateMiddleList_Yield_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Await_Tagged.
	VisitTemplateMiddleList_Await_Tagged(ctx *TemplateMiddleList_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#templateMiddleList_Yield_Await_Tagged.
	VisitTemplateMiddleList_Yield_Await_Tagged(ctx *TemplateMiddleList_Yield_Await_TaggedContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#memberExpression.
	VisitMemberExpression(ctx *MemberExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#memberExpression_Yield.
	VisitMemberExpression_Yield(ctx *MemberExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#memberExpression_Await.
	VisitMemberExpression_Await(ctx *MemberExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#memberExpression_Yield_Await.
	VisitMemberExpression_Yield_Await(ctx *MemberExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superProperty.
	VisitSuperProperty(ctx *SuperPropertyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superProperty_Yield.
	VisitSuperProperty_Yield(ctx *SuperProperty_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superProperty_Await.
	VisitSuperProperty_Await(ctx *SuperProperty_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superProperty_Yield_Await.
	VisitSuperProperty_Yield_Await(ctx *SuperProperty_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#metaProperty.
	VisitMetaProperty(ctx *MetaPropertyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#newTarget.
	VisitNewTarget(ctx *NewTargetContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#newExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#newExpression_Yield.
	VisitNewExpression_Yield(ctx *NewExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#newExpression_Await.
	VisitNewExpression_Await(ctx *NewExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#newExpression_Yield_Await.
	VisitNewExpression_Yield_Await(ctx *NewExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#callExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#callExpression_Yield.
	VisitCallExpression_Yield(ctx *CallExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#callExpression_Await.
	VisitCallExpression_Await(ctx *CallExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#callExpression_Yield_Await.
	VisitCallExpression_Yield_Await(ctx *CallExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superCall.
	VisitSuperCall(ctx *SuperCallContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superCall_Yield.
	VisitSuperCall_Yield(ctx *SuperCall_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superCall_Await.
	VisitSuperCall_Await(ctx *SuperCall_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#superCall_Yield_Await.
	VisitSuperCall_Yield_Await(ctx *SuperCall_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arguments_Yield.
	VisitArguments_Yield(ctx *Arguments_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arguments_Await.
	VisitArguments_Await(ctx *Arguments_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arguments_Yield_Await.
	VisitArguments_Yield_Await(ctx *Arguments_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#argumentList_Yield.
	VisitArgumentList_Yield(ctx *ArgumentList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#argumentList_Await.
	VisitArgumentList_Await(ctx *ArgumentList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#argumentList_Yield_Await.
	VisitArgumentList_Yield_Await(ctx *ArgumentList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#leftHandSideExpression.
	VisitLeftHandSideExpression(ctx *LeftHandSideExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#leftHandSideExpression_Yield.
	VisitLeftHandSideExpression_Yield(ctx *LeftHandSideExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#leftHandSideExpression_Await.
	VisitLeftHandSideExpression_Await(ctx *LeftHandSideExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#leftHandSideExpression_Yield_Await.
	VisitLeftHandSideExpression_Yield_Await(ctx *LeftHandSideExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#updateExpression.
	VisitUpdateExpression(ctx *UpdateExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#updateExpression_Yield.
	VisitUpdateExpression_Yield(ctx *UpdateExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#updateExpression_Await.
	VisitUpdateExpression_Await(ctx *UpdateExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#updateExpression_Yield_Await.
	VisitUpdateExpression_Yield_Await(ctx *UpdateExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#unaryExpression_Yield.
	VisitUnaryExpression_Yield(ctx *UnaryExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#unaryExpression_Await.
	VisitUnaryExpression_Await(ctx *UnaryExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#unaryExpression_Yield_Await.
	VisitUnaryExpression_Yield_Await(ctx *UnaryExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exponentationExpression.
	VisitExponentationExpression(ctx *ExponentationExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exponentationExpression_Yield.
	VisitExponentationExpression_Yield(ctx *ExponentationExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exponentationExpression_Await.
	VisitExponentationExpression_Await(ctx *ExponentationExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exponentationExpression_Yield_Await.
	VisitExponentationExpression_Yield_Await(ctx *ExponentationExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#multiplicativeExpression_Yield.
	VisitMultiplicativeExpression_Yield(ctx *MultiplicativeExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#multiplicativeExpression_Await.
	VisitMultiplicativeExpression_Await(ctx *MultiplicativeExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#multiplicativeExpression_Yield_Await.
	VisitMultiplicativeExpression_Yield_Await(ctx *MultiplicativeExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#additiveExpression_Yield.
	VisitAdditiveExpression_Yield(ctx *AdditiveExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#additiveExpression_Await.
	VisitAdditiveExpression_Await(ctx *AdditiveExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#additiveExpression_Yield_Await.
	VisitAdditiveExpression_Yield_Await(ctx *AdditiveExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#shiftExpression_Yield.
	VisitShiftExpression_Yield(ctx *ShiftExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#shiftExpression_Await.
	VisitShiftExpression_Await(ctx *ShiftExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#shiftExpression_Yield_Await.
	VisitShiftExpression_Yield_Await(ctx *ShiftExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_In.
	VisitRelationalExpression_In(ctx *RelationalExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_Yield.
	VisitRelationalExpression_Yield(ctx *RelationalExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_In_Yield.
	VisitRelationalExpression_In_Yield(ctx *RelationalExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_Await.
	VisitRelationalExpression_Await(ctx *RelationalExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_In_Await.
	VisitRelationalExpression_In_Await(ctx *RelationalExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_Yield_Await.
	VisitRelationalExpression_Yield_Await(ctx *RelationalExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#relationalExpression_In_Yield_Await.
	VisitRelationalExpression_In_Yield_Await(ctx *RelationalExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_In.
	VisitEqualityExpression_In(ctx *EqualityExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_Yield.
	VisitEqualityExpression_Yield(ctx *EqualityExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_In_Yield.
	VisitEqualityExpression_In_Yield(ctx *EqualityExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_Await.
	VisitEqualityExpression_Await(ctx *EqualityExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_In_Await.
	VisitEqualityExpression_In_Await(ctx *EqualityExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_Yield_Await.
	VisitEqualityExpression_Yield_Await(ctx *EqualityExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#equalityExpression_In_Yield_Await.
	VisitEqualityExpression_In_Yield_Await(ctx *EqualityExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression.
	VisitBitwiseANDExpression(ctx *BitwiseANDExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_In.
	VisitBitwiseANDExpression_In(ctx *BitwiseANDExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_Yield.
	VisitBitwiseANDExpression_Yield(ctx *BitwiseANDExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_In_Yield.
	VisitBitwiseANDExpression_In_Yield(ctx *BitwiseANDExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_Await.
	VisitBitwiseANDExpression_Await(ctx *BitwiseANDExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_In_Await.
	VisitBitwiseANDExpression_In_Await(ctx *BitwiseANDExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_Yield_Await.
	VisitBitwiseANDExpression_Yield_Await(ctx *BitwiseANDExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseANDExpression_In_Yield_Await.
	VisitBitwiseANDExpression_In_Yield_Await(ctx *BitwiseANDExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression.
	VisitBitwiseXORExpression(ctx *BitwiseXORExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_In.
	VisitBitwiseXORExpression_In(ctx *BitwiseXORExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_Yield.
	VisitBitwiseXORExpression_Yield(ctx *BitwiseXORExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_In_Yield.
	VisitBitwiseXORExpression_In_Yield(ctx *BitwiseXORExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_Await.
	VisitBitwiseXORExpression_Await(ctx *BitwiseXORExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_In_Await.
	VisitBitwiseXORExpression_In_Await(ctx *BitwiseXORExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_Yield_Await.
	VisitBitwiseXORExpression_Yield_Await(ctx *BitwiseXORExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseXORExpression_In_Yield_Await.
	VisitBitwiseXORExpression_In_Yield_Await(ctx *BitwiseXORExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression.
	VisitBitwiseORExpression(ctx *BitwiseORExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_In.
	VisitBitwiseORExpression_In(ctx *BitwiseORExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_Yield.
	VisitBitwiseORExpression_Yield(ctx *BitwiseORExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_In_Yield.
	VisitBitwiseORExpression_In_Yield(ctx *BitwiseORExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_Await.
	VisitBitwiseORExpression_Await(ctx *BitwiseORExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_In_Await.
	VisitBitwiseORExpression_In_Await(ctx *BitwiseORExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_Yield_Await.
	VisitBitwiseORExpression_Yield_Await(ctx *BitwiseORExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bitwiseORExpression_In_Yield_Await.
	VisitBitwiseORExpression_In_Yield_Await(ctx *BitwiseORExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression.
	VisitLogicalANDExpression(ctx *LogicalANDExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_In.
	VisitLogicalANDExpression_In(ctx *LogicalANDExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_Yield.
	VisitLogicalANDExpression_Yield(ctx *LogicalANDExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_In_Yield.
	VisitLogicalANDExpression_In_Yield(ctx *LogicalANDExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_Await.
	VisitLogicalANDExpression_Await(ctx *LogicalANDExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_In_Await.
	VisitLogicalANDExpression_In_Await(ctx *LogicalANDExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_Yield_Await.
	VisitLogicalANDExpression_Yield_Await(ctx *LogicalANDExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalANDExpression_In_Yield_Await.
	VisitLogicalANDExpression_In_Yield_Await(ctx *LogicalANDExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression.
	VisitLogicalORExpression(ctx *LogicalORExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_In.
	VisitLogicalORExpression_In(ctx *LogicalORExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_Yield.
	VisitLogicalORExpression_Yield(ctx *LogicalORExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_In_Yield.
	VisitLogicalORExpression_In_Yield(ctx *LogicalORExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_Await.
	VisitLogicalORExpression_Await(ctx *LogicalORExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_In_Await.
	VisitLogicalORExpression_In_Await(ctx *LogicalORExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_Yield_Await.
	VisitLogicalORExpression_Yield_Await(ctx *LogicalORExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#logicalORExpression_In_Yield_Await.
	VisitLogicalORExpression_In_Yield_Await(ctx *LogicalORExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_In.
	VisitConditionalExpression_In(ctx *ConditionalExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_Yield.
	VisitConditionalExpression_Yield(ctx *ConditionalExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_In_Yield.
	VisitConditionalExpression_In_Yield(ctx *ConditionalExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_Await.
	VisitConditionalExpression_Await(ctx *ConditionalExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_In_Await.
	VisitConditionalExpression_In_Await(ctx *ConditionalExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_Yield_Await.
	VisitConditionalExpression_Yield_Await(ctx *ConditionalExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conditionalExpression_In_Yield_Await.
	VisitConditionalExpression_In_Yield_Await(ctx *ConditionalExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_In.
	VisitAssignmentExpression_In(ctx *AssignmentExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_Yield.
	VisitAssignmentExpression_Yield(ctx *AssignmentExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_In_Yield.
	VisitAssignmentExpression_In_Yield(ctx *AssignmentExpression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_Await.
	VisitAssignmentExpression_Await(ctx *AssignmentExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_In_Await.
	VisitAssignmentExpression_In_Await(ctx *AssignmentExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_Yield_Await.
	VisitAssignmentExpression_Yield_Await(ctx *AssignmentExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentExpression_In_Yield_Await.
	VisitAssignmentExpression_In_Yield_Await(ctx *AssignmentExpression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_In.
	VisitExpression_In(ctx *Expression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_Yield.
	VisitExpression_Yield(ctx *Expression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_In_Yield.
	VisitExpression_In_Yield(ctx *Expression_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_Await.
	VisitExpression_Await(ctx *Expression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_In_Await.
	VisitExpression_In_Await(ctx *Expression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_Yield_Await.
	VisitExpression_Yield_Await(ctx *Expression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expression_In_Yield_Await.
	VisitExpression_In_Yield_Await(ctx *Expression_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Yield.
	VisitStatement_Yield(ctx *Statement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Await.
	VisitStatement_Await(ctx *Statement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Yield_Await.
	VisitStatement_Yield_Await(ctx *Statement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Return.
	VisitStatement_Return(ctx *Statement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Yield_Return.
	VisitStatement_Yield_Return(ctx *Statement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Await_Return.
	VisitStatement_Await_Return(ctx *Statement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement_Yield_Await_Return.
	VisitStatement_Yield_Await_Return(ctx *Statement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#declaration_Yield.
	VisitDeclaration_Yield(ctx *Declaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#declaration_Await.
	VisitDeclaration_Await(ctx *Declaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#declaration_Yield_Await.
	VisitDeclaration_Yield_Await(ctx *Declaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration.
	VisitHoistableDeclaration(ctx *HoistableDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Yield.
	VisitHoistableDeclaration_Yield(ctx *HoistableDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Await.
	VisitHoistableDeclaration_Await(ctx *HoistableDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Yield_Await.
	VisitHoistableDeclaration_Yield_Await(ctx *HoistableDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Default.
	VisitHoistableDeclaration_Default(ctx *HoistableDeclaration_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Yield_Default.
	VisitHoistableDeclaration_Yield_Default(ctx *HoistableDeclaration_Yield_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Await_Default.
	VisitHoistableDeclaration_Await_Default(ctx *HoistableDeclaration_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#hoistableDeclaration_Yield_Await_Default.
	VisitHoistableDeclaration_Yield_Await_Default(ctx *HoistableDeclaration_Yield_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement.
	VisitBreakableStatement(ctx *BreakableStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Yield.
	VisitBreakableStatement_Yield(ctx *BreakableStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Await.
	VisitBreakableStatement_Await(ctx *BreakableStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Yield_Await.
	VisitBreakableStatement_Yield_Await(ctx *BreakableStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Return.
	VisitBreakableStatement_Return(ctx *BreakableStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Yield_Return.
	VisitBreakableStatement_Yield_Return(ctx *BreakableStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Await_Return.
	VisitBreakableStatement_Await_Return(ctx *BreakableStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakableStatement_Yield_Await_Return.
	VisitBreakableStatement_Yield_Await_Return(ctx *BreakableStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Yield.
	VisitBlockStatement_Yield(ctx *BlockStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Await.
	VisitBlockStatement_Await(ctx *BlockStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Yield_Await.
	VisitBlockStatement_Yield_Await(ctx *BlockStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Return.
	VisitBlockStatement_Return(ctx *BlockStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Yield_Return.
	VisitBlockStatement_Yield_Return(ctx *BlockStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Await_Return.
	VisitBlockStatement_Await_Return(ctx *BlockStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#blockStatement_Yield_Await_Return.
	VisitBlockStatement_Yield_Await_Return(ctx *BlockStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Yield.
	VisitBlock_Yield(ctx *Block_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Await.
	VisitBlock_Await(ctx *Block_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Yield_Await.
	VisitBlock_Yield_Await(ctx *Block_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Return.
	VisitBlock_Return(ctx *Block_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Yield_Return.
	VisitBlock_Yield_Return(ctx *Block_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Await_Return.
	VisitBlock_Await_Return(ctx *Block_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block_Yield_Await_Return.
	VisitBlock_Yield_Await_Return(ctx *Block_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Yield.
	VisitStatementList_Yield(ctx *StatementList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Await.
	VisitStatementList_Await(ctx *StatementList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Yield_Await.
	VisitStatementList_Yield_Await(ctx *StatementList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Return.
	VisitStatementList_Return(ctx *StatementList_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Yield_Return.
	VisitStatementList_Yield_Return(ctx *StatementList_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Await_Return.
	VisitStatementList_Await_Return(ctx *StatementList_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList_Yield_Await_Return.
	VisitStatementList_Yield_Await_Return(ctx *StatementList_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem.
	VisitStatementListItem(ctx *StatementListItemContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Yield.
	VisitStatementListItem_Yield(ctx *StatementListItem_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Await.
	VisitStatementListItem_Await(ctx *StatementListItem_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Yield_Await.
	VisitStatementListItem_Yield_Await(ctx *StatementListItem_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Return.
	VisitStatementListItem_Return(ctx *StatementListItem_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Yield_Return.
	VisitStatementListItem_Yield_Return(ctx *StatementListItem_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Await_Return.
	VisitStatementListItem_Await_Return(ctx *StatementListItem_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementListItem_Yield_Await_Return.
	VisitStatementListItem_Yield_Await_Return(ctx *StatementListItem_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration.
	VisitLexicalDeclaration(ctx *LexicalDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_In.
	VisitLexicalDeclaration_In(ctx *LexicalDeclaration_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_Yield.
	VisitLexicalDeclaration_Yield(ctx *LexicalDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_In_Yield.
	VisitLexicalDeclaration_In_Yield(ctx *LexicalDeclaration_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_Await.
	VisitLexicalDeclaration_Await(ctx *LexicalDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_In_Await.
	VisitLexicalDeclaration_In_Await(ctx *LexicalDeclaration_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_Yield_Await.
	VisitLexicalDeclaration_Yield_Await(ctx *LexicalDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalDeclaration_In_Yield_Await.
	VisitLexicalDeclaration_In_Yield_Await(ctx *LexicalDeclaration_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#letOrConst.
	VisitLetOrConst(ctx *LetOrConstContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList.
	VisitBindingList(ctx *BindingListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_In.
	VisitBindingList_In(ctx *BindingList_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_Yield.
	VisitBindingList_Yield(ctx *BindingList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_In_Yield.
	VisitBindingList_In_Yield(ctx *BindingList_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_Await.
	VisitBindingList_Await(ctx *BindingList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_In_Await.
	VisitBindingList_In_Await(ctx *BindingList_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_Yield_Await.
	VisitBindingList_Yield_Await(ctx *BindingList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingList_In_Yield_Await.
	VisitBindingList_In_Yield_Await(ctx *BindingList_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding.
	VisitLexicalBinding(ctx *LexicalBindingContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_In.
	VisitLexicalBinding_In(ctx *LexicalBinding_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_Yield.
	VisitLexicalBinding_Yield(ctx *LexicalBinding_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_In_Yield.
	VisitLexicalBinding_In_Yield(ctx *LexicalBinding_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_Await.
	VisitLexicalBinding_Await(ctx *LexicalBinding_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_In_Await.
	VisitLexicalBinding_In_Await(ctx *LexicalBinding_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_Yield_Await.
	VisitLexicalBinding_Yield_Await(ctx *LexicalBinding_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#lexicalBinding_In_Yield_Await.
	VisitLexicalBinding_In_Yield_Await(ctx *LexicalBinding_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableStatement_Yield.
	VisitVariableStatement_Yield(ctx *VariableStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableStatement_Await.
	VisitVariableStatement_Await(ctx *VariableStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableStatement_Yield_Await.
	VisitVariableStatement_Yield_Await(ctx *VariableStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_In.
	VisitVariableDeclarationList_In(ctx *VariableDeclarationList_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_Yield.
	VisitVariableDeclarationList_Yield(ctx *VariableDeclarationList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_In_Yield.
	VisitVariableDeclarationList_In_Yield(ctx *VariableDeclarationList_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_Await.
	VisitVariableDeclarationList_Await(ctx *VariableDeclarationList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_In_Await.
	VisitVariableDeclarationList_In_Await(ctx *VariableDeclarationList_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_Yield_Await.
	VisitVariableDeclarationList_Yield_Await(ctx *VariableDeclarationList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList_In_Yield_Await.
	VisitVariableDeclarationList_In_Yield_Await(ctx *VariableDeclarationList_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_In.
	VisitVariableDeclaration_In(ctx *VariableDeclaration_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_Yield.
	VisitVariableDeclaration_Yield(ctx *VariableDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_In_Yield.
	VisitVariableDeclaration_In_Yield(ctx *VariableDeclaration_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_Await.
	VisitVariableDeclaration_Await(ctx *VariableDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_In_Await.
	VisitVariableDeclaration_In_Await(ctx *VariableDeclaration_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_Yield_Await.
	VisitVariableDeclaration_Yield_Await(ctx *VariableDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration_In_Yield_Await.
	VisitVariableDeclaration_In_Yield_Await(ctx *VariableDeclaration_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPattern.
	VisitBindingPattern(ctx *BindingPatternContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPattern_Yield.
	VisitBindingPattern_Yield(ctx *BindingPattern_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPattern_Await.
	VisitBindingPattern_Await(ctx *BindingPattern_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPattern_Yield_Await.
	VisitBindingPattern_Yield_Await(ctx *BindingPattern_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectBindingPattern.
	VisitObjectBindingPattern(ctx *ObjectBindingPatternContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectBindingPattern_Yield.
	VisitObjectBindingPattern_Yield(ctx *ObjectBindingPattern_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectBindingPattern_Await.
	VisitObjectBindingPattern_Await(ctx *ObjectBindingPattern_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectBindingPattern_Yield_Await.
	VisitObjectBindingPattern_Yield_Await(ctx *ObjectBindingPattern_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayBindingPattern.
	VisitArrayBindingPattern(ctx *ArrayBindingPatternContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayBindingPattern_Yield.
	VisitArrayBindingPattern_Yield(ctx *ArrayBindingPattern_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayBindingPattern_Await.
	VisitArrayBindingPattern_Await(ctx *ArrayBindingPattern_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayBindingPattern_Yield_Await.
	VisitArrayBindingPattern_Yield_Await(ctx *ArrayBindingPattern_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestProperty.
	VisitBindingRestProperty(ctx *BindingRestPropertyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestProperty_Yield.
	VisitBindingRestProperty_Yield(ctx *BindingRestProperty_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestProperty_Await.
	VisitBindingRestProperty_Await(ctx *BindingRestProperty_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestProperty_Yield_Await.
	VisitBindingRestProperty_Yield_Await(ctx *BindingRestProperty_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPropertyList.
	VisitBindingPropertyList(ctx *BindingPropertyListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPropertyList_Yield.
	VisitBindingPropertyList_Yield(ctx *BindingPropertyList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPropertyList_Await.
	VisitBindingPropertyList_Await(ctx *BindingPropertyList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingPropertyList_Yield_Await.
	VisitBindingPropertyList_Yield_Await(ctx *BindingPropertyList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElementList.
	VisitBindingElementList(ctx *BindingElementListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElementList_Yield.
	VisitBindingElementList_Yield(ctx *BindingElementList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElementList_Await.
	VisitBindingElementList_Await(ctx *BindingElementList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElementList_Yield_Await.
	VisitBindingElementList_Yield_Await(ctx *BindingElementList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElisionElement.
	VisitBindingElisionElement(ctx *BindingElisionElementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElisionElement_Yield.
	VisitBindingElisionElement_Yield(ctx *BindingElisionElement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElisionElement_Await.
	VisitBindingElisionElement_Await(ctx *BindingElisionElement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElisionElement_Yield_Await.
	VisitBindingElisionElement_Yield_Await(ctx *BindingElisionElement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingProperty.
	VisitBindingProperty(ctx *BindingPropertyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingProperty_Yield.
	VisitBindingProperty_Yield(ctx *BindingProperty_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingProperty_Await.
	VisitBindingProperty_Await(ctx *BindingProperty_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingProperty_Yield_Await.
	VisitBindingProperty_Yield_Await(ctx *BindingProperty_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElement.
	VisitBindingElement(ctx *BindingElementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElement_Yield.
	VisitBindingElement_Yield(ctx *BindingElement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElement_Await.
	VisitBindingElement_Await(ctx *BindingElement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingElement_Yield_Await.
	VisitBindingElement_Yield_Await(ctx *BindingElement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#singleNameBinding.
	VisitSingleNameBinding(ctx *SingleNameBindingContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#singleNameBinding_Yield.
	VisitSingleNameBinding_Yield(ctx *SingleNameBinding_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#singleNameBinding_Await.
	VisitSingleNameBinding_Await(ctx *SingleNameBinding_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#singleNameBinding_Yield_Await.
	VisitSingleNameBinding_Yield_Await(ctx *SingleNameBinding_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestElement.
	VisitBindingRestElement(ctx *BindingRestElementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestElement_Yield.
	VisitBindingRestElement_Yield(ctx *BindingRestElement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestElement_Await.
	VisitBindingRestElement_Await(ctx *BindingRestElement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#bindingRestElement_Yield_Await.
	VisitBindingRestElement_Yield_Await(ctx *BindingRestElement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expressionStatement_Yield.
	VisitExpressionStatement_Yield(ctx *ExpressionStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expressionStatement_Await.
	VisitExpressionStatement_Await(ctx *ExpressionStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expressionStatement_Yield_Await.
	VisitExpressionStatement_Yield_Await(ctx *ExpressionStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Yield.
	VisitIfStatement_Yield(ctx *IfStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Await.
	VisitIfStatement_Await(ctx *IfStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Yield_Await.
	VisitIfStatement_Yield_Await(ctx *IfStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Return.
	VisitIfStatement_Return(ctx *IfStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Yield_Return.
	VisitIfStatement_Yield_Return(ctx *IfStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Await_Return.
	VisitIfStatement_Await_Return(ctx *IfStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement_Yield_Await_Return.
	VisitIfStatement_Yield_Await_Return(ctx *IfStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement.
	VisitIterationStatement(ctx *IterationStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Yield.
	VisitIterationStatement_Yield(ctx *IterationStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Await.
	VisitIterationStatement_Await(ctx *IterationStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Yield_Await.
	VisitIterationStatement_Yield_Await(ctx *IterationStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Return.
	VisitIterationStatement_Return(ctx *IterationStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Yield_Return.
	VisitIterationStatement_Yield_Return(ctx *IterationStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Await_Return.
	VisitIterationStatement_Await_Return(ctx *IterationStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#iterationStatement_Yield_Await_Return.
	VisitIterationStatement_Yield_Await_Return(ctx *IterationStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forDeclaration.
	VisitForDeclaration(ctx *ForDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forDeclaration_Yield.
	VisitForDeclaration_Yield(ctx *ForDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forDeclaration_Await.
	VisitForDeclaration_Await(ctx *ForDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forDeclaration_Yield_Await.
	VisitForDeclaration_Yield_Await(ctx *ForDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forBinding.
	VisitForBinding(ctx *ForBindingContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forBinding_Yield.
	VisitForBinding_Yield(ctx *ForBinding_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forBinding_Await.
	VisitForBinding_Await(ctx *ForBinding_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#forBinding_Yield_Await.
	VisitForBinding_Yield_Await(ctx *ForBinding_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#continueStatement_Yield.
	VisitContinueStatement_Yield(ctx *ContinueStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#continueStatement_Await.
	VisitContinueStatement_Await(ctx *ContinueStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#continueStatement_Yield_Await.
	VisitContinueStatement_Yield_Await(ctx *ContinueStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakStatement_Yield.
	VisitBreakStatement_Yield(ctx *BreakStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakStatement_Await.
	VisitBreakStatement_Await(ctx *BreakStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakStatement_Yield_Await.
	VisitBreakStatement_Yield_Await(ctx *BreakStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#returnStatement_Yield.
	VisitReturnStatement_Yield(ctx *ReturnStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#returnStatement_Await.
	VisitReturnStatement_Await(ctx *ReturnStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#returnStatement_Yield_Await.
	VisitReturnStatement_Yield_Await(ctx *ReturnStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement.
	VisitWithStatement(ctx *WithStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Yield.
	VisitWithStatement_Yield(ctx *WithStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Await.
	VisitWithStatement_Await(ctx *WithStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Yield_Await.
	VisitWithStatement_Yield_Await(ctx *WithStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Return.
	VisitWithStatement_Return(ctx *WithStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Yield_Return.
	VisitWithStatement_Yield_Return(ctx *WithStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Await_Return.
	VisitWithStatement_Await_Return(ctx *WithStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement_Yield_Await_Return.
	VisitWithStatement_Yield_Await_Return(ctx *WithStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Yield.
	VisitSwitchStatement_Yield(ctx *SwitchStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Await.
	VisitSwitchStatement_Await(ctx *SwitchStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Yield_Await.
	VisitSwitchStatement_Yield_Await(ctx *SwitchStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Return.
	VisitSwitchStatement_Return(ctx *SwitchStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Yield_Return.
	VisitSwitchStatement_Yield_Return(ctx *SwitchStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Await_Return.
	VisitSwitchStatement_Await_Return(ctx *SwitchStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement_Yield_Await_Return.
	VisitSwitchStatement_Yield_Await_Return(ctx *SwitchStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Yield.
	VisitCaseBlock_Yield(ctx *CaseBlock_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Await.
	VisitCaseBlock_Await(ctx *CaseBlock_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Yield_Await.
	VisitCaseBlock_Yield_Await(ctx *CaseBlock_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Return.
	VisitCaseBlock_Return(ctx *CaseBlock_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Yield_Return.
	VisitCaseBlock_Yield_Return(ctx *CaseBlock_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Await_Return.
	VisitCaseBlock_Await_Return(ctx *CaseBlock_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock_Yield_Await_Return.
	VisitCaseBlock_Yield_Await_Return(ctx *CaseBlock_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Yield.
	VisitCaseClause_Yield(ctx *CaseClause_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Await.
	VisitCaseClause_Await(ctx *CaseClause_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Yield_Await.
	VisitCaseClause_Yield_Await(ctx *CaseClause_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Return.
	VisitCaseClause_Return(ctx *CaseClause_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Yield_Return.
	VisitCaseClause_Yield_Return(ctx *CaseClause_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Await_Return.
	VisitCaseClause_Await_Return(ctx *CaseClause_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause_Yield_Await_Return.
	VisitCaseClause_Yield_Await_Return(ctx *CaseClause_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Yield.
	VisitDefaultClause_Yield(ctx *DefaultClause_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Await.
	VisitDefaultClause_Await(ctx *DefaultClause_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Yield_Await.
	VisitDefaultClause_Yield_Await(ctx *DefaultClause_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Return.
	VisitDefaultClause_Return(ctx *DefaultClause_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Yield_Return.
	VisitDefaultClause_Yield_Return(ctx *DefaultClause_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Await_Return.
	VisitDefaultClause_Await_Return(ctx *DefaultClause_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause_Yield_Await_Return.
	VisitDefaultClause_Yield_Await_Return(ctx *DefaultClause_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement.
	VisitLabelledStatement(ctx *LabelledStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Yield.
	VisitLabelledStatement_Yield(ctx *LabelledStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Await.
	VisitLabelledStatement_Await(ctx *LabelledStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Yield_Await.
	VisitLabelledStatement_Yield_Await(ctx *LabelledStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Return.
	VisitLabelledStatement_Return(ctx *LabelledStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Yield_Return.
	VisitLabelledStatement_Yield_Return(ctx *LabelledStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Await_Return.
	VisitLabelledStatement_Await_Return(ctx *LabelledStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement_Yield_Await_Return.
	VisitLabelledStatement_Yield_Await_Return(ctx *LabelledStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem.
	VisitLabelledItem(ctx *LabelledItemContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Yield.
	VisitLabelledItem_Yield(ctx *LabelledItem_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Await.
	VisitLabelledItem_Await(ctx *LabelledItem_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Yield_Await.
	VisitLabelledItem_Yield_Await(ctx *LabelledItem_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Return.
	VisitLabelledItem_Return(ctx *LabelledItem_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Yield_Return.
	VisitLabelledItem_Yield_Return(ctx *LabelledItem_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Await_Return.
	VisitLabelledItem_Await_Return(ctx *LabelledItem_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledItem_Yield_Await_Return.
	VisitLabelledItem_Yield_Await_Return(ctx *LabelledItem_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#throwStatement_Yield.
	VisitThrowStatement_Yield(ctx *ThrowStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#throwStatement_Await.
	VisitThrowStatement_Await(ctx *ThrowStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#throwStatement_Yield_Await.
	VisitThrowStatement_Yield_Await(ctx *ThrowStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Yield.
	VisitTryStatement_Yield(ctx *TryStatement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Await.
	VisitTryStatement_Await(ctx *TryStatement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Yield_Await.
	VisitTryStatement_Yield_Await(ctx *TryStatement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Return.
	VisitTryStatement_Return(ctx *TryStatement_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Yield_Return.
	VisitTryStatement_Yield_Return(ctx *TryStatement_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Await_Return.
	VisitTryStatement_Await_Return(ctx *TryStatement_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement_Yield_Await_Return.
	VisitTryStatement_Yield_Await_Return(ctx *TryStatement_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_.
	VisitCatch_(ctx *Catch_Context) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Yield.
	VisitCatch_Yield(ctx *Catch_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Await.
	VisitCatch_Await(ctx *Catch_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Yield_Await.
	VisitCatch_Yield_Await(ctx *Catch_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Return.
	VisitCatch_Return(ctx *Catch_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Yield_Return.
	VisitCatch_Yield_Return(ctx *Catch_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Await_Return.
	VisitCatch_Await_Return(ctx *Catch_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catch_Yield_Await_Return.
	VisitCatch_Yield_Await_Return(ctx *Catch_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_.
	VisitFinally_(ctx *Finally_Context) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Yield.
	VisitFinally_Yield(ctx *Finally_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Await.
	VisitFinally_Await(ctx *Finally_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Yield_Await.
	VisitFinally_Yield_Await(ctx *Finally_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Return.
	VisitFinally_Return(ctx *Finally_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Yield_Return.
	VisitFinally_Yield_Return(ctx *Finally_Yield_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Await_Return.
	VisitFinally_Await_Return(ctx *Finally_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finally_Yield_Await_Return.
	VisitFinally_Yield_Await_Return(ctx *Finally_Yield_Await_ReturnContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catchParameter.
	VisitCatchParameter(ctx *CatchParameterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catchParameter_Yield.
	VisitCatchParameter_Yield(ctx *CatchParameter_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catchParameter_Await.
	VisitCatchParameter_Await(ctx *CatchParameter_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catchParameter_Yield_Await.
	VisitCatchParameter_Yield_Await(ctx *CatchParameter_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#debuggerStatement.
	VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Yield.
	VisitFunctionDeclaration_Yield(ctx *FunctionDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Await.
	VisitFunctionDeclaration_Await(ctx *FunctionDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Yield_Await.
	VisitFunctionDeclaration_Yield_Await(ctx *FunctionDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Default.
	VisitFunctionDeclaration_Default(ctx *FunctionDeclaration_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Yield_Default.
	VisitFunctionDeclaration_Yield_Default(ctx *FunctionDeclaration_Yield_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Await_Default.
	VisitFunctionDeclaration_Await_Default(ctx *FunctionDeclaration_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration_Yield_Await_Default.
	VisitFunctionDeclaration_Yield_Await_Default(ctx *FunctionDeclaration_Yield_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#uniqueFormalParameters.
	VisitUniqueFormalParameters(ctx *UniqueFormalParametersContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#uniqueFormalParameters_Yield.
	VisitUniqueFormalParameters_Yield(ctx *UniqueFormalParameters_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#uniqueFormalParameters_Await.
	VisitUniqueFormalParameters_Await(ctx *UniqueFormalParameters_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#uniqueFormalParameters_Yield_Await.
	VisitUniqueFormalParameters_Yield_Await(ctx *UniqueFormalParameters_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameters_Yield.
	VisitFormalParameters_Yield(ctx *FormalParameters_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameters_Await.
	VisitFormalParameters_Await(ctx *FormalParameters_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameters_Yield_Await.
	VisitFormalParameters_Yield_Await(ctx *FormalParameters_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameterList_Yield.
	VisitFormalParameterList_Yield(ctx *FormalParameterList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameterList_Await.
	VisitFormalParameterList_Await(ctx *FormalParameterList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameterList_Yield_Await.
	VisitFormalParameterList_Yield_Await(ctx *FormalParameterList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionRestParameter.
	VisitFunctionRestParameter(ctx *FunctionRestParameterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionRestParameter_Yield.
	VisitFunctionRestParameter_Yield(ctx *FunctionRestParameter_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionRestParameter_Await.
	VisitFunctionRestParameter_Await(ctx *FunctionRestParameter_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionRestParameter_Yield_Await.
	VisitFunctionRestParameter_Yield_Await(ctx *FunctionRestParameter_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameter_Yield.
	VisitFormalParameter_Yield(ctx *FormalParameter_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameter_Await.
	VisitFormalParameter_Await(ctx *FormalParameter_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameter_Yield_Await.
	VisitFormalParameter_Yield_Await(ctx *FormalParameter_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionBody_Yield.
	VisitFunctionBody_Yield(ctx *FunctionBody_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionBody_Await.
	VisitFunctionBody_Await(ctx *FunctionBody_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionBody_Yield_Await.
	VisitFunctionBody_Yield_Await(ctx *FunctionBody_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionStatementList.
	VisitFunctionStatementList(ctx *FunctionStatementListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionStatementList_Yield.
	VisitFunctionStatementList_Yield(ctx *FunctionStatementList_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionStatementList_Await.
	VisitFunctionStatementList_Await(ctx *FunctionStatementList_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionStatementList_Yield_Await.
	VisitFunctionStatementList_Yield_Await(ctx *FunctionStatementList_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction.
	VisitArrowFunction(ctx *ArrowFunctionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_In.
	VisitArrowFunction_In(ctx *ArrowFunction_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_Yield.
	VisitArrowFunction_Yield(ctx *ArrowFunction_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_In_Yield.
	VisitArrowFunction_In_Yield(ctx *ArrowFunction_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_Await.
	VisitArrowFunction_Await(ctx *ArrowFunction_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_In_Await.
	VisitArrowFunction_In_Await(ctx *ArrowFunction_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_Yield_Await.
	VisitArrowFunction_Yield_Await(ctx *ArrowFunction_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowFunction_In_Yield_Await.
	VisitArrowFunction_In_Yield_Await(ctx *ArrowFunction_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowParameters.
	VisitArrowParameters(ctx *ArrowParametersContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowParameters_Yield.
	VisitArrowParameters_Yield(ctx *ArrowParameters_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowParameters_Await.
	VisitArrowParameters_Await(ctx *ArrowParameters_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrowParameters_Yield_Await.
	VisitArrowParameters_Yield_Await(ctx *ArrowParameters_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conciseBody.
	VisitConciseBody(ctx *ConciseBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#conciseBody_In.
	VisitConciseBody_In(ctx *ConciseBody_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#methodDefinition.
	VisitMethodDefinition(ctx *MethodDefinitionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#methodDefinition_Yield.
	VisitMethodDefinition_Yield(ctx *MethodDefinition_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#methodDefinition_Await.
	VisitMethodDefinition_Await(ctx *MethodDefinition_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#methodDefinition_Yield_Await.
	VisitMethodDefinition_Yield_Await(ctx *MethodDefinition_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertySetParameterList.
	VisitPropertySetParameterList(ctx *PropertySetParameterListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorMethod.
	VisitGeneratorMethod(ctx *GeneratorMethodContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorMethod_Yield.
	VisitGeneratorMethod_Yield(ctx *GeneratorMethod_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorMethod_Await.
	VisitGeneratorMethod_Await(ctx *GeneratorMethod_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorMethod_Yield_Await.
	VisitGeneratorMethod_Yield_Await(ctx *GeneratorMethod_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration.
	VisitGeneratorDeclaration(ctx *GeneratorDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Yield.
	VisitGeneratorDeclaration_Yield(ctx *GeneratorDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Await.
	VisitGeneratorDeclaration_Await(ctx *GeneratorDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Yield_Await.
	VisitGeneratorDeclaration_Yield_Await(ctx *GeneratorDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Default.
	VisitGeneratorDeclaration_Default(ctx *GeneratorDeclaration_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Yield_Default.
	VisitGeneratorDeclaration_Yield_Default(ctx *GeneratorDeclaration_Yield_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Await_Default.
	VisitGeneratorDeclaration_Await_Default(ctx *GeneratorDeclaration_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorDeclaration_Yield_Await_Default.
	VisitGeneratorDeclaration_Yield_Await_Default(ctx *GeneratorDeclaration_Yield_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorExpression.
	VisitGeneratorExpression(ctx *GeneratorExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#generatorBody.
	VisitGeneratorBody(ctx *GeneratorBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#yieldExpression.
	VisitYieldExpression(ctx *YieldExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#yieldExpression_In.
	VisitYieldExpression_In(ctx *YieldExpression_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#yieldExpression_Await.
	VisitYieldExpression_Await(ctx *YieldExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#yieldExpression_In_Await.
	VisitYieldExpression_In_Await(ctx *YieldExpression_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorMethod.
	VisitAsyncGeneratorMethod(ctx *AsyncGeneratorMethodContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorMethod_Yield.
	VisitAsyncGeneratorMethod_Yield(ctx *AsyncGeneratorMethod_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorMethod_Await.
	VisitAsyncGeneratorMethod_Await(ctx *AsyncGeneratorMethod_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorMethod_Yield_Await.
	VisitAsyncGeneratorMethod_Yield_Await(ctx *AsyncGeneratorMethod_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration.
	VisitAsyncGeneratorDeclaration(ctx *AsyncGeneratorDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Yield.
	VisitAsyncGeneratorDeclaration_Yield(ctx *AsyncGeneratorDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Await.
	VisitAsyncGeneratorDeclaration_Await(ctx *AsyncGeneratorDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Yield_Await.
	VisitAsyncGeneratorDeclaration_Yield_Await(ctx *AsyncGeneratorDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Default.
	VisitAsyncGeneratorDeclaration_Default(ctx *AsyncGeneratorDeclaration_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Yield_Default.
	VisitAsyncGeneratorDeclaration_Yield_Default(ctx *AsyncGeneratorDeclaration_Yield_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Await_Default.
	VisitAsyncGeneratorDeclaration_Await_Default(ctx *AsyncGeneratorDeclaration_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorDeclaration_Yield_Await_Default.
	VisitAsyncGeneratorDeclaration_Yield_Await_Default(ctx *AsyncGeneratorDeclaration_Yield_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorExpression.
	VisitAsyncGeneratorExpression(ctx *AsyncGeneratorExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncGeneratorBody.
	VisitAsyncGeneratorBody(ctx *AsyncGeneratorBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Yield.
	VisitClassDeclaration_Yield(ctx *ClassDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Await.
	VisitClassDeclaration_Await(ctx *ClassDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Yield_Await.
	VisitClassDeclaration_Yield_Await(ctx *ClassDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Default.
	VisitClassDeclaration_Default(ctx *ClassDeclaration_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Yield_Default.
	VisitClassDeclaration_Yield_Default(ctx *ClassDeclaration_Yield_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Await_Default.
	VisitClassDeclaration_Await_Default(ctx *ClassDeclaration_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classDeclaration_Yield_Await_Default.
	VisitClassDeclaration_Yield_Await_Default(ctx *ClassDeclaration_Yield_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classExpression.
	VisitClassExpression(ctx *ClassExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classExpression_Yield.
	VisitClassExpression_Yield(ctx *ClassExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classExpression_Await.
	VisitClassExpression_Await(ctx *ClassExpression_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classExpression_Yield_Await.
	VisitClassExpression_Yield_Await(ctx *ClassExpression_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classTail.
	VisitClassTail(ctx *ClassTailContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classTail_Yield.
	VisitClassTail_Yield(ctx *ClassTail_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classTail_Await.
	VisitClassTail_Await(ctx *ClassTail_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classTail_Yield_Await.
	VisitClassTail_Yield_Await(ctx *ClassTail_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classHeritage.
	VisitClassHeritage(ctx *ClassHeritageContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classHeritage_Yield.
	VisitClassHeritage_Yield(ctx *ClassHeritage_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classHeritage_Await.
	VisitClassHeritage_Await(ctx *ClassHeritage_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classHeritage_Yield_Await.
	VisitClassHeritage_Yield_Await(ctx *ClassHeritage_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classBody_Yield.
	VisitClassBody_Yield(ctx *ClassBody_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classBody_Await.
	VisitClassBody_Await(ctx *ClassBody_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classBody_Yield_Await.
	VisitClassBody_Yield_Await(ctx *ClassBody_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classElement.
	VisitClassElement(ctx *ClassElementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classElement_Yield.
	VisitClassElement_Yield(ctx *ClassElement_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classElement_Await.
	VisitClassElement_Await(ctx *ClassElement_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#classElement_Yield_Await.
	VisitClassElement_Yield_Await(ctx *ClassElement_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration.
	VisitAsyncFunctionDeclaration(ctx *AsyncFunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Yield.
	VisitAsyncFunctionDeclaration_Yield(ctx *AsyncFunctionDeclaration_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Await.
	VisitAsyncFunctionDeclaration_Await(ctx *AsyncFunctionDeclaration_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Yield_Await.
	VisitAsyncFunctionDeclaration_Yield_Await(ctx *AsyncFunctionDeclaration_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Default.
	VisitAsyncFunctionDeclaration_Default(ctx *AsyncFunctionDeclaration_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Yield_Default.
	VisitAsyncFunctionDeclaration_Yield_Default(ctx *AsyncFunctionDeclaration_Yield_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Await_Default.
	VisitAsyncFunctionDeclaration_Await_Default(ctx *AsyncFunctionDeclaration_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionDeclaration_Yield_Await_Default.
	VisitAsyncFunctionDeclaration_Yield_Await_Default(ctx *AsyncFunctionDeclaration_Yield_Await_DefaultContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionExpression.
	VisitAsyncFunctionExpression(ctx *AsyncFunctionExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncMethod.
	VisitAsyncMethod(ctx *AsyncMethodContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncMethod_Yield.
	VisitAsyncMethod_Yield(ctx *AsyncMethod_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncMethod_Await.
	VisitAsyncMethod_Await(ctx *AsyncMethod_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncMethod_Yield_Await.
	VisitAsyncMethod_Yield_Await(ctx *AsyncMethod_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncFunctionBody.
	VisitAsyncFunctionBody(ctx *AsyncFunctionBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#awaitExpression.
	VisitAwaitExpression(ctx *AwaitExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#awaitExpression_Yield.
	VisitAwaitExpression_Yield(ctx *AwaitExpression_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction.
	VisitAsyncArrowFunction(ctx *AsyncArrowFunctionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_In.
	VisitAsyncArrowFunction_In(ctx *AsyncArrowFunction_InContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_Yield.
	VisitAsyncArrowFunction_Yield(ctx *AsyncArrowFunction_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_In_Yield.
	VisitAsyncArrowFunction_In_Yield(ctx *AsyncArrowFunction_In_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_Await.
	VisitAsyncArrowFunction_Await(ctx *AsyncArrowFunction_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_In_Await.
	VisitAsyncArrowFunction_In_Await(ctx *AsyncArrowFunction_In_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_Yield_Await.
	VisitAsyncArrowFunction_Yield_Await(ctx *AsyncArrowFunction_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowFunction_In_Yield_Await.
	VisitAsyncArrowFunction_In_Yield_Await(ctx *AsyncArrowFunction_In_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowBindingIdentifier.
	VisitAsyncArrowBindingIdentifier(ctx *AsyncArrowBindingIdentifierContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncArrowBindingIdentifier_Yield.
	VisitAsyncArrowBindingIdentifier_Yield(ctx *AsyncArrowBindingIdentifier_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverCallExpressionAndAsyncArrowHead.
	VisitCoverCallExpressionAndAsyncArrowHead(ctx *CoverCallExpressionAndAsyncArrowHeadContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverCallExpressionAndAsyncArrowHead_Yield.
	VisitCoverCallExpressionAndAsyncArrowHead_Yield(ctx *CoverCallExpressionAndAsyncArrowHead_YieldContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverCallExpressionAndAsyncArrowHead_Await.
	VisitCoverCallExpressionAndAsyncArrowHead_Await(ctx *CoverCallExpressionAndAsyncArrowHead_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#coverCallExpressionAndAsyncArrowHead_Yield_Await.
	VisitCoverCallExpressionAndAsyncArrowHead_Yield_Await(ctx *CoverCallExpressionAndAsyncArrowHead_Yield_AwaitContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#scriptBody.
	VisitScriptBody(ctx *ScriptBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#moduleBody.
	VisitModuleBody(ctx *ModuleBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#moduleItem.
	VisitModuleItem(ctx *ModuleItemContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#importClause.
	VisitImportClause(ctx *ImportClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#importedDefaultBinding.
	VisitImportedDefaultBinding(ctx *ImportedDefaultBindingContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#nameSpaceImport.
	VisitNameSpaceImport(ctx *NameSpaceImportContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#namedImports.
	VisitNamedImports(ctx *NamedImportsContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#importsList.
	VisitImportsList(ctx *ImportsListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#importSpecifier.
	VisitImportSpecifier(ctx *ImportSpecifierContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#moduleSpecifier.
	VisitModuleSpecifier(ctx *ModuleSpecifierContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#importedBinding.
	VisitImportedBinding(ctx *ImportedBindingContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exportDeclaration.
	VisitExportDeclaration(ctx *ExportDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exportClause.
	VisitExportClause(ctx *ExportClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exportsList.
	VisitExportsList(ctx *ExportsListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#exportSpecifier.
	VisitExportSpecifier(ctx *ExportSpecifierContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncConciseBody.
	VisitAsyncConciseBody(ctx *AsyncConciseBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#asyncConciseBody_In.
	VisitAsyncConciseBody_In(ctx *AsyncConciseBody_InContext) interface{}
}
