package token

import "fmt"

// Token represents a token with a type and a value, which is represented as
// string. It also holds its starting position in the source.
type Token struct {
	Type  Type
	Value string
	Pos   int
}

// New creates a new token with the given type, value, start position and token
// length.
func New(t Type, v string, start, length int) Token {
	return Token{
		Type:  t,
		Value: v,
		Pos:   start,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%v(%v), offset %v", t.Type.String(), t.Value, t.Pos)
}
