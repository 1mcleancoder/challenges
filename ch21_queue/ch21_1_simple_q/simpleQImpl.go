package ch21_1_simple_q

type Solution1QUsingSlices struct {
	vals []any
}

func (s1 *Solution1QUsingSlices) Enqueue(v any) {
	s1.vals = append(s1.vals, v)
}

func (s1 *Solution1QUsingSlices) Dequeue() any {
	if s1.IsEmpty() {
		return nil
	}
	d := s1.vals[0]
	s1.vals = s1.vals[1:]

	return d
}

func (s1 *Solution1QUsingSlices) Peek() any {
	if s1.IsEmpty() {
		return nil
	}
	return s1.vals[0]
}

func (s1 *Solution1QUsingSlices) IsEmpty() bool {
	if len(s1.vals) == 0 {
		return true
	} else {
		return false
	}
}
