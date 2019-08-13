package golden_test

import (
	"testing"

	"github.com/gojisvm/gojis/tools/golden"
)

func TestEqual(t *testing.T) {
	golden.Equal(t, "Equal", []byte("This is the expected content.\n"))
}
