package ints

import (
	"container/list"
	"testing"
)

func BenchmarkList_Push(b *testing.B) {
	lst := list.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lst.PushFront(i)
	}
}

func BenchmarkStack_Push(b *testing.B) {
	st := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		st.Push(i)
	}
}

func BenchmarkArray_Push(b *testing.B) {
	st := make([]int, b.N)
	head := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		st[head] = i
		head++
	}
}
