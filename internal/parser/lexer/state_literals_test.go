package lexer

import (
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
)

func Test_lexNull(t *testing.T) {
	SingleTokenTests{
		name:    "null literal",
		initial: lexNull,
		typ:     token.Null,
		tests: []SingleTokenTest{
			{"null"},
		},
	}.Execute(t)
}

func Test_lexBoolean(t *testing.T) {
	SingleTokenTests{
		name:    "boolean literals",
		initial: lexBoolean,
		typ:     token.Boolean,
		tests: []SingleTokenTest{
			{"true"},
			{"false"},
		},
	}.Execute(t)
}

func Test_lexNumericLiteral(t *testing.T) {
	SingleTokenTests{
		name:    "binary integer literals",
		initial: lexNumericLiteral,
		typ:     token.BinaryIntegerLiteral,
		tests: []SingleTokenTest{
			{"0b0"},
			{"0b1"},
			{"0b00"},
			{"0b01"},
			{"0b10"},
			{"0b11"},
			{"0B0"},
			{"0B1"},
			{"0B00"},
			{"0B01"},
			{"0B10"},
			{"0B11"},
		},
	}.Execute(t)

	SingleTokenTests{
		name:    "octal integer literals",
		initial: lexNumericLiteral,
		typ:     token.OctalIntegerLiteral,
		tests: []SingleTokenTest{
			{"0o0"},
			{"0o1"},
			{"0o00"},
			{"0o01"},
			{"0o10"},
			{"0o11"},
			{"0o01234567"},
			{"0O0"},
			{"0O1"},
			{"0O00"},
			{"0O01"},
			{"0O10"},
			{"0O11"},
			{"0O01234567"},
		},
	}.Execute(t)

	SingleTokenTests{
		name:    "hex integer literals",
		initial: lexNumericLiteral,
		typ:     token.HexIntegerLiteral,
		tests: []SingleTokenTest{
			{"0x0"},
			{"0x1"},
			{"0x00"},
			{"0x01"},
			{"0x10"},
			{"0x11"},
			{"0x0123456789abcdef"},
			{"0x0123456789ABCDEF"},
			{"0X0"},
			{"0X1"},
			{"0X00"},
			{"0X01"},
			{"0X10"},
			{"0X11"},
			{"0X0123456789abcdef"},
			{"0X0123456789ABCDEF"},
		},
	}.Execute(t)

	SingleTokenTests{
		name:    "decimal integer literals",
		initial: lexNumericLiteral,
		typ:     token.DecimalLiteral,
		tests: []SingleTokenTest{
			{"0"},
			{"1"},
			{"2"},
			{"3"},
			{"4"},
			{"5"},
			{"6"},
			{"7"},
			{"8"},
			{"9"},
			{"01"},
			{"02"},
			{"03"},
			{"04"},
			{"05"},
			{"06"},
			{"07"},
			{"08"},
			{"09"},
			{"0.5"},
			{"0.05"},
			{"0.005"},
			{"0.5e3"},
			{"0.5e+3"},
			{"0.5e-3"},
			{".5"},
			{".05"},
			{".005"},
			{".5e3"},
			{".5e+3"},
			{".5e-3"},
			{"5e3"},
			{"5e+3"},
			{"5e-3"},
		},
	}.Execute(t)
}

func Test_lexStringLiteral(t *testing.T) {
	SingleTokenTests{
		name:    "string literals",
		initial: lexStringLiteral,
		typ:     token.StringLiteral,
		tests: []SingleTokenTest{
			{`"0"`},
			{`"0.5e3"`},
			{`"0.5e+3"`},
			{`"0.5e-3"`},
			{`".5"`},
			{`".05"`},
			{`"5e-3"`},
			{`"true"`},
			{`"false"`},
			{`"foobar"`},
			{`"foobar'"`},
			{`"'foobar"`},
			{`"'foobar'"`},
			{`"\x1234"`},
			{`"\x123456z"`},
			{`"\x12345z"`},
			{`"\u1234"`},
			{`"\u123456z"`},
			{`"\u12345z"`},
			{`"\u{1}z23456"`},
			{`"\u{123456}"`},
			{`"\u{1234}56"`},
			{`"\
12{3456}"`},
			{`"\\ \a \b \c"`},
			{`"\0"`},
			{`" "`},
			{`"a b"`},
		},
	}.Execute(t)
}
