package fixed

import "testing"

func TestNew(t *testing.T) {
	if stack := new(Stack); !stack.Empty() {
		t.Fatalf("New stack not empty.")
	}
}

func TestStack_PushPop(t *testing.T) {
	vals := make([]int, 100)
	for i := 0; i < 100; i++ {
		vals[i] = i - 50
	}
	stack := new(Stack)

	for i := 0; i < 100; i++ {
		stack.Push(vals[i])
		if stack.Empty() {
			t.Fatalf("Pushed stack is empty.")
		}
	}

	for i := 99; i >= 0; i-- {
		if pop := stack.Pop(); pop != vals[i] {
			t.Fatalf("Popped unexpected value: %d", pop)
		}
	}

	if !stack.Empty() {
		t.Fatalf("Popped stack not empty.")
	}
}

func TestStack_Peek(t *testing.T) {
	val := 100
	stack := new(Stack)

	stack.Push(val)
	if peek := stack.Peek(); peek != val {
		t.Fatalf("Peeked unexpected value: %d", peek)
	}
}
