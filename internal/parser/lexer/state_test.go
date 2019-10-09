package lexer

import (
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
)

func Test_lexToken(t *testing.T) {
	_punctuators().Execute(t)
}

func _punctuators() TokenSequenceTests {
	return TokenSequenceTests{
		name: "punctuator token sequence tests",
		tests: []TokenSequenceTest{
			{">", []token.Type{token.GreaterThan}},
			{">>", []token.Type{token.RightShift}},
			{">>>", []token.Type{token.UnsignedRightShift}},
			{">>>=", []token.Type{token.UnsignedRightShiftAssign}},
			{">>>>", []token.Type{token.UnsignedRightShift, token.GreaterThan}},
		},
	}
}
