package ast

type Literal struct {
	NullLiteral    string
	BooleanLiteral string
	NumericLiteral *NumericLiteral
	StringLiteral  string
}

type NumericLiteral struct {
	DecimalLiteral       string
	BinaryIntegerLiteral string
	OctalIntegerLiteral  string
	HexIntegerLiteral    string
}

type ArrayLiteral struct {
	Elision     *Elision
	ElementList *ElementList
	Comma       bool
}

type ElementList struct {
	Elision              *Elision
	AssignmentExpression *AssignmentExpression
	SpreadElement        *SpreadElement
	ElementList          *ElementList
	Comma                bool
}

type SpreadElement struct {
	AssignmentExpression *AssignmentExpression
}

type ObjectLiteral struct {
	PropertyDefinitionList *PropertyDefinitionList
	Comma                  bool
}

type PropertyDefinitionList struct {
	PropertyDefinitions []*PropertyDefinition
}

type PropertyDefinition struct {
	Colon                bool
	Ellipsis             bool
	IdentifierReference  *IdentifierReference
	CoverInitializedName *CoverInitializedName
	PropertyName         *PropertyName
	AssignmentExpression *AssignmentExpression
	MethodDefinition     *MethodDefinition
}

type PropertyName struct {
	LiteralPropertyName  *LiteralPropertyName
	ComputedPropertyName *ComputedPropertyName
}

type LiteralPropertyName struct {
	IdentifierName string
	StringLiteral  string
	NumericLiteral *NumericLiteral
}

type ComputedPropertyName struct {
	AssignmentExpression *AssignmentExpression
}

type CoverInitializedName struct {
	IdentifierReference *IdentifierReference
	Initializer         *Initializer
}

type TemplateLiteral struct {
	NoSubstitutionTemplate string
	SubstitutionTemplate   *SubstitutionTemplate
}

type SubstitutionTemplate struct {
	TemplateHead  string
	Expression    *Expression
	TemplateSpans *TemplateSpans
}

type TemplateSpans struct {
	TemplateTail       string
	TemplateMiddleList *TemplateMiddleList
}

type TemplateMiddleList struct {
	TemplateMiddle     string
	Expression         *Expression
	TemplateMiddleList *TemplateMiddleList
}

type RegularExpressionLiteral struct {
	RegularExpressionBody  string
	RegularExpressionFlags string
}
