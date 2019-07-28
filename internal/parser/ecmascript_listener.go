// Code generated from ECMAScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ECMAScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ECMAScriptListener is a complete listener for a parse tree produced by ECMAScriptParser.
type ECMAScriptListener interface {
	antlr.ParseTreeListener

	// EnterInputElementDiv is called when entering the inputElementDiv production.
	EnterInputElementDiv(c *InputElementDivContext)

	// EnterInputElementRegExp is called when entering the inputElementRegExp production.
	EnterInputElementRegExp(c *InputElementRegExpContext)

	// EnterInputElementRegExpOrTemplateTail is called when entering the inputElementRegExpOrTemplateTail production.
	EnterInputElementRegExpOrTemplateTail(c *InputElementRegExpOrTemplateTailContext)

	// EnterInputElementTemplateTail is called when entering the inputElementTemplateTail production.
	EnterInputElementTemplateTail(c *InputElementTemplateTailContext)

	// EnterRegularExpressionLiteral is called when entering the regularExpressionLiteral production.
	EnterRegularExpressionLiteral(c *RegularExpressionLiteralContext)

	// EnterIdentifierReference is called when entering the identifierReference production.
	EnterIdentifierReference(c *IdentifierReferenceContext)

	// EnterIdentifierReference_Yield is called when entering the identifierReference_Yield production.
	EnterIdentifierReference_Yield(c *IdentifierReference_YieldContext)

	// EnterIdentifierReference_Await is called when entering the identifierReference_Await production.
	EnterIdentifierReference_Await(c *IdentifierReference_AwaitContext)

	// EnterIdentifierReference_Yield_Await is called when entering the identifierReference_Yield_Await production.
	EnterIdentifierReference_Yield_Await(c *IdentifierReference_Yield_AwaitContext)

	// EnterBindingIdentifier is called when entering the bindingIdentifier production.
	EnterBindingIdentifier(c *BindingIdentifierContext)

	// EnterBindingIdentifier_Yield is called when entering the bindingIdentifier_Yield production.
	EnterBindingIdentifier_Yield(c *BindingIdentifier_YieldContext)

	// EnterBindingIdentifier_Await is called when entering the bindingIdentifier_Await production.
	EnterBindingIdentifier_Await(c *BindingIdentifier_AwaitContext)

	// EnterBindingIdentifier_Yield_Await is called when entering the bindingIdentifier_Yield_Await production.
	EnterBindingIdentifier_Yield_Await(c *BindingIdentifier_Yield_AwaitContext)

	// EnterLabelIdentifier is called when entering the labelIdentifier production.
	EnterLabelIdentifier(c *LabelIdentifierContext)

	// EnterLabelIdentifier_Yield is called when entering the labelIdentifier_Yield production.
	EnterLabelIdentifier_Yield(c *LabelIdentifier_YieldContext)

	// EnterLabelIdentifier_Await is called when entering the labelIdentifier_Await production.
	EnterLabelIdentifier_Await(c *LabelIdentifier_AwaitContext)

	// EnterLabelIdentifier_Yield_Await is called when entering the labelIdentifier_Yield_Await production.
	EnterLabelIdentifier_Yield_Await(c *LabelIdentifier_Yield_AwaitContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterPrimaryExpression_Yield is called when entering the primaryExpression_Yield production.
	EnterPrimaryExpression_Yield(c *PrimaryExpression_YieldContext)

	// EnterPrimaryExpression_Await is called when entering the primaryExpression_Await production.
	EnterPrimaryExpression_Await(c *PrimaryExpression_AwaitContext)

	// EnterPrimaryExpression_Yield_Await is called when entering the primaryExpression_Yield_Await production.
	EnterPrimaryExpression_Yield_Await(c *PrimaryExpression_Yield_AwaitContext)

	// EnterCoverParenthesizedExpressionAndArrowParameterList is called when entering the coverParenthesizedExpressionAndArrowParameterList production.
	EnterCoverParenthesizedExpressionAndArrowParameterList(c *CoverParenthesizedExpressionAndArrowParameterListContext)

	// EnterCoverParenthesizedExpressionAndArrowParameterList_Yield is called when entering the coverParenthesizedExpressionAndArrowParameterList_Yield production.
	EnterCoverParenthesizedExpressionAndArrowParameterList_Yield(c *CoverParenthesizedExpressionAndArrowParameterList_YieldContext)

	// EnterCoverParenthesizedExpressionAndArrowParameterList_Await is called when entering the coverParenthesizedExpressionAndArrowParameterList_Await production.
	EnterCoverParenthesizedExpressionAndArrowParameterList_Await(c *CoverParenthesizedExpressionAndArrowParameterList_AwaitContext)

	// EnterCoverParenthesizedExpressionAndArrowParameterList_Yield_Await is called when entering the coverParenthesizedExpressionAndArrowParameterList_Yield_Await production.
	EnterCoverParenthesizedExpressionAndArrowParameterList_Yield_Await(c *CoverParenthesizedExpressionAndArrowParameterList_Yield_AwaitContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterArrayLiteral_Yield is called when entering the arrayLiteral_Yield production.
	EnterArrayLiteral_Yield(c *ArrayLiteral_YieldContext)

	// EnterArrayLiteral_Await is called when entering the arrayLiteral_Await production.
	EnterArrayLiteral_Await(c *ArrayLiteral_AwaitContext)

	// EnterArrayLiteral_Yield_Await is called when entering the arrayLiteral_Yield_Await production.
	EnterArrayLiteral_Yield_Await(c *ArrayLiteral_Yield_AwaitContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterElementList_Yield is called when entering the elementList_Yield production.
	EnterElementList_Yield(c *ElementList_YieldContext)

	// EnterElementList_Await is called when entering the elementList_Await production.
	EnterElementList_Await(c *ElementList_AwaitContext)

	// EnterElementList_Yield_Await is called when entering the elementList_Yield_Await production.
	EnterElementList_Yield_Await(c *ElementList_Yield_AwaitContext)

	// EnterElision is called when entering the elision production.
	EnterElision(c *ElisionContext)

	// EnterSpreadElement is called when entering the spreadElement production.
	EnterSpreadElement(c *SpreadElementContext)

	// EnterSpreadElement_Yield is called when entering the spreadElement_Yield production.
	EnterSpreadElement_Yield(c *SpreadElement_YieldContext)

	// EnterSpreadElement_Await is called when entering the spreadElement_Await production.
	EnterSpreadElement_Await(c *SpreadElement_AwaitContext)

	// EnterSpreadElement_Yield_Await is called when entering the spreadElement_Yield_Await production.
	EnterSpreadElement_Yield_Await(c *SpreadElement_Yield_AwaitContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterObjectLiteral_Yield is called when entering the objectLiteral_Yield production.
	EnterObjectLiteral_Yield(c *ObjectLiteral_YieldContext)

	// EnterObjectLiteral_Await is called when entering the objectLiteral_Await production.
	EnterObjectLiteral_Await(c *ObjectLiteral_AwaitContext)

	// EnterObjectLiteral_Yield_Await is called when entering the objectLiteral_Yield_Await production.
	EnterObjectLiteral_Yield_Await(c *ObjectLiteral_Yield_AwaitContext)

	// EnterPropertyDefinitionList is called when entering the propertyDefinitionList production.
	EnterPropertyDefinitionList(c *PropertyDefinitionListContext)

	// EnterPropertyDefinitionList_Yield is called when entering the propertyDefinitionList_Yield production.
	EnterPropertyDefinitionList_Yield(c *PropertyDefinitionList_YieldContext)

	// EnterPropertyDefinitionList_Await is called when entering the propertyDefinitionList_Await production.
	EnterPropertyDefinitionList_Await(c *PropertyDefinitionList_AwaitContext)

	// EnterPropertyDefinitionList_Yield_Await is called when entering the propertyDefinitionList_Yield_Await production.
	EnterPropertyDefinitionList_Yield_Await(c *PropertyDefinitionList_Yield_AwaitContext)

	// EnterPropertyDefinition is called when entering the propertyDefinition production.
	EnterPropertyDefinition(c *PropertyDefinitionContext)

	// EnterPropertyDefinition_Yield is called when entering the propertyDefinition_Yield production.
	EnterPropertyDefinition_Yield(c *PropertyDefinition_YieldContext)

	// EnterPropertyDefinition_Await is called when entering the propertyDefinition_Await production.
	EnterPropertyDefinition_Await(c *PropertyDefinition_AwaitContext)

	// EnterPropertyDefinition_Yield_Await is called when entering the propertyDefinition_Yield_Await production.
	EnterPropertyDefinition_Yield_Await(c *PropertyDefinition_Yield_AwaitContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterPropertyName_Yield is called when entering the propertyName_Yield production.
	EnterPropertyName_Yield(c *PropertyName_YieldContext)

	// EnterPropertyName_Await is called when entering the propertyName_Await production.
	EnterPropertyName_Await(c *PropertyName_AwaitContext)

	// EnterPropertyName_Yield_Await is called when entering the propertyName_Yield_Await production.
	EnterPropertyName_Yield_Await(c *PropertyName_Yield_AwaitContext)

	// EnterLiteralPropertyName is called when entering the literalPropertyName production.
	EnterLiteralPropertyName(c *LiteralPropertyNameContext)

	// EnterComputedPropertyName is called when entering the computedPropertyName production.
	EnterComputedPropertyName(c *ComputedPropertyNameContext)

	// EnterComputedPropertyName_Yield is called when entering the computedPropertyName_Yield production.
	EnterComputedPropertyName_Yield(c *ComputedPropertyName_YieldContext)

	// EnterComputedPropertyName_Await is called when entering the computedPropertyName_Await production.
	EnterComputedPropertyName_Await(c *ComputedPropertyName_AwaitContext)

	// EnterComputedPropertyName_Yield_Await is called when entering the computedPropertyName_Yield_Await production.
	EnterComputedPropertyName_Yield_Await(c *ComputedPropertyName_Yield_AwaitContext)

	// EnterCoverInitializedName is called when entering the coverInitializedName production.
	EnterCoverInitializedName(c *CoverInitializedNameContext)

	// EnterCoverInitializedName_Yield is called when entering the coverInitializedName_Yield production.
	EnterCoverInitializedName_Yield(c *CoverInitializedName_YieldContext)

	// EnterCoverInitializedName_Await is called when entering the coverInitializedName_Await production.
	EnterCoverInitializedName_Await(c *CoverInitializedName_AwaitContext)

	// EnterCoverInitializedName_Yield_Await is called when entering the coverInitializedName_Yield_Await production.
	EnterCoverInitializedName_Yield_Await(c *CoverInitializedName_Yield_AwaitContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterInitializer_In is called when entering the initializer_In production.
	EnterInitializer_In(c *Initializer_InContext)

	// EnterInitializer_Yield is called when entering the initializer_Yield production.
	EnterInitializer_Yield(c *Initializer_YieldContext)

	// EnterInitializer_In_Yield is called when entering the initializer_In_Yield production.
	EnterInitializer_In_Yield(c *Initializer_In_YieldContext)

	// EnterInitializer_Await is called when entering the initializer_Await production.
	EnterInitializer_Await(c *Initializer_AwaitContext)

	// EnterInitializer_In_Await is called when entering the initializer_In_Await production.
	EnterInitializer_In_Await(c *Initializer_In_AwaitContext)

	// EnterInitializer_Yield_Await is called when entering the initializer_Yield_Await production.
	EnterInitializer_Yield_Await(c *Initializer_Yield_AwaitContext)

	// EnterInitializer_In_Yield_Await is called when entering the initializer_In_Yield_Await production.
	EnterInitializer_In_Yield_Await(c *Initializer_In_Yield_AwaitContext)

	// EnterTemplateLiteral is called when entering the templateLiteral production.
	EnterTemplateLiteral(c *TemplateLiteralContext)

	// EnterTemplateLiteral_Yield is called when entering the templateLiteral_Yield production.
	EnterTemplateLiteral_Yield(c *TemplateLiteral_YieldContext)

	// EnterTemplateLiteral_Await is called when entering the templateLiteral_Await production.
	EnterTemplateLiteral_Await(c *TemplateLiteral_AwaitContext)

	// EnterTemplateLiteral_Yield_Await is called when entering the templateLiteral_Yield_Await production.
	EnterTemplateLiteral_Yield_Await(c *TemplateLiteral_Yield_AwaitContext)

	// EnterTemplateLiteral_Tagged is called when entering the templateLiteral_Tagged production.
	EnterTemplateLiteral_Tagged(c *TemplateLiteral_TaggedContext)

	// EnterTemplateLiteral_Yield_Tagged is called when entering the templateLiteral_Yield_Tagged production.
	EnterTemplateLiteral_Yield_Tagged(c *TemplateLiteral_Yield_TaggedContext)

	// EnterTemplateLiteral_Await_Tagged is called when entering the templateLiteral_Await_Tagged production.
	EnterTemplateLiteral_Await_Tagged(c *TemplateLiteral_Await_TaggedContext)

	// EnterTemplateLiteral_Yield_Await_Tagged is called when entering the templateLiteral_Yield_Await_Tagged production.
	EnterTemplateLiteral_Yield_Await_Tagged(c *TemplateLiteral_Yield_Await_TaggedContext)

	// EnterSubstitutionTemplate is called when entering the substitutionTemplate production.
	EnterSubstitutionTemplate(c *SubstitutionTemplateContext)

	// EnterSubstitutionTemplate_Yield is called when entering the substitutionTemplate_Yield production.
	EnterSubstitutionTemplate_Yield(c *SubstitutionTemplate_YieldContext)

	// EnterSubstitutionTemplate_Await is called when entering the substitutionTemplate_Await production.
	EnterSubstitutionTemplate_Await(c *SubstitutionTemplate_AwaitContext)

	// EnterSubstitutionTemplate_Yield_Await is called when entering the substitutionTemplate_Yield_Await production.
	EnterSubstitutionTemplate_Yield_Await(c *SubstitutionTemplate_Yield_AwaitContext)

	// EnterSubstitutionTemplate_Tagged is called when entering the substitutionTemplate_Tagged production.
	EnterSubstitutionTemplate_Tagged(c *SubstitutionTemplate_TaggedContext)

	// EnterSubstitutionTemplate_Yield_Tagged is called when entering the substitutionTemplate_Yield_Tagged production.
	EnterSubstitutionTemplate_Yield_Tagged(c *SubstitutionTemplate_Yield_TaggedContext)

	// EnterSubstitutionTemplate_Await_Tagged is called when entering the substitutionTemplate_Await_Tagged production.
	EnterSubstitutionTemplate_Await_Tagged(c *SubstitutionTemplate_Await_TaggedContext)

	// EnterSubstitutionTemplate_Yield_Await_Tagged is called when entering the substitutionTemplate_Yield_Await_Tagged production.
	EnterSubstitutionTemplate_Yield_Await_Tagged(c *SubstitutionTemplate_Yield_Await_TaggedContext)

	// EnterTemplateSpans is called when entering the templateSpans production.
	EnterTemplateSpans(c *TemplateSpansContext)

	// EnterTemplateSpans_Yield is called when entering the templateSpans_Yield production.
	EnterTemplateSpans_Yield(c *TemplateSpans_YieldContext)

	// EnterTemplateSpans_Await is called when entering the templateSpans_Await production.
	EnterTemplateSpans_Await(c *TemplateSpans_AwaitContext)

	// EnterTemplateSpans_Yield_Await is called when entering the templateSpans_Yield_Await production.
	EnterTemplateSpans_Yield_Await(c *TemplateSpans_Yield_AwaitContext)

	// EnterTemplateSpans_Tagged is called when entering the templateSpans_Tagged production.
	EnterTemplateSpans_Tagged(c *TemplateSpans_TaggedContext)

	// EnterTemplateSpans_Yield_Tagged is called when entering the templateSpans_Yield_Tagged production.
	EnterTemplateSpans_Yield_Tagged(c *TemplateSpans_Yield_TaggedContext)

	// EnterTemplateSpans_Await_Tagged is called when entering the templateSpans_Await_Tagged production.
	EnterTemplateSpans_Await_Tagged(c *TemplateSpans_Await_TaggedContext)

	// EnterTemplateSpans_Yield_Await_Tagged is called when entering the templateSpans_Yield_Await_Tagged production.
	EnterTemplateSpans_Yield_Await_Tagged(c *TemplateSpans_Yield_Await_TaggedContext)

	// EnterTemplateMiddleList is called when entering the templateMiddleList production.
	EnterTemplateMiddleList(c *TemplateMiddleListContext)

	// EnterTemplateMiddleList_Yield is called when entering the templateMiddleList_Yield production.
	EnterTemplateMiddleList_Yield(c *TemplateMiddleList_YieldContext)

	// EnterTemplateMiddleList_Await is called when entering the templateMiddleList_Await production.
	EnterTemplateMiddleList_Await(c *TemplateMiddleList_AwaitContext)

	// EnterTemplateMiddleList_Yield_Await is called when entering the templateMiddleList_Yield_Await production.
	EnterTemplateMiddleList_Yield_Await(c *TemplateMiddleList_Yield_AwaitContext)

	// EnterTemplateMiddleList_Tagged is called when entering the templateMiddleList_Tagged production.
	EnterTemplateMiddleList_Tagged(c *TemplateMiddleList_TaggedContext)

	// EnterTemplateMiddleList_Yield_Tagged is called when entering the templateMiddleList_Yield_Tagged production.
	EnterTemplateMiddleList_Yield_Tagged(c *TemplateMiddleList_Yield_TaggedContext)

	// EnterTemplateMiddleList_Await_Tagged is called when entering the templateMiddleList_Await_Tagged production.
	EnterTemplateMiddleList_Await_Tagged(c *TemplateMiddleList_Await_TaggedContext)

	// EnterTemplateMiddleList_Yield_Await_Tagged is called when entering the templateMiddleList_Yield_Await_Tagged production.
	EnterTemplateMiddleList_Yield_Await_Tagged(c *TemplateMiddleList_Yield_Await_TaggedContext)

	// EnterMemberExpression is called when entering the memberExpression production.
	EnterMemberExpression(c *MemberExpressionContext)

	// EnterMemberExpression_Yield is called when entering the memberExpression_Yield production.
	EnterMemberExpression_Yield(c *MemberExpression_YieldContext)

	// EnterMemberExpression_Await is called when entering the memberExpression_Await production.
	EnterMemberExpression_Await(c *MemberExpression_AwaitContext)

	// EnterMemberExpression_Yield_Await is called when entering the memberExpression_Yield_Await production.
	EnterMemberExpression_Yield_Await(c *MemberExpression_Yield_AwaitContext)

	// EnterSuperProperty is called when entering the superProperty production.
	EnterSuperProperty(c *SuperPropertyContext)

	// EnterSuperProperty_Yield is called when entering the superProperty_Yield production.
	EnterSuperProperty_Yield(c *SuperProperty_YieldContext)

	// EnterSuperProperty_Await is called when entering the superProperty_Await production.
	EnterSuperProperty_Await(c *SuperProperty_AwaitContext)

	// EnterSuperProperty_Yield_Await is called when entering the superProperty_Yield_Await production.
	EnterSuperProperty_Yield_Await(c *SuperProperty_Yield_AwaitContext)

	// EnterMetaProperty is called when entering the metaProperty production.
	EnterMetaProperty(c *MetaPropertyContext)

	// EnterNewTarget is called when entering the newTarget production.
	EnterNewTarget(c *NewTargetContext)

	// EnterNewExpression is called when entering the newExpression production.
	EnterNewExpression(c *NewExpressionContext)

	// EnterNewExpression_Yield is called when entering the newExpression_Yield production.
	EnterNewExpression_Yield(c *NewExpression_YieldContext)

	// EnterNewExpression_Await is called when entering the newExpression_Await production.
	EnterNewExpression_Await(c *NewExpression_AwaitContext)

	// EnterNewExpression_Yield_Await is called when entering the newExpression_Yield_Await production.
	EnterNewExpression_Yield_Await(c *NewExpression_Yield_AwaitContext)

	// EnterCallExpression is called when entering the callExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterCallExpression_Yield is called when entering the callExpression_Yield production.
	EnterCallExpression_Yield(c *CallExpression_YieldContext)

	// EnterCallExpression_Await is called when entering the callExpression_Await production.
	EnterCallExpression_Await(c *CallExpression_AwaitContext)

	// EnterCallExpression_Yield_Await is called when entering the callExpression_Yield_Await production.
	EnterCallExpression_Yield_Await(c *CallExpression_Yield_AwaitContext)

	// EnterSuperCall is called when entering the superCall production.
	EnterSuperCall(c *SuperCallContext)

	// EnterSuperCall_Yield is called when entering the superCall_Yield production.
	EnterSuperCall_Yield(c *SuperCall_YieldContext)

	// EnterSuperCall_Await is called when entering the superCall_Await production.
	EnterSuperCall_Await(c *SuperCall_AwaitContext)

	// EnterSuperCall_Yield_Await is called when entering the superCall_Yield_Await production.
	EnterSuperCall_Yield_Await(c *SuperCall_Yield_AwaitContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArguments_Yield is called when entering the arguments_Yield production.
	EnterArguments_Yield(c *Arguments_YieldContext)

	// EnterArguments_Await is called when entering the arguments_Await production.
	EnterArguments_Await(c *Arguments_AwaitContext)

	// EnterArguments_Yield_Await is called when entering the arguments_Yield_Await production.
	EnterArguments_Yield_Await(c *Arguments_Yield_AwaitContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArgumentList_Yield is called when entering the argumentList_Yield production.
	EnterArgumentList_Yield(c *ArgumentList_YieldContext)

	// EnterArgumentList_Await is called when entering the argumentList_Await production.
	EnterArgumentList_Await(c *ArgumentList_AwaitContext)

	// EnterArgumentList_Yield_Await is called when entering the argumentList_Yield_Await production.
	EnterArgumentList_Yield_Await(c *ArgumentList_Yield_AwaitContext)

	// EnterLeftHandSideExpression is called when entering the leftHandSideExpression production.
	EnterLeftHandSideExpression(c *LeftHandSideExpressionContext)

	// EnterLeftHandSideExpression_Yield is called when entering the leftHandSideExpression_Yield production.
	EnterLeftHandSideExpression_Yield(c *LeftHandSideExpression_YieldContext)

	// EnterLeftHandSideExpression_Await is called when entering the leftHandSideExpression_Await production.
	EnterLeftHandSideExpression_Await(c *LeftHandSideExpression_AwaitContext)

	// EnterLeftHandSideExpression_Yield_Await is called when entering the leftHandSideExpression_Yield_Await production.
	EnterLeftHandSideExpression_Yield_Await(c *LeftHandSideExpression_Yield_AwaitContext)

	// EnterUpdateExpression is called when entering the updateExpression production.
	EnterUpdateExpression(c *UpdateExpressionContext)

	// EnterUpdateExpression_Yield is called when entering the updateExpression_Yield production.
	EnterUpdateExpression_Yield(c *UpdateExpression_YieldContext)

	// EnterUpdateExpression_Await is called when entering the updateExpression_Await production.
	EnterUpdateExpression_Await(c *UpdateExpression_AwaitContext)

	// EnterUpdateExpression_Yield_Await is called when entering the updateExpression_Yield_Await production.
	EnterUpdateExpression_Yield_Await(c *UpdateExpression_Yield_AwaitContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryExpression_Yield is called when entering the unaryExpression_Yield production.
	EnterUnaryExpression_Yield(c *UnaryExpression_YieldContext)

	// EnterUnaryExpression_Await is called when entering the unaryExpression_Await production.
	EnterUnaryExpression_Await(c *UnaryExpression_AwaitContext)

	// EnterUnaryExpression_Yield_Await is called when entering the unaryExpression_Yield_Await production.
	EnterUnaryExpression_Yield_Await(c *UnaryExpression_Yield_AwaitContext)

	// EnterExponentationExpression is called when entering the exponentationExpression production.
	EnterExponentationExpression(c *ExponentationExpressionContext)

	// EnterExponentationExpression_Yield is called when entering the exponentationExpression_Yield production.
	EnterExponentationExpression_Yield(c *ExponentationExpression_YieldContext)

	// EnterExponentationExpression_Await is called when entering the exponentationExpression_Await production.
	EnterExponentationExpression_Await(c *ExponentationExpression_AwaitContext)

	// EnterExponentationExpression_Yield_Await is called when entering the exponentationExpression_Yield_Await production.
	EnterExponentationExpression_Yield_Await(c *ExponentationExpression_Yield_AwaitContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterMultiplicativeExpression_Yield is called when entering the multiplicativeExpression_Yield production.
	EnterMultiplicativeExpression_Yield(c *MultiplicativeExpression_YieldContext)

	// EnterMultiplicativeExpression_Await is called when entering the multiplicativeExpression_Await production.
	EnterMultiplicativeExpression_Await(c *MultiplicativeExpression_AwaitContext)

	// EnterMultiplicativeExpression_Yield_Await is called when entering the multiplicativeExpression_Yield_Await production.
	EnterMultiplicativeExpression_Yield_Await(c *MultiplicativeExpression_Yield_AwaitContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterAdditiveExpression_Yield is called when entering the additiveExpression_Yield production.
	EnterAdditiveExpression_Yield(c *AdditiveExpression_YieldContext)

	// EnterAdditiveExpression_Await is called when entering the additiveExpression_Await production.
	EnterAdditiveExpression_Await(c *AdditiveExpression_AwaitContext)

	// EnterAdditiveExpression_Yield_Await is called when entering the additiveExpression_Yield_Await production.
	EnterAdditiveExpression_Yield_Await(c *AdditiveExpression_Yield_AwaitContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterShiftExpression_Yield is called when entering the shiftExpression_Yield production.
	EnterShiftExpression_Yield(c *ShiftExpression_YieldContext)

	// EnterShiftExpression_Await is called when entering the shiftExpression_Await production.
	EnterShiftExpression_Await(c *ShiftExpression_AwaitContext)

	// EnterShiftExpression_Yield_Await is called when entering the shiftExpression_Yield_Await production.
	EnterShiftExpression_Yield_Await(c *ShiftExpression_Yield_AwaitContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterRelationalExpression_In is called when entering the relationalExpression_In production.
	EnterRelationalExpression_In(c *RelationalExpression_InContext)

	// EnterRelationalExpression_Yield is called when entering the relationalExpression_Yield production.
	EnterRelationalExpression_Yield(c *RelationalExpression_YieldContext)

	// EnterRelationalExpression_In_Yield is called when entering the relationalExpression_In_Yield production.
	EnterRelationalExpression_In_Yield(c *RelationalExpression_In_YieldContext)

	// EnterRelationalExpression_Await is called when entering the relationalExpression_Await production.
	EnterRelationalExpression_Await(c *RelationalExpression_AwaitContext)

	// EnterRelationalExpression_In_Await is called when entering the relationalExpression_In_Await production.
	EnterRelationalExpression_In_Await(c *RelationalExpression_In_AwaitContext)

	// EnterRelationalExpression_Yield_Await is called when entering the relationalExpression_Yield_Await production.
	EnterRelationalExpression_Yield_Await(c *RelationalExpression_Yield_AwaitContext)

	// EnterRelationalExpression_In_Yield_Await is called when entering the relationalExpression_In_Yield_Await production.
	EnterRelationalExpression_In_Yield_Await(c *RelationalExpression_In_Yield_AwaitContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterEqualityExpression_In is called when entering the equalityExpression_In production.
	EnterEqualityExpression_In(c *EqualityExpression_InContext)

	// EnterEqualityExpression_Yield is called when entering the equalityExpression_Yield production.
	EnterEqualityExpression_Yield(c *EqualityExpression_YieldContext)

	// EnterEqualityExpression_In_Yield is called when entering the equalityExpression_In_Yield production.
	EnterEqualityExpression_In_Yield(c *EqualityExpression_In_YieldContext)

	// EnterEqualityExpression_Await is called when entering the equalityExpression_Await production.
	EnterEqualityExpression_Await(c *EqualityExpression_AwaitContext)

	// EnterEqualityExpression_In_Await is called when entering the equalityExpression_In_Await production.
	EnterEqualityExpression_In_Await(c *EqualityExpression_In_AwaitContext)

	// EnterEqualityExpression_Yield_Await is called when entering the equalityExpression_Yield_Await production.
	EnterEqualityExpression_Yield_Await(c *EqualityExpression_Yield_AwaitContext)

	// EnterEqualityExpression_In_Yield_Await is called when entering the equalityExpression_In_Yield_Await production.
	EnterEqualityExpression_In_Yield_Await(c *EqualityExpression_In_Yield_AwaitContext)

	// EnterBitwiseANDExpression is called when entering the bitwiseANDExpression production.
	EnterBitwiseANDExpression(c *BitwiseANDExpressionContext)

	// EnterBitwiseANDExpression_In is called when entering the bitwiseANDExpression_In production.
	EnterBitwiseANDExpression_In(c *BitwiseANDExpression_InContext)

	// EnterBitwiseANDExpression_Yield is called when entering the bitwiseANDExpression_Yield production.
	EnterBitwiseANDExpression_Yield(c *BitwiseANDExpression_YieldContext)

	// EnterBitwiseANDExpression_In_Yield is called when entering the bitwiseANDExpression_In_Yield production.
	EnterBitwiseANDExpression_In_Yield(c *BitwiseANDExpression_In_YieldContext)

	// EnterBitwiseANDExpression_Await is called when entering the bitwiseANDExpression_Await production.
	EnterBitwiseANDExpression_Await(c *BitwiseANDExpression_AwaitContext)

	// EnterBitwiseANDExpression_In_Await is called when entering the bitwiseANDExpression_In_Await production.
	EnterBitwiseANDExpression_In_Await(c *BitwiseANDExpression_In_AwaitContext)

	// EnterBitwiseANDExpression_Yield_Await is called when entering the bitwiseANDExpression_Yield_Await production.
	EnterBitwiseANDExpression_Yield_Await(c *BitwiseANDExpression_Yield_AwaitContext)

	// EnterBitwiseANDExpression_In_Yield_Await is called when entering the bitwiseANDExpression_In_Yield_Await production.
	EnterBitwiseANDExpression_In_Yield_Await(c *BitwiseANDExpression_In_Yield_AwaitContext)

	// EnterBitwiseXORExpression is called when entering the bitwiseXORExpression production.
	EnterBitwiseXORExpression(c *BitwiseXORExpressionContext)

	// EnterBitwiseXORExpression_In is called when entering the bitwiseXORExpression_In production.
	EnterBitwiseXORExpression_In(c *BitwiseXORExpression_InContext)

	// EnterBitwiseXORExpression_Yield is called when entering the bitwiseXORExpression_Yield production.
	EnterBitwiseXORExpression_Yield(c *BitwiseXORExpression_YieldContext)

	// EnterBitwiseXORExpression_In_Yield is called when entering the bitwiseXORExpression_In_Yield production.
	EnterBitwiseXORExpression_In_Yield(c *BitwiseXORExpression_In_YieldContext)

	// EnterBitwiseXORExpression_Await is called when entering the bitwiseXORExpression_Await production.
	EnterBitwiseXORExpression_Await(c *BitwiseXORExpression_AwaitContext)

	// EnterBitwiseXORExpression_In_Await is called when entering the bitwiseXORExpression_In_Await production.
	EnterBitwiseXORExpression_In_Await(c *BitwiseXORExpression_In_AwaitContext)

	// EnterBitwiseXORExpression_Yield_Await is called when entering the bitwiseXORExpression_Yield_Await production.
	EnterBitwiseXORExpression_Yield_Await(c *BitwiseXORExpression_Yield_AwaitContext)

	// EnterBitwiseXORExpression_In_Yield_Await is called when entering the bitwiseXORExpression_In_Yield_Await production.
	EnterBitwiseXORExpression_In_Yield_Await(c *BitwiseXORExpression_In_Yield_AwaitContext)

	// EnterBitwiseORExpression is called when entering the bitwiseORExpression production.
	EnterBitwiseORExpression(c *BitwiseORExpressionContext)

	// EnterBitwiseORExpression_In is called when entering the bitwiseORExpression_In production.
	EnterBitwiseORExpression_In(c *BitwiseORExpression_InContext)

	// EnterBitwiseORExpression_Yield is called when entering the bitwiseORExpression_Yield production.
	EnterBitwiseORExpression_Yield(c *BitwiseORExpression_YieldContext)

	// EnterBitwiseORExpression_In_Yield is called when entering the bitwiseORExpression_In_Yield production.
	EnterBitwiseORExpression_In_Yield(c *BitwiseORExpression_In_YieldContext)

	// EnterBitwiseORExpression_Await is called when entering the bitwiseORExpression_Await production.
	EnterBitwiseORExpression_Await(c *BitwiseORExpression_AwaitContext)

	// EnterBitwiseORExpression_In_Await is called when entering the bitwiseORExpression_In_Await production.
	EnterBitwiseORExpression_In_Await(c *BitwiseORExpression_In_AwaitContext)

	// EnterBitwiseORExpression_Yield_Await is called when entering the bitwiseORExpression_Yield_Await production.
	EnterBitwiseORExpression_Yield_Await(c *BitwiseORExpression_Yield_AwaitContext)

	// EnterBitwiseORExpression_In_Yield_Await is called when entering the bitwiseORExpression_In_Yield_Await production.
	EnterBitwiseORExpression_In_Yield_Await(c *BitwiseORExpression_In_Yield_AwaitContext)

	// EnterLogicalANDExpression is called when entering the logicalANDExpression production.
	EnterLogicalANDExpression(c *LogicalANDExpressionContext)

	// EnterLogicalANDExpression_In is called when entering the logicalANDExpression_In production.
	EnterLogicalANDExpression_In(c *LogicalANDExpression_InContext)

	// EnterLogicalANDExpression_Yield is called when entering the logicalANDExpression_Yield production.
	EnterLogicalANDExpression_Yield(c *LogicalANDExpression_YieldContext)

	// EnterLogicalANDExpression_In_Yield is called when entering the logicalANDExpression_In_Yield production.
	EnterLogicalANDExpression_In_Yield(c *LogicalANDExpression_In_YieldContext)

	// EnterLogicalANDExpression_Await is called when entering the logicalANDExpression_Await production.
	EnterLogicalANDExpression_Await(c *LogicalANDExpression_AwaitContext)

	// EnterLogicalANDExpression_In_Await is called when entering the logicalANDExpression_In_Await production.
	EnterLogicalANDExpression_In_Await(c *LogicalANDExpression_In_AwaitContext)

	// EnterLogicalANDExpression_Yield_Await is called when entering the logicalANDExpression_Yield_Await production.
	EnterLogicalANDExpression_Yield_Await(c *LogicalANDExpression_Yield_AwaitContext)

	// EnterLogicalANDExpression_In_Yield_Await is called when entering the logicalANDExpression_In_Yield_Await production.
	EnterLogicalANDExpression_In_Yield_Await(c *LogicalANDExpression_In_Yield_AwaitContext)

	// EnterLogicalORExpression is called when entering the logicalORExpression production.
	EnterLogicalORExpression(c *LogicalORExpressionContext)

	// EnterLogicalORExpression_In is called when entering the logicalORExpression_In production.
	EnterLogicalORExpression_In(c *LogicalORExpression_InContext)

	// EnterLogicalORExpression_Yield is called when entering the logicalORExpression_Yield production.
	EnterLogicalORExpression_Yield(c *LogicalORExpression_YieldContext)

	// EnterLogicalORExpression_In_Yield is called when entering the logicalORExpression_In_Yield production.
	EnterLogicalORExpression_In_Yield(c *LogicalORExpression_In_YieldContext)

	// EnterLogicalORExpression_Await is called when entering the logicalORExpression_Await production.
	EnterLogicalORExpression_Await(c *LogicalORExpression_AwaitContext)

	// EnterLogicalORExpression_In_Await is called when entering the logicalORExpression_In_Await production.
	EnterLogicalORExpression_In_Await(c *LogicalORExpression_In_AwaitContext)

	// EnterLogicalORExpression_Yield_Await is called when entering the logicalORExpression_Yield_Await production.
	EnterLogicalORExpression_Yield_Await(c *LogicalORExpression_Yield_AwaitContext)

	// EnterLogicalORExpression_In_Yield_Await is called when entering the logicalORExpression_In_Yield_Await production.
	EnterLogicalORExpression_In_Yield_Await(c *LogicalORExpression_In_Yield_AwaitContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterConditionalExpression_In is called when entering the conditionalExpression_In production.
	EnterConditionalExpression_In(c *ConditionalExpression_InContext)

	// EnterConditionalExpression_Yield is called when entering the conditionalExpression_Yield production.
	EnterConditionalExpression_Yield(c *ConditionalExpression_YieldContext)

	// EnterConditionalExpression_In_Yield is called when entering the conditionalExpression_In_Yield production.
	EnterConditionalExpression_In_Yield(c *ConditionalExpression_In_YieldContext)

	// EnterConditionalExpression_Await is called when entering the conditionalExpression_Await production.
	EnterConditionalExpression_Await(c *ConditionalExpression_AwaitContext)

	// EnterConditionalExpression_In_Await is called when entering the conditionalExpression_In_Await production.
	EnterConditionalExpression_In_Await(c *ConditionalExpression_In_AwaitContext)

	// EnterConditionalExpression_Yield_Await is called when entering the conditionalExpression_Yield_Await production.
	EnterConditionalExpression_Yield_Await(c *ConditionalExpression_Yield_AwaitContext)

	// EnterConditionalExpression_In_Yield_Await is called when entering the conditionalExpression_In_Yield_Await production.
	EnterConditionalExpression_In_Yield_Await(c *ConditionalExpression_In_Yield_AwaitContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentExpression_In is called when entering the assignmentExpression_In production.
	EnterAssignmentExpression_In(c *AssignmentExpression_InContext)

	// EnterAssignmentExpression_Yield is called when entering the assignmentExpression_Yield production.
	EnterAssignmentExpression_Yield(c *AssignmentExpression_YieldContext)

	// EnterAssignmentExpression_In_Yield is called when entering the assignmentExpression_In_Yield production.
	EnterAssignmentExpression_In_Yield(c *AssignmentExpression_In_YieldContext)

	// EnterAssignmentExpression_Await is called when entering the assignmentExpression_Await production.
	EnterAssignmentExpression_Await(c *AssignmentExpression_AwaitContext)

	// EnterAssignmentExpression_In_Await is called when entering the assignmentExpression_In_Await production.
	EnterAssignmentExpression_In_Await(c *AssignmentExpression_In_AwaitContext)

	// EnterAssignmentExpression_Yield_Await is called when entering the assignmentExpression_Yield_Await production.
	EnterAssignmentExpression_Yield_Await(c *AssignmentExpression_Yield_AwaitContext)

	// EnterAssignmentExpression_In_Yield_Await is called when entering the assignmentExpression_In_Yield_Await production.
	EnterAssignmentExpression_In_Yield_Await(c *AssignmentExpression_In_Yield_AwaitContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpression_In is called when entering the expression_In production.
	EnterExpression_In(c *Expression_InContext)

	// EnterExpression_Yield is called when entering the expression_Yield production.
	EnterExpression_Yield(c *Expression_YieldContext)

	// EnterExpression_In_Yield is called when entering the expression_In_Yield production.
	EnterExpression_In_Yield(c *Expression_In_YieldContext)

	// EnterExpression_Await is called when entering the expression_Await production.
	EnterExpression_Await(c *Expression_AwaitContext)

	// EnterExpression_In_Await is called when entering the expression_In_Await production.
	EnterExpression_In_Await(c *Expression_In_AwaitContext)

	// EnterExpression_Yield_Await is called when entering the expression_Yield_Await production.
	EnterExpression_Yield_Await(c *Expression_Yield_AwaitContext)

	// EnterExpression_In_Yield_Await is called when entering the expression_In_Yield_Await production.
	EnterExpression_In_Yield_Await(c *Expression_In_Yield_AwaitContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatement_Yield is called when entering the statement_Yield production.
	EnterStatement_Yield(c *Statement_YieldContext)

	// EnterStatement_Await is called when entering the statement_Await production.
	EnterStatement_Await(c *Statement_AwaitContext)

	// EnterStatement_Yield_Await is called when entering the statement_Yield_Await production.
	EnterStatement_Yield_Await(c *Statement_Yield_AwaitContext)

	// EnterStatement_Return is called when entering the statement_Return production.
	EnterStatement_Return(c *Statement_ReturnContext)

	// EnterStatement_Yield_Return is called when entering the statement_Yield_Return production.
	EnterStatement_Yield_Return(c *Statement_Yield_ReturnContext)

	// EnterStatement_Await_Return is called when entering the statement_Await_Return production.
	EnterStatement_Await_Return(c *Statement_Await_ReturnContext)

	// EnterStatement_Yield_Await_Return is called when entering the statement_Yield_Await_Return production.
	EnterStatement_Yield_Await_Return(c *Statement_Yield_Await_ReturnContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDeclaration_Yield is called when entering the declaration_Yield production.
	EnterDeclaration_Yield(c *Declaration_YieldContext)

	// EnterDeclaration_Await is called when entering the declaration_Await production.
	EnterDeclaration_Await(c *Declaration_AwaitContext)

	// EnterDeclaration_Yield_Await is called when entering the declaration_Yield_Await production.
	EnterDeclaration_Yield_Await(c *Declaration_Yield_AwaitContext)

	// EnterHoistableDeclaration is called when entering the hoistableDeclaration production.
	EnterHoistableDeclaration(c *HoistableDeclarationContext)

	// EnterHoistableDeclaration_Yield is called when entering the hoistableDeclaration_Yield production.
	EnterHoistableDeclaration_Yield(c *HoistableDeclaration_YieldContext)

	// EnterHoistableDeclaration_Await is called when entering the hoistableDeclaration_Await production.
	EnterHoistableDeclaration_Await(c *HoistableDeclaration_AwaitContext)

	// EnterHoistableDeclaration_Yield_Await is called when entering the hoistableDeclaration_Yield_Await production.
	EnterHoistableDeclaration_Yield_Await(c *HoistableDeclaration_Yield_AwaitContext)

	// EnterHoistableDeclaration_Default is called when entering the hoistableDeclaration_Default production.
	EnterHoistableDeclaration_Default(c *HoistableDeclaration_DefaultContext)

	// EnterHoistableDeclaration_Yield_Default is called when entering the hoistableDeclaration_Yield_Default production.
	EnterHoistableDeclaration_Yield_Default(c *HoistableDeclaration_Yield_DefaultContext)

	// EnterHoistableDeclaration_Await_Default is called when entering the hoistableDeclaration_Await_Default production.
	EnterHoistableDeclaration_Await_Default(c *HoistableDeclaration_Await_DefaultContext)

	// EnterHoistableDeclaration_Yield_Await_Default is called when entering the hoistableDeclaration_Yield_Await_Default production.
	EnterHoistableDeclaration_Yield_Await_Default(c *HoistableDeclaration_Yield_Await_DefaultContext)

	// EnterBreakableStatement is called when entering the breakableStatement production.
	EnterBreakableStatement(c *BreakableStatementContext)

	// EnterBreakableStatement_Yield is called when entering the breakableStatement_Yield production.
	EnterBreakableStatement_Yield(c *BreakableStatement_YieldContext)

	// EnterBreakableStatement_Await is called when entering the breakableStatement_Await production.
	EnterBreakableStatement_Await(c *BreakableStatement_AwaitContext)

	// EnterBreakableStatement_Yield_Await is called when entering the breakableStatement_Yield_Await production.
	EnterBreakableStatement_Yield_Await(c *BreakableStatement_Yield_AwaitContext)

	// EnterBreakableStatement_Return is called when entering the breakableStatement_Return production.
	EnterBreakableStatement_Return(c *BreakableStatement_ReturnContext)

	// EnterBreakableStatement_Yield_Return is called when entering the breakableStatement_Yield_Return production.
	EnterBreakableStatement_Yield_Return(c *BreakableStatement_Yield_ReturnContext)

	// EnterBreakableStatement_Await_Return is called when entering the breakableStatement_Await_Return production.
	EnterBreakableStatement_Await_Return(c *BreakableStatement_Await_ReturnContext)

	// EnterBreakableStatement_Yield_Await_Return is called when entering the breakableStatement_Yield_Await_Return production.
	EnterBreakableStatement_Yield_Await_Return(c *BreakableStatement_Yield_Await_ReturnContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterBlockStatement_Yield is called when entering the blockStatement_Yield production.
	EnterBlockStatement_Yield(c *BlockStatement_YieldContext)

	// EnterBlockStatement_Await is called when entering the blockStatement_Await production.
	EnterBlockStatement_Await(c *BlockStatement_AwaitContext)

	// EnterBlockStatement_Yield_Await is called when entering the blockStatement_Yield_Await production.
	EnterBlockStatement_Yield_Await(c *BlockStatement_Yield_AwaitContext)

	// EnterBlockStatement_Return is called when entering the blockStatement_Return production.
	EnterBlockStatement_Return(c *BlockStatement_ReturnContext)

	// EnterBlockStatement_Yield_Return is called when entering the blockStatement_Yield_Return production.
	EnterBlockStatement_Yield_Return(c *BlockStatement_Yield_ReturnContext)

	// EnterBlockStatement_Await_Return is called when entering the blockStatement_Await_Return production.
	EnterBlockStatement_Await_Return(c *BlockStatement_Await_ReturnContext)

	// EnterBlockStatement_Yield_Await_Return is called when entering the blockStatement_Yield_Await_Return production.
	EnterBlockStatement_Yield_Await_Return(c *BlockStatement_Yield_Await_ReturnContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlock_Yield is called when entering the block_Yield production.
	EnterBlock_Yield(c *Block_YieldContext)

	// EnterBlock_Await is called when entering the block_Await production.
	EnterBlock_Await(c *Block_AwaitContext)

	// EnterBlock_Yield_Await is called when entering the block_Yield_Await production.
	EnterBlock_Yield_Await(c *Block_Yield_AwaitContext)

	// EnterBlock_Return is called when entering the block_Return production.
	EnterBlock_Return(c *Block_ReturnContext)

	// EnterBlock_Yield_Return is called when entering the block_Yield_Return production.
	EnterBlock_Yield_Return(c *Block_Yield_ReturnContext)

	// EnterBlock_Await_Return is called when entering the block_Await_Return production.
	EnterBlock_Await_Return(c *Block_Await_ReturnContext)

	// EnterBlock_Yield_Await_Return is called when entering the block_Yield_Await_Return production.
	EnterBlock_Yield_Await_Return(c *Block_Yield_Await_ReturnContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatementList_Yield is called when entering the statementList_Yield production.
	EnterStatementList_Yield(c *StatementList_YieldContext)

	// EnterStatementList_Await is called when entering the statementList_Await production.
	EnterStatementList_Await(c *StatementList_AwaitContext)

	// EnterStatementList_Yield_Await is called when entering the statementList_Yield_Await production.
	EnterStatementList_Yield_Await(c *StatementList_Yield_AwaitContext)

	// EnterStatementList_Return is called when entering the statementList_Return production.
	EnterStatementList_Return(c *StatementList_ReturnContext)

	// EnterStatementList_Yield_Return is called when entering the statementList_Yield_Return production.
	EnterStatementList_Yield_Return(c *StatementList_Yield_ReturnContext)

	// EnterStatementList_Await_Return is called when entering the statementList_Await_Return production.
	EnterStatementList_Await_Return(c *StatementList_Await_ReturnContext)

	// EnterStatementList_Yield_Await_Return is called when entering the statementList_Yield_Await_Return production.
	EnterStatementList_Yield_Await_Return(c *StatementList_Yield_Await_ReturnContext)

	// EnterStatementListItem is called when entering the statementListItem production.
	EnterStatementListItem(c *StatementListItemContext)

	// EnterStatementListItem_Yield is called when entering the statementListItem_Yield production.
	EnterStatementListItem_Yield(c *StatementListItem_YieldContext)

	// EnterStatementListItem_Await is called when entering the statementListItem_Await production.
	EnterStatementListItem_Await(c *StatementListItem_AwaitContext)

	// EnterStatementListItem_Yield_Await is called when entering the statementListItem_Yield_Await production.
	EnterStatementListItem_Yield_Await(c *StatementListItem_Yield_AwaitContext)

	// EnterStatementListItem_Return is called when entering the statementListItem_Return production.
	EnterStatementListItem_Return(c *StatementListItem_ReturnContext)

	// EnterStatementListItem_Yield_Return is called when entering the statementListItem_Yield_Return production.
	EnterStatementListItem_Yield_Return(c *StatementListItem_Yield_ReturnContext)

	// EnterStatementListItem_Await_Return is called when entering the statementListItem_Await_Return production.
	EnterStatementListItem_Await_Return(c *StatementListItem_Await_ReturnContext)

	// EnterStatementListItem_Yield_Await_Return is called when entering the statementListItem_Yield_Await_Return production.
	EnterStatementListItem_Yield_Await_Return(c *StatementListItem_Yield_Await_ReturnContext)

	// EnterLexicalDeclaration is called when entering the lexicalDeclaration production.
	EnterLexicalDeclaration(c *LexicalDeclarationContext)

	// EnterLexicalDeclaration_In is called when entering the lexicalDeclaration_In production.
	EnterLexicalDeclaration_In(c *LexicalDeclaration_InContext)

	// EnterLexicalDeclaration_Yield is called when entering the lexicalDeclaration_Yield production.
	EnterLexicalDeclaration_Yield(c *LexicalDeclaration_YieldContext)

	// EnterLexicalDeclaration_In_Yield is called when entering the lexicalDeclaration_In_Yield production.
	EnterLexicalDeclaration_In_Yield(c *LexicalDeclaration_In_YieldContext)

	// EnterLexicalDeclaration_Await is called when entering the lexicalDeclaration_Await production.
	EnterLexicalDeclaration_Await(c *LexicalDeclaration_AwaitContext)

	// EnterLexicalDeclaration_In_Await is called when entering the lexicalDeclaration_In_Await production.
	EnterLexicalDeclaration_In_Await(c *LexicalDeclaration_In_AwaitContext)

	// EnterLexicalDeclaration_Yield_Await is called when entering the lexicalDeclaration_Yield_Await production.
	EnterLexicalDeclaration_Yield_Await(c *LexicalDeclaration_Yield_AwaitContext)

	// EnterLexicalDeclaration_In_Yield_Await is called when entering the lexicalDeclaration_In_Yield_Await production.
	EnterLexicalDeclaration_In_Yield_Await(c *LexicalDeclaration_In_Yield_AwaitContext)

	// EnterLetOrConst is called when entering the letOrConst production.
	EnterLetOrConst(c *LetOrConstContext)

	// EnterBindingList is called when entering the bindingList production.
	EnterBindingList(c *BindingListContext)

	// EnterBindingList_In is called when entering the bindingList_In production.
	EnterBindingList_In(c *BindingList_InContext)

	// EnterBindingList_Yield is called when entering the bindingList_Yield production.
	EnterBindingList_Yield(c *BindingList_YieldContext)

	// EnterBindingList_In_Yield is called when entering the bindingList_In_Yield production.
	EnterBindingList_In_Yield(c *BindingList_In_YieldContext)

	// EnterBindingList_Await is called when entering the bindingList_Await production.
	EnterBindingList_Await(c *BindingList_AwaitContext)

	// EnterBindingList_In_Await is called when entering the bindingList_In_Await production.
	EnterBindingList_In_Await(c *BindingList_In_AwaitContext)

	// EnterBindingList_Yield_Await is called when entering the bindingList_Yield_Await production.
	EnterBindingList_Yield_Await(c *BindingList_Yield_AwaitContext)

	// EnterBindingList_In_Yield_Await is called when entering the bindingList_In_Yield_Await production.
	EnterBindingList_In_Yield_Await(c *BindingList_In_Yield_AwaitContext)

	// EnterLexicalBinding is called when entering the lexicalBinding production.
	EnterLexicalBinding(c *LexicalBindingContext)

	// EnterLexicalBinding_In is called when entering the lexicalBinding_In production.
	EnterLexicalBinding_In(c *LexicalBinding_InContext)

	// EnterLexicalBinding_Yield is called when entering the lexicalBinding_Yield production.
	EnterLexicalBinding_Yield(c *LexicalBinding_YieldContext)

	// EnterLexicalBinding_In_Yield is called when entering the lexicalBinding_In_Yield production.
	EnterLexicalBinding_In_Yield(c *LexicalBinding_In_YieldContext)

	// EnterLexicalBinding_Await is called when entering the lexicalBinding_Await production.
	EnterLexicalBinding_Await(c *LexicalBinding_AwaitContext)

	// EnterLexicalBinding_In_Await is called when entering the lexicalBinding_In_Await production.
	EnterLexicalBinding_In_Await(c *LexicalBinding_In_AwaitContext)

	// EnterLexicalBinding_Yield_Await is called when entering the lexicalBinding_Yield_Await production.
	EnterLexicalBinding_Yield_Await(c *LexicalBinding_Yield_AwaitContext)

	// EnterLexicalBinding_In_Yield_Await is called when entering the lexicalBinding_In_Yield_Await production.
	EnterLexicalBinding_In_Yield_Await(c *LexicalBinding_In_Yield_AwaitContext)

	// EnterVariableStatement is called when entering the variableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterVariableStatement_Yield is called when entering the variableStatement_Yield production.
	EnterVariableStatement_Yield(c *VariableStatement_YieldContext)

	// EnterVariableStatement_Await is called when entering the variableStatement_Await production.
	EnterVariableStatement_Await(c *VariableStatement_AwaitContext)

	// EnterVariableStatement_Yield_Await is called when entering the variableStatement_Yield_Await production.
	EnterVariableStatement_Yield_Await(c *VariableStatement_Yield_AwaitContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterVariableDeclarationList_In is called when entering the variableDeclarationList_In production.
	EnterVariableDeclarationList_In(c *VariableDeclarationList_InContext)

	// EnterVariableDeclarationList_Yield is called when entering the variableDeclarationList_Yield production.
	EnterVariableDeclarationList_Yield(c *VariableDeclarationList_YieldContext)

	// EnterVariableDeclarationList_In_Yield is called when entering the variableDeclarationList_In_Yield production.
	EnterVariableDeclarationList_In_Yield(c *VariableDeclarationList_In_YieldContext)

	// EnterVariableDeclarationList_Await is called when entering the variableDeclarationList_Await production.
	EnterVariableDeclarationList_Await(c *VariableDeclarationList_AwaitContext)

	// EnterVariableDeclarationList_In_Await is called when entering the variableDeclarationList_In_Await production.
	EnterVariableDeclarationList_In_Await(c *VariableDeclarationList_In_AwaitContext)

	// EnterVariableDeclarationList_Yield_Await is called when entering the variableDeclarationList_Yield_Await production.
	EnterVariableDeclarationList_Yield_Await(c *VariableDeclarationList_Yield_AwaitContext)

	// EnterVariableDeclarationList_In_Yield_Await is called when entering the variableDeclarationList_In_Yield_Await production.
	EnterVariableDeclarationList_In_Yield_Await(c *VariableDeclarationList_In_Yield_AwaitContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterVariableDeclaration_In is called when entering the variableDeclaration_In production.
	EnterVariableDeclaration_In(c *VariableDeclaration_InContext)

	// EnterVariableDeclaration_Yield is called when entering the variableDeclaration_Yield production.
	EnterVariableDeclaration_Yield(c *VariableDeclaration_YieldContext)

	// EnterVariableDeclaration_In_Yield is called when entering the variableDeclaration_In_Yield production.
	EnterVariableDeclaration_In_Yield(c *VariableDeclaration_In_YieldContext)

	// EnterVariableDeclaration_Await is called when entering the variableDeclaration_Await production.
	EnterVariableDeclaration_Await(c *VariableDeclaration_AwaitContext)

	// EnterVariableDeclaration_In_Await is called when entering the variableDeclaration_In_Await production.
	EnterVariableDeclaration_In_Await(c *VariableDeclaration_In_AwaitContext)

	// EnterVariableDeclaration_Yield_Await is called when entering the variableDeclaration_Yield_Await production.
	EnterVariableDeclaration_Yield_Await(c *VariableDeclaration_Yield_AwaitContext)

	// EnterVariableDeclaration_In_Yield_Await is called when entering the variableDeclaration_In_Yield_Await production.
	EnterVariableDeclaration_In_Yield_Await(c *VariableDeclaration_In_Yield_AwaitContext)

	// EnterBindingPattern is called when entering the bindingPattern production.
	EnterBindingPattern(c *BindingPatternContext)

	// EnterBindingPattern_Yield is called when entering the bindingPattern_Yield production.
	EnterBindingPattern_Yield(c *BindingPattern_YieldContext)

	// EnterBindingPattern_Await is called when entering the bindingPattern_Await production.
	EnterBindingPattern_Await(c *BindingPattern_AwaitContext)

	// EnterBindingPattern_Yield_Await is called when entering the bindingPattern_Yield_Await production.
	EnterBindingPattern_Yield_Await(c *BindingPattern_Yield_AwaitContext)

	// EnterObjectBindingPattern is called when entering the objectBindingPattern production.
	EnterObjectBindingPattern(c *ObjectBindingPatternContext)

	// EnterObjectBindingPattern_Yield is called when entering the objectBindingPattern_Yield production.
	EnterObjectBindingPattern_Yield(c *ObjectBindingPattern_YieldContext)

	// EnterObjectBindingPattern_Await is called when entering the objectBindingPattern_Await production.
	EnterObjectBindingPattern_Await(c *ObjectBindingPattern_AwaitContext)

	// EnterObjectBindingPattern_Yield_Await is called when entering the objectBindingPattern_Yield_Await production.
	EnterObjectBindingPattern_Yield_Await(c *ObjectBindingPattern_Yield_AwaitContext)

	// EnterArrayBindingPattern is called when entering the arrayBindingPattern production.
	EnterArrayBindingPattern(c *ArrayBindingPatternContext)

	// EnterArrayBindingPattern_Yield is called when entering the arrayBindingPattern_Yield production.
	EnterArrayBindingPattern_Yield(c *ArrayBindingPattern_YieldContext)

	// EnterArrayBindingPattern_Await is called when entering the arrayBindingPattern_Await production.
	EnterArrayBindingPattern_Await(c *ArrayBindingPattern_AwaitContext)

	// EnterArrayBindingPattern_Yield_Await is called when entering the arrayBindingPattern_Yield_Await production.
	EnterArrayBindingPattern_Yield_Await(c *ArrayBindingPattern_Yield_AwaitContext)

	// EnterBindingRestProperty is called when entering the bindingRestProperty production.
	EnterBindingRestProperty(c *BindingRestPropertyContext)

	// EnterBindingRestProperty_Yield is called when entering the bindingRestProperty_Yield production.
	EnterBindingRestProperty_Yield(c *BindingRestProperty_YieldContext)

	// EnterBindingRestProperty_Await is called when entering the bindingRestProperty_Await production.
	EnterBindingRestProperty_Await(c *BindingRestProperty_AwaitContext)

	// EnterBindingRestProperty_Yield_Await is called when entering the bindingRestProperty_Yield_Await production.
	EnterBindingRestProperty_Yield_Await(c *BindingRestProperty_Yield_AwaitContext)

	// EnterBindingPropertyList is called when entering the bindingPropertyList production.
	EnterBindingPropertyList(c *BindingPropertyListContext)

	// EnterBindingPropertyList_Yield is called when entering the bindingPropertyList_Yield production.
	EnterBindingPropertyList_Yield(c *BindingPropertyList_YieldContext)

	// EnterBindingPropertyList_Await is called when entering the bindingPropertyList_Await production.
	EnterBindingPropertyList_Await(c *BindingPropertyList_AwaitContext)

	// EnterBindingPropertyList_Yield_Await is called when entering the bindingPropertyList_Yield_Await production.
	EnterBindingPropertyList_Yield_Await(c *BindingPropertyList_Yield_AwaitContext)

	// EnterBindingElementList is called when entering the bindingElementList production.
	EnterBindingElementList(c *BindingElementListContext)

	// EnterBindingElementList_Yield is called when entering the bindingElementList_Yield production.
	EnterBindingElementList_Yield(c *BindingElementList_YieldContext)

	// EnterBindingElementList_Await is called when entering the bindingElementList_Await production.
	EnterBindingElementList_Await(c *BindingElementList_AwaitContext)

	// EnterBindingElementList_Yield_Await is called when entering the bindingElementList_Yield_Await production.
	EnterBindingElementList_Yield_Await(c *BindingElementList_Yield_AwaitContext)

	// EnterBindingElisionElement is called when entering the bindingElisionElement production.
	EnterBindingElisionElement(c *BindingElisionElementContext)

	// EnterBindingElisionElement_Yield is called when entering the bindingElisionElement_Yield production.
	EnterBindingElisionElement_Yield(c *BindingElisionElement_YieldContext)

	// EnterBindingElisionElement_Await is called when entering the bindingElisionElement_Await production.
	EnterBindingElisionElement_Await(c *BindingElisionElement_AwaitContext)

	// EnterBindingElisionElement_Yield_Await is called when entering the bindingElisionElement_Yield_Await production.
	EnterBindingElisionElement_Yield_Await(c *BindingElisionElement_Yield_AwaitContext)

	// EnterBindingProperty is called when entering the bindingProperty production.
	EnterBindingProperty(c *BindingPropertyContext)

	// EnterBindingProperty_Yield is called when entering the bindingProperty_Yield production.
	EnterBindingProperty_Yield(c *BindingProperty_YieldContext)

	// EnterBindingProperty_Await is called when entering the bindingProperty_Await production.
	EnterBindingProperty_Await(c *BindingProperty_AwaitContext)

	// EnterBindingProperty_Yield_Await is called when entering the bindingProperty_Yield_Await production.
	EnterBindingProperty_Yield_Await(c *BindingProperty_Yield_AwaitContext)

	// EnterBindingElement is called when entering the bindingElement production.
	EnterBindingElement(c *BindingElementContext)

	// EnterBindingElement_Yield is called when entering the bindingElement_Yield production.
	EnterBindingElement_Yield(c *BindingElement_YieldContext)

	// EnterBindingElement_Await is called when entering the bindingElement_Await production.
	EnterBindingElement_Await(c *BindingElement_AwaitContext)

	// EnterBindingElement_Yield_Await is called when entering the bindingElement_Yield_Await production.
	EnterBindingElement_Yield_Await(c *BindingElement_Yield_AwaitContext)

	// EnterSingleNameBinding is called when entering the singleNameBinding production.
	EnterSingleNameBinding(c *SingleNameBindingContext)

	// EnterSingleNameBinding_Yield is called when entering the singleNameBinding_Yield production.
	EnterSingleNameBinding_Yield(c *SingleNameBinding_YieldContext)

	// EnterSingleNameBinding_Await is called when entering the singleNameBinding_Await production.
	EnterSingleNameBinding_Await(c *SingleNameBinding_AwaitContext)

	// EnterSingleNameBinding_Yield_Await is called when entering the singleNameBinding_Yield_Await production.
	EnterSingleNameBinding_Yield_Await(c *SingleNameBinding_Yield_AwaitContext)

	// EnterBindingRestElement is called when entering the bindingRestElement production.
	EnterBindingRestElement(c *BindingRestElementContext)

	// EnterBindingRestElement_Yield is called when entering the bindingRestElement_Yield production.
	EnterBindingRestElement_Yield(c *BindingRestElement_YieldContext)

	// EnterBindingRestElement_Await is called when entering the bindingRestElement_Await production.
	EnterBindingRestElement_Await(c *BindingRestElement_AwaitContext)

	// EnterBindingRestElement_Yield_Await is called when entering the bindingRestElement_Yield_Await production.
	EnterBindingRestElement_Yield_Await(c *BindingRestElement_Yield_AwaitContext)

	// EnterEmptyStatement is called when entering the emptyStatement production.
	EnterEmptyStatement(c *EmptyStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterExpressionStatement_Yield is called when entering the expressionStatement_Yield production.
	EnterExpressionStatement_Yield(c *ExpressionStatement_YieldContext)

	// EnterExpressionStatement_Await is called when entering the expressionStatement_Await production.
	EnterExpressionStatement_Await(c *ExpressionStatement_AwaitContext)

	// EnterExpressionStatement_Yield_Await is called when entering the expressionStatement_Yield_Await production.
	EnterExpressionStatement_Yield_Await(c *ExpressionStatement_Yield_AwaitContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfStatement_Yield is called when entering the ifStatement_Yield production.
	EnterIfStatement_Yield(c *IfStatement_YieldContext)

	// EnterIfStatement_Await is called when entering the ifStatement_Await production.
	EnterIfStatement_Await(c *IfStatement_AwaitContext)

	// EnterIfStatement_Yield_Await is called when entering the ifStatement_Yield_Await production.
	EnterIfStatement_Yield_Await(c *IfStatement_Yield_AwaitContext)

	// EnterIfStatement_Return is called when entering the ifStatement_Return production.
	EnterIfStatement_Return(c *IfStatement_ReturnContext)

	// EnterIfStatement_Yield_Return is called when entering the ifStatement_Yield_Return production.
	EnterIfStatement_Yield_Return(c *IfStatement_Yield_ReturnContext)

	// EnterIfStatement_Await_Return is called when entering the ifStatement_Await_Return production.
	EnterIfStatement_Await_Return(c *IfStatement_Await_ReturnContext)

	// EnterIfStatement_Yield_Await_Return is called when entering the ifStatement_Yield_Await_Return production.
	EnterIfStatement_Yield_Await_Return(c *IfStatement_Yield_Await_ReturnContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterIterationStatement_Yield is called when entering the iterationStatement_Yield production.
	EnterIterationStatement_Yield(c *IterationStatement_YieldContext)

	// EnterIterationStatement_Await is called when entering the iterationStatement_Await production.
	EnterIterationStatement_Await(c *IterationStatement_AwaitContext)

	// EnterIterationStatement_Yield_Await is called when entering the iterationStatement_Yield_Await production.
	EnterIterationStatement_Yield_Await(c *IterationStatement_Yield_AwaitContext)

	// EnterIterationStatement_Return is called when entering the iterationStatement_Return production.
	EnterIterationStatement_Return(c *IterationStatement_ReturnContext)

	// EnterIterationStatement_Yield_Return is called when entering the iterationStatement_Yield_Return production.
	EnterIterationStatement_Yield_Return(c *IterationStatement_Yield_ReturnContext)

	// EnterIterationStatement_Await_Return is called when entering the iterationStatement_Await_Return production.
	EnterIterationStatement_Await_Return(c *IterationStatement_Await_ReturnContext)

	// EnterIterationStatement_Yield_Await_Return is called when entering the iterationStatement_Yield_Await_Return production.
	EnterIterationStatement_Yield_Await_Return(c *IterationStatement_Yield_Await_ReturnContext)

	// EnterForDeclaration is called when entering the forDeclaration production.
	EnterForDeclaration(c *ForDeclarationContext)

	// EnterForDeclaration_Yield is called when entering the forDeclaration_Yield production.
	EnterForDeclaration_Yield(c *ForDeclaration_YieldContext)

	// EnterForDeclaration_Await is called when entering the forDeclaration_Await production.
	EnterForDeclaration_Await(c *ForDeclaration_AwaitContext)

	// EnterForDeclaration_Yield_Await is called when entering the forDeclaration_Yield_Await production.
	EnterForDeclaration_Yield_Await(c *ForDeclaration_Yield_AwaitContext)

	// EnterForBinding is called when entering the forBinding production.
	EnterForBinding(c *ForBindingContext)

	// EnterForBinding_Yield is called when entering the forBinding_Yield production.
	EnterForBinding_Yield(c *ForBinding_YieldContext)

	// EnterForBinding_Await is called when entering the forBinding_Await production.
	EnterForBinding_Await(c *ForBinding_AwaitContext)

	// EnterForBinding_Yield_Await is called when entering the forBinding_Yield_Await production.
	EnterForBinding_Yield_Await(c *ForBinding_Yield_AwaitContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterContinueStatement_Yield is called when entering the continueStatement_Yield production.
	EnterContinueStatement_Yield(c *ContinueStatement_YieldContext)

	// EnterContinueStatement_Await is called when entering the continueStatement_Await production.
	EnterContinueStatement_Await(c *ContinueStatement_AwaitContext)

	// EnterContinueStatement_Yield_Await is called when entering the continueStatement_Yield_Await production.
	EnterContinueStatement_Yield_Await(c *ContinueStatement_Yield_AwaitContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterBreakStatement_Yield is called when entering the breakStatement_Yield production.
	EnterBreakStatement_Yield(c *BreakStatement_YieldContext)

	// EnterBreakStatement_Await is called when entering the breakStatement_Await production.
	EnterBreakStatement_Await(c *BreakStatement_AwaitContext)

	// EnterBreakStatement_Yield_Await is called when entering the breakStatement_Yield_Await production.
	EnterBreakStatement_Yield_Await(c *BreakStatement_Yield_AwaitContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterReturnStatement_Yield is called when entering the returnStatement_Yield production.
	EnterReturnStatement_Yield(c *ReturnStatement_YieldContext)

	// EnterReturnStatement_Await is called when entering the returnStatement_Await production.
	EnterReturnStatement_Await(c *ReturnStatement_AwaitContext)

	// EnterReturnStatement_Yield_Await is called when entering the returnStatement_Yield_Await production.
	EnterReturnStatement_Yield_Await(c *ReturnStatement_Yield_AwaitContext)

	// EnterWithStatement is called when entering the withStatement production.
	EnterWithStatement(c *WithStatementContext)

	// EnterWithStatement_Yield is called when entering the withStatement_Yield production.
	EnterWithStatement_Yield(c *WithStatement_YieldContext)

	// EnterWithStatement_Await is called when entering the withStatement_Await production.
	EnterWithStatement_Await(c *WithStatement_AwaitContext)

	// EnterWithStatement_Yield_Await is called when entering the withStatement_Yield_Await production.
	EnterWithStatement_Yield_Await(c *WithStatement_Yield_AwaitContext)

	// EnterWithStatement_Return is called when entering the withStatement_Return production.
	EnterWithStatement_Return(c *WithStatement_ReturnContext)

	// EnterWithStatement_Yield_Return is called when entering the withStatement_Yield_Return production.
	EnterWithStatement_Yield_Return(c *WithStatement_Yield_ReturnContext)

	// EnterWithStatement_Await_Return is called when entering the withStatement_Await_Return production.
	EnterWithStatement_Await_Return(c *WithStatement_Await_ReturnContext)

	// EnterWithStatement_Yield_Await_Return is called when entering the withStatement_Yield_Await_Return production.
	EnterWithStatement_Yield_Await_Return(c *WithStatement_Yield_Await_ReturnContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchStatement_Yield is called when entering the switchStatement_Yield production.
	EnterSwitchStatement_Yield(c *SwitchStatement_YieldContext)

	// EnterSwitchStatement_Await is called when entering the switchStatement_Await production.
	EnterSwitchStatement_Await(c *SwitchStatement_AwaitContext)

	// EnterSwitchStatement_Yield_Await is called when entering the switchStatement_Yield_Await production.
	EnterSwitchStatement_Yield_Await(c *SwitchStatement_Yield_AwaitContext)

	// EnterSwitchStatement_Return is called when entering the switchStatement_Return production.
	EnterSwitchStatement_Return(c *SwitchStatement_ReturnContext)

	// EnterSwitchStatement_Yield_Return is called when entering the switchStatement_Yield_Return production.
	EnterSwitchStatement_Yield_Return(c *SwitchStatement_Yield_ReturnContext)

	// EnterSwitchStatement_Await_Return is called when entering the switchStatement_Await_Return production.
	EnterSwitchStatement_Await_Return(c *SwitchStatement_Await_ReturnContext)

	// EnterSwitchStatement_Yield_Await_Return is called when entering the switchStatement_Yield_Await_Return production.
	EnterSwitchStatement_Yield_Await_Return(c *SwitchStatement_Yield_Await_ReturnContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterCaseBlock_Yield is called when entering the caseBlock_Yield production.
	EnterCaseBlock_Yield(c *CaseBlock_YieldContext)

	// EnterCaseBlock_Await is called when entering the caseBlock_Await production.
	EnterCaseBlock_Await(c *CaseBlock_AwaitContext)

	// EnterCaseBlock_Yield_Await is called when entering the caseBlock_Yield_Await production.
	EnterCaseBlock_Yield_Await(c *CaseBlock_Yield_AwaitContext)

	// EnterCaseBlock_Return is called when entering the caseBlock_Return production.
	EnterCaseBlock_Return(c *CaseBlock_ReturnContext)

	// EnterCaseBlock_Yield_Return is called when entering the caseBlock_Yield_Return production.
	EnterCaseBlock_Yield_Return(c *CaseBlock_Yield_ReturnContext)

	// EnterCaseBlock_Await_Return is called when entering the caseBlock_Await_Return production.
	EnterCaseBlock_Await_Return(c *CaseBlock_Await_ReturnContext)

	// EnterCaseBlock_Yield_Await_Return is called when entering the caseBlock_Yield_Await_Return production.
	EnterCaseBlock_Yield_Await_Return(c *CaseBlock_Yield_Await_ReturnContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterCaseClause_Yield is called when entering the caseClause_Yield production.
	EnterCaseClause_Yield(c *CaseClause_YieldContext)

	// EnterCaseClause_Await is called when entering the caseClause_Await production.
	EnterCaseClause_Await(c *CaseClause_AwaitContext)

	// EnterCaseClause_Yield_Await is called when entering the caseClause_Yield_Await production.
	EnterCaseClause_Yield_Await(c *CaseClause_Yield_AwaitContext)

	// EnterCaseClause_Return is called when entering the caseClause_Return production.
	EnterCaseClause_Return(c *CaseClause_ReturnContext)

	// EnterCaseClause_Yield_Return is called when entering the caseClause_Yield_Return production.
	EnterCaseClause_Yield_Return(c *CaseClause_Yield_ReturnContext)

	// EnterCaseClause_Await_Return is called when entering the caseClause_Await_Return production.
	EnterCaseClause_Await_Return(c *CaseClause_Await_ReturnContext)

	// EnterCaseClause_Yield_Await_Return is called when entering the caseClause_Yield_Await_Return production.
	EnterCaseClause_Yield_Await_Return(c *CaseClause_Yield_Await_ReturnContext)

	// EnterDefaultClause is called when entering the defaultClause production.
	EnterDefaultClause(c *DefaultClauseContext)

	// EnterDefaultClause_Yield is called when entering the defaultClause_Yield production.
	EnterDefaultClause_Yield(c *DefaultClause_YieldContext)

	// EnterDefaultClause_Await is called when entering the defaultClause_Await production.
	EnterDefaultClause_Await(c *DefaultClause_AwaitContext)

	// EnterDefaultClause_Yield_Await is called when entering the defaultClause_Yield_Await production.
	EnterDefaultClause_Yield_Await(c *DefaultClause_Yield_AwaitContext)

	// EnterDefaultClause_Return is called when entering the defaultClause_Return production.
	EnterDefaultClause_Return(c *DefaultClause_ReturnContext)

	// EnterDefaultClause_Yield_Return is called when entering the defaultClause_Yield_Return production.
	EnterDefaultClause_Yield_Return(c *DefaultClause_Yield_ReturnContext)

	// EnterDefaultClause_Await_Return is called when entering the defaultClause_Await_Return production.
	EnterDefaultClause_Await_Return(c *DefaultClause_Await_ReturnContext)

	// EnterDefaultClause_Yield_Await_Return is called when entering the defaultClause_Yield_Await_Return production.
	EnterDefaultClause_Yield_Await_Return(c *DefaultClause_Yield_Await_ReturnContext)

	// EnterLabelledStatement is called when entering the labelledStatement production.
	EnterLabelledStatement(c *LabelledStatementContext)

	// EnterLabelledStatement_Yield is called when entering the labelledStatement_Yield production.
	EnterLabelledStatement_Yield(c *LabelledStatement_YieldContext)

	// EnterLabelledStatement_Await is called when entering the labelledStatement_Await production.
	EnterLabelledStatement_Await(c *LabelledStatement_AwaitContext)

	// EnterLabelledStatement_Yield_Await is called when entering the labelledStatement_Yield_Await production.
	EnterLabelledStatement_Yield_Await(c *LabelledStatement_Yield_AwaitContext)

	// EnterLabelledStatement_Return is called when entering the labelledStatement_Return production.
	EnterLabelledStatement_Return(c *LabelledStatement_ReturnContext)

	// EnterLabelledStatement_Yield_Return is called when entering the labelledStatement_Yield_Return production.
	EnterLabelledStatement_Yield_Return(c *LabelledStatement_Yield_ReturnContext)

	// EnterLabelledStatement_Await_Return is called when entering the labelledStatement_Await_Return production.
	EnterLabelledStatement_Await_Return(c *LabelledStatement_Await_ReturnContext)

	// EnterLabelledStatement_Yield_Await_Return is called when entering the labelledStatement_Yield_Await_Return production.
	EnterLabelledStatement_Yield_Await_Return(c *LabelledStatement_Yield_Await_ReturnContext)

	// EnterLabelledItem is called when entering the labelledItem production.
	EnterLabelledItem(c *LabelledItemContext)

	// EnterLabelledItem_Yield is called when entering the labelledItem_Yield production.
	EnterLabelledItem_Yield(c *LabelledItem_YieldContext)

	// EnterLabelledItem_Await is called when entering the labelledItem_Await production.
	EnterLabelledItem_Await(c *LabelledItem_AwaitContext)

	// EnterLabelledItem_Yield_Await is called when entering the labelledItem_Yield_Await production.
	EnterLabelledItem_Yield_Await(c *LabelledItem_Yield_AwaitContext)

	// EnterLabelledItem_Return is called when entering the labelledItem_Return production.
	EnterLabelledItem_Return(c *LabelledItem_ReturnContext)

	// EnterLabelledItem_Yield_Return is called when entering the labelledItem_Yield_Return production.
	EnterLabelledItem_Yield_Return(c *LabelledItem_Yield_ReturnContext)

	// EnterLabelledItem_Await_Return is called when entering the labelledItem_Await_Return production.
	EnterLabelledItem_Await_Return(c *LabelledItem_Await_ReturnContext)

	// EnterLabelledItem_Yield_Await_Return is called when entering the labelledItem_Yield_Await_Return production.
	EnterLabelledItem_Yield_Await_Return(c *LabelledItem_Yield_Await_ReturnContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterThrowStatement_Yield is called when entering the throwStatement_Yield production.
	EnterThrowStatement_Yield(c *ThrowStatement_YieldContext)

	// EnterThrowStatement_Await is called when entering the throwStatement_Await production.
	EnterThrowStatement_Await(c *ThrowStatement_AwaitContext)

	// EnterThrowStatement_Yield_Await is called when entering the throwStatement_Yield_Await production.
	EnterThrowStatement_Yield_Await(c *ThrowStatement_Yield_AwaitContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterTryStatement_Yield is called when entering the tryStatement_Yield production.
	EnterTryStatement_Yield(c *TryStatement_YieldContext)

	// EnterTryStatement_Await is called when entering the tryStatement_Await production.
	EnterTryStatement_Await(c *TryStatement_AwaitContext)

	// EnterTryStatement_Yield_Await is called when entering the tryStatement_Yield_Await production.
	EnterTryStatement_Yield_Await(c *TryStatement_Yield_AwaitContext)

	// EnterTryStatement_Return is called when entering the tryStatement_Return production.
	EnterTryStatement_Return(c *TryStatement_ReturnContext)

	// EnterTryStatement_Yield_Return is called when entering the tryStatement_Yield_Return production.
	EnterTryStatement_Yield_Return(c *TryStatement_Yield_ReturnContext)

	// EnterTryStatement_Await_Return is called when entering the tryStatement_Await_Return production.
	EnterTryStatement_Await_Return(c *TryStatement_Await_ReturnContext)

	// EnterTryStatement_Yield_Await_Return is called when entering the tryStatement_Yield_Await_Return production.
	EnterTryStatement_Yield_Await_Return(c *TryStatement_Yield_Await_ReturnContext)

	// EnterCatch_ is called when entering the catch_ production.
	EnterCatch_(c *Catch_Context)

	// EnterCatch_Yield is called when entering the catch_Yield production.
	EnterCatch_Yield(c *Catch_YieldContext)

	// EnterCatch_Await is called when entering the catch_Await production.
	EnterCatch_Await(c *Catch_AwaitContext)

	// EnterCatch_Yield_Await is called when entering the catch_Yield_Await production.
	EnterCatch_Yield_Await(c *Catch_Yield_AwaitContext)

	// EnterCatch_Return is called when entering the catch_Return production.
	EnterCatch_Return(c *Catch_ReturnContext)

	// EnterCatch_Yield_Return is called when entering the catch_Yield_Return production.
	EnterCatch_Yield_Return(c *Catch_Yield_ReturnContext)

	// EnterCatch_Await_Return is called when entering the catch_Await_Return production.
	EnterCatch_Await_Return(c *Catch_Await_ReturnContext)

	// EnterCatch_Yield_Await_Return is called when entering the catch_Yield_Await_Return production.
	EnterCatch_Yield_Await_Return(c *Catch_Yield_Await_ReturnContext)

	// EnterFinally_ is called when entering the finally_ production.
	EnterFinally_(c *Finally_Context)

	// EnterFinally_Yield is called when entering the finally_Yield production.
	EnterFinally_Yield(c *Finally_YieldContext)

	// EnterFinally_Await is called when entering the finally_Await production.
	EnterFinally_Await(c *Finally_AwaitContext)

	// EnterFinally_Yield_Await is called when entering the finally_Yield_Await production.
	EnterFinally_Yield_Await(c *Finally_Yield_AwaitContext)

	// EnterFinally_Return is called when entering the finally_Return production.
	EnterFinally_Return(c *Finally_ReturnContext)

	// EnterFinally_Yield_Return is called when entering the finally_Yield_Return production.
	EnterFinally_Yield_Return(c *Finally_Yield_ReturnContext)

	// EnterFinally_Await_Return is called when entering the finally_Await_Return production.
	EnterFinally_Await_Return(c *Finally_Await_ReturnContext)

	// EnterFinally_Yield_Await_Return is called when entering the finally_Yield_Await_Return production.
	EnterFinally_Yield_Await_Return(c *Finally_Yield_Await_ReturnContext)

	// EnterCatchParameter is called when entering the catchParameter production.
	EnterCatchParameter(c *CatchParameterContext)

	// EnterCatchParameter_Yield is called when entering the catchParameter_Yield production.
	EnterCatchParameter_Yield(c *CatchParameter_YieldContext)

	// EnterCatchParameter_Await is called when entering the catchParameter_Await production.
	EnterCatchParameter_Await(c *CatchParameter_AwaitContext)

	// EnterCatchParameter_Yield_Await is called when entering the catchParameter_Yield_Await production.
	EnterCatchParameter_Yield_Await(c *CatchParameter_Yield_AwaitContext)

	// EnterDebuggerStatement is called when entering the debuggerStatement production.
	EnterDebuggerStatement(c *DebuggerStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionDeclaration_Yield is called when entering the functionDeclaration_Yield production.
	EnterFunctionDeclaration_Yield(c *FunctionDeclaration_YieldContext)

	// EnterFunctionDeclaration_Await is called when entering the functionDeclaration_Await production.
	EnterFunctionDeclaration_Await(c *FunctionDeclaration_AwaitContext)

	// EnterFunctionDeclaration_Yield_Await is called when entering the functionDeclaration_Yield_Await production.
	EnterFunctionDeclaration_Yield_Await(c *FunctionDeclaration_Yield_AwaitContext)

	// EnterFunctionDeclaration_Default is called when entering the functionDeclaration_Default production.
	EnterFunctionDeclaration_Default(c *FunctionDeclaration_DefaultContext)

	// EnterFunctionDeclaration_Yield_Default is called when entering the functionDeclaration_Yield_Default production.
	EnterFunctionDeclaration_Yield_Default(c *FunctionDeclaration_Yield_DefaultContext)

	// EnterFunctionDeclaration_Await_Default is called when entering the functionDeclaration_Await_Default production.
	EnterFunctionDeclaration_Await_Default(c *FunctionDeclaration_Await_DefaultContext)

	// EnterFunctionDeclaration_Yield_Await_Default is called when entering the functionDeclaration_Yield_Await_Default production.
	EnterFunctionDeclaration_Yield_Await_Default(c *FunctionDeclaration_Yield_Await_DefaultContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterUniqueFormalParameters is called when entering the uniqueFormalParameters production.
	EnterUniqueFormalParameters(c *UniqueFormalParametersContext)

	// EnterUniqueFormalParameters_Yield is called when entering the uniqueFormalParameters_Yield production.
	EnterUniqueFormalParameters_Yield(c *UniqueFormalParameters_YieldContext)

	// EnterUniqueFormalParameters_Await is called when entering the uniqueFormalParameters_Await production.
	EnterUniqueFormalParameters_Await(c *UniqueFormalParameters_AwaitContext)

	// EnterUniqueFormalParameters_Yield_Await is called when entering the uniqueFormalParameters_Yield_Await production.
	EnterUniqueFormalParameters_Yield_Await(c *UniqueFormalParameters_Yield_AwaitContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameters_Yield is called when entering the formalParameters_Yield production.
	EnterFormalParameters_Yield(c *FormalParameters_YieldContext)

	// EnterFormalParameters_Await is called when entering the formalParameters_Await production.
	EnterFormalParameters_Await(c *FormalParameters_AwaitContext)

	// EnterFormalParameters_Yield_Await is called when entering the formalParameters_Yield_Await production.
	EnterFormalParameters_Yield_Await(c *FormalParameters_Yield_AwaitContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameterList_Yield is called when entering the formalParameterList_Yield production.
	EnterFormalParameterList_Yield(c *FormalParameterList_YieldContext)

	// EnterFormalParameterList_Await is called when entering the formalParameterList_Await production.
	EnterFormalParameterList_Await(c *FormalParameterList_AwaitContext)

	// EnterFormalParameterList_Yield_Await is called when entering the formalParameterList_Yield_Await production.
	EnterFormalParameterList_Yield_Await(c *FormalParameterList_Yield_AwaitContext)

	// EnterFunctionRestParameter is called when entering the functionRestParameter production.
	EnterFunctionRestParameter(c *FunctionRestParameterContext)

	// EnterFunctionRestParameter_Yield is called when entering the functionRestParameter_Yield production.
	EnterFunctionRestParameter_Yield(c *FunctionRestParameter_YieldContext)

	// EnterFunctionRestParameter_Await is called when entering the functionRestParameter_Await production.
	EnterFunctionRestParameter_Await(c *FunctionRestParameter_AwaitContext)

	// EnterFunctionRestParameter_Yield_Await is called when entering the functionRestParameter_Yield_Await production.
	EnterFunctionRestParameter_Yield_Await(c *FunctionRestParameter_Yield_AwaitContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterFormalParameter_Yield is called when entering the formalParameter_Yield production.
	EnterFormalParameter_Yield(c *FormalParameter_YieldContext)

	// EnterFormalParameter_Await is called when entering the formalParameter_Await production.
	EnterFormalParameter_Await(c *FormalParameter_AwaitContext)

	// EnterFormalParameter_Yield_Await is called when entering the formalParameter_Yield_Await production.
	EnterFormalParameter_Yield_Await(c *FormalParameter_Yield_AwaitContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterFunctionBody_Yield is called when entering the functionBody_Yield production.
	EnterFunctionBody_Yield(c *FunctionBody_YieldContext)

	// EnterFunctionBody_Await is called when entering the functionBody_Await production.
	EnterFunctionBody_Await(c *FunctionBody_AwaitContext)

	// EnterFunctionBody_Yield_Await is called when entering the functionBody_Yield_Await production.
	EnterFunctionBody_Yield_Await(c *FunctionBody_Yield_AwaitContext)

	// EnterFunctionStatementList is called when entering the functionStatementList production.
	EnterFunctionStatementList(c *FunctionStatementListContext)

	// EnterFunctionStatementList_Yield is called when entering the functionStatementList_Yield production.
	EnterFunctionStatementList_Yield(c *FunctionStatementList_YieldContext)

	// EnterFunctionStatementList_Await is called when entering the functionStatementList_Await production.
	EnterFunctionStatementList_Await(c *FunctionStatementList_AwaitContext)

	// EnterFunctionStatementList_Yield_Await is called when entering the functionStatementList_Yield_Await production.
	EnterFunctionStatementList_Yield_Await(c *FunctionStatementList_Yield_AwaitContext)

	// EnterArrowFunction is called when entering the arrowFunction production.
	EnterArrowFunction(c *ArrowFunctionContext)

	// EnterArrowFunction_In is called when entering the arrowFunction_In production.
	EnterArrowFunction_In(c *ArrowFunction_InContext)

	// EnterArrowFunction_Yield is called when entering the arrowFunction_Yield production.
	EnterArrowFunction_Yield(c *ArrowFunction_YieldContext)

	// EnterArrowFunction_In_Yield is called when entering the arrowFunction_In_Yield production.
	EnterArrowFunction_In_Yield(c *ArrowFunction_In_YieldContext)

	// EnterArrowFunction_Await is called when entering the arrowFunction_Await production.
	EnterArrowFunction_Await(c *ArrowFunction_AwaitContext)

	// EnterArrowFunction_In_Await is called when entering the arrowFunction_In_Await production.
	EnterArrowFunction_In_Await(c *ArrowFunction_In_AwaitContext)

	// EnterArrowFunction_Yield_Await is called when entering the arrowFunction_Yield_Await production.
	EnterArrowFunction_Yield_Await(c *ArrowFunction_Yield_AwaitContext)

	// EnterArrowFunction_In_Yield_Await is called when entering the arrowFunction_In_Yield_Await production.
	EnterArrowFunction_In_Yield_Await(c *ArrowFunction_In_Yield_AwaitContext)

	// EnterArrowParameters is called when entering the arrowParameters production.
	EnterArrowParameters(c *ArrowParametersContext)

	// EnterArrowParameters_Yield is called when entering the arrowParameters_Yield production.
	EnterArrowParameters_Yield(c *ArrowParameters_YieldContext)

	// EnterArrowParameters_Await is called when entering the arrowParameters_Await production.
	EnterArrowParameters_Await(c *ArrowParameters_AwaitContext)

	// EnterArrowParameters_Yield_Await is called when entering the arrowParameters_Yield_Await production.
	EnterArrowParameters_Yield_Await(c *ArrowParameters_Yield_AwaitContext)

	// EnterConciseBody is called when entering the conciseBody production.
	EnterConciseBody(c *ConciseBodyContext)

	// EnterConciseBody_In is called when entering the conciseBody_In production.
	EnterConciseBody_In(c *ConciseBody_InContext)

	// EnterMethodDefinition is called when entering the methodDefinition production.
	EnterMethodDefinition(c *MethodDefinitionContext)

	// EnterMethodDefinition_Yield is called when entering the methodDefinition_Yield production.
	EnterMethodDefinition_Yield(c *MethodDefinition_YieldContext)

	// EnterMethodDefinition_Await is called when entering the methodDefinition_Await production.
	EnterMethodDefinition_Await(c *MethodDefinition_AwaitContext)

	// EnterMethodDefinition_Yield_Await is called when entering the methodDefinition_Yield_Await production.
	EnterMethodDefinition_Yield_Await(c *MethodDefinition_Yield_AwaitContext)

	// EnterPropertySetParameterList is called when entering the propertySetParameterList production.
	EnterPropertySetParameterList(c *PropertySetParameterListContext)

	// EnterGeneratorMethod is called when entering the generatorMethod production.
	EnterGeneratorMethod(c *GeneratorMethodContext)

	// EnterGeneratorMethod_Yield is called when entering the generatorMethod_Yield production.
	EnterGeneratorMethod_Yield(c *GeneratorMethod_YieldContext)

	// EnterGeneratorMethod_Await is called when entering the generatorMethod_Await production.
	EnterGeneratorMethod_Await(c *GeneratorMethod_AwaitContext)

	// EnterGeneratorMethod_Yield_Await is called when entering the generatorMethod_Yield_Await production.
	EnterGeneratorMethod_Yield_Await(c *GeneratorMethod_Yield_AwaitContext)

	// EnterGeneratorDeclaration is called when entering the generatorDeclaration production.
	EnterGeneratorDeclaration(c *GeneratorDeclarationContext)

	// EnterGeneratorDeclaration_Yield is called when entering the generatorDeclaration_Yield production.
	EnterGeneratorDeclaration_Yield(c *GeneratorDeclaration_YieldContext)

	// EnterGeneratorDeclaration_Await is called when entering the generatorDeclaration_Await production.
	EnterGeneratorDeclaration_Await(c *GeneratorDeclaration_AwaitContext)

	// EnterGeneratorDeclaration_Yield_Await is called when entering the generatorDeclaration_Yield_Await production.
	EnterGeneratorDeclaration_Yield_Await(c *GeneratorDeclaration_Yield_AwaitContext)

	// EnterGeneratorDeclaration_Default is called when entering the generatorDeclaration_Default production.
	EnterGeneratorDeclaration_Default(c *GeneratorDeclaration_DefaultContext)

	// EnterGeneratorDeclaration_Yield_Default is called when entering the generatorDeclaration_Yield_Default production.
	EnterGeneratorDeclaration_Yield_Default(c *GeneratorDeclaration_Yield_DefaultContext)

	// EnterGeneratorDeclaration_Await_Default is called when entering the generatorDeclaration_Await_Default production.
	EnterGeneratorDeclaration_Await_Default(c *GeneratorDeclaration_Await_DefaultContext)

	// EnterGeneratorDeclaration_Yield_Await_Default is called when entering the generatorDeclaration_Yield_Await_Default production.
	EnterGeneratorDeclaration_Yield_Await_Default(c *GeneratorDeclaration_Yield_Await_DefaultContext)

	// EnterGeneratorExpression is called when entering the generatorExpression production.
	EnterGeneratorExpression(c *GeneratorExpressionContext)

	// EnterGeneratorBody is called when entering the generatorBody production.
	EnterGeneratorBody(c *GeneratorBodyContext)

	// EnterYieldExpression is called when entering the yieldExpression production.
	EnterYieldExpression(c *YieldExpressionContext)

	// EnterYieldExpression_In is called when entering the yieldExpression_In production.
	EnterYieldExpression_In(c *YieldExpression_InContext)

	// EnterYieldExpression_Await is called when entering the yieldExpression_Await production.
	EnterYieldExpression_Await(c *YieldExpression_AwaitContext)

	// EnterYieldExpression_In_Await is called when entering the yieldExpression_In_Await production.
	EnterYieldExpression_In_Await(c *YieldExpression_In_AwaitContext)

	// EnterAsyncGeneratorMethod is called when entering the asyncGeneratorMethod production.
	EnterAsyncGeneratorMethod(c *AsyncGeneratorMethodContext)

	// EnterAsyncGeneratorMethod_Yield is called when entering the asyncGeneratorMethod_Yield production.
	EnterAsyncGeneratorMethod_Yield(c *AsyncGeneratorMethod_YieldContext)

	// EnterAsyncGeneratorMethod_Await is called when entering the asyncGeneratorMethod_Await production.
	EnterAsyncGeneratorMethod_Await(c *AsyncGeneratorMethod_AwaitContext)

	// EnterAsyncGeneratorMethod_Yield_Await is called when entering the asyncGeneratorMethod_Yield_Await production.
	EnterAsyncGeneratorMethod_Yield_Await(c *AsyncGeneratorMethod_Yield_AwaitContext)

	// EnterAsyncGeneratorDeclaration is called when entering the asyncGeneratorDeclaration production.
	EnterAsyncGeneratorDeclaration(c *AsyncGeneratorDeclarationContext)

	// EnterAsyncGeneratorDeclaration_Yield is called when entering the asyncGeneratorDeclaration_Yield production.
	EnterAsyncGeneratorDeclaration_Yield(c *AsyncGeneratorDeclaration_YieldContext)

	// EnterAsyncGeneratorDeclaration_Await is called when entering the asyncGeneratorDeclaration_Await production.
	EnterAsyncGeneratorDeclaration_Await(c *AsyncGeneratorDeclaration_AwaitContext)

	// EnterAsyncGeneratorDeclaration_Yield_Await is called when entering the asyncGeneratorDeclaration_Yield_Await production.
	EnterAsyncGeneratorDeclaration_Yield_Await(c *AsyncGeneratorDeclaration_Yield_AwaitContext)

	// EnterAsyncGeneratorDeclaration_Default is called when entering the asyncGeneratorDeclaration_Default production.
	EnterAsyncGeneratorDeclaration_Default(c *AsyncGeneratorDeclaration_DefaultContext)

	// EnterAsyncGeneratorDeclaration_Yield_Default is called when entering the asyncGeneratorDeclaration_Yield_Default production.
	EnterAsyncGeneratorDeclaration_Yield_Default(c *AsyncGeneratorDeclaration_Yield_DefaultContext)

	// EnterAsyncGeneratorDeclaration_Await_Default is called when entering the asyncGeneratorDeclaration_Await_Default production.
	EnterAsyncGeneratorDeclaration_Await_Default(c *AsyncGeneratorDeclaration_Await_DefaultContext)

	// EnterAsyncGeneratorDeclaration_Yield_Await_Default is called when entering the asyncGeneratorDeclaration_Yield_Await_Default production.
	EnterAsyncGeneratorDeclaration_Yield_Await_Default(c *AsyncGeneratorDeclaration_Yield_Await_DefaultContext)

	// EnterAsyncGeneratorExpression is called when entering the asyncGeneratorExpression production.
	EnterAsyncGeneratorExpression(c *AsyncGeneratorExpressionContext)

	// EnterAsyncGeneratorBody is called when entering the asyncGeneratorBody production.
	EnterAsyncGeneratorBody(c *AsyncGeneratorBodyContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassDeclaration_Yield is called when entering the classDeclaration_Yield production.
	EnterClassDeclaration_Yield(c *ClassDeclaration_YieldContext)

	// EnterClassDeclaration_Await is called when entering the classDeclaration_Await production.
	EnterClassDeclaration_Await(c *ClassDeclaration_AwaitContext)

	// EnterClassDeclaration_Yield_Await is called when entering the classDeclaration_Yield_Await production.
	EnterClassDeclaration_Yield_Await(c *ClassDeclaration_Yield_AwaitContext)

	// EnterClassDeclaration_Default is called when entering the classDeclaration_Default production.
	EnterClassDeclaration_Default(c *ClassDeclaration_DefaultContext)

	// EnterClassDeclaration_Yield_Default is called when entering the classDeclaration_Yield_Default production.
	EnterClassDeclaration_Yield_Default(c *ClassDeclaration_Yield_DefaultContext)

	// EnterClassDeclaration_Await_Default is called when entering the classDeclaration_Await_Default production.
	EnterClassDeclaration_Await_Default(c *ClassDeclaration_Await_DefaultContext)

	// EnterClassDeclaration_Yield_Await_Default is called when entering the classDeclaration_Yield_Await_Default production.
	EnterClassDeclaration_Yield_Await_Default(c *ClassDeclaration_Yield_Await_DefaultContext)

	// EnterClassExpression is called when entering the classExpression production.
	EnterClassExpression(c *ClassExpressionContext)

	// EnterClassExpression_Yield is called when entering the classExpression_Yield production.
	EnterClassExpression_Yield(c *ClassExpression_YieldContext)

	// EnterClassExpression_Await is called when entering the classExpression_Await production.
	EnterClassExpression_Await(c *ClassExpression_AwaitContext)

	// EnterClassExpression_Yield_Await is called when entering the classExpression_Yield_Await production.
	EnterClassExpression_Yield_Await(c *ClassExpression_Yield_AwaitContext)

	// EnterClassTail is called when entering the classTail production.
	EnterClassTail(c *ClassTailContext)

	// EnterClassTail_Yield is called when entering the classTail_Yield production.
	EnterClassTail_Yield(c *ClassTail_YieldContext)

	// EnterClassTail_Await is called when entering the classTail_Await production.
	EnterClassTail_Await(c *ClassTail_AwaitContext)

	// EnterClassTail_Yield_Await is called when entering the classTail_Yield_Await production.
	EnterClassTail_Yield_Await(c *ClassTail_Yield_AwaitContext)

	// EnterClassHeritage is called when entering the classHeritage production.
	EnterClassHeritage(c *ClassHeritageContext)

	// EnterClassHeritage_Yield is called when entering the classHeritage_Yield production.
	EnterClassHeritage_Yield(c *ClassHeritage_YieldContext)

	// EnterClassHeritage_Await is called when entering the classHeritage_Await production.
	EnterClassHeritage_Await(c *ClassHeritage_AwaitContext)

	// EnterClassHeritage_Yield_Await is called when entering the classHeritage_Yield_Await production.
	EnterClassHeritage_Yield_Await(c *ClassHeritage_Yield_AwaitContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassBody_Yield is called when entering the classBody_Yield production.
	EnterClassBody_Yield(c *ClassBody_YieldContext)

	// EnterClassBody_Await is called when entering the classBody_Await production.
	EnterClassBody_Await(c *ClassBody_AwaitContext)

	// EnterClassBody_Yield_Await is called when entering the classBody_Yield_Await production.
	EnterClassBody_Yield_Await(c *ClassBody_Yield_AwaitContext)

	// EnterClassElement is called when entering the classElement production.
	EnterClassElement(c *ClassElementContext)

	// EnterClassElement_Yield is called when entering the classElement_Yield production.
	EnterClassElement_Yield(c *ClassElement_YieldContext)

	// EnterClassElement_Await is called when entering the classElement_Await production.
	EnterClassElement_Await(c *ClassElement_AwaitContext)

	// EnterClassElement_Yield_Await is called when entering the classElement_Yield_Await production.
	EnterClassElement_Yield_Await(c *ClassElement_Yield_AwaitContext)

	// EnterAsyncFunctionDeclaration is called when entering the asyncFunctionDeclaration production.
	EnterAsyncFunctionDeclaration(c *AsyncFunctionDeclarationContext)

	// EnterAsyncFunctionDeclaration_Yield is called when entering the asyncFunctionDeclaration_Yield production.
	EnterAsyncFunctionDeclaration_Yield(c *AsyncFunctionDeclaration_YieldContext)

	// EnterAsyncFunctionDeclaration_Await is called when entering the asyncFunctionDeclaration_Await production.
	EnterAsyncFunctionDeclaration_Await(c *AsyncFunctionDeclaration_AwaitContext)

	// EnterAsyncFunctionDeclaration_Yield_Await is called when entering the asyncFunctionDeclaration_Yield_Await production.
	EnterAsyncFunctionDeclaration_Yield_Await(c *AsyncFunctionDeclaration_Yield_AwaitContext)

	// EnterAsyncFunctionDeclaration_Default is called when entering the asyncFunctionDeclaration_Default production.
	EnterAsyncFunctionDeclaration_Default(c *AsyncFunctionDeclaration_DefaultContext)

	// EnterAsyncFunctionDeclaration_Yield_Default is called when entering the asyncFunctionDeclaration_Yield_Default production.
	EnterAsyncFunctionDeclaration_Yield_Default(c *AsyncFunctionDeclaration_Yield_DefaultContext)

	// EnterAsyncFunctionDeclaration_Await_Default is called when entering the asyncFunctionDeclaration_Await_Default production.
	EnterAsyncFunctionDeclaration_Await_Default(c *AsyncFunctionDeclaration_Await_DefaultContext)

	// EnterAsyncFunctionDeclaration_Yield_Await_Default is called when entering the asyncFunctionDeclaration_Yield_Await_Default production.
	EnterAsyncFunctionDeclaration_Yield_Await_Default(c *AsyncFunctionDeclaration_Yield_Await_DefaultContext)

	// EnterAsyncFunctionExpression is called when entering the asyncFunctionExpression production.
	EnterAsyncFunctionExpression(c *AsyncFunctionExpressionContext)

	// EnterAsyncMethod is called when entering the asyncMethod production.
	EnterAsyncMethod(c *AsyncMethodContext)

	// EnterAsyncMethod_Yield is called when entering the asyncMethod_Yield production.
	EnterAsyncMethod_Yield(c *AsyncMethod_YieldContext)

	// EnterAsyncMethod_Await is called when entering the asyncMethod_Await production.
	EnterAsyncMethod_Await(c *AsyncMethod_AwaitContext)

	// EnterAsyncMethod_Yield_Await is called when entering the asyncMethod_Yield_Await production.
	EnterAsyncMethod_Yield_Await(c *AsyncMethod_Yield_AwaitContext)

	// EnterAsyncFunctionBody is called when entering the asyncFunctionBody production.
	EnterAsyncFunctionBody(c *AsyncFunctionBodyContext)

	// EnterAwaitExpression is called when entering the awaitExpression production.
	EnterAwaitExpression(c *AwaitExpressionContext)

	// EnterAwaitExpression_Yield is called when entering the awaitExpression_Yield production.
	EnterAwaitExpression_Yield(c *AwaitExpression_YieldContext)

	// EnterAsyncArrowFunction is called when entering the asyncArrowFunction production.
	EnterAsyncArrowFunction(c *AsyncArrowFunctionContext)

	// EnterAsyncArrowFunction_In is called when entering the asyncArrowFunction_In production.
	EnterAsyncArrowFunction_In(c *AsyncArrowFunction_InContext)

	// EnterAsyncArrowFunction_Yield is called when entering the asyncArrowFunction_Yield production.
	EnterAsyncArrowFunction_Yield(c *AsyncArrowFunction_YieldContext)

	// EnterAsyncArrowFunction_In_Yield is called when entering the asyncArrowFunction_In_Yield production.
	EnterAsyncArrowFunction_In_Yield(c *AsyncArrowFunction_In_YieldContext)

	// EnterAsyncArrowFunction_Await is called when entering the asyncArrowFunction_Await production.
	EnterAsyncArrowFunction_Await(c *AsyncArrowFunction_AwaitContext)

	// EnterAsyncArrowFunction_In_Await is called when entering the asyncArrowFunction_In_Await production.
	EnterAsyncArrowFunction_In_Await(c *AsyncArrowFunction_In_AwaitContext)

	// EnterAsyncArrowFunction_Yield_Await is called when entering the asyncArrowFunction_Yield_Await production.
	EnterAsyncArrowFunction_Yield_Await(c *AsyncArrowFunction_Yield_AwaitContext)

	// EnterAsyncArrowFunction_In_Yield_Await is called when entering the asyncArrowFunction_In_Yield_Await production.
	EnterAsyncArrowFunction_In_Yield_Await(c *AsyncArrowFunction_In_Yield_AwaitContext)

	// EnterAsyncArrowBindingIdentifier is called when entering the asyncArrowBindingIdentifier production.
	EnterAsyncArrowBindingIdentifier(c *AsyncArrowBindingIdentifierContext)

	// EnterAsyncArrowBindingIdentifier_Yield is called when entering the asyncArrowBindingIdentifier_Yield production.
	EnterAsyncArrowBindingIdentifier_Yield(c *AsyncArrowBindingIdentifier_YieldContext)

	// EnterCoverCallExpressionAndAsyncArrowHead is called when entering the coverCallExpressionAndAsyncArrowHead production.
	EnterCoverCallExpressionAndAsyncArrowHead(c *CoverCallExpressionAndAsyncArrowHeadContext)

	// EnterCoverCallExpressionAndAsyncArrowHead_Yield is called when entering the coverCallExpressionAndAsyncArrowHead_Yield production.
	EnterCoverCallExpressionAndAsyncArrowHead_Yield(c *CoverCallExpressionAndAsyncArrowHead_YieldContext)

	// EnterCoverCallExpressionAndAsyncArrowHead_Await is called when entering the coverCallExpressionAndAsyncArrowHead_Await production.
	EnterCoverCallExpressionAndAsyncArrowHead_Await(c *CoverCallExpressionAndAsyncArrowHead_AwaitContext)

	// EnterCoverCallExpressionAndAsyncArrowHead_Yield_Await is called when entering the coverCallExpressionAndAsyncArrowHead_Yield_Await production.
	EnterCoverCallExpressionAndAsyncArrowHead_Yield_Await(c *CoverCallExpressionAndAsyncArrowHead_Yield_AwaitContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterScriptBody is called when entering the scriptBody production.
	EnterScriptBody(c *ScriptBodyContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterModuleBody is called when entering the moduleBody production.
	EnterModuleBody(c *ModuleBodyContext)

	// EnterModuleItem is called when entering the moduleItem production.
	EnterModuleItem(c *ModuleItemContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterImportClause is called when entering the importClause production.
	EnterImportClause(c *ImportClauseContext)

	// EnterImportedDefaultBinding is called when entering the importedDefaultBinding production.
	EnterImportedDefaultBinding(c *ImportedDefaultBindingContext)

	// EnterNameSpaceImport is called when entering the nameSpaceImport production.
	EnterNameSpaceImport(c *NameSpaceImportContext)

	// EnterNamedImports is called when entering the namedImports production.
	EnterNamedImports(c *NamedImportsContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterImportsList is called when entering the importsList production.
	EnterImportsList(c *ImportsListContext)

	// EnterImportSpecifier is called when entering the importSpecifier production.
	EnterImportSpecifier(c *ImportSpecifierContext)

	// EnterModuleSpecifier is called when entering the moduleSpecifier production.
	EnterModuleSpecifier(c *ModuleSpecifierContext)

	// EnterImportedBinding is called when entering the importedBinding production.
	EnterImportedBinding(c *ImportedBindingContext)

	// EnterExportDeclaration is called when entering the exportDeclaration production.
	EnterExportDeclaration(c *ExportDeclarationContext)

	// EnterExportClause is called when entering the exportClause production.
	EnterExportClause(c *ExportClauseContext)

	// EnterExportsList is called when entering the exportsList production.
	EnterExportsList(c *ExportsListContext)

	// EnterExportSpecifier is called when entering the exportSpecifier production.
	EnterExportSpecifier(c *ExportSpecifierContext)

	// EnterAsyncConciseBody is called when entering the asyncConciseBody production.
	EnterAsyncConciseBody(c *AsyncConciseBodyContext)

	// EnterAsyncConciseBody_In is called when entering the asyncConciseBody_In production.
	EnterAsyncConciseBody_In(c *AsyncConciseBody_InContext)

	// ExitInputElementDiv is called when exiting the inputElementDiv production.
	ExitInputElementDiv(c *InputElementDivContext)

	// ExitInputElementRegExp is called when exiting the inputElementRegExp production.
	ExitInputElementRegExp(c *InputElementRegExpContext)

	// ExitInputElementRegExpOrTemplateTail is called when exiting the inputElementRegExpOrTemplateTail production.
	ExitInputElementRegExpOrTemplateTail(c *InputElementRegExpOrTemplateTailContext)

	// ExitInputElementTemplateTail is called when exiting the inputElementTemplateTail production.
	ExitInputElementTemplateTail(c *InputElementTemplateTailContext)

	// ExitRegularExpressionLiteral is called when exiting the regularExpressionLiteral production.
	ExitRegularExpressionLiteral(c *RegularExpressionLiteralContext)

	// ExitIdentifierReference is called when exiting the identifierReference production.
	ExitIdentifierReference(c *IdentifierReferenceContext)

	// ExitIdentifierReference_Yield is called when exiting the identifierReference_Yield production.
	ExitIdentifierReference_Yield(c *IdentifierReference_YieldContext)

	// ExitIdentifierReference_Await is called when exiting the identifierReference_Await production.
	ExitIdentifierReference_Await(c *IdentifierReference_AwaitContext)

	// ExitIdentifierReference_Yield_Await is called when exiting the identifierReference_Yield_Await production.
	ExitIdentifierReference_Yield_Await(c *IdentifierReference_Yield_AwaitContext)

	// ExitBindingIdentifier is called when exiting the bindingIdentifier production.
	ExitBindingIdentifier(c *BindingIdentifierContext)

	// ExitBindingIdentifier_Yield is called when exiting the bindingIdentifier_Yield production.
	ExitBindingIdentifier_Yield(c *BindingIdentifier_YieldContext)

	// ExitBindingIdentifier_Await is called when exiting the bindingIdentifier_Await production.
	ExitBindingIdentifier_Await(c *BindingIdentifier_AwaitContext)

	// ExitBindingIdentifier_Yield_Await is called when exiting the bindingIdentifier_Yield_Await production.
	ExitBindingIdentifier_Yield_Await(c *BindingIdentifier_Yield_AwaitContext)

	// ExitLabelIdentifier is called when exiting the labelIdentifier production.
	ExitLabelIdentifier(c *LabelIdentifierContext)

	// ExitLabelIdentifier_Yield is called when exiting the labelIdentifier_Yield production.
	ExitLabelIdentifier_Yield(c *LabelIdentifier_YieldContext)

	// ExitLabelIdentifier_Await is called when exiting the labelIdentifier_Await production.
	ExitLabelIdentifier_Await(c *LabelIdentifier_AwaitContext)

	// ExitLabelIdentifier_Yield_Await is called when exiting the labelIdentifier_Yield_Await production.
	ExitLabelIdentifier_Yield_Await(c *LabelIdentifier_Yield_AwaitContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitPrimaryExpression_Yield is called when exiting the primaryExpression_Yield production.
	ExitPrimaryExpression_Yield(c *PrimaryExpression_YieldContext)

	// ExitPrimaryExpression_Await is called when exiting the primaryExpression_Await production.
	ExitPrimaryExpression_Await(c *PrimaryExpression_AwaitContext)

	// ExitPrimaryExpression_Yield_Await is called when exiting the primaryExpression_Yield_Await production.
	ExitPrimaryExpression_Yield_Await(c *PrimaryExpression_Yield_AwaitContext)

	// ExitCoverParenthesizedExpressionAndArrowParameterList is called when exiting the coverParenthesizedExpressionAndArrowParameterList production.
	ExitCoverParenthesizedExpressionAndArrowParameterList(c *CoverParenthesizedExpressionAndArrowParameterListContext)

	// ExitCoverParenthesizedExpressionAndArrowParameterList_Yield is called when exiting the coverParenthesizedExpressionAndArrowParameterList_Yield production.
	ExitCoverParenthesizedExpressionAndArrowParameterList_Yield(c *CoverParenthesizedExpressionAndArrowParameterList_YieldContext)

	// ExitCoverParenthesizedExpressionAndArrowParameterList_Await is called when exiting the coverParenthesizedExpressionAndArrowParameterList_Await production.
	ExitCoverParenthesizedExpressionAndArrowParameterList_Await(c *CoverParenthesizedExpressionAndArrowParameterList_AwaitContext)

	// ExitCoverParenthesizedExpressionAndArrowParameterList_Yield_Await is called when exiting the coverParenthesizedExpressionAndArrowParameterList_Yield_Await production.
	ExitCoverParenthesizedExpressionAndArrowParameterList_Yield_Await(c *CoverParenthesizedExpressionAndArrowParameterList_Yield_AwaitContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitArrayLiteral_Yield is called when exiting the arrayLiteral_Yield production.
	ExitArrayLiteral_Yield(c *ArrayLiteral_YieldContext)

	// ExitArrayLiteral_Await is called when exiting the arrayLiteral_Await production.
	ExitArrayLiteral_Await(c *ArrayLiteral_AwaitContext)

	// ExitArrayLiteral_Yield_Await is called when exiting the arrayLiteral_Yield_Await production.
	ExitArrayLiteral_Yield_Await(c *ArrayLiteral_Yield_AwaitContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitElementList_Yield is called when exiting the elementList_Yield production.
	ExitElementList_Yield(c *ElementList_YieldContext)

	// ExitElementList_Await is called when exiting the elementList_Await production.
	ExitElementList_Await(c *ElementList_AwaitContext)

	// ExitElementList_Yield_Await is called when exiting the elementList_Yield_Await production.
	ExitElementList_Yield_Await(c *ElementList_Yield_AwaitContext)

	// ExitElision is called when exiting the elision production.
	ExitElision(c *ElisionContext)

	// ExitSpreadElement is called when exiting the spreadElement production.
	ExitSpreadElement(c *SpreadElementContext)

	// ExitSpreadElement_Yield is called when exiting the spreadElement_Yield production.
	ExitSpreadElement_Yield(c *SpreadElement_YieldContext)

	// ExitSpreadElement_Await is called when exiting the spreadElement_Await production.
	ExitSpreadElement_Await(c *SpreadElement_AwaitContext)

	// ExitSpreadElement_Yield_Await is called when exiting the spreadElement_Yield_Await production.
	ExitSpreadElement_Yield_Await(c *SpreadElement_Yield_AwaitContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitObjectLiteral_Yield is called when exiting the objectLiteral_Yield production.
	ExitObjectLiteral_Yield(c *ObjectLiteral_YieldContext)

	// ExitObjectLiteral_Await is called when exiting the objectLiteral_Await production.
	ExitObjectLiteral_Await(c *ObjectLiteral_AwaitContext)

	// ExitObjectLiteral_Yield_Await is called when exiting the objectLiteral_Yield_Await production.
	ExitObjectLiteral_Yield_Await(c *ObjectLiteral_Yield_AwaitContext)

	// ExitPropertyDefinitionList is called when exiting the propertyDefinitionList production.
	ExitPropertyDefinitionList(c *PropertyDefinitionListContext)

	// ExitPropertyDefinitionList_Yield is called when exiting the propertyDefinitionList_Yield production.
	ExitPropertyDefinitionList_Yield(c *PropertyDefinitionList_YieldContext)

	// ExitPropertyDefinitionList_Await is called when exiting the propertyDefinitionList_Await production.
	ExitPropertyDefinitionList_Await(c *PropertyDefinitionList_AwaitContext)

	// ExitPropertyDefinitionList_Yield_Await is called when exiting the propertyDefinitionList_Yield_Await production.
	ExitPropertyDefinitionList_Yield_Await(c *PropertyDefinitionList_Yield_AwaitContext)

	// ExitPropertyDefinition is called when exiting the propertyDefinition production.
	ExitPropertyDefinition(c *PropertyDefinitionContext)

	// ExitPropertyDefinition_Yield is called when exiting the propertyDefinition_Yield production.
	ExitPropertyDefinition_Yield(c *PropertyDefinition_YieldContext)

	// ExitPropertyDefinition_Await is called when exiting the propertyDefinition_Await production.
	ExitPropertyDefinition_Await(c *PropertyDefinition_AwaitContext)

	// ExitPropertyDefinition_Yield_Await is called when exiting the propertyDefinition_Yield_Await production.
	ExitPropertyDefinition_Yield_Await(c *PropertyDefinition_Yield_AwaitContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitPropertyName_Yield is called when exiting the propertyName_Yield production.
	ExitPropertyName_Yield(c *PropertyName_YieldContext)

	// ExitPropertyName_Await is called when exiting the propertyName_Await production.
	ExitPropertyName_Await(c *PropertyName_AwaitContext)

	// ExitPropertyName_Yield_Await is called when exiting the propertyName_Yield_Await production.
	ExitPropertyName_Yield_Await(c *PropertyName_Yield_AwaitContext)

	// ExitLiteralPropertyName is called when exiting the literalPropertyName production.
	ExitLiteralPropertyName(c *LiteralPropertyNameContext)

	// ExitComputedPropertyName is called when exiting the computedPropertyName production.
	ExitComputedPropertyName(c *ComputedPropertyNameContext)

	// ExitComputedPropertyName_Yield is called when exiting the computedPropertyName_Yield production.
	ExitComputedPropertyName_Yield(c *ComputedPropertyName_YieldContext)

	// ExitComputedPropertyName_Await is called when exiting the computedPropertyName_Await production.
	ExitComputedPropertyName_Await(c *ComputedPropertyName_AwaitContext)

	// ExitComputedPropertyName_Yield_Await is called when exiting the computedPropertyName_Yield_Await production.
	ExitComputedPropertyName_Yield_Await(c *ComputedPropertyName_Yield_AwaitContext)

	// ExitCoverInitializedName is called when exiting the coverInitializedName production.
	ExitCoverInitializedName(c *CoverInitializedNameContext)

	// ExitCoverInitializedName_Yield is called when exiting the coverInitializedName_Yield production.
	ExitCoverInitializedName_Yield(c *CoverInitializedName_YieldContext)

	// ExitCoverInitializedName_Await is called when exiting the coverInitializedName_Await production.
	ExitCoverInitializedName_Await(c *CoverInitializedName_AwaitContext)

	// ExitCoverInitializedName_Yield_Await is called when exiting the coverInitializedName_Yield_Await production.
	ExitCoverInitializedName_Yield_Await(c *CoverInitializedName_Yield_AwaitContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitInitializer_In is called when exiting the initializer_In production.
	ExitInitializer_In(c *Initializer_InContext)

	// ExitInitializer_Yield is called when exiting the initializer_Yield production.
	ExitInitializer_Yield(c *Initializer_YieldContext)

	// ExitInitializer_In_Yield is called when exiting the initializer_In_Yield production.
	ExitInitializer_In_Yield(c *Initializer_In_YieldContext)

	// ExitInitializer_Await is called when exiting the initializer_Await production.
	ExitInitializer_Await(c *Initializer_AwaitContext)

	// ExitInitializer_In_Await is called when exiting the initializer_In_Await production.
	ExitInitializer_In_Await(c *Initializer_In_AwaitContext)

	// ExitInitializer_Yield_Await is called when exiting the initializer_Yield_Await production.
	ExitInitializer_Yield_Await(c *Initializer_Yield_AwaitContext)

	// ExitInitializer_In_Yield_Await is called when exiting the initializer_In_Yield_Await production.
	ExitInitializer_In_Yield_Await(c *Initializer_In_Yield_AwaitContext)

	// ExitTemplateLiteral is called when exiting the templateLiteral production.
	ExitTemplateLiteral(c *TemplateLiteralContext)

	// ExitTemplateLiteral_Yield is called when exiting the templateLiteral_Yield production.
	ExitTemplateLiteral_Yield(c *TemplateLiteral_YieldContext)

	// ExitTemplateLiteral_Await is called when exiting the templateLiteral_Await production.
	ExitTemplateLiteral_Await(c *TemplateLiteral_AwaitContext)

	// ExitTemplateLiteral_Yield_Await is called when exiting the templateLiteral_Yield_Await production.
	ExitTemplateLiteral_Yield_Await(c *TemplateLiteral_Yield_AwaitContext)

	// ExitTemplateLiteral_Tagged is called when exiting the templateLiteral_Tagged production.
	ExitTemplateLiteral_Tagged(c *TemplateLiteral_TaggedContext)

	// ExitTemplateLiteral_Yield_Tagged is called when exiting the templateLiteral_Yield_Tagged production.
	ExitTemplateLiteral_Yield_Tagged(c *TemplateLiteral_Yield_TaggedContext)

	// ExitTemplateLiteral_Await_Tagged is called when exiting the templateLiteral_Await_Tagged production.
	ExitTemplateLiteral_Await_Tagged(c *TemplateLiteral_Await_TaggedContext)

	// ExitTemplateLiteral_Yield_Await_Tagged is called when exiting the templateLiteral_Yield_Await_Tagged production.
	ExitTemplateLiteral_Yield_Await_Tagged(c *TemplateLiteral_Yield_Await_TaggedContext)

	// ExitSubstitutionTemplate is called when exiting the substitutionTemplate production.
	ExitSubstitutionTemplate(c *SubstitutionTemplateContext)

	// ExitSubstitutionTemplate_Yield is called when exiting the substitutionTemplate_Yield production.
	ExitSubstitutionTemplate_Yield(c *SubstitutionTemplate_YieldContext)

	// ExitSubstitutionTemplate_Await is called when exiting the substitutionTemplate_Await production.
	ExitSubstitutionTemplate_Await(c *SubstitutionTemplate_AwaitContext)

	// ExitSubstitutionTemplate_Yield_Await is called when exiting the substitutionTemplate_Yield_Await production.
	ExitSubstitutionTemplate_Yield_Await(c *SubstitutionTemplate_Yield_AwaitContext)

	// ExitSubstitutionTemplate_Tagged is called when exiting the substitutionTemplate_Tagged production.
	ExitSubstitutionTemplate_Tagged(c *SubstitutionTemplate_TaggedContext)

	// ExitSubstitutionTemplate_Yield_Tagged is called when exiting the substitutionTemplate_Yield_Tagged production.
	ExitSubstitutionTemplate_Yield_Tagged(c *SubstitutionTemplate_Yield_TaggedContext)

	// ExitSubstitutionTemplate_Await_Tagged is called when exiting the substitutionTemplate_Await_Tagged production.
	ExitSubstitutionTemplate_Await_Tagged(c *SubstitutionTemplate_Await_TaggedContext)

	// ExitSubstitutionTemplate_Yield_Await_Tagged is called when exiting the substitutionTemplate_Yield_Await_Tagged production.
	ExitSubstitutionTemplate_Yield_Await_Tagged(c *SubstitutionTemplate_Yield_Await_TaggedContext)

	// ExitTemplateSpans is called when exiting the templateSpans production.
	ExitTemplateSpans(c *TemplateSpansContext)

	// ExitTemplateSpans_Yield is called when exiting the templateSpans_Yield production.
	ExitTemplateSpans_Yield(c *TemplateSpans_YieldContext)

	// ExitTemplateSpans_Await is called when exiting the templateSpans_Await production.
	ExitTemplateSpans_Await(c *TemplateSpans_AwaitContext)

	// ExitTemplateSpans_Yield_Await is called when exiting the templateSpans_Yield_Await production.
	ExitTemplateSpans_Yield_Await(c *TemplateSpans_Yield_AwaitContext)

	// ExitTemplateSpans_Tagged is called when exiting the templateSpans_Tagged production.
	ExitTemplateSpans_Tagged(c *TemplateSpans_TaggedContext)

	// ExitTemplateSpans_Yield_Tagged is called when exiting the templateSpans_Yield_Tagged production.
	ExitTemplateSpans_Yield_Tagged(c *TemplateSpans_Yield_TaggedContext)

	// ExitTemplateSpans_Await_Tagged is called when exiting the templateSpans_Await_Tagged production.
	ExitTemplateSpans_Await_Tagged(c *TemplateSpans_Await_TaggedContext)

	// ExitTemplateSpans_Yield_Await_Tagged is called when exiting the templateSpans_Yield_Await_Tagged production.
	ExitTemplateSpans_Yield_Await_Tagged(c *TemplateSpans_Yield_Await_TaggedContext)

	// ExitTemplateMiddleList is called when exiting the templateMiddleList production.
	ExitTemplateMiddleList(c *TemplateMiddleListContext)

	// ExitTemplateMiddleList_Yield is called when exiting the templateMiddleList_Yield production.
	ExitTemplateMiddleList_Yield(c *TemplateMiddleList_YieldContext)

	// ExitTemplateMiddleList_Await is called when exiting the templateMiddleList_Await production.
	ExitTemplateMiddleList_Await(c *TemplateMiddleList_AwaitContext)

	// ExitTemplateMiddleList_Yield_Await is called when exiting the templateMiddleList_Yield_Await production.
	ExitTemplateMiddleList_Yield_Await(c *TemplateMiddleList_Yield_AwaitContext)

	// ExitTemplateMiddleList_Tagged is called when exiting the templateMiddleList_Tagged production.
	ExitTemplateMiddleList_Tagged(c *TemplateMiddleList_TaggedContext)

	// ExitTemplateMiddleList_Yield_Tagged is called when exiting the templateMiddleList_Yield_Tagged production.
	ExitTemplateMiddleList_Yield_Tagged(c *TemplateMiddleList_Yield_TaggedContext)

	// ExitTemplateMiddleList_Await_Tagged is called when exiting the templateMiddleList_Await_Tagged production.
	ExitTemplateMiddleList_Await_Tagged(c *TemplateMiddleList_Await_TaggedContext)

	// ExitTemplateMiddleList_Yield_Await_Tagged is called when exiting the templateMiddleList_Yield_Await_Tagged production.
	ExitTemplateMiddleList_Yield_Await_Tagged(c *TemplateMiddleList_Yield_Await_TaggedContext)

	// ExitMemberExpression is called when exiting the memberExpression production.
	ExitMemberExpression(c *MemberExpressionContext)

	// ExitMemberExpression_Yield is called when exiting the memberExpression_Yield production.
	ExitMemberExpression_Yield(c *MemberExpression_YieldContext)

	// ExitMemberExpression_Await is called when exiting the memberExpression_Await production.
	ExitMemberExpression_Await(c *MemberExpression_AwaitContext)

	// ExitMemberExpression_Yield_Await is called when exiting the memberExpression_Yield_Await production.
	ExitMemberExpression_Yield_Await(c *MemberExpression_Yield_AwaitContext)

	// ExitSuperProperty is called when exiting the superProperty production.
	ExitSuperProperty(c *SuperPropertyContext)

	// ExitSuperProperty_Yield is called when exiting the superProperty_Yield production.
	ExitSuperProperty_Yield(c *SuperProperty_YieldContext)

	// ExitSuperProperty_Await is called when exiting the superProperty_Await production.
	ExitSuperProperty_Await(c *SuperProperty_AwaitContext)

	// ExitSuperProperty_Yield_Await is called when exiting the superProperty_Yield_Await production.
	ExitSuperProperty_Yield_Await(c *SuperProperty_Yield_AwaitContext)

	// ExitMetaProperty is called when exiting the metaProperty production.
	ExitMetaProperty(c *MetaPropertyContext)

	// ExitNewTarget is called when exiting the newTarget production.
	ExitNewTarget(c *NewTargetContext)

	// ExitNewExpression is called when exiting the newExpression production.
	ExitNewExpression(c *NewExpressionContext)

	// ExitNewExpression_Yield is called when exiting the newExpression_Yield production.
	ExitNewExpression_Yield(c *NewExpression_YieldContext)

	// ExitNewExpression_Await is called when exiting the newExpression_Await production.
	ExitNewExpression_Await(c *NewExpression_AwaitContext)

	// ExitNewExpression_Yield_Await is called when exiting the newExpression_Yield_Await production.
	ExitNewExpression_Yield_Await(c *NewExpression_Yield_AwaitContext)

	// ExitCallExpression is called when exiting the callExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitCallExpression_Yield is called when exiting the callExpression_Yield production.
	ExitCallExpression_Yield(c *CallExpression_YieldContext)

	// ExitCallExpression_Await is called when exiting the callExpression_Await production.
	ExitCallExpression_Await(c *CallExpression_AwaitContext)

	// ExitCallExpression_Yield_Await is called when exiting the callExpression_Yield_Await production.
	ExitCallExpression_Yield_Await(c *CallExpression_Yield_AwaitContext)

	// ExitSuperCall is called when exiting the superCall production.
	ExitSuperCall(c *SuperCallContext)

	// ExitSuperCall_Yield is called when exiting the superCall_Yield production.
	ExitSuperCall_Yield(c *SuperCall_YieldContext)

	// ExitSuperCall_Await is called when exiting the superCall_Await production.
	ExitSuperCall_Await(c *SuperCall_AwaitContext)

	// ExitSuperCall_Yield_Await is called when exiting the superCall_Yield_Await production.
	ExitSuperCall_Yield_Await(c *SuperCall_Yield_AwaitContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArguments_Yield is called when exiting the arguments_Yield production.
	ExitArguments_Yield(c *Arguments_YieldContext)

	// ExitArguments_Await is called when exiting the arguments_Await production.
	ExitArguments_Await(c *Arguments_AwaitContext)

	// ExitArguments_Yield_Await is called when exiting the arguments_Yield_Await production.
	ExitArguments_Yield_Await(c *Arguments_Yield_AwaitContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArgumentList_Yield is called when exiting the argumentList_Yield production.
	ExitArgumentList_Yield(c *ArgumentList_YieldContext)

	// ExitArgumentList_Await is called when exiting the argumentList_Await production.
	ExitArgumentList_Await(c *ArgumentList_AwaitContext)

	// ExitArgumentList_Yield_Await is called when exiting the argumentList_Yield_Await production.
	ExitArgumentList_Yield_Await(c *ArgumentList_Yield_AwaitContext)

	// ExitLeftHandSideExpression is called when exiting the leftHandSideExpression production.
	ExitLeftHandSideExpression(c *LeftHandSideExpressionContext)

	// ExitLeftHandSideExpression_Yield is called when exiting the leftHandSideExpression_Yield production.
	ExitLeftHandSideExpression_Yield(c *LeftHandSideExpression_YieldContext)

	// ExitLeftHandSideExpression_Await is called when exiting the leftHandSideExpression_Await production.
	ExitLeftHandSideExpression_Await(c *LeftHandSideExpression_AwaitContext)

	// ExitLeftHandSideExpression_Yield_Await is called when exiting the leftHandSideExpression_Yield_Await production.
	ExitLeftHandSideExpression_Yield_Await(c *LeftHandSideExpression_Yield_AwaitContext)

	// ExitUpdateExpression is called when exiting the updateExpression production.
	ExitUpdateExpression(c *UpdateExpressionContext)

	// ExitUpdateExpression_Yield is called when exiting the updateExpression_Yield production.
	ExitUpdateExpression_Yield(c *UpdateExpression_YieldContext)

	// ExitUpdateExpression_Await is called when exiting the updateExpression_Await production.
	ExitUpdateExpression_Await(c *UpdateExpression_AwaitContext)

	// ExitUpdateExpression_Yield_Await is called when exiting the updateExpression_Yield_Await production.
	ExitUpdateExpression_Yield_Await(c *UpdateExpression_Yield_AwaitContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryExpression_Yield is called when exiting the unaryExpression_Yield production.
	ExitUnaryExpression_Yield(c *UnaryExpression_YieldContext)

	// ExitUnaryExpression_Await is called when exiting the unaryExpression_Await production.
	ExitUnaryExpression_Await(c *UnaryExpression_AwaitContext)

	// ExitUnaryExpression_Yield_Await is called when exiting the unaryExpression_Yield_Await production.
	ExitUnaryExpression_Yield_Await(c *UnaryExpression_Yield_AwaitContext)

	// ExitExponentationExpression is called when exiting the exponentationExpression production.
	ExitExponentationExpression(c *ExponentationExpressionContext)

	// ExitExponentationExpression_Yield is called when exiting the exponentationExpression_Yield production.
	ExitExponentationExpression_Yield(c *ExponentationExpression_YieldContext)

	// ExitExponentationExpression_Await is called when exiting the exponentationExpression_Await production.
	ExitExponentationExpression_Await(c *ExponentationExpression_AwaitContext)

	// ExitExponentationExpression_Yield_Await is called when exiting the exponentationExpression_Yield_Await production.
	ExitExponentationExpression_Yield_Await(c *ExponentationExpression_Yield_AwaitContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitMultiplicativeExpression_Yield is called when exiting the multiplicativeExpression_Yield production.
	ExitMultiplicativeExpression_Yield(c *MultiplicativeExpression_YieldContext)

	// ExitMultiplicativeExpression_Await is called when exiting the multiplicativeExpression_Await production.
	ExitMultiplicativeExpression_Await(c *MultiplicativeExpression_AwaitContext)

	// ExitMultiplicativeExpression_Yield_Await is called when exiting the multiplicativeExpression_Yield_Await production.
	ExitMultiplicativeExpression_Yield_Await(c *MultiplicativeExpression_Yield_AwaitContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitAdditiveExpression_Yield is called when exiting the additiveExpression_Yield production.
	ExitAdditiveExpression_Yield(c *AdditiveExpression_YieldContext)

	// ExitAdditiveExpression_Await is called when exiting the additiveExpression_Await production.
	ExitAdditiveExpression_Await(c *AdditiveExpression_AwaitContext)

	// ExitAdditiveExpression_Yield_Await is called when exiting the additiveExpression_Yield_Await production.
	ExitAdditiveExpression_Yield_Await(c *AdditiveExpression_Yield_AwaitContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitShiftExpression_Yield is called when exiting the shiftExpression_Yield production.
	ExitShiftExpression_Yield(c *ShiftExpression_YieldContext)

	// ExitShiftExpression_Await is called when exiting the shiftExpression_Await production.
	ExitShiftExpression_Await(c *ShiftExpression_AwaitContext)

	// ExitShiftExpression_Yield_Await is called when exiting the shiftExpression_Yield_Await production.
	ExitShiftExpression_Yield_Await(c *ShiftExpression_Yield_AwaitContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitRelationalExpression_In is called when exiting the relationalExpression_In production.
	ExitRelationalExpression_In(c *RelationalExpression_InContext)

	// ExitRelationalExpression_Yield is called when exiting the relationalExpression_Yield production.
	ExitRelationalExpression_Yield(c *RelationalExpression_YieldContext)

	// ExitRelationalExpression_In_Yield is called when exiting the relationalExpression_In_Yield production.
	ExitRelationalExpression_In_Yield(c *RelationalExpression_In_YieldContext)

	// ExitRelationalExpression_Await is called when exiting the relationalExpression_Await production.
	ExitRelationalExpression_Await(c *RelationalExpression_AwaitContext)

	// ExitRelationalExpression_In_Await is called when exiting the relationalExpression_In_Await production.
	ExitRelationalExpression_In_Await(c *RelationalExpression_In_AwaitContext)

	// ExitRelationalExpression_Yield_Await is called when exiting the relationalExpression_Yield_Await production.
	ExitRelationalExpression_Yield_Await(c *RelationalExpression_Yield_AwaitContext)

	// ExitRelationalExpression_In_Yield_Await is called when exiting the relationalExpression_In_Yield_Await production.
	ExitRelationalExpression_In_Yield_Await(c *RelationalExpression_In_Yield_AwaitContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitEqualityExpression_In is called when exiting the equalityExpression_In production.
	ExitEqualityExpression_In(c *EqualityExpression_InContext)

	// ExitEqualityExpression_Yield is called when exiting the equalityExpression_Yield production.
	ExitEqualityExpression_Yield(c *EqualityExpression_YieldContext)

	// ExitEqualityExpression_In_Yield is called when exiting the equalityExpression_In_Yield production.
	ExitEqualityExpression_In_Yield(c *EqualityExpression_In_YieldContext)

	// ExitEqualityExpression_Await is called when exiting the equalityExpression_Await production.
	ExitEqualityExpression_Await(c *EqualityExpression_AwaitContext)

	// ExitEqualityExpression_In_Await is called when exiting the equalityExpression_In_Await production.
	ExitEqualityExpression_In_Await(c *EqualityExpression_In_AwaitContext)

	// ExitEqualityExpression_Yield_Await is called when exiting the equalityExpression_Yield_Await production.
	ExitEqualityExpression_Yield_Await(c *EqualityExpression_Yield_AwaitContext)

	// ExitEqualityExpression_In_Yield_Await is called when exiting the equalityExpression_In_Yield_Await production.
	ExitEqualityExpression_In_Yield_Await(c *EqualityExpression_In_Yield_AwaitContext)

	// ExitBitwiseANDExpression is called when exiting the bitwiseANDExpression production.
	ExitBitwiseANDExpression(c *BitwiseANDExpressionContext)

	// ExitBitwiseANDExpression_In is called when exiting the bitwiseANDExpression_In production.
	ExitBitwiseANDExpression_In(c *BitwiseANDExpression_InContext)

	// ExitBitwiseANDExpression_Yield is called when exiting the bitwiseANDExpression_Yield production.
	ExitBitwiseANDExpression_Yield(c *BitwiseANDExpression_YieldContext)

	// ExitBitwiseANDExpression_In_Yield is called when exiting the bitwiseANDExpression_In_Yield production.
	ExitBitwiseANDExpression_In_Yield(c *BitwiseANDExpression_In_YieldContext)

	// ExitBitwiseANDExpression_Await is called when exiting the bitwiseANDExpression_Await production.
	ExitBitwiseANDExpression_Await(c *BitwiseANDExpression_AwaitContext)

	// ExitBitwiseANDExpression_In_Await is called when exiting the bitwiseANDExpression_In_Await production.
	ExitBitwiseANDExpression_In_Await(c *BitwiseANDExpression_In_AwaitContext)

	// ExitBitwiseANDExpression_Yield_Await is called when exiting the bitwiseANDExpression_Yield_Await production.
	ExitBitwiseANDExpression_Yield_Await(c *BitwiseANDExpression_Yield_AwaitContext)

	// ExitBitwiseANDExpression_In_Yield_Await is called when exiting the bitwiseANDExpression_In_Yield_Await production.
	ExitBitwiseANDExpression_In_Yield_Await(c *BitwiseANDExpression_In_Yield_AwaitContext)

	// ExitBitwiseXORExpression is called when exiting the bitwiseXORExpression production.
	ExitBitwiseXORExpression(c *BitwiseXORExpressionContext)

	// ExitBitwiseXORExpression_In is called when exiting the bitwiseXORExpression_In production.
	ExitBitwiseXORExpression_In(c *BitwiseXORExpression_InContext)

	// ExitBitwiseXORExpression_Yield is called when exiting the bitwiseXORExpression_Yield production.
	ExitBitwiseXORExpression_Yield(c *BitwiseXORExpression_YieldContext)

	// ExitBitwiseXORExpression_In_Yield is called when exiting the bitwiseXORExpression_In_Yield production.
	ExitBitwiseXORExpression_In_Yield(c *BitwiseXORExpression_In_YieldContext)

	// ExitBitwiseXORExpression_Await is called when exiting the bitwiseXORExpression_Await production.
	ExitBitwiseXORExpression_Await(c *BitwiseXORExpression_AwaitContext)

	// ExitBitwiseXORExpression_In_Await is called when exiting the bitwiseXORExpression_In_Await production.
	ExitBitwiseXORExpression_In_Await(c *BitwiseXORExpression_In_AwaitContext)

	// ExitBitwiseXORExpression_Yield_Await is called when exiting the bitwiseXORExpression_Yield_Await production.
	ExitBitwiseXORExpression_Yield_Await(c *BitwiseXORExpression_Yield_AwaitContext)

	// ExitBitwiseXORExpression_In_Yield_Await is called when exiting the bitwiseXORExpression_In_Yield_Await production.
	ExitBitwiseXORExpression_In_Yield_Await(c *BitwiseXORExpression_In_Yield_AwaitContext)

	// ExitBitwiseORExpression is called when exiting the bitwiseORExpression production.
	ExitBitwiseORExpression(c *BitwiseORExpressionContext)

	// ExitBitwiseORExpression_In is called when exiting the bitwiseORExpression_In production.
	ExitBitwiseORExpression_In(c *BitwiseORExpression_InContext)

	// ExitBitwiseORExpression_Yield is called when exiting the bitwiseORExpression_Yield production.
	ExitBitwiseORExpression_Yield(c *BitwiseORExpression_YieldContext)

	// ExitBitwiseORExpression_In_Yield is called when exiting the bitwiseORExpression_In_Yield production.
	ExitBitwiseORExpression_In_Yield(c *BitwiseORExpression_In_YieldContext)

	// ExitBitwiseORExpression_Await is called when exiting the bitwiseORExpression_Await production.
	ExitBitwiseORExpression_Await(c *BitwiseORExpression_AwaitContext)

	// ExitBitwiseORExpression_In_Await is called when exiting the bitwiseORExpression_In_Await production.
	ExitBitwiseORExpression_In_Await(c *BitwiseORExpression_In_AwaitContext)

	// ExitBitwiseORExpression_Yield_Await is called when exiting the bitwiseORExpression_Yield_Await production.
	ExitBitwiseORExpression_Yield_Await(c *BitwiseORExpression_Yield_AwaitContext)

	// ExitBitwiseORExpression_In_Yield_Await is called when exiting the bitwiseORExpression_In_Yield_Await production.
	ExitBitwiseORExpression_In_Yield_Await(c *BitwiseORExpression_In_Yield_AwaitContext)

	// ExitLogicalANDExpression is called when exiting the logicalANDExpression production.
	ExitLogicalANDExpression(c *LogicalANDExpressionContext)

	// ExitLogicalANDExpression_In is called when exiting the logicalANDExpression_In production.
	ExitLogicalANDExpression_In(c *LogicalANDExpression_InContext)

	// ExitLogicalANDExpression_Yield is called when exiting the logicalANDExpression_Yield production.
	ExitLogicalANDExpression_Yield(c *LogicalANDExpression_YieldContext)

	// ExitLogicalANDExpression_In_Yield is called when exiting the logicalANDExpression_In_Yield production.
	ExitLogicalANDExpression_In_Yield(c *LogicalANDExpression_In_YieldContext)

	// ExitLogicalANDExpression_Await is called when exiting the logicalANDExpression_Await production.
	ExitLogicalANDExpression_Await(c *LogicalANDExpression_AwaitContext)

	// ExitLogicalANDExpression_In_Await is called when exiting the logicalANDExpression_In_Await production.
	ExitLogicalANDExpression_In_Await(c *LogicalANDExpression_In_AwaitContext)

	// ExitLogicalANDExpression_Yield_Await is called when exiting the logicalANDExpression_Yield_Await production.
	ExitLogicalANDExpression_Yield_Await(c *LogicalANDExpression_Yield_AwaitContext)

	// ExitLogicalANDExpression_In_Yield_Await is called when exiting the logicalANDExpression_In_Yield_Await production.
	ExitLogicalANDExpression_In_Yield_Await(c *LogicalANDExpression_In_Yield_AwaitContext)

	// ExitLogicalORExpression is called when exiting the logicalORExpression production.
	ExitLogicalORExpression(c *LogicalORExpressionContext)

	// ExitLogicalORExpression_In is called when exiting the logicalORExpression_In production.
	ExitLogicalORExpression_In(c *LogicalORExpression_InContext)

	// ExitLogicalORExpression_Yield is called when exiting the logicalORExpression_Yield production.
	ExitLogicalORExpression_Yield(c *LogicalORExpression_YieldContext)

	// ExitLogicalORExpression_In_Yield is called when exiting the logicalORExpression_In_Yield production.
	ExitLogicalORExpression_In_Yield(c *LogicalORExpression_In_YieldContext)

	// ExitLogicalORExpression_Await is called when exiting the logicalORExpression_Await production.
	ExitLogicalORExpression_Await(c *LogicalORExpression_AwaitContext)

	// ExitLogicalORExpression_In_Await is called when exiting the logicalORExpression_In_Await production.
	ExitLogicalORExpression_In_Await(c *LogicalORExpression_In_AwaitContext)

	// ExitLogicalORExpression_Yield_Await is called when exiting the logicalORExpression_Yield_Await production.
	ExitLogicalORExpression_Yield_Await(c *LogicalORExpression_Yield_AwaitContext)

	// ExitLogicalORExpression_In_Yield_Await is called when exiting the logicalORExpression_In_Yield_Await production.
	ExitLogicalORExpression_In_Yield_Await(c *LogicalORExpression_In_Yield_AwaitContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitConditionalExpression_In is called when exiting the conditionalExpression_In production.
	ExitConditionalExpression_In(c *ConditionalExpression_InContext)

	// ExitConditionalExpression_Yield is called when exiting the conditionalExpression_Yield production.
	ExitConditionalExpression_Yield(c *ConditionalExpression_YieldContext)

	// ExitConditionalExpression_In_Yield is called when exiting the conditionalExpression_In_Yield production.
	ExitConditionalExpression_In_Yield(c *ConditionalExpression_In_YieldContext)

	// ExitConditionalExpression_Await is called when exiting the conditionalExpression_Await production.
	ExitConditionalExpression_Await(c *ConditionalExpression_AwaitContext)

	// ExitConditionalExpression_In_Await is called when exiting the conditionalExpression_In_Await production.
	ExitConditionalExpression_In_Await(c *ConditionalExpression_In_AwaitContext)

	// ExitConditionalExpression_Yield_Await is called when exiting the conditionalExpression_Yield_Await production.
	ExitConditionalExpression_Yield_Await(c *ConditionalExpression_Yield_AwaitContext)

	// ExitConditionalExpression_In_Yield_Await is called when exiting the conditionalExpression_In_Yield_Await production.
	ExitConditionalExpression_In_Yield_Await(c *ConditionalExpression_In_Yield_AwaitContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentExpression_In is called when exiting the assignmentExpression_In production.
	ExitAssignmentExpression_In(c *AssignmentExpression_InContext)

	// ExitAssignmentExpression_Yield is called when exiting the assignmentExpression_Yield production.
	ExitAssignmentExpression_Yield(c *AssignmentExpression_YieldContext)

	// ExitAssignmentExpression_In_Yield is called when exiting the assignmentExpression_In_Yield production.
	ExitAssignmentExpression_In_Yield(c *AssignmentExpression_In_YieldContext)

	// ExitAssignmentExpression_Await is called when exiting the assignmentExpression_Await production.
	ExitAssignmentExpression_Await(c *AssignmentExpression_AwaitContext)

	// ExitAssignmentExpression_In_Await is called when exiting the assignmentExpression_In_Await production.
	ExitAssignmentExpression_In_Await(c *AssignmentExpression_In_AwaitContext)

	// ExitAssignmentExpression_Yield_Await is called when exiting the assignmentExpression_Yield_Await production.
	ExitAssignmentExpression_Yield_Await(c *AssignmentExpression_Yield_AwaitContext)

	// ExitAssignmentExpression_In_Yield_Await is called when exiting the assignmentExpression_In_Yield_Await production.
	ExitAssignmentExpression_In_Yield_Await(c *AssignmentExpression_In_Yield_AwaitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpression_In is called when exiting the expression_In production.
	ExitExpression_In(c *Expression_InContext)

	// ExitExpression_Yield is called when exiting the expression_Yield production.
	ExitExpression_Yield(c *Expression_YieldContext)

	// ExitExpression_In_Yield is called when exiting the expression_In_Yield production.
	ExitExpression_In_Yield(c *Expression_In_YieldContext)

	// ExitExpression_Await is called when exiting the expression_Await production.
	ExitExpression_Await(c *Expression_AwaitContext)

	// ExitExpression_In_Await is called when exiting the expression_In_Await production.
	ExitExpression_In_Await(c *Expression_In_AwaitContext)

	// ExitExpression_Yield_Await is called when exiting the expression_Yield_Await production.
	ExitExpression_Yield_Await(c *Expression_Yield_AwaitContext)

	// ExitExpression_In_Yield_Await is called when exiting the expression_In_Yield_Await production.
	ExitExpression_In_Yield_Await(c *Expression_In_Yield_AwaitContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatement_Yield is called when exiting the statement_Yield production.
	ExitStatement_Yield(c *Statement_YieldContext)

	// ExitStatement_Await is called when exiting the statement_Await production.
	ExitStatement_Await(c *Statement_AwaitContext)

	// ExitStatement_Yield_Await is called when exiting the statement_Yield_Await production.
	ExitStatement_Yield_Await(c *Statement_Yield_AwaitContext)

	// ExitStatement_Return is called when exiting the statement_Return production.
	ExitStatement_Return(c *Statement_ReturnContext)

	// ExitStatement_Yield_Return is called when exiting the statement_Yield_Return production.
	ExitStatement_Yield_Return(c *Statement_Yield_ReturnContext)

	// ExitStatement_Await_Return is called when exiting the statement_Await_Return production.
	ExitStatement_Await_Return(c *Statement_Await_ReturnContext)

	// ExitStatement_Yield_Await_Return is called when exiting the statement_Yield_Await_Return production.
	ExitStatement_Yield_Await_Return(c *Statement_Yield_Await_ReturnContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDeclaration_Yield is called when exiting the declaration_Yield production.
	ExitDeclaration_Yield(c *Declaration_YieldContext)

	// ExitDeclaration_Await is called when exiting the declaration_Await production.
	ExitDeclaration_Await(c *Declaration_AwaitContext)

	// ExitDeclaration_Yield_Await is called when exiting the declaration_Yield_Await production.
	ExitDeclaration_Yield_Await(c *Declaration_Yield_AwaitContext)

	// ExitHoistableDeclaration is called when exiting the hoistableDeclaration production.
	ExitHoistableDeclaration(c *HoistableDeclarationContext)

	// ExitHoistableDeclaration_Yield is called when exiting the hoistableDeclaration_Yield production.
	ExitHoistableDeclaration_Yield(c *HoistableDeclaration_YieldContext)

	// ExitHoistableDeclaration_Await is called when exiting the hoistableDeclaration_Await production.
	ExitHoistableDeclaration_Await(c *HoistableDeclaration_AwaitContext)

	// ExitHoistableDeclaration_Yield_Await is called when exiting the hoistableDeclaration_Yield_Await production.
	ExitHoistableDeclaration_Yield_Await(c *HoistableDeclaration_Yield_AwaitContext)

	// ExitHoistableDeclaration_Default is called when exiting the hoistableDeclaration_Default production.
	ExitHoistableDeclaration_Default(c *HoistableDeclaration_DefaultContext)

	// ExitHoistableDeclaration_Yield_Default is called when exiting the hoistableDeclaration_Yield_Default production.
	ExitHoistableDeclaration_Yield_Default(c *HoistableDeclaration_Yield_DefaultContext)

	// ExitHoistableDeclaration_Await_Default is called when exiting the hoistableDeclaration_Await_Default production.
	ExitHoistableDeclaration_Await_Default(c *HoistableDeclaration_Await_DefaultContext)

	// ExitHoistableDeclaration_Yield_Await_Default is called when exiting the hoistableDeclaration_Yield_Await_Default production.
	ExitHoistableDeclaration_Yield_Await_Default(c *HoistableDeclaration_Yield_Await_DefaultContext)

	// ExitBreakableStatement is called when exiting the breakableStatement production.
	ExitBreakableStatement(c *BreakableStatementContext)

	// ExitBreakableStatement_Yield is called when exiting the breakableStatement_Yield production.
	ExitBreakableStatement_Yield(c *BreakableStatement_YieldContext)

	// ExitBreakableStatement_Await is called when exiting the breakableStatement_Await production.
	ExitBreakableStatement_Await(c *BreakableStatement_AwaitContext)

	// ExitBreakableStatement_Yield_Await is called when exiting the breakableStatement_Yield_Await production.
	ExitBreakableStatement_Yield_Await(c *BreakableStatement_Yield_AwaitContext)

	// ExitBreakableStatement_Return is called when exiting the breakableStatement_Return production.
	ExitBreakableStatement_Return(c *BreakableStatement_ReturnContext)

	// ExitBreakableStatement_Yield_Return is called when exiting the breakableStatement_Yield_Return production.
	ExitBreakableStatement_Yield_Return(c *BreakableStatement_Yield_ReturnContext)

	// ExitBreakableStatement_Await_Return is called when exiting the breakableStatement_Await_Return production.
	ExitBreakableStatement_Await_Return(c *BreakableStatement_Await_ReturnContext)

	// ExitBreakableStatement_Yield_Await_Return is called when exiting the breakableStatement_Yield_Await_Return production.
	ExitBreakableStatement_Yield_Await_Return(c *BreakableStatement_Yield_Await_ReturnContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitBlockStatement_Yield is called when exiting the blockStatement_Yield production.
	ExitBlockStatement_Yield(c *BlockStatement_YieldContext)

	// ExitBlockStatement_Await is called when exiting the blockStatement_Await production.
	ExitBlockStatement_Await(c *BlockStatement_AwaitContext)

	// ExitBlockStatement_Yield_Await is called when exiting the blockStatement_Yield_Await production.
	ExitBlockStatement_Yield_Await(c *BlockStatement_Yield_AwaitContext)

	// ExitBlockStatement_Return is called when exiting the blockStatement_Return production.
	ExitBlockStatement_Return(c *BlockStatement_ReturnContext)

	// ExitBlockStatement_Yield_Return is called when exiting the blockStatement_Yield_Return production.
	ExitBlockStatement_Yield_Return(c *BlockStatement_Yield_ReturnContext)

	// ExitBlockStatement_Await_Return is called when exiting the blockStatement_Await_Return production.
	ExitBlockStatement_Await_Return(c *BlockStatement_Await_ReturnContext)

	// ExitBlockStatement_Yield_Await_Return is called when exiting the blockStatement_Yield_Await_Return production.
	ExitBlockStatement_Yield_Await_Return(c *BlockStatement_Yield_Await_ReturnContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlock_Yield is called when exiting the block_Yield production.
	ExitBlock_Yield(c *Block_YieldContext)

	// ExitBlock_Await is called when exiting the block_Await production.
	ExitBlock_Await(c *Block_AwaitContext)

	// ExitBlock_Yield_Await is called when exiting the block_Yield_Await production.
	ExitBlock_Yield_Await(c *Block_Yield_AwaitContext)

	// ExitBlock_Return is called when exiting the block_Return production.
	ExitBlock_Return(c *Block_ReturnContext)

	// ExitBlock_Yield_Return is called when exiting the block_Yield_Return production.
	ExitBlock_Yield_Return(c *Block_Yield_ReturnContext)

	// ExitBlock_Await_Return is called when exiting the block_Await_Return production.
	ExitBlock_Await_Return(c *Block_Await_ReturnContext)

	// ExitBlock_Yield_Await_Return is called when exiting the block_Yield_Await_Return production.
	ExitBlock_Yield_Await_Return(c *Block_Yield_Await_ReturnContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatementList_Yield is called when exiting the statementList_Yield production.
	ExitStatementList_Yield(c *StatementList_YieldContext)

	// ExitStatementList_Await is called when exiting the statementList_Await production.
	ExitStatementList_Await(c *StatementList_AwaitContext)

	// ExitStatementList_Yield_Await is called when exiting the statementList_Yield_Await production.
	ExitStatementList_Yield_Await(c *StatementList_Yield_AwaitContext)

	// ExitStatementList_Return is called when exiting the statementList_Return production.
	ExitStatementList_Return(c *StatementList_ReturnContext)

	// ExitStatementList_Yield_Return is called when exiting the statementList_Yield_Return production.
	ExitStatementList_Yield_Return(c *StatementList_Yield_ReturnContext)

	// ExitStatementList_Await_Return is called when exiting the statementList_Await_Return production.
	ExitStatementList_Await_Return(c *StatementList_Await_ReturnContext)

	// ExitStatementList_Yield_Await_Return is called when exiting the statementList_Yield_Await_Return production.
	ExitStatementList_Yield_Await_Return(c *StatementList_Yield_Await_ReturnContext)

	// ExitStatementListItem is called when exiting the statementListItem production.
	ExitStatementListItem(c *StatementListItemContext)

	// ExitStatementListItem_Yield is called when exiting the statementListItem_Yield production.
	ExitStatementListItem_Yield(c *StatementListItem_YieldContext)

	// ExitStatementListItem_Await is called when exiting the statementListItem_Await production.
	ExitStatementListItem_Await(c *StatementListItem_AwaitContext)

	// ExitStatementListItem_Yield_Await is called when exiting the statementListItem_Yield_Await production.
	ExitStatementListItem_Yield_Await(c *StatementListItem_Yield_AwaitContext)

	// ExitStatementListItem_Return is called when exiting the statementListItem_Return production.
	ExitStatementListItem_Return(c *StatementListItem_ReturnContext)

	// ExitStatementListItem_Yield_Return is called when exiting the statementListItem_Yield_Return production.
	ExitStatementListItem_Yield_Return(c *StatementListItem_Yield_ReturnContext)

	// ExitStatementListItem_Await_Return is called when exiting the statementListItem_Await_Return production.
	ExitStatementListItem_Await_Return(c *StatementListItem_Await_ReturnContext)

	// ExitStatementListItem_Yield_Await_Return is called when exiting the statementListItem_Yield_Await_Return production.
	ExitStatementListItem_Yield_Await_Return(c *StatementListItem_Yield_Await_ReturnContext)

	// ExitLexicalDeclaration is called when exiting the lexicalDeclaration production.
	ExitLexicalDeclaration(c *LexicalDeclarationContext)

	// ExitLexicalDeclaration_In is called when exiting the lexicalDeclaration_In production.
	ExitLexicalDeclaration_In(c *LexicalDeclaration_InContext)

	// ExitLexicalDeclaration_Yield is called when exiting the lexicalDeclaration_Yield production.
	ExitLexicalDeclaration_Yield(c *LexicalDeclaration_YieldContext)

	// ExitLexicalDeclaration_In_Yield is called when exiting the lexicalDeclaration_In_Yield production.
	ExitLexicalDeclaration_In_Yield(c *LexicalDeclaration_In_YieldContext)

	// ExitLexicalDeclaration_Await is called when exiting the lexicalDeclaration_Await production.
	ExitLexicalDeclaration_Await(c *LexicalDeclaration_AwaitContext)

	// ExitLexicalDeclaration_In_Await is called when exiting the lexicalDeclaration_In_Await production.
	ExitLexicalDeclaration_In_Await(c *LexicalDeclaration_In_AwaitContext)

	// ExitLexicalDeclaration_Yield_Await is called when exiting the lexicalDeclaration_Yield_Await production.
	ExitLexicalDeclaration_Yield_Await(c *LexicalDeclaration_Yield_AwaitContext)

	// ExitLexicalDeclaration_In_Yield_Await is called when exiting the lexicalDeclaration_In_Yield_Await production.
	ExitLexicalDeclaration_In_Yield_Await(c *LexicalDeclaration_In_Yield_AwaitContext)

	// ExitLetOrConst is called when exiting the letOrConst production.
	ExitLetOrConst(c *LetOrConstContext)

	// ExitBindingList is called when exiting the bindingList production.
	ExitBindingList(c *BindingListContext)

	// ExitBindingList_In is called when exiting the bindingList_In production.
	ExitBindingList_In(c *BindingList_InContext)

	// ExitBindingList_Yield is called when exiting the bindingList_Yield production.
	ExitBindingList_Yield(c *BindingList_YieldContext)

	// ExitBindingList_In_Yield is called when exiting the bindingList_In_Yield production.
	ExitBindingList_In_Yield(c *BindingList_In_YieldContext)

	// ExitBindingList_Await is called when exiting the bindingList_Await production.
	ExitBindingList_Await(c *BindingList_AwaitContext)

	// ExitBindingList_In_Await is called when exiting the bindingList_In_Await production.
	ExitBindingList_In_Await(c *BindingList_In_AwaitContext)

	// ExitBindingList_Yield_Await is called when exiting the bindingList_Yield_Await production.
	ExitBindingList_Yield_Await(c *BindingList_Yield_AwaitContext)

	// ExitBindingList_In_Yield_Await is called when exiting the bindingList_In_Yield_Await production.
	ExitBindingList_In_Yield_Await(c *BindingList_In_Yield_AwaitContext)

	// ExitLexicalBinding is called when exiting the lexicalBinding production.
	ExitLexicalBinding(c *LexicalBindingContext)

	// ExitLexicalBinding_In is called when exiting the lexicalBinding_In production.
	ExitLexicalBinding_In(c *LexicalBinding_InContext)

	// ExitLexicalBinding_Yield is called when exiting the lexicalBinding_Yield production.
	ExitLexicalBinding_Yield(c *LexicalBinding_YieldContext)

	// ExitLexicalBinding_In_Yield is called when exiting the lexicalBinding_In_Yield production.
	ExitLexicalBinding_In_Yield(c *LexicalBinding_In_YieldContext)

	// ExitLexicalBinding_Await is called when exiting the lexicalBinding_Await production.
	ExitLexicalBinding_Await(c *LexicalBinding_AwaitContext)

	// ExitLexicalBinding_In_Await is called when exiting the lexicalBinding_In_Await production.
	ExitLexicalBinding_In_Await(c *LexicalBinding_In_AwaitContext)

	// ExitLexicalBinding_Yield_Await is called when exiting the lexicalBinding_Yield_Await production.
	ExitLexicalBinding_Yield_Await(c *LexicalBinding_Yield_AwaitContext)

	// ExitLexicalBinding_In_Yield_Await is called when exiting the lexicalBinding_In_Yield_Await production.
	ExitLexicalBinding_In_Yield_Await(c *LexicalBinding_In_Yield_AwaitContext)

	// ExitVariableStatement is called when exiting the variableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitVariableStatement_Yield is called when exiting the variableStatement_Yield production.
	ExitVariableStatement_Yield(c *VariableStatement_YieldContext)

	// ExitVariableStatement_Await is called when exiting the variableStatement_Await production.
	ExitVariableStatement_Await(c *VariableStatement_AwaitContext)

	// ExitVariableStatement_Yield_Await is called when exiting the variableStatement_Yield_Await production.
	ExitVariableStatement_Yield_Await(c *VariableStatement_Yield_AwaitContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitVariableDeclarationList_In is called when exiting the variableDeclarationList_In production.
	ExitVariableDeclarationList_In(c *VariableDeclarationList_InContext)

	// ExitVariableDeclarationList_Yield is called when exiting the variableDeclarationList_Yield production.
	ExitVariableDeclarationList_Yield(c *VariableDeclarationList_YieldContext)

	// ExitVariableDeclarationList_In_Yield is called when exiting the variableDeclarationList_In_Yield production.
	ExitVariableDeclarationList_In_Yield(c *VariableDeclarationList_In_YieldContext)

	// ExitVariableDeclarationList_Await is called when exiting the variableDeclarationList_Await production.
	ExitVariableDeclarationList_Await(c *VariableDeclarationList_AwaitContext)

	// ExitVariableDeclarationList_In_Await is called when exiting the variableDeclarationList_In_Await production.
	ExitVariableDeclarationList_In_Await(c *VariableDeclarationList_In_AwaitContext)

	// ExitVariableDeclarationList_Yield_Await is called when exiting the variableDeclarationList_Yield_Await production.
	ExitVariableDeclarationList_Yield_Await(c *VariableDeclarationList_Yield_AwaitContext)

	// ExitVariableDeclarationList_In_Yield_Await is called when exiting the variableDeclarationList_In_Yield_Await production.
	ExitVariableDeclarationList_In_Yield_Await(c *VariableDeclarationList_In_Yield_AwaitContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitVariableDeclaration_In is called when exiting the variableDeclaration_In production.
	ExitVariableDeclaration_In(c *VariableDeclaration_InContext)

	// ExitVariableDeclaration_Yield is called when exiting the variableDeclaration_Yield production.
	ExitVariableDeclaration_Yield(c *VariableDeclaration_YieldContext)

	// ExitVariableDeclaration_In_Yield is called when exiting the variableDeclaration_In_Yield production.
	ExitVariableDeclaration_In_Yield(c *VariableDeclaration_In_YieldContext)

	// ExitVariableDeclaration_Await is called when exiting the variableDeclaration_Await production.
	ExitVariableDeclaration_Await(c *VariableDeclaration_AwaitContext)

	// ExitVariableDeclaration_In_Await is called when exiting the variableDeclaration_In_Await production.
	ExitVariableDeclaration_In_Await(c *VariableDeclaration_In_AwaitContext)

	// ExitVariableDeclaration_Yield_Await is called when exiting the variableDeclaration_Yield_Await production.
	ExitVariableDeclaration_Yield_Await(c *VariableDeclaration_Yield_AwaitContext)

	// ExitVariableDeclaration_In_Yield_Await is called when exiting the variableDeclaration_In_Yield_Await production.
	ExitVariableDeclaration_In_Yield_Await(c *VariableDeclaration_In_Yield_AwaitContext)

	// ExitBindingPattern is called when exiting the bindingPattern production.
	ExitBindingPattern(c *BindingPatternContext)

	// ExitBindingPattern_Yield is called when exiting the bindingPattern_Yield production.
	ExitBindingPattern_Yield(c *BindingPattern_YieldContext)

	// ExitBindingPattern_Await is called when exiting the bindingPattern_Await production.
	ExitBindingPattern_Await(c *BindingPattern_AwaitContext)

	// ExitBindingPattern_Yield_Await is called when exiting the bindingPattern_Yield_Await production.
	ExitBindingPattern_Yield_Await(c *BindingPattern_Yield_AwaitContext)

	// ExitObjectBindingPattern is called when exiting the objectBindingPattern production.
	ExitObjectBindingPattern(c *ObjectBindingPatternContext)

	// ExitObjectBindingPattern_Yield is called when exiting the objectBindingPattern_Yield production.
	ExitObjectBindingPattern_Yield(c *ObjectBindingPattern_YieldContext)

	// ExitObjectBindingPattern_Await is called when exiting the objectBindingPattern_Await production.
	ExitObjectBindingPattern_Await(c *ObjectBindingPattern_AwaitContext)

	// ExitObjectBindingPattern_Yield_Await is called when exiting the objectBindingPattern_Yield_Await production.
	ExitObjectBindingPattern_Yield_Await(c *ObjectBindingPattern_Yield_AwaitContext)

	// ExitArrayBindingPattern is called when exiting the arrayBindingPattern production.
	ExitArrayBindingPattern(c *ArrayBindingPatternContext)

	// ExitArrayBindingPattern_Yield is called when exiting the arrayBindingPattern_Yield production.
	ExitArrayBindingPattern_Yield(c *ArrayBindingPattern_YieldContext)

	// ExitArrayBindingPattern_Await is called when exiting the arrayBindingPattern_Await production.
	ExitArrayBindingPattern_Await(c *ArrayBindingPattern_AwaitContext)

	// ExitArrayBindingPattern_Yield_Await is called when exiting the arrayBindingPattern_Yield_Await production.
	ExitArrayBindingPattern_Yield_Await(c *ArrayBindingPattern_Yield_AwaitContext)

	// ExitBindingRestProperty is called when exiting the bindingRestProperty production.
	ExitBindingRestProperty(c *BindingRestPropertyContext)

	// ExitBindingRestProperty_Yield is called when exiting the bindingRestProperty_Yield production.
	ExitBindingRestProperty_Yield(c *BindingRestProperty_YieldContext)

	// ExitBindingRestProperty_Await is called when exiting the bindingRestProperty_Await production.
	ExitBindingRestProperty_Await(c *BindingRestProperty_AwaitContext)

	// ExitBindingRestProperty_Yield_Await is called when exiting the bindingRestProperty_Yield_Await production.
	ExitBindingRestProperty_Yield_Await(c *BindingRestProperty_Yield_AwaitContext)

	// ExitBindingPropertyList is called when exiting the bindingPropertyList production.
	ExitBindingPropertyList(c *BindingPropertyListContext)

	// ExitBindingPropertyList_Yield is called when exiting the bindingPropertyList_Yield production.
	ExitBindingPropertyList_Yield(c *BindingPropertyList_YieldContext)

	// ExitBindingPropertyList_Await is called when exiting the bindingPropertyList_Await production.
	ExitBindingPropertyList_Await(c *BindingPropertyList_AwaitContext)

	// ExitBindingPropertyList_Yield_Await is called when exiting the bindingPropertyList_Yield_Await production.
	ExitBindingPropertyList_Yield_Await(c *BindingPropertyList_Yield_AwaitContext)

	// ExitBindingElementList is called when exiting the bindingElementList production.
	ExitBindingElementList(c *BindingElementListContext)

	// ExitBindingElementList_Yield is called when exiting the bindingElementList_Yield production.
	ExitBindingElementList_Yield(c *BindingElementList_YieldContext)

	// ExitBindingElementList_Await is called when exiting the bindingElementList_Await production.
	ExitBindingElementList_Await(c *BindingElementList_AwaitContext)

	// ExitBindingElementList_Yield_Await is called when exiting the bindingElementList_Yield_Await production.
	ExitBindingElementList_Yield_Await(c *BindingElementList_Yield_AwaitContext)

	// ExitBindingElisionElement is called when exiting the bindingElisionElement production.
	ExitBindingElisionElement(c *BindingElisionElementContext)

	// ExitBindingElisionElement_Yield is called when exiting the bindingElisionElement_Yield production.
	ExitBindingElisionElement_Yield(c *BindingElisionElement_YieldContext)

	// ExitBindingElisionElement_Await is called when exiting the bindingElisionElement_Await production.
	ExitBindingElisionElement_Await(c *BindingElisionElement_AwaitContext)

	// ExitBindingElisionElement_Yield_Await is called when exiting the bindingElisionElement_Yield_Await production.
	ExitBindingElisionElement_Yield_Await(c *BindingElisionElement_Yield_AwaitContext)

	// ExitBindingProperty is called when exiting the bindingProperty production.
	ExitBindingProperty(c *BindingPropertyContext)

	// ExitBindingProperty_Yield is called when exiting the bindingProperty_Yield production.
	ExitBindingProperty_Yield(c *BindingProperty_YieldContext)

	// ExitBindingProperty_Await is called when exiting the bindingProperty_Await production.
	ExitBindingProperty_Await(c *BindingProperty_AwaitContext)

	// ExitBindingProperty_Yield_Await is called when exiting the bindingProperty_Yield_Await production.
	ExitBindingProperty_Yield_Await(c *BindingProperty_Yield_AwaitContext)

	// ExitBindingElement is called when exiting the bindingElement production.
	ExitBindingElement(c *BindingElementContext)

	// ExitBindingElement_Yield is called when exiting the bindingElement_Yield production.
	ExitBindingElement_Yield(c *BindingElement_YieldContext)

	// ExitBindingElement_Await is called when exiting the bindingElement_Await production.
	ExitBindingElement_Await(c *BindingElement_AwaitContext)

	// ExitBindingElement_Yield_Await is called when exiting the bindingElement_Yield_Await production.
	ExitBindingElement_Yield_Await(c *BindingElement_Yield_AwaitContext)

	// ExitSingleNameBinding is called when exiting the singleNameBinding production.
	ExitSingleNameBinding(c *SingleNameBindingContext)

	// ExitSingleNameBinding_Yield is called when exiting the singleNameBinding_Yield production.
	ExitSingleNameBinding_Yield(c *SingleNameBinding_YieldContext)

	// ExitSingleNameBinding_Await is called when exiting the singleNameBinding_Await production.
	ExitSingleNameBinding_Await(c *SingleNameBinding_AwaitContext)

	// ExitSingleNameBinding_Yield_Await is called when exiting the singleNameBinding_Yield_Await production.
	ExitSingleNameBinding_Yield_Await(c *SingleNameBinding_Yield_AwaitContext)

	// ExitBindingRestElement is called when exiting the bindingRestElement production.
	ExitBindingRestElement(c *BindingRestElementContext)

	// ExitBindingRestElement_Yield is called when exiting the bindingRestElement_Yield production.
	ExitBindingRestElement_Yield(c *BindingRestElement_YieldContext)

	// ExitBindingRestElement_Await is called when exiting the bindingRestElement_Await production.
	ExitBindingRestElement_Await(c *BindingRestElement_AwaitContext)

	// ExitBindingRestElement_Yield_Await is called when exiting the bindingRestElement_Yield_Await production.
	ExitBindingRestElement_Yield_Await(c *BindingRestElement_Yield_AwaitContext)

	// ExitEmptyStatement is called when exiting the emptyStatement production.
	ExitEmptyStatement(c *EmptyStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitExpressionStatement_Yield is called when exiting the expressionStatement_Yield production.
	ExitExpressionStatement_Yield(c *ExpressionStatement_YieldContext)

	// ExitExpressionStatement_Await is called when exiting the expressionStatement_Await production.
	ExitExpressionStatement_Await(c *ExpressionStatement_AwaitContext)

	// ExitExpressionStatement_Yield_Await is called when exiting the expressionStatement_Yield_Await production.
	ExitExpressionStatement_Yield_Await(c *ExpressionStatement_Yield_AwaitContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfStatement_Yield is called when exiting the ifStatement_Yield production.
	ExitIfStatement_Yield(c *IfStatement_YieldContext)

	// ExitIfStatement_Await is called when exiting the ifStatement_Await production.
	ExitIfStatement_Await(c *IfStatement_AwaitContext)

	// ExitIfStatement_Yield_Await is called when exiting the ifStatement_Yield_Await production.
	ExitIfStatement_Yield_Await(c *IfStatement_Yield_AwaitContext)

	// ExitIfStatement_Return is called when exiting the ifStatement_Return production.
	ExitIfStatement_Return(c *IfStatement_ReturnContext)

	// ExitIfStatement_Yield_Return is called when exiting the ifStatement_Yield_Return production.
	ExitIfStatement_Yield_Return(c *IfStatement_Yield_ReturnContext)

	// ExitIfStatement_Await_Return is called when exiting the ifStatement_Await_Return production.
	ExitIfStatement_Await_Return(c *IfStatement_Await_ReturnContext)

	// ExitIfStatement_Yield_Await_Return is called when exiting the ifStatement_Yield_Await_Return production.
	ExitIfStatement_Yield_Await_Return(c *IfStatement_Yield_Await_ReturnContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitIterationStatement_Yield is called when exiting the iterationStatement_Yield production.
	ExitIterationStatement_Yield(c *IterationStatement_YieldContext)

	// ExitIterationStatement_Await is called when exiting the iterationStatement_Await production.
	ExitIterationStatement_Await(c *IterationStatement_AwaitContext)

	// ExitIterationStatement_Yield_Await is called when exiting the iterationStatement_Yield_Await production.
	ExitIterationStatement_Yield_Await(c *IterationStatement_Yield_AwaitContext)

	// ExitIterationStatement_Return is called when exiting the iterationStatement_Return production.
	ExitIterationStatement_Return(c *IterationStatement_ReturnContext)

	// ExitIterationStatement_Yield_Return is called when exiting the iterationStatement_Yield_Return production.
	ExitIterationStatement_Yield_Return(c *IterationStatement_Yield_ReturnContext)

	// ExitIterationStatement_Await_Return is called when exiting the iterationStatement_Await_Return production.
	ExitIterationStatement_Await_Return(c *IterationStatement_Await_ReturnContext)

	// ExitIterationStatement_Yield_Await_Return is called when exiting the iterationStatement_Yield_Await_Return production.
	ExitIterationStatement_Yield_Await_Return(c *IterationStatement_Yield_Await_ReturnContext)

	// ExitForDeclaration is called when exiting the forDeclaration production.
	ExitForDeclaration(c *ForDeclarationContext)

	// ExitForDeclaration_Yield is called when exiting the forDeclaration_Yield production.
	ExitForDeclaration_Yield(c *ForDeclaration_YieldContext)

	// ExitForDeclaration_Await is called when exiting the forDeclaration_Await production.
	ExitForDeclaration_Await(c *ForDeclaration_AwaitContext)

	// ExitForDeclaration_Yield_Await is called when exiting the forDeclaration_Yield_Await production.
	ExitForDeclaration_Yield_Await(c *ForDeclaration_Yield_AwaitContext)

	// ExitForBinding is called when exiting the forBinding production.
	ExitForBinding(c *ForBindingContext)

	// ExitForBinding_Yield is called when exiting the forBinding_Yield production.
	ExitForBinding_Yield(c *ForBinding_YieldContext)

	// ExitForBinding_Await is called when exiting the forBinding_Await production.
	ExitForBinding_Await(c *ForBinding_AwaitContext)

	// ExitForBinding_Yield_Await is called when exiting the forBinding_Yield_Await production.
	ExitForBinding_Yield_Await(c *ForBinding_Yield_AwaitContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitContinueStatement_Yield is called when exiting the continueStatement_Yield production.
	ExitContinueStatement_Yield(c *ContinueStatement_YieldContext)

	// ExitContinueStatement_Await is called when exiting the continueStatement_Await production.
	ExitContinueStatement_Await(c *ContinueStatement_AwaitContext)

	// ExitContinueStatement_Yield_Await is called when exiting the continueStatement_Yield_Await production.
	ExitContinueStatement_Yield_Await(c *ContinueStatement_Yield_AwaitContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitBreakStatement_Yield is called when exiting the breakStatement_Yield production.
	ExitBreakStatement_Yield(c *BreakStatement_YieldContext)

	// ExitBreakStatement_Await is called when exiting the breakStatement_Await production.
	ExitBreakStatement_Await(c *BreakStatement_AwaitContext)

	// ExitBreakStatement_Yield_Await is called when exiting the breakStatement_Yield_Await production.
	ExitBreakStatement_Yield_Await(c *BreakStatement_Yield_AwaitContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitReturnStatement_Yield is called when exiting the returnStatement_Yield production.
	ExitReturnStatement_Yield(c *ReturnStatement_YieldContext)

	// ExitReturnStatement_Await is called when exiting the returnStatement_Await production.
	ExitReturnStatement_Await(c *ReturnStatement_AwaitContext)

	// ExitReturnStatement_Yield_Await is called when exiting the returnStatement_Yield_Await production.
	ExitReturnStatement_Yield_Await(c *ReturnStatement_Yield_AwaitContext)

	// ExitWithStatement is called when exiting the withStatement production.
	ExitWithStatement(c *WithStatementContext)

	// ExitWithStatement_Yield is called when exiting the withStatement_Yield production.
	ExitWithStatement_Yield(c *WithStatement_YieldContext)

	// ExitWithStatement_Await is called when exiting the withStatement_Await production.
	ExitWithStatement_Await(c *WithStatement_AwaitContext)

	// ExitWithStatement_Yield_Await is called when exiting the withStatement_Yield_Await production.
	ExitWithStatement_Yield_Await(c *WithStatement_Yield_AwaitContext)

	// ExitWithStatement_Return is called when exiting the withStatement_Return production.
	ExitWithStatement_Return(c *WithStatement_ReturnContext)

	// ExitWithStatement_Yield_Return is called when exiting the withStatement_Yield_Return production.
	ExitWithStatement_Yield_Return(c *WithStatement_Yield_ReturnContext)

	// ExitWithStatement_Await_Return is called when exiting the withStatement_Await_Return production.
	ExitWithStatement_Await_Return(c *WithStatement_Await_ReturnContext)

	// ExitWithStatement_Yield_Await_Return is called when exiting the withStatement_Yield_Await_Return production.
	ExitWithStatement_Yield_Await_Return(c *WithStatement_Yield_Await_ReturnContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchStatement_Yield is called when exiting the switchStatement_Yield production.
	ExitSwitchStatement_Yield(c *SwitchStatement_YieldContext)

	// ExitSwitchStatement_Await is called when exiting the switchStatement_Await production.
	ExitSwitchStatement_Await(c *SwitchStatement_AwaitContext)

	// ExitSwitchStatement_Yield_Await is called when exiting the switchStatement_Yield_Await production.
	ExitSwitchStatement_Yield_Await(c *SwitchStatement_Yield_AwaitContext)

	// ExitSwitchStatement_Return is called when exiting the switchStatement_Return production.
	ExitSwitchStatement_Return(c *SwitchStatement_ReturnContext)

	// ExitSwitchStatement_Yield_Return is called when exiting the switchStatement_Yield_Return production.
	ExitSwitchStatement_Yield_Return(c *SwitchStatement_Yield_ReturnContext)

	// ExitSwitchStatement_Await_Return is called when exiting the switchStatement_Await_Return production.
	ExitSwitchStatement_Await_Return(c *SwitchStatement_Await_ReturnContext)

	// ExitSwitchStatement_Yield_Await_Return is called when exiting the switchStatement_Yield_Await_Return production.
	ExitSwitchStatement_Yield_Await_Return(c *SwitchStatement_Yield_Await_ReturnContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitCaseBlock_Yield is called when exiting the caseBlock_Yield production.
	ExitCaseBlock_Yield(c *CaseBlock_YieldContext)

	// ExitCaseBlock_Await is called when exiting the caseBlock_Await production.
	ExitCaseBlock_Await(c *CaseBlock_AwaitContext)

	// ExitCaseBlock_Yield_Await is called when exiting the caseBlock_Yield_Await production.
	ExitCaseBlock_Yield_Await(c *CaseBlock_Yield_AwaitContext)

	// ExitCaseBlock_Return is called when exiting the caseBlock_Return production.
	ExitCaseBlock_Return(c *CaseBlock_ReturnContext)

	// ExitCaseBlock_Yield_Return is called when exiting the caseBlock_Yield_Return production.
	ExitCaseBlock_Yield_Return(c *CaseBlock_Yield_ReturnContext)

	// ExitCaseBlock_Await_Return is called when exiting the caseBlock_Await_Return production.
	ExitCaseBlock_Await_Return(c *CaseBlock_Await_ReturnContext)

	// ExitCaseBlock_Yield_Await_Return is called when exiting the caseBlock_Yield_Await_Return production.
	ExitCaseBlock_Yield_Await_Return(c *CaseBlock_Yield_Await_ReturnContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitCaseClause_Yield is called when exiting the caseClause_Yield production.
	ExitCaseClause_Yield(c *CaseClause_YieldContext)

	// ExitCaseClause_Await is called when exiting the caseClause_Await production.
	ExitCaseClause_Await(c *CaseClause_AwaitContext)

	// ExitCaseClause_Yield_Await is called when exiting the caseClause_Yield_Await production.
	ExitCaseClause_Yield_Await(c *CaseClause_Yield_AwaitContext)

	// ExitCaseClause_Return is called when exiting the caseClause_Return production.
	ExitCaseClause_Return(c *CaseClause_ReturnContext)

	// ExitCaseClause_Yield_Return is called when exiting the caseClause_Yield_Return production.
	ExitCaseClause_Yield_Return(c *CaseClause_Yield_ReturnContext)

	// ExitCaseClause_Await_Return is called when exiting the caseClause_Await_Return production.
	ExitCaseClause_Await_Return(c *CaseClause_Await_ReturnContext)

	// ExitCaseClause_Yield_Await_Return is called when exiting the caseClause_Yield_Await_Return production.
	ExitCaseClause_Yield_Await_Return(c *CaseClause_Yield_Await_ReturnContext)

	// ExitDefaultClause is called when exiting the defaultClause production.
	ExitDefaultClause(c *DefaultClauseContext)

	// ExitDefaultClause_Yield is called when exiting the defaultClause_Yield production.
	ExitDefaultClause_Yield(c *DefaultClause_YieldContext)

	// ExitDefaultClause_Await is called when exiting the defaultClause_Await production.
	ExitDefaultClause_Await(c *DefaultClause_AwaitContext)

	// ExitDefaultClause_Yield_Await is called when exiting the defaultClause_Yield_Await production.
	ExitDefaultClause_Yield_Await(c *DefaultClause_Yield_AwaitContext)

	// ExitDefaultClause_Return is called when exiting the defaultClause_Return production.
	ExitDefaultClause_Return(c *DefaultClause_ReturnContext)

	// ExitDefaultClause_Yield_Return is called when exiting the defaultClause_Yield_Return production.
	ExitDefaultClause_Yield_Return(c *DefaultClause_Yield_ReturnContext)

	// ExitDefaultClause_Await_Return is called when exiting the defaultClause_Await_Return production.
	ExitDefaultClause_Await_Return(c *DefaultClause_Await_ReturnContext)

	// ExitDefaultClause_Yield_Await_Return is called when exiting the defaultClause_Yield_Await_Return production.
	ExitDefaultClause_Yield_Await_Return(c *DefaultClause_Yield_Await_ReturnContext)

	// ExitLabelledStatement is called when exiting the labelledStatement production.
	ExitLabelledStatement(c *LabelledStatementContext)

	// ExitLabelledStatement_Yield is called when exiting the labelledStatement_Yield production.
	ExitLabelledStatement_Yield(c *LabelledStatement_YieldContext)

	// ExitLabelledStatement_Await is called when exiting the labelledStatement_Await production.
	ExitLabelledStatement_Await(c *LabelledStatement_AwaitContext)

	// ExitLabelledStatement_Yield_Await is called when exiting the labelledStatement_Yield_Await production.
	ExitLabelledStatement_Yield_Await(c *LabelledStatement_Yield_AwaitContext)

	// ExitLabelledStatement_Return is called when exiting the labelledStatement_Return production.
	ExitLabelledStatement_Return(c *LabelledStatement_ReturnContext)

	// ExitLabelledStatement_Yield_Return is called when exiting the labelledStatement_Yield_Return production.
	ExitLabelledStatement_Yield_Return(c *LabelledStatement_Yield_ReturnContext)

	// ExitLabelledStatement_Await_Return is called when exiting the labelledStatement_Await_Return production.
	ExitLabelledStatement_Await_Return(c *LabelledStatement_Await_ReturnContext)

	// ExitLabelledStatement_Yield_Await_Return is called when exiting the labelledStatement_Yield_Await_Return production.
	ExitLabelledStatement_Yield_Await_Return(c *LabelledStatement_Yield_Await_ReturnContext)

	// ExitLabelledItem is called when exiting the labelledItem production.
	ExitLabelledItem(c *LabelledItemContext)

	// ExitLabelledItem_Yield is called when exiting the labelledItem_Yield production.
	ExitLabelledItem_Yield(c *LabelledItem_YieldContext)

	// ExitLabelledItem_Await is called when exiting the labelledItem_Await production.
	ExitLabelledItem_Await(c *LabelledItem_AwaitContext)

	// ExitLabelledItem_Yield_Await is called when exiting the labelledItem_Yield_Await production.
	ExitLabelledItem_Yield_Await(c *LabelledItem_Yield_AwaitContext)

	// ExitLabelledItem_Return is called when exiting the labelledItem_Return production.
	ExitLabelledItem_Return(c *LabelledItem_ReturnContext)

	// ExitLabelledItem_Yield_Return is called when exiting the labelledItem_Yield_Return production.
	ExitLabelledItem_Yield_Return(c *LabelledItem_Yield_ReturnContext)

	// ExitLabelledItem_Await_Return is called when exiting the labelledItem_Await_Return production.
	ExitLabelledItem_Await_Return(c *LabelledItem_Await_ReturnContext)

	// ExitLabelledItem_Yield_Await_Return is called when exiting the labelledItem_Yield_Await_Return production.
	ExitLabelledItem_Yield_Await_Return(c *LabelledItem_Yield_Await_ReturnContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitThrowStatement_Yield is called when exiting the throwStatement_Yield production.
	ExitThrowStatement_Yield(c *ThrowStatement_YieldContext)

	// ExitThrowStatement_Await is called when exiting the throwStatement_Await production.
	ExitThrowStatement_Await(c *ThrowStatement_AwaitContext)

	// ExitThrowStatement_Yield_Await is called when exiting the throwStatement_Yield_Await production.
	ExitThrowStatement_Yield_Await(c *ThrowStatement_Yield_AwaitContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitTryStatement_Yield is called when exiting the tryStatement_Yield production.
	ExitTryStatement_Yield(c *TryStatement_YieldContext)

	// ExitTryStatement_Await is called when exiting the tryStatement_Await production.
	ExitTryStatement_Await(c *TryStatement_AwaitContext)

	// ExitTryStatement_Yield_Await is called when exiting the tryStatement_Yield_Await production.
	ExitTryStatement_Yield_Await(c *TryStatement_Yield_AwaitContext)

	// ExitTryStatement_Return is called when exiting the tryStatement_Return production.
	ExitTryStatement_Return(c *TryStatement_ReturnContext)

	// ExitTryStatement_Yield_Return is called when exiting the tryStatement_Yield_Return production.
	ExitTryStatement_Yield_Return(c *TryStatement_Yield_ReturnContext)

	// ExitTryStatement_Await_Return is called when exiting the tryStatement_Await_Return production.
	ExitTryStatement_Await_Return(c *TryStatement_Await_ReturnContext)

	// ExitTryStatement_Yield_Await_Return is called when exiting the tryStatement_Yield_Await_Return production.
	ExitTryStatement_Yield_Await_Return(c *TryStatement_Yield_Await_ReturnContext)

	// ExitCatch_ is called when exiting the catch_ production.
	ExitCatch_(c *Catch_Context)

	// ExitCatch_Yield is called when exiting the catch_Yield production.
	ExitCatch_Yield(c *Catch_YieldContext)

	// ExitCatch_Await is called when exiting the catch_Await production.
	ExitCatch_Await(c *Catch_AwaitContext)

	// ExitCatch_Yield_Await is called when exiting the catch_Yield_Await production.
	ExitCatch_Yield_Await(c *Catch_Yield_AwaitContext)

	// ExitCatch_Return is called when exiting the catch_Return production.
	ExitCatch_Return(c *Catch_ReturnContext)

	// ExitCatch_Yield_Return is called when exiting the catch_Yield_Return production.
	ExitCatch_Yield_Return(c *Catch_Yield_ReturnContext)

	// ExitCatch_Await_Return is called when exiting the catch_Await_Return production.
	ExitCatch_Await_Return(c *Catch_Await_ReturnContext)

	// ExitCatch_Yield_Await_Return is called when exiting the catch_Yield_Await_Return production.
	ExitCatch_Yield_Await_Return(c *Catch_Yield_Await_ReturnContext)

	// ExitFinally_ is called when exiting the finally_ production.
	ExitFinally_(c *Finally_Context)

	// ExitFinally_Yield is called when exiting the finally_Yield production.
	ExitFinally_Yield(c *Finally_YieldContext)

	// ExitFinally_Await is called when exiting the finally_Await production.
	ExitFinally_Await(c *Finally_AwaitContext)

	// ExitFinally_Yield_Await is called when exiting the finally_Yield_Await production.
	ExitFinally_Yield_Await(c *Finally_Yield_AwaitContext)

	// ExitFinally_Return is called when exiting the finally_Return production.
	ExitFinally_Return(c *Finally_ReturnContext)

	// ExitFinally_Yield_Return is called when exiting the finally_Yield_Return production.
	ExitFinally_Yield_Return(c *Finally_Yield_ReturnContext)

	// ExitFinally_Await_Return is called when exiting the finally_Await_Return production.
	ExitFinally_Await_Return(c *Finally_Await_ReturnContext)

	// ExitFinally_Yield_Await_Return is called when exiting the finally_Yield_Await_Return production.
	ExitFinally_Yield_Await_Return(c *Finally_Yield_Await_ReturnContext)

	// ExitCatchParameter is called when exiting the catchParameter production.
	ExitCatchParameter(c *CatchParameterContext)

	// ExitCatchParameter_Yield is called when exiting the catchParameter_Yield production.
	ExitCatchParameter_Yield(c *CatchParameter_YieldContext)

	// ExitCatchParameter_Await is called when exiting the catchParameter_Await production.
	ExitCatchParameter_Await(c *CatchParameter_AwaitContext)

	// ExitCatchParameter_Yield_Await is called when exiting the catchParameter_Yield_Await production.
	ExitCatchParameter_Yield_Await(c *CatchParameter_Yield_AwaitContext)

	// ExitDebuggerStatement is called when exiting the debuggerStatement production.
	ExitDebuggerStatement(c *DebuggerStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionDeclaration_Yield is called when exiting the functionDeclaration_Yield production.
	ExitFunctionDeclaration_Yield(c *FunctionDeclaration_YieldContext)

	// ExitFunctionDeclaration_Await is called when exiting the functionDeclaration_Await production.
	ExitFunctionDeclaration_Await(c *FunctionDeclaration_AwaitContext)

	// ExitFunctionDeclaration_Yield_Await is called when exiting the functionDeclaration_Yield_Await production.
	ExitFunctionDeclaration_Yield_Await(c *FunctionDeclaration_Yield_AwaitContext)

	// ExitFunctionDeclaration_Default is called when exiting the functionDeclaration_Default production.
	ExitFunctionDeclaration_Default(c *FunctionDeclaration_DefaultContext)

	// ExitFunctionDeclaration_Yield_Default is called when exiting the functionDeclaration_Yield_Default production.
	ExitFunctionDeclaration_Yield_Default(c *FunctionDeclaration_Yield_DefaultContext)

	// ExitFunctionDeclaration_Await_Default is called when exiting the functionDeclaration_Await_Default production.
	ExitFunctionDeclaration_Await_Default(c *FunctionDeclaration_Await_DefaultContext)

	// ExitFunctionDeclaration_Yield_Await_Default is called when exiting the functionDeclaration_Yield_Await_Default production.
	ExitFunctionDeclaration_Yield_Await_Default(c *FunctionDeclaration_Yield_Await_DefaultContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitUniqueFormalParameters is called when exiting the uniqueFormalParameters production.
	ExitUniqueFormalParameters(c *UniqueFormalParametersContext)

	// ExitUniqueFormalParameters_Yield is called when exiting the uniqueFormalParameters_Yield production.
	ExitUniqueFormalParameters_Yield(c *UniqueFormalParameters_YieldContext)

	// ExitUniqueFormalParameters_Await is called when exiting the uniqueFormalParameters_Await production.
	ExitUniqueFormalParameters_Await(c *UniqueFormalParameters_AwaitContext)

	// ExitUniqueFormalParameters_Yield_Await is called when exiting the uniqueFormalParameters_Yield_Await production.
	ExitUniqueFormalParameters_Yield_Await(c *UniqueFormalParameters_Yield_AwaitContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameters_Yield is called when exiting the formalParameters_Yield production.
	ExitFormalParameters_Yield(c *FormalParameters_YieldContext)

	// ExitFormalParameters_Await is called when exiting the formalParameters_Await production.
	ExitFormalParameters_Await(c *FormalParameters_AwaitContext)

	// ExitFormalParameters_Yield_Await is called when exiting the formalParameters_Yield_Await production.
	ExitFormalParameters_Yield_Await(c *FormalParameters_Yield_AwaitContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameterList_Yield is called when exiting the formalParameterList_Yield production.
	ExitFormalParameterList_Yield(c *FormalParameterList_YieldContext)

	// ExitFormalParameterList_Await is called when exiting the formalParameterList_Await production.
	ExitFormalParameterList_Await(c *FormalParameterList_AwaitContext)

	// ExitFormalParameterList_Yield_Await is called when exiting the formalParameterList_Yield_Await production.
	ExitFormalParameterList_Yield_Await(c *FormalParameterList_Yield_AwaitContext)

	// ExitFunctionRestParameter is called when exiting the functionRestParameter production.
	ExitFunctionRestParameter(c *FunctionRestParameterContext)

	// ExitFunctionRestParameter_Yield is called when exiting the functionRestParameter_Yield production.
	ExitFunctionRestParameter_Yield(c *FunctionRestParameter_YieldContext)

	// ExitFunctionRestParameter_Await is called when exiting the functionRestParameter_Await production.
	ExitFunctionRestParameter_Await(c *FunctionRestParameter_AwaitContext)

	// ExitFunctionRestParameter_Yield_Await is called when exiting the functionRestParameter_Yield_Await production.
	ExitFunctionRestParameter_Yield_Await(c *FunctionRestParameter_Yield_AwaitContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitFormalParameter_Yield is called when exiting the formalParameter_Yield production.
	ExitFormalParameter_Yield(c *FormalParameter_YieldContext)

	// ExitFormalParameter_Await is called when exiting the formalParameter_Await production.
	ExitFormalParameter_Await(c *FormalParameter_AwaitContext)

	// ExitFormalParameter_Yield_Await is called when exiting the formalParameter_Yield_Await production.
	ExitFormalParameter_Yield_Await(c *FormalParameter_Yield_AwaitContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitFunctionBody_Yield is called when exiting the functionBody_Yield production.
	ExitFunctionBody_Yield(c *FunctionBody_YieldContext)

	// ExitFunctionBody_Await is called when exiting the functionBody_Await production.
	ExitFunctionBody_Await(c *FunctionBody_AwaitContext)

	// ExitFunctionBody_Yield_Await is called when exiting the functionBody_Yield_Await production.
	ExitFunctionBody_Yield_Await(c *FunctionBody_Yield_AwaitContext)

	// ExitFunctionStatementList is called when exiting the functionStatementList production.
	ExitFunctionStatementList(c *FunctionStatementListContext)

	// ExitFunctionStatementList_Yield is called when exiting the functionStatementList_Yield production.
	ExitFunctionStatementList_Yield(c *FunctionStatementList_YieldContext)

	// ExitFunctionStatementList_Await is called when exiting the functionStatementList_Await production.
	ExitFunctionStatementList_Await(c *FunctionStatementList_AwaitContext)

	// ExitFunctionStatementList_Yield_Await is called when exiting the functionStatementList_Yield_Await production.
	ExitFunctionStatementList_Yield_Await(c *FunctionStatementList_Yield_AwaitContext)

	// ExitArrowFunction is called when exiting the arrowFunction production.
	ExitArrowFunction(c *ArrowFunctionContext)

	// ExitArrowFunction_In is called when exiting the arrowFunction_In production.
	ExitArrowFunction_In(c *ArrowFunction_InContext)

	// ExitArrowFunction_Yield is called when exiting the arrowFunction_Yield production.
	ExitArrowFunction_Yield(c *ArrowFunction_YieldContext)

	// ExitArrowFunction_In_Yield is called when exiting the arrowFunction_In_Yield production.
	ExitArrowFunction_In_Yield(c *ArrowFunction_In_YieldContext)

	// ExitArrowFunction_Await is called when exiting the arrowFunction_Await production.
	ExitArrowFunction_Await(c *ArrowFunction_AwaitContext)

	// ExitArrowFunction_In_Await is called when exiting the arrowFunction_In_Await production.
	ExitArrowFunction_In_Await(c *ArrowFunction_In_AwaitContext)

	// ExitArrowFunction_Yield_Await is called when exiting the arrowFunction_Yield_Await production.
	ExitArrowFunction_Yield_Await(c *ArrowFunction_Yield_AwaitContext)

	// ExitArrowFunction_In_Yield_Await is called when exiting the arrowFunction_In_Yield_Await production.
	ExitArrowFunction_In_Yield_Await(c *ArrowFunction_In_Yield_AwaitContext)

	// ExitArrowParameters is called when exiting the arrowParameters production.
	ExitArrowParameters(c *ArrowParametersContext)

	// ExitArrowParameters_Yield is called when exiting the arrowParameters_Yield production.
	ExitArrowParameters_Yield(c *ArrowParameters_YieldContext)

	// ExitArrowParameters_Await is called when exiting the arrowParameters_Await production.
	ExitArrowParameters_Await(c *ArrowParameters_AwaitContext)

	// ExitArrowParameters_Yield_Await is called when exiting the arrowParameters_Yield_Await production.
	ExitArrowParameters_Yield_Await(c *ArrowParameters_Yield_AwaitContext)

	// ExitConciseBody is called when exiting the conciseBody production.
	ExitConciseBody(c *ConciseBodyContext)

	// ExitConciseBody_In is called when exiting the conciseBody_In production.
	ExitConciseBody_In(c *ConciseBody_InContext)

	// ExitMethodDefinition is called when exiting the methodDefinition production.
	ExitMethodDefinition(c *MethodDefinitionContext)

	// ExitMethodDefinition_Yield is called when exiting the methodDefinition_Yield production.
	ExitMethodDefinition_Yield(c *MethodDefinition_YieldContext)

	// ExitMethodDefinition_Await is called when exiting the methodDefinition_Await production.
	ExitMethodDefinition_Await(c *MethodDefinition_AwaitContext)

	// ExitMethodDefinition_Yield_Await is called when exiting the methodDefinition_Yield_Await production.
	ExitMethodDefinition_Yield_Await(c *MethodDefinition_Yield_AwaitContext)

	// ExitPropertySetParameterList is called when exiting the propertySetParameterList production.
	ExitPropertySetParameterList(c *PropertySetParameterListContext)

	// ExitGeneratorMethod is called when exiting the generatorMethod production.
	ExitGeneratorMethod(c *GeneratorMethodContext)

	// ExitGeneratorMethod_Yield is called when exiting the generatorMethod_Yield production.
	ExitGeneratorMethod_Yield(c *GeneratorMethod_YieldContext)

	// ExitGeneratorMethod_Await is called when exiting the generatorMethod_Await production.
	ExitGeneratorMethod_Await(c *GeneratorMethod_AwaitContext)

	// ExitGeneratorMethod_Yield_Await is called when exiting the generatorMethod_Yield_Await production.
	ExitGeneratorMethod_Yield_Await(c *GeneratorMethod_Yield_AwaitContext)

	// ExitGeneratorDeclaration is called when exiting the generatorDeclaration production.
	ExitGeneratorDeclaration(c *GeneratorDeclarationContext)

	// ExitGeneratorDeclaration_Yield is called when exiting the generatorDeclaration_Yield production.
	ExitGeneratorDeclaration_Yield(c *GeneratorDeclaration_YieldContext)

	// ExitGeneratorDeclaration_Await is called when exiting the generatorDeclaration_Await production.
	ExitGeneratorDeclaration_Await(c *GeneratorDeclaration_AwaitContext)

	// ExitGeneratorDeclaration_Yield_Await is called when exiting the generatorDeclaration_Yield_Await production.
	ExitGeneratorDeclaration_Yield_Await(c *GeneratorDeclaration_Yield_AwaitContext)

	// ExitGeneratorDeclaration_Default is called when exiting the generatorDeclaration_Default production.
	ExitGeneratorDeclaration_Default(c *GeneratorDeclaration_DefaultContext)

	// ExitGeneratorDeclaration_Yield_Default is called when exiting the generatorDeclaration_Yield_Default production.
	ExitGeneratorDeclaration_Yield_Default(c *GeneratorDeclaration_Yield_DefaultContext)

	// ExitGeneratorDeclaration_Await_Default is called when exiting the generatorDeclaration_Await_Default production.
	ExitGeneratorDeclaration_Await_Default(c *GeneratorDeclaration_Await_DefaultContext)

	// ExitGeneratorDeclaration_Yield_Await_Default is called when exiting the generatorDeclaration_Yield_Await_Default production.
	ExitGeneratorDeclaration_Yield_Await_Default(c *GeneratorDeclaration_Yield_Await_DefaultContext)

	// ExitGeneratorExpression is called when exiting the generatorExpression production.
	ExitGeneratorExpression(c *GeneratorExpressionContext)

	// ExitGeneratorBody is called when exiting the generatorBody production.
	ExitGeneratorBody(c *GeneratorBodyContext)

	// ExitYieldExpression is called when exiting the yieldExpression production.
	ExitYieldExpression(c *YieldExpressionContext)

	// ExitYieldExpression_In is called when exiting the yieldExpression_In production.
	ExitYieldExpression_In(c *YieldExpression_InContext)

	// ExitYieldExpression_Await is called when exiting the yieldExpression_Await production.
	ExitYieldExpression_Await(c *YieldExpression_AwaitContext)

	// ExitYieldExpression_In_Await is called when exiting the yieldExpression_In_Await production.
	ExitYieldExpression_In_Await(c *YieldExpression_In_AwaitContext)

	// ExitAsyncGeneratorMethod is called when exiting the asyncGeneratorMethod production.
	ExitAsyncGeneratorMethod(c *AsyncGeneratorMethodContext)

	// ExitAsyncGeneratorMethod_Yield is called when exiting the asyncGeneratorMethod_Yield production.
	ExitAsyncGeneratorMethod_Yield(c *AsyncGeneratorMethod_YieldContext)

	// ExitAsyncGeneratorMethod_Await is called when exiting the asyncGeneratorMethod_Await production.
	ExitAsyncGeneratorMethod_Await(c *AsyncGeneratorMethod_AwaitContext)

	// ExitAsyncGeneratorMethod_Yield_Await is called when exiting the asyncGeneratorMethod_Yield_Await production.
	ExitAsyncGeneratorMethod_Yield_Await(c *AsyncGeneratorMethod_Yield_AwaitContext)

	// ExitAsyncGeneratorDeclaration is called when exiting the asyncGeneratorDeclaration production.
	ExitAsyncGeneratorDeclaration(c *AsyncGeneratorDeclarationContext)

	// ExitAsyncGeneratorDeclaration_Yield is called when exiting the asyncGeneratorDeclaration_Yield production.
	ExitAsyncGeneratorDeclaration_Yield(c *AsyncGeneratorDeclaration_YieldContext)

	// ExitAsyncGeneratorDeclaration_Await is called when exiting the asyncGeneratorDeclaration_Await production.
	ExitAsyncGeneratorDeclaration_Await(c *AsyncGeneratorDeclaration_AwaitContext)

	// ExitAsyncGeneratorDeclaration_Yield_Await is called when exiting the asyncGeneratorDeclaration_Yield_Await production.
	ExitAsyncGeneratorDeclaration_Yield_Await(c *AsyncGeneratorDeclaration_Yield_AwaitContext)

	// ExitAsyncGeneratorDeclaration_Default is called when exiting the asyncGeneratorDeclaration_Default production.
	ExitAsyncGeneratorDeclaration_Default(c *AsyncGeneratorDeclaration_DefaultContext)

	// ExitAsyncGeneratorDeclaration_Yield_Default is called when exiting the asyncGeneratorDeclaration_Yield_Default production.
	ExitAsyncGeneratorDeclaration_Yield_Default(c *AsyncGeneratorDeclaration_Yield_DefaultContext)

	// ExitAsyncGeneratorDeclaration_Await_Default is called when exiting the asyncGeneratorDeclaration_Await_Default production.
	ExitAsyncGeneratorDeclaration_Await_Default(c *AsyncGeneratorDeclaration_Await_DefaultContext)

	// ExitAsyncGeneratorDeclaration_Yield_Await_Default is called when exiting the asyncGeneratorDeclaration_Yield_Await_Default production.
	ExitAsyncGeneratorDeclaration_Yield_Await_Default(c *AsyncGeneratorDeclaration_Yield_Await_DefaultContext)

	// ExitAsyncGeneratorExpression is called when exiting the asyncGeneratorExpression production.
	ExitAsyncGeneratorExpression(c *AsyncGeneratorExpressionContext)

	// ExitAsyncGeneratorBody is called when exiting the asyncGeneratorBody production.
	ExitAsyncGeneratorBody(c *AsyncGeneratorBodyContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassDeclaration_Yield is called when exiting the classDeclaration_Yield production.
	ExitClassDeclaration_Yield(c *ClassDeclaration_YieldContext)

	// ExitClassDeclaration_Await is called when exiting the classDeclaration_Await production.
	ExitClassDeclaration_Await(c *ClassDeclaration_AwaitContext)

	// ExitClassDeclaration_Yield_Await is called when exiting the classDeclaration_Yield_Await production.
	ExitClassDeclaration_Yield_Await(c *ClassDeclaration_Yield_AwaitContext)

	// ExitClassDeclaration_Default is called when exiting the classDeclaration_Default production.
	ExitClassDeclaration_Default(c *ClassDeclaration_DefaultContext)

	// ExitClassDeclaration_Yield_Default is called when exiting the classDeclaration_Yield_Default production.
	ExitClassDeclaration_Yield_Default(c *ClassDeclaration_Yield_DefaultContext)

	// ExitClassDeclaration_Await_Default is called when exiting the classDeclaration_Await_Default production.
	ExitClassDeclaration_Await_Default(c *ClassDeclaration_Await_DefaultContext)

	// ExitClassDeclaration_Yield_Await_Default is called when exiting the classDeclaration_Yield_Await_Default production.
	ExitClassDeclaration_Yield_Await_Default(c *ClassDeclaration_Yield_Await_DefaultContext)

	// ExitClassExpression is called when exiting the classExpression production.
	ExitClassExpression(c *ClassExpressionContext)

	// ExitClassExpression_Yield is called when exiting the classExpression_Yield production.
	ExitClassExpression_Yield(c *ClassExpression_YieldContext)

	// ExitClassExpression_Await is called when exiting the classExpression_Await production.
	ExitClassExpression_Await(c *ClassExpression_AwaitContext)

	// ExitClassExpression_Yield_Await is called when exiting the classExpression_Yield_Await production.
	ExitClassExpression_Yield_Await(c *ClassExpression_Yield_AwaitContext)

	// ExitClassTail is called when exiting the classTail production.
	ExitClassTail(c *ClassTailContext)

	// ExitClassTail_Yield is called when exiting the classTail_Yield production.
	ExitClassTail_Yield(c *ClassTail_YieldContext)

	// ExitClassTail_Await is called when exiting the classTail_Await production.
	ExitClassTail_Await(c *ClassTail_AwaitContext)

	// ExitClassTail_Yield_Await is called when exiting the classTail_Yield_Await production.
	ExitClassTail_Yield_Await(c *ClassTail_Yield_AwaitContext)

	// ExitClassHeritage is called when exiting the classHeritage production.
	ExitClassHeritage(c *ClassHeritageContext)

	// ExitClassHeritage_Yield is called when exiting the classHeritage_Yield production.
	ExitClassHeritage_Yield(c *ClassHeritage_YieldContext)

	// ExitClassHeritage_Await is called when exiting the classHeritage_Await production.
	ExitClassHeritage_Await(c *ClassHeritage_AwaitContext)

	// ExitClassHeritage_Yield_Await is called when exiting the classHeritage_Yield_Await production.
	ExitClassHeritage_Yield_Await(c *ClassHeritage_Yield_AwaitContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassBody_Yield is called when exiting the classBody_Yield production.
	ExitClassBody_Yield(c *ClassBody_YieldContext)

	// ExitClassBody_Await is called when exiting the classBody_Await production.
	ExitClassBody_Await(c *ClassBody_AwaitContext)

	// ExitClassBody_Yield_Await is called when exiting the classBody_Yield_Await production.
	ExitClassBody_Yield_Await(c *ClassBody_Yield_AwaitContext)

	// ExitClassElement is called when exiting the classElement production.
	ExitClassElement(c *ClassElementContext)

	// ExitClassElement_Yield is called when exiting the classElement_Yield production.
	ExitClassElement_Yield(c *ClassElement_YieldContext)

	// ExitClassElement_Await is called when exiting the classElement_Await production.
	ExitClassElement_Await(c *ClassElement_AwaitContext)

	// ExitClassElement_Yield_Await is called when exiting the classElement_Yield_Await production.
	ExitClassElement_Yield_Await(c *ClassElement_Yield_AwaitContext)

	// ExitAsyncFunctionDeclaration is called when exiting the asyncFunctionDeclaration production.
	ExitAsyncFunctionDeclaration(c *AsyncFunctionDeclarationContext)

	// ExitAsyncFunctionDeclaration_Yield is called when exiting the asyncFunctionDeclaration_Yield production.
	ExitAsyncFunctionDeclaration_Yield(c *AsyncFunctionDeclaration_YieldContext)

	// ExitAsyncFunctionDeclaration_Await is called when exiting the asyncFunctionDeclaration_Await production.
	ExitAsyncFunctionDeclaration_Await(c *AsyncFunctionDeclaration_AwaitContext)

	// ExitAsyncFunctionDeclaration_Yield_Await is called when exiting the asyncFunctionDeclaration_Yield_Await production.
	ExitAsyncFunctionDeclaration_Yield_Await(c *AsyncFunctionDeclaration_Yield_AwaitContext)

	// ExitAsyncFunctionDeclaration_Default is called when exiting the asyncFunctionDeclaration_Default production.
	ExitAsyncFunctionDeclaration_Default(c *AsyncFunctionDeclaration_DefaultContext)

	// ExitAsyncFunctionDeclaration_Yield_Default is called when exiting the asyncFunctionDeclaration_Yield_Default production.
	ExitAsyncFunctionDeclaration_Yield_Default(c *AsyncFunctionDeclaration_Yield_DefaultContext)

	// ExitAsyncFunctionDeclaration_Await_Default is called when exiting the asyncFunctionDeclaration_Await_Default production.
	ExitAsyncFunctionDeclaration_Await_Default(c *AsyncFunctionDeclaration_Await_DefaultContext)

	// ExitAsyncFunctionDeclaration_Yield_Await_Default is called when exiting the asyncFunctionDeclaration_Yield_Await_Default production.
	ExitAsyncFunctionDeclaration_Yield_Await_Default(c *AsyncFunctionDeclaration_Yield_Await_DefaultContext)

	// ExitAsyncFunctionExpression is called when exiting the asyncFunctionExpression production.
	ExitAsyncFunctionExpression(c *AsyncFunctionExpressionContext)

	// ExitAsyncMethod is called when exiting the asyncMethod production.
	ExitAsyncMethod(c *AsyncMethodContext)

	// ExitAsyncMethod_Yield is called when exiting the asyncMethod_Yield production.
	ExitAsyncMethod_Yield(c *AsyncMethod_YieldContext)

	// ExitAsyncMethod_Await is called when exiting the asyncMethod_Await production.
	ExitAsyncMethod_Await(c *AsyncMethod_AwaitContext)

	// ExitAsyncMethod_Yield_Await is called when exiting the asyncMethod_Yield_Await production.
	ExitAsyncMethod_Yield_Await(c *AsyncMethod_Yield_AwaitContext)

	// ExitAsyncFunctionBody is called when exiting the asyncFunctionBody production.
	ExitAsyncFunctionBody(c *AsyncFunctionBodyContext)

	// ExitAwaitExpression is called when exiting the awaitExpression production.
	ExitAwaitExpression(c *AwaitExpressionContext)

	// ExitAwaitExpression_Yield is called when exiting the awaitExpression_Yield production.
	ExitAwaitExpression_Yield(c *AwaitExpression_YieldContext)

	// ExitAsyncArrowFunction is called when exiting the asyncArrowFunction production.
	ExitAsyncArrowFunction(c *AsyncArrowFunctionContext)

	// ExitAsyncArrowFunction_In is called when exiting the asyncArrowFunction_In production.
	ExitAsyncArrowFunction_In(c *AsyncArrowFunction_InContext)

	// ExitAsyncArrowFunction_Yield is called when exiting the asyncArrowFunction_Yield production.
	ExitAsyncArrowFunction_Yield(c *AsyncArrowFunction_YieldContext)

	// ExitAsyncArrowFunction_In_Yield is called when exiting the asyncArrowFunction_In_Yield production.
	ExitAsyncArrowFunction_In_Yield(c *AsyncArrowFunction_In_YieldContext)

	// ExitAsyncArrowFunction_Await is called when exiting the asyncArrowFunction_Await production.
	ExitAsyncArrowFunction_Await(c *AsyncArrowFunction_AwaitContext)

	// ExitAsyncArrowFunction_In_Await is called when exiting the asyncArrowFunction_In_Await production.
	ExitAsyncArrowFunction_In_Await(c *AsyncArrowFunction_In_AwaitContext)

	// ExitAsyncArrowFunction_Yield_Await is called when exiting the asyncArrowFunction_Yield_Await production.
	ExitAsyncArrowFunction_Yield_Await(c *AsyncArrowFunction_Yield_AwaitContext)

	// ExitAsyncArrowFunction_In_Yield_Await is called when exiting the asyncArrowFunction_In_Yield_Await production.
	ExitAsyncArrowFunction_In_Yield_Await(c *AsyncArrowFunction_In_Yield_AwaitContext)

	// ExitAsyncArrowBindingIdentifier is called when exiting the asyncArrowBindingIdentifier production.
	ExitAsyncArrowBindingIdentifier(c *AsyncArrowBindingIdentifierContext)

	// ExitAsyncArrowBindingIdentifier_Yield is called when exiting the asyncArrowBindingIdentifier_Yield production.
	ExitAsyncArrowBindingIdentifier_Yield(c *AsyncArrowBindingIdentifier_YieldContext)

	// ExitCoverCallExpressionAndAsyncArrowHead is called when exiting the coverCallExpressionAndAsyncArrowHead production.
	ExitCoverCallExpressionAndAsyncArrowHead(c *CoverCallExpressionAndAsyncArrowHeadContext)

	// ExitCoverCallExpressionAndAsyncArrowHead_Yield is called when exiting the coverCallExpressionAndAsyncArrowHead_Yield production.
	ExitCoverCallExpressionAndAsyncArrowHead_Yield(c *CoverCallExpressionAndAsyncArrowHead_YieldContext)

	// ExitCoverCallExpressionAndAsyncArrowHead_Await is called when exiting the coverCallExpressionAndAsyncArrowHead_Await production.
	ExitCoverCallExpressionAndAsyncArrowHead_Await(c *CoverCallExpressionAndAsyncArrowHead_AwaitContext)

	// ExitCoverCallExpressionAndAsyncArrowHead_Yield_Await is called when exiting the coverCallExpressionAndAsyncArrowHead_Yield_Await production.
	ExitCoverCallExpressionAndAsyncArrowHead_Yield_Await(c *CoverCallExpressionAndAsyncArrowHead_Yield_AwaitContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitScriptBody is called when exiting the scriptBody production.
	ExitScriptBody(c *ScriptBodyContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitModuleBody is called when exiting the moduleBody production.
	ExitModuleBody(c *ModuleBodyContext)

	// ExitModuleItem is called when exiting the moduleItem production.
	ExitModuleItem(c *ModuleItemContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitImportClause is called when exiting the importClause production.
	ExitImportClause(c *ImportClauseContext)

	// ExitImportedDefaultBinding is called when exiting the importedDefaultBinding production.
	ExitImportedDefaultBinding(c *ImportedDefaultBindingContext)

	// ExitNameSpaceImport is called when exiting the nameSpaceImport production.
	ExitNameSpaceImport(c *NameSpaceImportContext)

	// ExitNamedImports is called when exiting the namedImports production.
	ExitNamedImports(c *NamedImportsContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitImportsList is called when exiting the importsList production.
	ExitImportsList(c *ImportsListContext)

	// ExitImportSpecifier is called when exiting the importSpecifier production.
	ExitImportSpecifier(c *ImportSpecifierContext)

	// ExitModuleSpecifier is called when exiting the moduleSpecifier production.
	ExitModuleSpecifier(c *ModuleSpecifierContext)

	// ExitImportedBinding is called when exiting the importedBinding production.
	ExitImportedBinding(c *ImportedBindingContext)

	// ExitExportDeclaration is called when exiting the exportDeclaration production.
	ExitExportDeclaration(c *ExportDeclarationContext)

	// ExitExportClause is called when exiting the exportClause production.
	ExitExportClause(c *ExportClauseContext)

	// ExitExportsList is called when exiting the exportsList production.
	ExitExportsList(c *ExportsListContext)

	// ExitExportSpecifier is called when exiting the exportSpecifier production.
	ExitExportSpecifier(c *ExportSpecifierContext)

	// ExitAsyncConciseBody is called when exiting the asyncConciseBody production.
	ExitAsyncConciseBody(c *AsyncConciseBodyContext)

	// ExitAsyncConciseBody_In is called when exiting the asyncConciseBody_In production.
	ExitAsyncConciseBody_In(c *AsyncConciseBody_InContext)
}
