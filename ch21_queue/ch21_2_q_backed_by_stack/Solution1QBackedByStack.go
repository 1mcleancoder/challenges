package ch21_2_q_backed_by_stack

import "challenges/ch19_parenthesis_valid"

type SolutionQBackedByStack struct {
	stack ch19_parenthesis_valid.Stack
}

func (s *SolutionQBackedByStack) Enqueue(val any) {
	if s.stack.IsEmpty() {
		s.stack.Push(val)
	} else {
		vals := []any{}
		for !s.stack.IsEmpty() {
			vals = append(vals, s.stack.Pop())
		}
		s.stack.Push(val)
		for i := len(vals) - 1; i >= 0; i-- {
			s.stack.Push(vals[i])
		}
	}
}

func (s *SolutionQBackedByStack) Dequeue() any {
	return s.stack.Pop()
}

func (s *SolutionQBackedByStack) Peek() any {
	return s.stack.Peek()
}

func (s *SolutionQBackedByStack) IsEmpty() bool {
	return s.stack.IsEmpty()
}
