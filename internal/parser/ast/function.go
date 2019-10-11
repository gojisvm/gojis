package ast

type FormalParameters struct {
	FunctionRestParameter *FunctionRestParameter
	FormalParameterList   *FormalParameterList
	Comma                 bool
}

type FunctionRestParameter struct {
	BindingRestElement *BindingRestElement
}

type FormalParameterList struct {
	FormalParameters []*FormalParameter
}

type FormalParameter struct {
	BindingElement *BindingElement
}

type FunctionBody struct {
	FunctionStatementList *FunctionStatementList
}

type FunctionStatementList struct {
	StatementList *StatementList
}

type GeneratorBody struct {
	FunctionBody *FunctionBody
}

type ArrowFunction struct {
	ArrowParameters *ArrowParameters
	ConciseBody     *ConciseBody
}

type ArrowParameters struct {
	BindingIdentifier                                 *BindingIdentifier
	CoverParenthesizedExpressionAndArrowParameterList *CoverParenthesizedExpressionAndArrowParameterList
}

type ConciseBody struct {
	AssignmentExpression *AssignmentExpression
	FunctionBody         *FunctionBody
}

type CoverParenthesizedExpressionAndArrowParameterList struct {
	Expression        *Expression
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
	Comma             bool
	Ellipsis          bool
}
