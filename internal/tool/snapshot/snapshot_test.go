package snapshot_test

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/gojisvm/gojis/internal/tool/snapshot"
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
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
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
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
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
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestEmbedded(t *testing.T) {
	require := require.New(t)

	type T1 struct {
		content string
	}
	type T struct {
		T1
		content string
		another int
	}
	obj := T{
		T1{"World"},
		"Hello",
		24,
	}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
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
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
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
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
	require.NoError(err)

	require.EqualValues(obj, result)
}

func TestPointer(t *testing.T) {
	require := require.New(t)

	type T1 struct {
		content int64
	}
	type T struct {
		another int64
		object  *T1
		content int64
	}
	obj := T{
		24,
		&T1{32},
		64,
	}

	var buf bytes.Buffer
	err := snapshot.Store(&buf, &obj)
	require.NoError(err)

	var result T
	err = snapshot.Load(buf.Bytes(), &result, unsafe.Sizeof(result))
	require.NoError(err)

	require.EqualValues(obj, result)
	require.NotEqual(unsafe.Pointer(obj.object), unsafe.Pointer(result.object), ".object of the loaded object must not point to the .object of the original object, since a pointer should transitively be copied")
}
