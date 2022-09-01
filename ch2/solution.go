package ch2

// conditions?
// lists are non-empty
// they do not contain any -ve numbers
// each node contains single digit (0, 1, 2, ... 9)
// digits are stored in reverse order
// max: nodes: 100, min: 0

//SolutionAddIntNums solution
type SolutionAddIntNums struct {

}

func (s SolutionAddIntNums) addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := make([]int, 0) // make a new slice for result; it contains sum
	var c int // carry

	var nonEmptyList *ListNode
	n1, n2 := l1, l2
	for { // iterate both the lists
		// add the two numbers
		s := n1.Val + n2.Val + c
		if s < 10 {
			c = 0
		} else {
			s = s - 10
			c = 1
		}

		r = append(r, s)

		n1 = n1.Next
		n2 = n2.Next
		// if anyone of the list is empty then we reached the end,
		// then add the other lists numbers in the end
		if n1 == nil || n2 == nil {
			if n1 != nil {
				nonEmptyList = n1
			} else if n2 != nil {
				nonEmptyList = n2
			}
			break;
		}
	}
	// when the loop ends
	if nonEmptyList != nil {
		for {
			v := nonEmptyList.Val
			if c > 0 {
				if (v + c) < 10 {
					v = v + c
					c = 0
				} else {
					v = (v + c) - 10
					c = 1
				}
			}

			r = append(r, v)

			if nonEmptyList.Next != nil{
				nonEmptyList = nonEmptyList.Next
			} else {
				break
			}
		}
	}

	if c > 0 {
		r = append(r, c)
	}

	var result *ListNode

	// reverse and create a list from slice
	for i:=len(r)-1; i>=0 ; i--{
		result = prepend(r[i], result)
	}

	return result
}

func prepend(v int, l *ListNode) *ListNode {
	return &ListNode{Val: v, Next: l}
}
