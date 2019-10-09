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
			{">=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>", []token.Type{token.GreaterThanOrEqualTo, token.RightShift}},
			{">=>>=", []token.Type{token.GreaterThanOrEqualTo, token.RightShiftAssign}},
			{">=>>=>>>", []token.Type{token.GreaterThanOrEqualTo, token.RightShiftAssign, token.UnsignedRightShift}},
			{">=>>=>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>=>>>>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>=>>>", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>=>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>=>>>>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>>", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">=>>>>>>>>=", []token.Type{token.GreaterThanOrEqualTo}},
			{">>", []token.Type{}},
			{">>=", []token.Type{}},
			{">>=>>", []token.Type{}},
			{">>=>>=", []token.Type{}},
			{">>=>>=>>>", []token.Type{}},
			{">>=>>=>>>=", []token.Type{}},
			{">>=>>=>>>>>>=", []token.Type{}},
			{">>=>>>", []token.Type{}},
			{">>=>>>=", []token.Type{}},
			{">>=>>>>=", []token.Type{}},
			{">>=>>>>=>>>", []token.Type{}},
			{">>=>>>>=>>>=", []token.Type{}},
			{">>=>>>>=>>>>>>=", []token.Type{}},
			{">>=>>>>>", []token.Type{}},
			{">>=>>>>>=", []token.Type{}},
			{">>=>>>>>>=", []token.Type{}},
			{">>=>>>>>>>>=", []token.Type{}},
			{">>>", []token.Type{}},
			{">>>=", []token.Type{}},
			{">>>=>>>", []token.Type{}},
			{">>>=>>>=", []token.Type{}},
			{">>>=>>>>>>=", []token.Type{}},
			{">>>>", []token.Type{}},
			{">>>>=", []token.Type{}},
			{">>>>=>>>", []token.Type{}},
			{">>>>=>>>=", []token.Type{}},
			{">>>>=>>>>>>=", []token.Type{}},
			{">>>>>", []token.Type{}},
			{">>>>>=", []token.Type{}},
			{">>>>>=>>>", []token.Type{}},
			{">>>>>=>>>=", []token.Type{}},
			{">>>>>=>>>>>>=", []token.Type{}},
			{">>>>>>", []token.Type{}},
			{">>>>>>=", []token.Type{}},
			{">>>>>>>=", []token.Type{}},
			{">>>>>>>>=", []token.Type{}},
			{">>>>>>>>>=", []token.Type{}},
		},
	}
}
