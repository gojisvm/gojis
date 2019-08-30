package conformance

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gojisvm/gojis/tools/test262"
	"github.com/stretchr/testify/require"
)

func TestConformance(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	require := require.New(t)
	err := filepath.Walk(filepath.Join(test262folder, "test"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, "_FIXTURE.js") {
			t.Logf("%v is fixture, skipping the file", path)
			return nil
		}

		t.Run(path, func(t *testing.T) {
			// maybe parallel testing makes sense here?
			executeTest(t, path)
		})

		return nil
	})
	require.NoError(err)
}

func executeTest(t *testing.T, path string) {
	require := require.New(t)

	file, err := os.Open(path)
	require.NoError(err)
	defer file.Close()

	req, err := test262.ParseHeader(file)
	if err != nil {
		t.Skipf("Error while parsing metadata of '%v': %v", path, err)
		return
	}
	_ = req
}
