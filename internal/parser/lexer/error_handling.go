package lexer

import (
	"fmt"
)

func unexpectedToken(l *Lexer) state {
	return errorf("Unexpected token '%s'", string(l.Peek()))
}

func errorf(msg string, args ...interface{}) state {
	return func(l *Lexer) state {
		l.error(fmt.Sprintf(msg, args...))
		return nil
	}
}
