package snapshot_test

import (
	"bytes"
	"testing"

	"github.com/gojisvm/gojis/tools/snapshot"
	"github.com/stretchr/testify/require"
)

func TestEmpty(t *testing.T) {
	require := require.New(t)

	type T struct{}
	obj := T{}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(&buf, &result)
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestSimple(t *testing.T) {
	require := require.New(t)

	type T struct {
		content string
	}
	obj := T{"Hello"}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(&buf, &result)
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestDifficult(t *testing.T) {
	require := require.New(t)

	type T struct {
		content string
		number  float64
		another int
	}
	obj := T{"Hello", 2.47e25, 12}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(&buf, &result)
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestComplex(t *testing.T) {
	require := require.New(t)

	type T1 struct {
		content string
	}
	type T struct {
		content string
		object  T1
		another int
	}
	obj := T{
		"Hello",
		T1{"World"},
		24,
	}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(&buf, &result)
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestAlias(t *testing.T) {
	require := require.New(t)

	type T string
	obj := T("Hello")

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(&buf, &result)
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestPointer(t *testing.T) {
	require := require.New(t)

	type T1 struct {
		content string
	}
	type T struct {
		content string
		object  *T1
		another int
	}
	obj := T{
		"Hello",
		&T1{"World"},
		24,
	}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(&buf, &result)
	require.NoError(err)

	require.EqualValues(obj, result)
}
