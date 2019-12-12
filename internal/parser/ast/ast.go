package ast

import (
	"fmt"
	"strings"
)

//go:generate go run ../../../tools/aststring .

// Merge merges multiple root nodes to one new root node, that can then be
// evaluated.
func Merge(asts ...*Script) *Script {
	panic("TODO")
}

// PrefixToString converts the given val to a string using fmt's %v and prepends
// the prefix. The result is then returned.
func PrefixToString(val interface{}, prefix string) string {
	v := fmt.Sprintf("%v", val)

	if v[len(v)-1:] == "\n" {
		result := ""
		for _, j := range strings.Split(v[:len(v)-1], "\n") {
			result += prefix + j + "\n"
		}
		return result
	}
	result := ""
	for _, j := range strings.Split(strings.TrimRight(v, "\n"), "\n") {
		result += prefix + j + "\n"
	}
	return result[:len(result)-1]
}
