package parser

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/tools/golden"
	"github.com/stretchr/testify/require"
)

func TestGolden(t *testing.T) {
	require := require.New(t)

	files, err := filepath.Glob("testdata/*.js")
	require.NoError(err)

	for _, file := range files {
		t.Run(file, testGoldenSingleFile(file))
	}
}

func testGoldenSingleFile(file string) func(*testing.T) {
	return func(t *testing.T) {
		require := require.New(t)

		data, err := ioutil.ReadFile(file)
		require.NoError(err)

		p := New()
		// reading the file and then creating a byte reader removes the possibility that the parser fails at reading the file
		tree, err := p.Parse(file, bytes.NewReader(data))
		require.NoError(err)

		astString := ast.ToString(tree)
		golden.Equal(t, filepath.Base(file), []byte(astString))
	}
}

// ============================================
// == as of this point, test262 parser tests ==
// ============================================

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
