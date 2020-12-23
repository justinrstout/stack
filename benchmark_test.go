package stack

import (
	"container/list"
	"github.com/justinrstout/stack/fixed"
	"github.com/justinrstout/stack/ints"
	"testing"
)

func Benchmark_LargePush(b *testing.B) {
	b.Run("container/list", func(b *testing.B) {
		lst := list.New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lst.PushFront(i)
		}
	})
	b.Run("stack/ints", func(b *testing.B) {
		st := ints.New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			st.Push(i)
		}
	})
	b.Run("stack/fixed", func(b *testing.B) {
		st := new(fixed.Stack)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			st.Push(i)
		}
	})
	b.Run("handwritten", func(b *testing.B) {
		st := make([]int, b.N)
		head := 0
		push := func(i int) {
			st[head] = i
			head++
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			push(i)
		}
	})
}

func Benchmark_SmallPushPop(b *testing.B) {
	b.Run("container/list", func(b *testing.B) {
		lst := list.New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lst.PushFront(i)
			lst.PushFront(i + 1)
			lst.PushFront(i + 2)
			lst.Remove(lst.Front())
			lst.Remove(lst.Front())
			lst.Remove(lst.Front())
		}
	})
	b.Run("stack/ints", func(b *testing.B) {
		st := ints.New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			st.Push(i)
			st.Push(i + 1)
			st.Push(i + 2)
			st.Pop()
			st.Pop()
			st.Pop()
		}
	})
	b.Run("stack/fixed", func(b *testing.B) {
		st := new(fixed.Stack)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			st.Push(i)
			st.Push(i + 1)
			st.Push(i + 2)
			st.Pop()
			st.Pop()
			st.Pop()
		}
	})
	b.Run("handwritten", func(b *testing.B) {
		st := make([]int, b.N+2)
		head := 0
		push := func(i int) {
			st[head] = i
			head++
		}
		pop := func() int {
			head--
			return st[head]
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			push(i)
			push(i + 1)
			push(i + 2)
			pop()
			pop()
			pop()
		}
	})
}
