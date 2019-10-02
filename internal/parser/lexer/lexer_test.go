package lexer

import (
	"fmt"
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
	"github.com/stretchr/testify/require"
)

func TestSingleLineComment(t *testing.T) {
	contents := []struct {
		data, content string
	}{
		{"// Hello World!", "// Hello World!"},
		{"// Hello World!\n", "// Hello World!"},
		{"// Hello World!\n\n", "// Hello World!"},
		{"//foo", "//foo"},
		{"//f o / bar // snafu //// abcdef ", "//f o / bar // snafu //// abcdef "},
	}
	for _, c := range contents {
		testToken(t, c.data, c.content, token.SingleLineComment)
	}
}

func testToken(_t *testing.T, data, content string, types ...token.Type) {
	_t.Run(data, func(t *testing.T) {
		require := require.New(t)

		l := New([]byte(data))
		go l.StartLexing()
		next, ok := l.TokenStream().Next()
		require.True(ok)
		require.ElementsMatch(types, next.Types)
		require.Equal(content, next.Value)

		for range l.TokenStream().Tokens() {
			// drain token stream so that the lexer closes the token stream
		}
	})
}

func TestSmall(t *testing.T) {
	data := "/* foo bar * / \nfoobar /** */"
	l := New([]byte(data))
	go l.StartLexing()
	for t := range l.TokenStream().Tokens() {
		fmt.Printf("token: %v\n", t.String())
	}

	t.Fail()
}
