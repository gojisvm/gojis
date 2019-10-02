package token

import "fmt"

// Token represents a token with a type and a value, which is represented as
// string. It also holds its starting position in the source.
type Token struct {
	Types []Type
	Value string
	Pos   int
}

// New creates a new token with the given type, value, start position and token
// length.
func New(t []Type, v string, start, length int) Token {
	return Token{
		Types: t,
		Value: v,
		Pos:   start,
	}
}

// IsError determines whether this token is an error token. At the moment, this
// only checks if the type of this token is Error.
func (t Token) IsError() bool {
	for _, typ := range t.Types {
		if typ == Error {
			return true
		}
	}
	return false
}

func (t Token) String() string {
	if len(t.Types) > 1 {
		return fmt.Sprintf("%v(%v), pos %v (actually has types %v)", t.Types[0].String(), t.Value, t.Pos, t.Types)
	}
	return fmt.Sprintf("%v(%v), pos %v", t.Types[0].String(), t.Value, t.Pos)
}
