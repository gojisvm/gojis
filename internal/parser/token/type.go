package token

// Type is a type of a token.
type Type uint64

//go:generate stringer -type=Type

// Available token type.
const (
	Unknown Type = iota
	Error

	// as of this point, token types should be sorted alphabetically

	LineTerminator    // 11.3
	MultiLineComment  // 11.4
	SingleLineComment // 11.4
	Whitespace        // 11.2
)
