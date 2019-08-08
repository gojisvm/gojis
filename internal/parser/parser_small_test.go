package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallDevFile(t *testing.T) {
	require := require.New(t)

	p := New()
	err := p.ParseFile("testdata/small.js")
	require.NoError(err)
}
