package matcher

import (
	"regexp"
	"strings"
	"unicode"
)

// M is the definition of a character class, which can tell whether a rune is
// part of that character definition or not.
type M interface {
	// Matches determines, whether the given rune is part of the defined
	// character class.
	Matches(rune) bool
	String() string
}

type functionMatcher struct {
	fn   func(rune) bool
	desc string
}

func (m functionMatcher) Matches(r rune) bool { return m.fn(r) }
func (m functionMatcher) String() string      { return m.desc }

type regexpMatcher struct {
	regexp *regexp.Regexp
	desc   string
}

func (m *regexpMatcher) Matches(r rune) bool {
	return m.regexp.MatchString(string(r))
}

func (m *regexpMatcher) String() string { return m.desc }

// Matcher constructors

// New creates a new matcher from a given match function.
func New(desc string, matchFn func(r rune) bool) functionMatcher {
	return functionMatcher{matchFn, desc}
}

// Merge creates a new matcher, that accepts runes that are matched by one or
// more of the given matchers.
func Merge(ms ...M) functionMatcher {
	descs := make([]string, len(ms))
	for i, m := range ms {
		descs[i] = m.String()
	}

	return functionMatcher{
		fn: func(r rune) bool {
			for _, m := range ms {
				if m.Matches(r) {
					return true
				}
			}
			return false
		},
		desc: strings.Join(descs, " or "),
	}
}

// Rune creates a matcher that matches only the given rune.
func Rune(exp rune) functionMatcher {
	return RuneWithDesc("'"+string(exp)+"'", exp)
}

// RuneWithDesc creates a matcher that matches only the given rune. The
// description is the string representation of this matcher. This is useful when
// dealing with whitespace characters.
func RuneWithDesc(desc string, exp rune) functionMatcher {
	return functionMatcher{
		fn:   func(r rune) bool { return r == exp },
		desc: desc,
	}
}

// Negate negates whatever matcher is passed into it, meaning that the new
// matcher will return false when the passed matcher returned true and vice
// versa.
func Negate(m M) functionMatcher {
	return functionMatcher{
		fn:   func(r rune) bool { return !m.Matches(r) },
		desc: "not (" + m.String() + ")",
	}
}

// String creates a matcher that checks whether a rune is part of the given
// string.
func String(s string) functionMatcher {
	return functionMatcher{
		fn:   func(r rune) bool { return strings.IndexRune(s, r) >= 0 },
		desc: s,
	}
}

// Regexp creates a matcher that checks whether a rune is matched by the given
// regular expression.
func Regexp(expr string) *regexpMatcher {
	return &regexpMatcher{
		regexp: regexp.MustCompile(expr),
		desc:   "g/" + expr + "/",
	}
}

// Diff returns a matcher that matches runes, that are matched by shouldMatch,
// but are not matched by butNot.
//
//	Diff(A, B).Matches(r) => r element (A \ B)
func Diff(shouldMatch M, butNot M) functionMatcher {
	return functionMatcher{
		fn: func(r rune) bool {
			return !butNot.Matches(r) && shouldMatch.Matches(r)
		},
		desc: shouldMatch.String() + " but not (" + butNot.String() + ")",
	}
}

// RangeTable creates a matcher that matches runes that are contained in the
// given range table.
func RangeTable(desc string, rt *unicode.RangeTable) functionMatcher {
	return functionMatcher{
		fn: func(r rune) bool {
			return unicode.Is(rt, r)
		},
		desc: desc,
	}
}
