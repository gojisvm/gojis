package lexer

import (
	"strings"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
)

func unexpectedWord(expected ...string) state {
	return func(l *Lexer) state {
		row, col := rowAndCol(l.pos, l)
		return errorf("Unexpected token at pos %v:%v, expected one of ['%v'], but next rune was '%s'", row, col, strings.Join(expected, "', '"), string(l.peek()))
	}
}

func tokenMismatch(expected ...matcher.M) state {
	var descs []string
	for _, m := range expected {
		descs = append(descs, m.String())
	}

	return func(l *Lexer) state {
		row, col := rowAndCol(l.pos, l)
		return errorf("Unexpected token at pos %v:%v, got '%s' but expected %s", row, col, string(l.peek()), strings.Join(descs, " or "))
	}
}

func unexpectedToken(l *Lexer) state {
	row, col := rowAndCol(l.pos, l)
	return errorf("Unexpected token '%s' at pos %v:%v", string(l.peek()), row, col)
}

func errorf(msg string, args ...interface{}) state {
	return func(l *Lexer) state {
		l.fatalf(msg, args...)
		return nil
	}
}

func rowAndCol(offset int, l *Lexer) (row, col int) {
	row = 1 // lines start at 1
	col = 1 // cols start at 1

	for i, r := range l.input {
		if i >= offset {
			break
		}
		if r == '\t' {
			col += 4
		} else {
			col++
		}
		if lineTerminator.Matches(r) {
			col = 1
			row++
		}
	}
	return
}
