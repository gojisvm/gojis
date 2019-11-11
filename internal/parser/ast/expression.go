package ast

// Expression represents a Expression node.
type Expression struct {
	AssignmentExpressions []*AssignmentExpression
}

// AssignmentExpression represents a AssignmentExpression node.
type AssignmentExpression struct {
	ConditionalExpression  *ConditionalExpression
	YieldExpression        *YieldExpression
	ArrowFunction          *ArrowFunction
	AsyncArrowFunction     *AsyncArrowFunction
	LeftHandSideExpression *LeftHandSideExpression
	Assign                 bool
	AssignmentOperator     string
	AssignmentExpression   *AssignmentExpression
}

// ConditionalExpression represents a ConditionalExpression node.
type ConditionalExpression struct {
	LogicalORExpression   *LogicalORExpression
	AssignmentExpression1 *AssignmentExpression
	AssignmentExpression2 *AssignmentExpression
}

// LogicalORExpression represents a LogicalORExpression node.
type LogicalORExpression struct {
	LogicalANDExpression *LogicalANDExpression
	LogicalORExpression  *LogicalORExpression
}

// LogicalANDExpression represents a LogicalANDExpression node.
type LogicalANDExpression struct {
	BitwiseORExpression  *BitwiseORExpression
	LogicalANDExpression *LogicalANDExpression
}

// BitwiseORExpression represents a BitwiseORExpression node.
type BitwiseORExpression struct {
	BitwiseORExpression  *BitwiseORExpression
	BitwiseXORExpression *BitwiseXORExpression
}

// BitwiseANDExpression represents a BitwiseANDExpression node.
type BitwiseANDExpression struct {
	BitwiseANDExpression *BitwiseANDExpression
	EqualityExpression   *EqualityExpression
}

// BitwiseXORExpression represents a BitwiseXORExpression node.
type BitwiseXORExpression struct {
	BitwiseANDExpression *BitwiseANDExpression
	BitwiseXORExpression *BitwiseXORExpression
}

// EqualityExpression represents a EqualityExpression node.
type EqualityExpression struct {
	EqualityExpression   *EqualityExpression
	RelationalExpression *RelationalExpression
	Equals               bool
	StrictEquals         bool
	NotEquals            bool
	StrictNotEquals      bool
}

// RelationalExpression represents a RelationalExpression node.
type RelationalExpression struct {
	ShiftExpression      *ShiftExpression
	RelationalExpression *RelationalExpression
	LessThan             bool
	GreaterThan          bool
	LessThanOrEqualTo    bool
	GreaterThanOrEqualTo bool
	Instanceof           bool
	In                   bool
}

// ShiftExpression represents a ShiftExpression node.
type ShiftExpression struct {
	ShiftExpression    *ShiftExpression
	AdditiveExpression *AdditiveExpression
	LeftShift          bool
	RightShift         bool
	UnsignedRightShift bool
}

// AdditiveExpression represents a AdditiveExpression node.
type AdditiveExpression struct {
	MultiplicativeExpression *MultiplicativeExpression
	AdditiveExpression       *AdditiveExpression
	Plus                     bool
	Minus                    bool
}

// MultiplicativeExpression represents a MultiplicativeExpression node.
type MultiplicativeExpression struct {
	ExponentiationExpression *ExponentiationExpression
	MultiplicativeExpression *MultiplicativeExpression
	Asterisk                 bool
	Slash                    bool
	Modulo                   bool
}

// ExponentiationExpression represents a ExponentiationExpression node.
type ExponentiationExpression struct {
	UnaryExpression          *UnaryExpression
	UpdateExpression         *UpdateExpression
	ExponentiationExpression *ExponentiationExpression
}

// UpdateExpression represents a UpdateExpression node.
type UpdateExpression struct {
	LeftHandSideExpression *LeftHandSideExpression
	UnaryExpression        *UnaryExpression
	PlusPlus               bool
	MinusMinus             bool
}

// UnaryExpression represents a UnaryExpression node.
type UnaryExpression struct {
	UpdateExpression *UpdateExpression
	UnaryExpression  *UnaryExpression
	AwaitExpression  *AwaitExpression
	Delete           bool
	Void             bool
	Typeof           bool
	Plus             bool
	Minus            bool
	Tilde            bool
	ExclamationMark  bool
}

