package sequence

type Stack struct {
	items []int
	top   int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
		top:   -1,
	}
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
	s.top++
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.top--
}

func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *Stack) Top() int {
	return s.items[s.top]
}
