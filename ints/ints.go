package ints

type Stack struct {
	seq  []int
	next int
}

func New() *Stack {
	stack := new(Stack)
	stack.seq = make([]int, 64)
	return stack
}

func (s *Stack) Push(v int) {
	if s.next == len(s.seq) {
		tmp := make([]int, s.next + s.next>>1)
		copy(tmp, s.seq)
		s.seq = tmp
	}
	s.seq[s.next] = v
	s.next++
}

func (s *Stack) Pop() int {
	s.next--
	return s.seq[s.next]
}

func (s *Stack) Peek() int {
	return s.seq[s.next-1]
}

func (s *Stack) Empty() bool {
	return s.next == 0
}