// AwaitExpression represents a AwaitExpression node.
type AwaitExpression struct {
	UnaryExpression *UnaryExpression
}

// YieldExpression represents a YieldExpression node.
type YieldExpression struct {
	AssignmentExpression *AssignmentExpression
	Asterisk             bool
}

// LeftHandSideExpression represents a LeftHandSideExpression node.
type LeftHandSideExpression struct {
	NewExpression  *NewExpression
	CallExpression *CallExpression
}

// NewExpression represents a NewExpression node.
type NewExpression struct {
	MemberExpression *MemberExpression
	NewExpression    *NewExpression
}

// MemberExpression represents a MemberExpression node.
type MemberExpression struct {
	PrimaryExpression *PrimaryExpression
	MemberExpression  *MemberExpression
	Expression        *Expression
	IdentifierName    string
	TemplateLiteral   *TemplateLiteral
	SuperProperty     *SuperProperty
	MetaProperty      *MetaProperty
	Arguments         *Arguments
}

// PrimaryExpression represents a PrimaryExpression node.
type PrimaryExpression struct {
	This                                              bool
	IdentifierReference                               *IdentifierReference
	Literal                                           *Literal
	ArrayLiteral                                      *ArrayLiteral
	ObjectLiteral                                     *ObjectLiteral
	FunctionExpression                                *FunctionExpression
	ClassExpression                                   *ClassExpression
	GeneratorExpression                               *GeneratorExpression
	AsyncFunctionExpression                           *AsyncFunctionExpression
	AsyncGeneratorExpression                          *AsyncGeneratorExpression
	RegularExpressionLiteral                          *RegularExpressionLiteral
	TemplateLiteral                                   *TemplateLiteral
	CoverParenthesizedExpressionAndArrowParameterList *CoverParenthesizedExpressionAndArrowParameterList
}

// FunctionExpression represents a FunctionExpression node.
type FunctionExpression struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	FunctionBody      *FunctionBody
}

// ClassExpression represents a ClassExpression node.
type ClassExpression struct {
	BindingIdentifier *BindingIdentifier
	ClassTail         *ClassTail
}

// ClassTail represents a ClassTail node.
type ClassTail struct {
	ClassHeritage *ClassHeritage
	ClassBody     *ClassBody
}

// ClassHeritage represents a ClassHeritage node.
type ClassHeritage struct {
	LeftHandSideExpression *LeftHandSideExpression
}

// ClassBody represents a ClassBody node.
type ClassBody struct {
	ClassElementList *ClassElementList
}

// ClassElementList represents a ClassElementList node.
type ClassElementList struct {
	ClassElements []*ClassElement
}

// ClassElement represents a ClassElement node.
type ClassElement struct {
	MethodDefinition *MethodDefinition
	Static           bool
}

// GeneratorExpression represents a GeneratorExpression node.
type GeneratorExpression struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	GeneratorBody     *GeneratorBody
}

// AsyncFunctionExpression represents a AsyncFunctionExpression node.
type AsyncFunctionExpression struct {
	FormalParameters  *FormalParameters
	AsyncFunctionBody *AsyncFunctionBody
	BindingIdentifier *BindingIdentifier
}

// AsyncGeneratorExpression represents a AsyncGeneratorExpression node.
type AsyncGeneratorExpression struct {
	BindingIdentifier  *BindingIdentifier
	FormalParameters   *FormalParameters
	AsyncGeneratorBody *AsyncGeneratorBody
}

// CallExpression represents a CallExpression node.
type CallExpression struct {
	CoverCallExpressionAndAsyncArrowHead *CoverCallExpressionAndAsyncArrowHead
	SuperCall                            *SuperCall
	CallExpression                       *CallExpression
	Arguments                            *Arguments
	Expression                           *Expression
	IdentifierName                       string
	TemplateLiteral                      *TemplateLiteral
}

// SuperCall represents a SuperCall node.
type SuperCall struct {
	Arguments *Arguments
}
