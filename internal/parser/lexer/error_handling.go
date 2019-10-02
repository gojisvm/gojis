package lexer

import (
	"fmt"
)

func unexpectedToken(l *lexer) state {
	return errorf("Unexpected token '%s'", string(l.Peek()))
}

func errorf(msg string, args ...interface{}) state {
	return func(l *lexer) state {
		l.error(fmt.Sprintf(msg, args...))
		return nil
	}
}
