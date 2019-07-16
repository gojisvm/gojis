package golden_test

import (
	"testing"

	"github.com/gojisvm/gojis/test/golden"
)

func TestEqual(t *testing.T) {
	golden.Equal(t, "Equal", []byte("This is the expected content.\n"))
}
