package lexer

import (
	"github.com/gojisvm/gojis/internal/parser/token"
)

func lexComment(l *Lexer) state {
	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	switch r := l.peek(); r {
	case '/':
		return lexSingleLineComment
	case '*':
		return lexMultiLineComment
	}
	return tokenMismatch(Asterisk, Slash)
}

func lexSingleLineComment(l *Lexer) state {
	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	l.acceptMultiple(SingleLineCommentChar)
	l.emit(token.SingleLineComment)
	return lexToken
}

func lexMultiLineComment(l *Lexer) state {
	if !l.accept(Asterisk) {
		return tokenMismatch(Asterisk)
	}

	for {
		l.acceptMultiple(MultiLineNotAsteriskChar)
		if !l.accept(Asterisk) {
			return tokenMismatch(Asterisk)
		}
		if l.accept(Slash) {
			break
		}
	}
	l.emit(token.MultiLineComment)
	return lexToken
}
