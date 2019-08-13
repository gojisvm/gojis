package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CollectingErrorListener is an ANTLR error listener that collects errors.
// Collected errors can be fetched with a call to Errors().
type CollectingErrorListener struct {
	errs []error
}

// NewCollectingErrorListener creates a new ready to use CollectingErrorListener.
func NewCollectingErrorListener() *CollectingErrorListener {
	l := new(CollectingErrorListener)
	return l
}

// SyntaxError see ErrorListener for documentation.
func (l *CollectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Append(fmt.Errorf("Syntax error at line %v column %v: %v", line, column, msg))
}

// ReportAmbiguity see ErrorListener for documentation.
func (l *CollectingErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

// ReportAttemptingFullContext see ErrorListener for documentation.
func (l *CollectingErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

// ReportContextSensitivity see ErrorListener for documentation.
func (l *CollectingErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

// Append adds a new error to the CollectingErrorListener.
func (l *CollectingErrorListener) Append(errs ...error) {
	l.errs = append(l.errs, errs...)
}

// Errors returns all collected errors and a flag, indicating whether there were any errors collected at all.
func (l *CollectingErrorListener) Errors() ([]error, bool) {
	return l.errs, l.errs != nil && len(l.errs) > 0
}
