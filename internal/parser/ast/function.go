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

type MethodDefinition struct {
	PropertyName             *PropertyName
	UniqueFormalPatameters   *UniqueFormalParameters
	FunctionBody             *FunctionBody
	GeneratorMethod          *GeneratorMethod
	AsyncMethod              *AsyncMethod
	AsyncGeneratorMethod     *AsyncGeneratorMethod
	PropertySetParameterList *PropertySetParameterList
	Get                      bool
	Set                      bool
}

type PropertySetParameterList struct {
	FormalParameter *FormalParameter
}

type UniqueFormalParameters struct {
	FormalParameters *FormalParameters
}

type GeneratorMethod struct {
	PropertyName           *PropertyName
	UniqueFormalParameters *UniqueFormalParameters
	GeneratorBody          *GeneratorBody
}

type AsyncMethod struct {
	PropertyName           *PropertyName
	UniqueFormalParameters *UniqueFormalParameters
	AsyncFunctionBody      *AsyncFunctionBody
}

type AsyncFunctionBody struct {
	FunctionBody *FunctionBody
}

type AsyncGeneratorMethod struct {
	PropertyName           *PropertyName
	UniqueFormalParameters *UniqueFormalParameters
	AsyncGeneratorBody     *AsyncGeneratorBody
}

type AsyncGeneratorBody struct {
	FunctionBody *FunctionBody
}

type AsyncArrowFunction struct {
	AsyncArrowBindingIdentifier          *AsyncArrowBindingIdentifier
	AsyncConciseBody                     *AsyncConciseBody
	CoverCallExpressionAndAsyncArrowHead *CoverCallExpressionAndAsyncArrowHead
}

type AsyncConciseBody struct {
	AssignmentExpression *AssignmentExpression
	AsyncFunctionBody    *AsyncFunctionBody
}

type AsyncArrowBindingIdentifier struct {
	BindingIdentifier *BindingIdentifier
}

type CoverCallExpressionAndAsyncArrowHead struct {
	MemberExpression *MemberExpression
	Arguments        *Arguments
}
