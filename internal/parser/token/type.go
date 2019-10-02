package token

// Type is a type of a token.
type Type uint64

//go:generate stringer -type=Type

// Available token type.
const (
	Unknown Type = iota
	Error

	// as of this point, token types should be sorted alphabetically

	BraceClose   // }
	BraceOpen    // {
	BracketClose // ]
	BracketOpen  // [
	Keyword      // any valid keyword, not a value
	Linefeed     // any valid linefeed characters
	QuoteDouble  // "
	QuoteSingle  // '
	String       // any string without surrounding quotes
	Whitespace   // any whitespace that is not a linefeed
)
