package snapshot_test

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/gojisvm/gojis/internal/tool/snapshot"
)

var result interface{}

func BenchmarkCreateObjectFromSnapshot(b *testing.B) {
	var r interface{}

	var buf bytes.Buffer
	_ = snapshot.Store(&buf, NewContainer())

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c, err := NewContainerFromSnapshot(buf.Bytes())
		if err != nil {
			b.Fail()
		}

		r = c
	}

	result = r
}

func BenchmarkCreateObject(b *testing.B) {
	var r interface{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := NewContainer()
		r = c
	}

	result = r
}

func NewContainerFromSnapshot(data []byte) (*Container, error) {
	c := new(Container)
	err := snapshot.Load(data, c, unsafe.Sizeof(Container{}))
	return c, err
}

func NewContainer() *Container {
	c := new(Container)
	c.value = fib(20)
	return c
}

type Container struct {
	value int64
}

func fib(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
