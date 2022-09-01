package ch21_4_q_backed_by_stack

import "challenges/ch19_parenthesis_valid"

type SolutionQBackedByTwoStacksLiterally struct {
	inStack  ch19_parenthesis_valid.Stack
	outStack ch19_parenthesis_valid.Stack
}

func (s *SolutionQBackedByTwoStacksLiterally) Enqueue(val any) {
	s.inStack.Push(val)
}

func (s *SolutionQBackedByTwoStacksLiterally) Dequeue() any {
	if s.outStack.IsEmpty() {
		for !s.inStack.IsEmpty() {
			s.outStack.Push(s.inStack.Pop())
		}
	}
	return s.outStack.Pop()
}

func (s *SolutionQBackedByTwoStacksLiterally) Peek() any {
	if s.outStack.IsEmpty() {
		for !s.inStack.IsEmpty() {
			s.outStack.Push(s.inStack.Pop())
		}
	}

	return s.outStack.Peek()
}

func (s *SolutionQBackedByTwoStacksLiterally) IsEmpty() bool {
	return s.inStack.IsEmpty() && s.outStack.IsEmpty()
}
