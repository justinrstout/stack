package fixed

type Stack struct {
	seq  [256]int
	next uint8
}

func (s *Stack) Push(v int) {
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
