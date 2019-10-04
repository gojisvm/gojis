package lexer

import (
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
)

func TestSingleLineComment(t *testing.T) {
	SingleTokenTests{
		name:    "single line comment",
		types:   []token.Type{token.SingleLineComment},
		initial: lexComment,
		tests: []SingleTokenTest{
			{"//"},
			{"// "},
			{"//*"},
			{"//**/"},
			{"//foobar"},
			{"//\u7fffabc"},
			{"// snafu"},
			{"// // abc"},
			{"// */ lorem"},
			{"// ipsum *// dolor"},
		},
	}.Execute(t)
}

func TestMultiLineComment(t *testing.T) {
	SingleTokenTests{
		name:    "multi line comment",
		types:   []token.Type{token.MultiLineComment},
		initial: lexComment,
		tests: []SingleTokenTest{
			{"/**/"},
			{"/* */"},
			{"/* f*/"},
			{"/* f */"},
			{"/*f */"},
			{"/* foobar */"},
			{"/* foobar * snafu */"},
			{"/* foobar * snafu **/"},
			{"/* foobar * snafu * */"},
			{"/* foobar * snafu * /*/"},
			{"/* foobar * snafu * / */"},
			{"/* foobar * snafu * /f*/"},
			{"/** foobar * snafu * /f*/"},
			{"/* *foobar * snafu * /f*/"},
			{"/* * foobar * snafu * /f*/"},
			{"/*/**/"},
			{"/*/* */"},
		},
	}.Execute(t)
}
