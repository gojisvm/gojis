package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const basePath = "test262-parser-tests"

func TestEarly(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	require := require.New(t)
	err := filepath.Walk(filepath.Join(basePath, "early"), genWalkerFunc(t, false))
	require.NoError(err)
}

func TestFail(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	require := require.New(t)
	err := filepath.Walk(filepath.Join(basePath, "fail"), genWalkerFunc(t, false))
	require.NoError(err)
}

func TestPass(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	require := require.New(t)
	err := filepath.Walk(filepath.Join(basePath, "pass"), genWalkerFunc(t, true))
	require.NoError(err)
}

func TestPassExplicit(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	require := require.New(t)
	err := filepath.Walk(filepath.Join(basePath, "pass-explicit"), genWalkerFunc(t, true))
	require.NoError(err)
}

func genWalkerFunc(t *testing.T, successfulParse bool) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		parse(t, path, successfulParse)

		return nil
	}
}

func parse(t *testing.T, path string, successfulParse bool) {
	t.Run(path, func(t *testing.T) {
		t.SkipNow()

		require := require.New(t)

		p := New()
		_, err := p.ParseFile(path)

		if successfulParse {
			require.NoError(err)
		} else {
			require.Error(err)
		}
	})
}
