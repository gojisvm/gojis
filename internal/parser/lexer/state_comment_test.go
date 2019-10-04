package lexer

import (
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
)

func Test_lexSingleLineComment(t *testing.T) {
	SingleTokenTests{
		name:    "single line comment",
		initial: lexComment,
		types:   []token.Type{token.SingleLineComment},
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

func Test_lexMultiLineComment(t *testing.T) {
	SingleTokenTests{
		name:    "multi line comment",
		initial: lexComment,
		types:   []token.Type{token.MultiLineComment},
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
