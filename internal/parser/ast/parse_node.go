package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ParseNode represents a generic node inside the AST.
type ParseNode interface{}

// ToString converts a ParseNode into a human readable string representation.
// The format may change over time.
func ToString(node ParseNode) string {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(node); err != nil {
		return fmt.Sprintf("encode: %v", err)
	}
	return buf.String()
}
