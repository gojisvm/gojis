package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexWhitespace(l *Lexer) state {
	if l.acceptMultiple(whiteSpace) > 0 {
		l.emit(token.Whitespace)
	}
	return lexToken
}

func lexLineTerminator(l *Lexer) state {
	if l.accept(lineTerminator) {
		l.emit(token.LineTerminator)
	}
	return lexToken
}
