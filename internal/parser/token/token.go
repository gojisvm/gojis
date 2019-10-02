package token

import "fmt"

// Token represents a token with a type and a value, which is represented as
// string. It also holds its starting position in the source.
type Token struct {
	Type  Type
	Value string
	Pos   Pos
}

// Pos describes the position of a token, with start and end index.
type Pos struct {
	Start int
	End   int
}

// New creates a new token with the given type, value, start position and token
// length.
func New(t Type, v string, start, length int) Token {
	return Token{
		Type:  t,
		Value: v,
		Pos: Pos{
			Start: start,
			End:   start + length,
		},
	}
}

// IsError determines whether this token is an error token. At the moment, this
// only checks if the type of this token is Error.
func (t Token) IsError() bool {
	return t.Type == Error
}

func (t Token) String() string {
	return fmt.Sprintf("%v(%v), pos %v", t.Type.String(), t.Value, t.Pos)
}
