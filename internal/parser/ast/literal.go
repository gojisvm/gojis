package ast

// Literal represents a Literal node.
type Literal struct {
	NullLiteral    string
	BooleanLiteral string
	NumericLiteral *NumericLiteral
	StringLiteral  string
}

// NumericLiteral represents a NumericLiteral node.
type NumericLiteral struct {
	DecimalLiteral       string
	BinaryIntegerLiteral string
	OctalIntegerLiteral  string
	HexIntegerLiteral    string
}

// ArrayLiteral represents a ArrayLiteral node.
type ArrayLiteral struct {
	Elision     *Elision
	ElementList *ElementList
	Comma       bool
}

// ElementList represents a ElementList node.
type ElementList struct {
	Elision              *Elision
	AssignmentExpression *AssignmentExpression
	SpreadElement        *SpreadElement
	ElementList          *ElementList
	Comma                bool
}

// SpreadElement represents a SpreadElement node.
type SpreadElement struct {
	AssignmentExpression *AssignmentExpression
}

// ObjectLiteral represents a ObjectLiteral node.
type ObjectLiteral struct {
	PropertyDefinitionList *PropertyDefinitionList
	Comma                  bool
}

// PropertyDefinitionList represents a PropertyDefinitionList node.
type PropertyDefinitionList struct {
	PropertyDefinitions []*PropertyDefinition
}

// PropertyDefinition represents a PropertyDefinition node.
type PropertyDefinition struct {
	Colon                bool
	Ellipsis             bool
	IdentifierReference  *IdentifierReference
	CoverInitializedName *CoverInitializedName
	PropertyName         *PropertyName
	AssignmentExpression *AssignmentExpression
	MethodDefinition     *MethodDefinition
}

// PropertyName represents a PropertyName node.
type PropertyName struct {
	LiteralPropertyName  *LiteralPropertyName
	ComputedPropertyName *ComputedPropertyName
}

// LiteralPropertyName represents a LiteralPropertyName node.
type LiteralPropertyName struct {
	IdentifierName string
	StringLiteral  string
	NumericLiteral *NumericLiteral
}

// ComputedPropertyName represents a ComputedPropertyName node.
type ComputedPropertyName struct {
	AssignmentExpression *AssignmentExpression
}

// CoverInitializedName represents a CoverInitializedName node.
type CoverInitializedName struct {
	IdentifierReference *IdentifierReference
	Initializer         *Initializer
}

// TemplateLiteral represents a TemplateLiteral node.
type TemplateLiteral struct {
	NoSubstitutionTemplate string
	SubstitutionTemplate   *SubstitutionTemplate
}

// SubstitutionTemplate represents a SubstitutionTemplate node.
type SubstitutionTemplate struct {
	TemplateHead  string
	Expression    *Expression
	TemplateSpans *TemplateSpans
}

// TemplateSpans represents a TemplateSpans node.
type TemplateSpans struct {
	TemplateTail       string
	TemplateMiddleList *TemplateMiddleList
}

// TemplateMiddleList represents a TemplateMiddleList node.
type TemplateMiddleList struct {
	TemplateMiddle     string
	Expression         *Expression
	TemplateMiddleList *TemplateMiddleList
}

// RegularExpressionLiteral represents a RegularExpressionLiteral node.
type RegularExpressionLiteral struct {
	RegularExpressionBody  string
	RegularExpressionFlags string
}
