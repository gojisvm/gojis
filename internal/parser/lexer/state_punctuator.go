package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexPunctuator(l *Lexer) state {
	for _, punctuator := range Punctuators {
		if l.acceptWord(punctuator) {
			// punctuator was matched
			l.emit(token.Punctuator)
			return lexToken
		}
	}
	return errorf("Unexpected token '%s', expected punctuator", string(l.peek()))
}

func lexDivPunctuator(l *Lexer) state {
	for _, punctuator := range DivPunctuators {
		if l.acceptWord(punctuator) {
			// div punctuator was matched
			l.emit(token.Punctuator)
			return lexToken
		}
	}
	return errorf("Unexpected token '%s', expected div punctuator", string(l.peek()))
}

func lexRightBracePunctuator(l *Lexer) state {
	if !l.accept(RightBracePunctuator) {
		return tokenMismatch(RightBracePunctuator)
	}
	return lexToken
}
