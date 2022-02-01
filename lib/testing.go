package lib

type Stack struct {
	storage []string
}

func NewStack(elements []string) *Stack {
	return &Stack{storage: elements}
}

func (s *Stack) Pop() string {
	last := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return last
}

func (s *Stack) Push(element string) {
	s.storage = append(s.storage, element)
}

func (s *Stack) ToArray() []string {
	return s.storage
}
