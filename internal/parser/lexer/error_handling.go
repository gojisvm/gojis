package lexer

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
)

func tokenMismatch(expected ...matcher.M) state {
	return func(l *Lexer) state {
		return errorf("Unexpected token, got '%s' but expected %s", string(l.peek()), matcher.Merge(expected...).String())
	}
}

func unexpectedToken(l *Lexer) state {
	return errorf("Unexpected token '%s'", string(l.peek()))
}

func errorf(msg string, args ...interface{}) state {
	return func(l *Lexer) state {
		l.error(fmt.Sprintf(msg, args...))
		return nil
	}
}
