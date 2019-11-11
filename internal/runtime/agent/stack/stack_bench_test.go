package stack

import (
	"crypto/rand"
	"testing"
)

var r interface{}

func benchmarkStack(b *testing.B, s Stack) {
	b.Run("Push 3B", benchmarkPush(s))
	b.Run("Push 8KB", benchmarkPushLarge(s))
	b.Run("Push Pop", benchmarkPushPop(s))
	b.Run("Peek", benchmarkPeek(s))
}

func benchmarkPush(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s.Push("foo")
		}
	}
}

func benchmarkPushLarge(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		var data [8192]byte
		_, _ = rand.Read(data[:])

		type t struct {
			data [8192]byte
		}
		newT := func() t {
			t := new(t)
			t.data = data
			return *t
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s.Push(newT())
		}
	}
}

func benchmarkPushPop(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		var elem interface{}
		val := "foo"

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s.Push(val)
			elem = s.Pop()
		}

		r = elem
	}
}

func benchmarkPeek(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		var elem interface{}

		s.Push("foo")

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			elem = s.Peek()
		}

		r = elem
	}
}
