package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

// state is a definition of a state. The state returned will be the one that is
// executed next.
type state func(*Lexer) state

// Below this point, all lexer states are defined.

func lexScript(l *Lexer) state {
	// return lexToken
	return lexComment
}

func lexToken(l *Lexer) state {
	switch r := l.peek(); r {
	default:
		// handle all cases that cannot be expressed in a switch block
	}
	return unexpectedToken
}

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
	return nil
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
	return nil
}
