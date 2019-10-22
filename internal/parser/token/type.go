package token

// Type is a type of a token.
type Type uint64

//go:generate stringer -type=Type

// Available token type.
const (
	Unknown Type = iota

	LineTerminator // 11.3
	Whitespace     // 11.2

	MultiLineComment  // 11.4
	SingleLineComment // 11.4

	CommonToken    // 11.5
	IdentifierName // 11.5, 11.6
	Punctuator     // 11.5, 11.7
	// NumericLiteral // 11.5
	StringLiteral // 11.5
	// Template       // 11.5
	TemplateHead
	TemplateMiddle
	TemplateTail
	NoSubstitutionTemplate

	Async
	Let

	Await      // 11.6.2.1
	Break      // 11.6.2.1
	Case       // 11.6.2.1
	Catch      // 11.6.2.1
	Class      // 11.6.2.1
	Const      // 11.6.2.1
	Continue   // 11.6.2.1
	Debugger   // 11.6.2.1
	Default    // 11.6.2.1
	Delete     // 11.6.2.1
	Do         // 11.6.2.1
	Else       // 11.6.2.1
	Export     // 11.6.2.1
	Extends    // 11.6.2.1
	Finally    // 11.6.2.1
	For        // 11.6.2.1
	Function   // 11.6.2.1
	If         // 11.6.2.1
	Import     // 11.6.2.1
	In         // 11.6.2.1
	Instanceof // 11.6.2.1
	New_       // 11.6.2.1
	Return     // 11.6.2.1
	Static
	Super  // 11.6.2.1
	Switch // 11.6.2.1
	This   // 11.6.2.1
	Throw  // 11.6.2.1
	Try    // 11.6.2.1
	Typeof // 11.6.2.1
	Var    // 11.6.2.1
	Void   // 11.6.2.1
	While  // 11.6.2.1
	With   // 11.6.2.1
	Yield  // 11.6.2.1

	Enum
	Implements
	Package
	Protected
	Interface
	Private
	Public

	Null                 // 11.8.1
	Boolean              // 11.8.2
	DecimalLiteral       // 11.8.3
	BinaryIntegerLiteral // 11.8.3
	OctalIntegerLiteral  // 11.8.3
	HexIntegerLiteral    // 11.8.3

	RegularExpressionLiteral // 11.8.5

	Target

	// punctuators

	AndAssign
	Arrow
	Assign
	BitwiseAnd
	BitwiseNot
	BitwiseOr
	BitwiseXor
	BraceClose
	BraceOpen
	BracketClose
	BracketOpen
	Colon
	Comma
	Div
	DivAssign
	Dot
	Ellipsis
	Equals
	GreaterThan
	GreaterThanOrEqualTo
	LeftShift
	LeftShiftAssign
	LessThan
	LessThanOrEqualTo
	LogicalAnd
	LogicalNot
	LogicalOr
	Minus
	MinusAssign
	Modulo
	ModuloAssign
	Multiply
	MultiplyAssign
	NotEquals
	OrAssign
	ParClose
	ParOpen
	Plus
	PlusAssign
	Power
	PowerAssign
	QuestionMark
	RightShift
	RightShiftAssign
	SemiColon
	StrictEquals
	StrictNotEquals
	Tilde
	UnsignedRightShift
	UnsignedRightShiftAssign
	UpdateMinus
	UpdatePlus
	XorAssign
)
