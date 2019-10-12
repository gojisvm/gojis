package lexer

import (
	"github.com/gojisvm/gojis/internal/parser/token"
)

func lexComment(l *Lexer) state {
	if !l.accept(slash) {
		return tokenMismatch(slash)
	}

	switch r := l.peek(); r {
	case '/':
		return lexSingleLineComment
	case '*':
		return lexMultiLineComment
	}
	return tokenMismatch(asterisk, slash)
}

func lexSingleLineComment(l *Lexer) state {
	if !l.accept(slash) {
		return tokenMismatch(slash)
	}

	l.acceptMultiple(singleLineCommentChar)
	l.emit(token.SingleLineComment)
	return lexToken
}

func lexMultiLineComment(l *Lexer) state {
	if !l.accept(asterisk) {
		return tokenMismatch(asterisk)
	}

	for {
		l.acceptMultiple(multiLineNotAsteriskChar)
		if !l.accept(asterisk) {
			return tokenMismatch(asterisk)
		}
		if l.accept(slash) {
			break
		}
	}
	l.emit(token.MultiLineComment)
	return lexToken
}
