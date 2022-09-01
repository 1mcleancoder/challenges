package ch21_3_q_backed_by_stack

import "challenges/ch19_parenthesis_valid"

type SolutionQBackedByTwoStacks struct {
	stack ch19_parenthesis_valid.Stack
}

func (s *SolutionQBackedByTwoStacks) Enqueue(val any) {
	if s.stack.IsEmpty() {
		s.stack.Push(val)
	} else {
		tStack := ch19_parenthesis_valid.Stack{}
		for !s.stack.IsEmpty() {
			tStack.Push(s.stack.Pop())
		}
		s.stack.Push(val)
		for !tStack.IsEmpty() {
			s.stack.Push(tStack.Pop())
		}
	}
}

func (s *SolutionQBackedByTwoStacks) Dequeue() any {
	return s.stack.Pop()
}

func (s *SolutionQBackedByTwoStacks) Peek() any {
	return s.stack.Peek()
}

func (s *SolutionQBackedByTwoStacks) IsEmpty() bool {
	return s.stack.IsEmpty()
}
