package token

// Type is a type of a token.
type Type uint64

func (t Type) Implies(other Type) bool {
	for _, implication := range implications[t] {
		if implication == other {
			return true
		}
	}
	return false
}

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
	NumericLiteral // 11.5
	StringLiteral  // 11.5
	Template       // 11.5

	ReservedWord // 11.6

	Null                 // 11.8.1
	Boolean              // 11.8.2
	DecimalLiteral       // 11.8.3
	BinaryIntegerLiteral // 11.8.3
	OctalIntegerLiteral  // 11.8.3
	HexIntegerLiteral    // 11.8.3

	RegularExpressionLiteral // 11.8.5

	Yield // 12.1
	Await // 12.1
	This  // 12.2

	NewType
	Super
	Target
	Delete
	Void
	Typeof

	// punctuators

	UnsignedRightShiftAssign
	Ellipsis
	LeftShiftAssign
	UnsignedRightShift
	RightShiftAssign
	StrictEquals
	StrictNotEquals
	PowerAssign
	LeftShift
	LessThanOrEqualTo
	RightShift
	GreaterThanOrEqualTo
	Equals
	NotEquals
	UpdatePlus
	UpdateMinus
	Power
	LogicalAnd
	LogicalOr
	PlusAssign
	MinusAssign
	MultiplyAssign
	ModuloAssign
	AndAssign
	OrAssign
	XorAssign
	Arrow
	GreaterThan
	Dot
	SemiColon
	Comma
	Tilde
	LessThan
	ParOpen
	ParClose
	BraceOpen
	BraceClose
	BracketOpen
	BracketClose
	LogicalNot
	Plus
	Minus
	Multiply
	Modulo
	BitwiseAnd
	BitwiseOr
	BitwiseXor
	BitwiseNot
	QuestionMark
	Colon
	Assign
	DivAssign
	Div
)
