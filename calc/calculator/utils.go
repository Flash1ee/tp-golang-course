package calculator

type Node interface{}

type Stack struct {
	data []Node
}

func NewStack() *Stack {
	return &Stack{data: []Node{}}
}

func (s *Stack) Push(elem Node) {
	if s.data == nil {
		s.data = []Node{}
	}
	s.data = append(s.data, elem)
}

func (s *Stack) Pop() Node {
	if s.isEmpty() {
		return nil
	}
	elem := s.data[len(s.data)-1:]
	s.data = s.data[0 : len(s.data)-1]

	return elem[0]
}

func (s *Stack) Peek() Node {
	if s.isEmpty() {
		return nil
	}
	elem := s.data[len(s.data)-1:]
	return elem[0]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}
