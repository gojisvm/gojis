package lexer

import (
	"fmt"
	"testing"
)

func TestSmall(t *testing.T) {
	data := "/* foo bar * / \nfoobar /** */"
	l := New([]byte(data))
	go l.StartLexing()
	for t := range l.TokenStream().Tokens() {
		fmt.Printf("token: %v\n", t.String())
	}

	t.Fail()
}
