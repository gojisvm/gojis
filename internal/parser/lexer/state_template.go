package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexTemplate(l *Lexer) state {
	panic("TODO")

	l.emit(token.Template)
	return lexToken
}
