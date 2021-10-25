package min_stack

type node struct {
	Value int
	Prev  *node
}

type MinStack struct {
	stackNodes *node
	minNodes   *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	n := &node{Value: val}
	if s.stackNodes != nil {
		n.Prev = s.stackNodes
	}
	s.stackNodes = n

	if s.minNodes == nil || s.minNodes.Value >= val {
		s.minNodes = &node{Value: val, Prev: s.minNodes}
	}
}

func (s *MinStack) Pop() {
	if s.minNodes != nil && s.minNodes.Value == s.stackNodes.Value {
		s.minNodes = s.minNodes.Prev
	}
	s.stackNodes = s.stackNodes.Prev
}

func (s *MinStack) Top() int {
	return s.stackNodes.Value
}

func (s *MinStack) GetMin() int {
	return s.minNodes.Value
}
