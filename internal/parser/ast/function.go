package ast

// FormalParameters represents a FormalParameters.
type FormalParameters struct {
	FunctionRestParameter *FunctionRestParameter
	FormalParameterList   *FormalParameterList
	Comma                 bool
}

// FunctionRestParameter represents a FunctionRestParameter.
type FunctionRestParameter struct {
	BindingRestElement *BindingRestElement
}

// FormalParameterList represents a FormalParameterList.
type FormalParameterList struct {
	FormalParameters []*FormalParameter
}

// FormalParameter represents a FormalParameter.
type FormalParameter struct {
	BindingElement *BindingElement
}

// FunctionBody represents a FunctionBody.
type FunctionBody struct {
	FunctionStatementList *FunctionStatementList
}

// FunctionStatementList represents a FunctionStatementList.
type FunctionStatementList struct {
	StatementList *StatementList
}

// GeneratorBody represents a GeneratorBody.
type GeneratorBody struct {
	FunctionBody *FunctionBody
}

// ArrowFunction represents a ArrowFunction.
type ArrowFunction struct {
	ArrowParameters *ArrowParameters
	ConciseBody     *ConciseBody
}

// ArrowParameters represents a ArrowParameters.
type ArrowParameters struct {
	BindingIdentifier                                 *BindingIdentifier
	CoverParenthesizedExpressionAndArrowParameterList *CoverParenthesizedExpressionAndArrowParameterList
}

// ConciseBody represents a ConciseBody.
type ConciseBody struct {
	AssignmentExpression *AssignmentExpression
	FunctionBody         *FunctionBody
}

// CoverParenthesizedExpressionAndArrowParameterList represents a CoverParenthesizedExpressionAndArrowParameterList.
type CoverParenthesizedExpressionAndArrowParameterList struct {
	Expression        *Expression
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
	Comma             bool
	Ellipsis          bool
}

// MethodDefinition represents a MethodDefinition.
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

// PropertySetParameterList represents a PropertySetParameterList.
type PropertySetParameterList struct {
	FormalParameter *FormalParameter
}

// UniqueFormalParameters represents a UniqueFormalParameters.
type UniqueFormalParameters struct {
	FormalParameters *FormalParameters
}

// GeneratorMethod represents a GeneratorMethod.
type GeneratorMethod struct {
	PropertyName           *PropertyName
	UniqueFormalParameters *UniqueFormalParameters
	GeneratorBody          *GeneratorBody
}

// AsyncMethod represents a AsyncMethod.
type AsyncMethod struct {
	PropertyName           *PropertyName
	UniqueFormalParameters *UniqueFormalParameters
	AsyncFunctionBody      *AsyncFunctionBody
}

// AsyncFunctionBody represents a AsyncFunctionBody.
type AsyncFunctionBody struct {
	FunctionBody *FunctionBody
}

// AsyncGeneratorMethod represents a AsyncGeneratorMethod.
type AsyncGeneratorMethod struct {
	PropertyName           *PropertyName
	UniqueFormalParameters *UniqueFormalParameters
	AsyncGeneratorBody     *AsyncGeneratorBody
}

// AsyncGeneratorBody represents a AsyncGeneratorBody.
type AsyncGeneratorBody struct {
	FunctionBody *FunctionBody
}

// AsyncArrowFunction represents a AsyncArrowFunction.
type AsyncArrowFunction struct {
	AsyncArrowBindingIdentifier          *AsyncArrowBindingIdentifier
	AsyncConciseBody                     *AsyncConciseBody
	CoverCallExpressionAndAsyncArrowHead *CoverCallExpressionAndAsyncArrowHead
}

// AsyncConciseBody represents a AsyncConciseBody.
type AsyncConciseBody struct {
	AssignmentExpression *AssignmentExpression
	AsyncFunctionBody    *AsyncFunctionBody
}

// AsyncArrowBindingIdentifier represents a AsyncArrowBindingIdentifier.
type AsyncArrowBindingIdentifier struct {
	BindingIdentifier *BindingIdentifier
}

// CoverCallExpressionAndAsyncArrowHead represents a CoverCallExpressionAndAsyncArrowHead.
type CoverCallExpressionAndAsyncArrowHead struct {
	MemberExpression *MemberExpression
	Arguments        *Arguments
}
