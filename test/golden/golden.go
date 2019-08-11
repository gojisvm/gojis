package golden

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	folder = "testdata"
	suffix = ".golden"
)

// var (
// 	update = flag.Bool("update", false, "Update golden test files.")
// )

// func init() {
// 	if !flag.Parsed() {
// 		flag.Parse()
// 	}
// }

// Equal asserts that the actual bytes are exactly the same as the bytes in the golden
// file with the given name.
//
// The golden file is located under 'testdata/' + name + '.golden'
func Equal(t *testing.T, name string, actual []byte) {
	goldenName := filepath.Join(folder, name+suffix)

	// if *update {
	if false {
		goldenUpdate(t, goldenName, actual)
	} else {
		goldenEqual(t, goldenName, actual)
	}
}

func goldenUpdate(t *testing.T, goldenName string, actual []byte) {
	require := require.New(t)

	err := ioutil.WriteFile(goldenName, actual, 0666)
	require.NoError(err, "Unable to update golden test file %v: %v", goldenName, err)
}

func goldenEqual(t *testing.T, goldenName string, actual []byte) {
	require := require.New(t)

	expected, err := ioutil.ReadFile(filepath.Clean(goldenName))
	require.NoError(err, "Could not open file %v: %v", goldenName, err)
	require.Equal(expected, actual, "Results differ, expected output was not equal to recorded output")
}
