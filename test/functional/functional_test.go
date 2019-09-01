package functional

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionality(t *testing.T) {
	dirs := getAllTestDirectories(t)
	for _, dir := range dirs {
		t.Run(dir, func(t *testing.T) {
			executeTestDir(t, dir)
		})
	}
}

func executeTestDir(t *testing.T, dir string) {
	t.Skipf("Skipping due to missing implementation")
}

func getAllTestDirectories(t *testing.T) []string {
	require := require.New(t)

	matches, err := filepath.Glob("testdata/*")
	require.NoError(err)
	return matches
}
