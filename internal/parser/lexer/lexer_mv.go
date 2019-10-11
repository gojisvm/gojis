package lexer

import (
	"fmt"
	"strconv"
)

// mvHex computes the mathematical value of the last n read runes as defined by
// the specification. The runes have to be hex digits, otherwise this will
// panic.
func (l *Lexer) mvHex(n uint) int64 {
	lastN, ok := l.lastN(n)
	if !ok {
		panic(Error{
			msg: fmt.Sprintf("Unable to access last %v accepted hex digits", n),
		})
	}
	mv, err := strconv.ParseInt(string(lastN), 16, 64)
	if err != nil {
		panic(Error{
			msg: fmt.Sprintf("Error computing MV of hex digits: %v", err),
		})
	}
	return mv
}
