package lexer

import (
	"testing"
)

func TestFoo(t *testing.T) {
	data := "_\\uabcdeffoobar"
	l := newWithInitialState([]byte(data), lexIdentifierName)
	go l.StartLexing()
	for tok := range l.TokenStream().Tokens() {
		t.Logf("token: %v\n", tok.String())
	}

	// t.Fail()
}
