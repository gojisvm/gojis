// Package matcher implements matchers that efficiently let you determine
// whether a rune is part of a previously defined rune set.
package matcher

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

// M is the definition of a character class, which can tell whether a rune is
// part of that character definition or not.
type M struct {
	rt   *unicode.RangeTable
	desc string
}

// Matches describes, whether the given rune is contained in this matcher's
// range table.
func (m M) Matches(r rune) bool { return unicode.Is(m.rt, r) }

// String returns a human readable string description of the runes that this
// matcher matches.
func (m M) String() string { return m.desc }

// Matcher constructors

// New creates a new matcher from a given match function.
func New(desc string, rt *unicode.RangeTable) M {
	return M{rt, desc}
}

// Merge creates a new matcher, that accepts runes that are matched by one or
// more of the given matchers.
func Merge(ms ...M) M {
	var rtms []*unicode.RangeTable
	descs := make([]string, len(ms))

	for i, m := range ms {
		descs[i] = m.String()
		rtms = append(rtms, m.rt)
	}

	return M{
		rt:   rangetable.Merge(rtms...),
		desc: strings.Join(descs, " or "),
	}
}

// RuneWithDesc creates a matcher that matches only the given rune. The
// description is the string representation of this matcher. This is useful when
// dealing with whitespace characters.
func RuneWithDesc(desc string, exp rune) M {
	return M{
		rt:   rangetable.New(exp),
		desc: desc,
	}
}

// String creates a matcher that checks whether a rune is part of the given
// string.
func String(s string) M {
	return M{
		rt:   rangetable.New([]rune(s)...),
		desc: "one of '" + s + "'",
	}
}

// Diff returns a matcher that matches runes, that are matched by shouldMatch,
// but are not matched by butNot.
//
//	Diff(A, B).Matches(r) => r element (A \ B)
func Diff(shouldMatch M, butNot M) M {
	r16s := []unicode.Range16{}
	r32s := []unicode.Range32{}
	// r16
	for _, rng := range shouldMatch.rt.R16 {
		for r := rng.Lo; r <= rng.Hi; r += rng.Stride {
			if !butNot.Matches(rune(r)) {
				r16s = append(r16s, unicode.Range16{r, r, 1})
			}
		}
	}

	// r32
	for _, rng := range shouldMatch.rt.R32 {
		for r := rng.Lo; r <= rng.Hi; r += rng.Stride {
			if !butNot.Matches(rune(r)) {
				r32s = append(r32s, unicode.Range32{r, r, 1})
			}
		}
	}

	return M{
		rt: rangetable.Merge(&unicode.RangeTable{
			R16: r16s,
			R32: r32s,
		}), // this optimizes the generated range table
		desc: shouldMatch.String() + " but not (" + butNot.String() + ")",
	}
}

// RangeTable creates a matcher that matches runes that are contained in the
// given range table.
func RangeTable(desc string, rt *unicode.RangeTable) M {
	return M{
		rt:   rt,
		desc: desc,
	}
}
