package lexer

import (
	"sync"
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
	"github.com/stretchr/testify/require"
)

type Tests interface {
	Execute(*testing.T)
}

type TokenSequenceTests struct {
	name string
	// initial must always be lexToken
	tests []TokenSequenceTest
}

type TokenSequenceTest struct {
	data  string
	types []token.Type
}

func (s TokenSequenceTests) Execute(t *testing.T) {
	for _, test := range s.tests {
		t.Run(s.name, func(t *testing.T) {
			require := require.New(t)

			t.Logf("data: %v", test.data)

			var wg sync.WaitGroup

			l := newWithInitialState([]byte(test.data), lexToken)
			wg.Add(1)
			go func() {
				defer wg.Done()

				err := l.StartLexing()
				require.NoError(err)
			}()

			var cnt int
			for next := range l.TokenStream().Tokens() {
				t.Logf("cnt: %v, recv: %v", cnt, next.Type.String())
				require.Less(cnt, len(test.types), "Received more tokens than expected")

				require.Equalf(test.types[cnt], next.Type, "Did not receive expected token type, expected %v, but got %v", test.types[cnt], next.Type)

				cnt++
			}

			require.Equal(len(test.types), cnt, "Received too few tokens")

			wg.Wait()
		})
	}
}

type SingleTokenTests struct {
	name    string
	initial state
	typ     token.Type

	tests []SingleTokenTest
}

// SingleTokenTest represents a test, where the lexer is expected to produce
// exactly one token with the value of the data passed in.
//
//	"12.04e13" -> NumericLiteral("12.04e13") [NumericLiteral, DecimalLiteral]
type SingleTokenTest struct {
	data string
}

func (s SingleTokenTests) Execute(t *testing.T) {
	for _, test := range s.tests {
		t.Run(s.name, func(t *testing.T) {
			require := require.New(t)

			t.Logf("data: %v", test.data)

			var wg sync.WaitGroup

			l := newWithInitialState([]byte(test.data), s.initial)
			wg.Add(1)
			go func() {
				defer wg.Done()

				err := l.StartLexing()
				require.NoError(err)
			}()

			next, ok := l.TokenStream().Next()
			require.True(ok, "Attempt to read a token failed, lexer did not produce one.")
			require.Equal(test.data, next.Value, "The token value did not match the expected value; the lexer produced an incorrect token.")
			require.Equal(next.Type, s.typ, "The produced token does not have the required token type, wanted %v, but got %v.", s.typ, next.Type)

			_, ok = l.TokenStream().Next()
			require.False(ok, "The lexer produced another token, which was not expected. SingleTokenTests must not produce more than one token.")

			wg.Wait()
		})
	}
}
