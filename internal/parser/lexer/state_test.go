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
			// >
			{">", []token.Type{token.GreaterThan}},
			{">>", []token.Type{token.RightShift}},
			{">>=", []token.Type{token.RightShiftAssign}},
			{">>=>", []token.Type{token.RightShiftAssign, token.GreaterThan}},
			{">>==", []token.Type{token.RightShiftAssign, token.Assign}},
			{">>==>", []token.Type{token.RightShiftAssign, token.Arrow}},
			{">>==>=", []token.Type{token.RightShiftAssign, token.Arrow, token.Assign}},
			{">>==>==", []token.Type{token.RightShiftAssign, token.Arrow, token.Equals}},
			{">>==>===", []token.Type{token.RightShiftAssign, token.Arrow, token.StrictEquals}},
			{">>==>===>", []token.Type{token.RightShiftAssign, token.Arrow, token.StrictEquals, token.GreaterThan}},
			{">>===>===>", []token.Type{token.RightShiftAssign, token.Equals, token.GreaterThanOrEqualTo, token.Equals, token.GreaterThan}},
			{">>====>===>", []token.Type{token.RightShiftAssign, token.StrictEquals, token.GreaterThanOrEqualTo, token.Equals, token.GreaterThan}},
			{">>>", []token.Type{token.UnsignedRightShift}},
			{">>>=", []token.Type{token.UnsignedRightShiftAssign}},
			{">>>>", []token.Type{token.UnsignedRightShift, token.GreaterThan}},
			{">>>>>", []token.Type{token.UnsignedRightShift, token.RightShift}},
			{">>>>>>", []token.Type{token.UnsignedRightShift, token.UnsignedRightShift}},
			// <
		},
	}
}
