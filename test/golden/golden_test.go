package golden_test

import (
	"flag"
	"testing"

	"github.com/gojisvm/gojis/test/golden"
)

func init() {
	if !flag.Parsed() {
		flag.Set("test.v", "false")
	}
}

func TestEqual(t *testing.T) {
	golden.Equal(t, "Equal", []byte("This is the expected content.\n"))
}
