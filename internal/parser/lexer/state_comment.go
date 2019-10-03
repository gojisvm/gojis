package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexComment(l *Lexer) state {
	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	switch r := l.next(); r {
	case '/':
		return lexSingleLineComment
	case '*':
		return lexMultiLineComment
	}
	l.unread()
	return tokenMismatch(Asterisk, Slash)
}

func lexSingleLineComment(l *Lexer) state {
	l.acceptMultiple(SingleLineCommentChar)
	l.emit(token.SingleLineComment)
	return lexToken
}

func lexMultiLineComment(l *Lexer) state {
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
