package token

// Type is a type of a token.
type Type uint64

//go:generate stringer -type=Type

// Available token type.
const (
	Unknown Type = iota
	Error

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
)
