package merge_multi_level_doubly_list

type DMNode struct {
	Val   any
	Prev  *DMNode
	Next  *DMNode
	Child *DMNode
}

type DoublyLinkedList interface {
	Print(h *DMNode)
	Create(vals []any) *DMNode
	Len(h *DMNode) int
}

// 1, 2, 2, 7, 8, 8, 10, 11, 8, 9, 2, 3, 4, 5, 5, 12, 13, 5, 6
func Print(h *DMNode) []any {
	var c int

	vals := make([]any, 0)
	for h != nil {
		vals = append(vals, h.Val)

		if h.Child != nil {
			vals = append(vals, Print(h.Child)...)
		}
		h = h.Next

		c++
	}

	return vals
}

func PrintWithoutChildNodes(h *DMNode) []any {
	vals := make([]any, 0)
	for h != nil {
		vals = append(vals, h.Val)
		h = h.Next
	}

	return vals
}

// 1 -> 2 -> 3 -> 4 ->  5, -> 6
// 		|				|
//		7 -> 8 -> 9    12 -> 13
//			|
//			10 -> 11
func FlattenWithoutRecursion(h *DMNode) *DMNode {
	if h == nil {
		return nil
	}

	c := h
	for c != nil {
		if c.Child == nil {
			c = c.Next
		} else {
			tail := c.Child // 7
			for tail.Next != nil {
				tail = tail.Next
			}
			// tail = 9
			tail.Next = c.Next // 9 -> 3
			if tail.Next != nil {
				tail.Next.Prev = tail // 9 <- 3
			}

			c.Next = c.Child // 2 -> 7
			c.Next.Prev = c  // 2 <- 7
			c.Child = nil

			c = c.Next
		}
	}

	return h
}

func FlattenUsingRecursion(h *DMNode) *DMNode {
	if h == nil {
		return nil
	}

	// []any{1, 2, 2, 7, 8, 8, 10, 11, 8, 9, 2, 3, 4, 5, 5, 12, 13, 5, 6},
	// []any{1, 2, 7, 8, 10, 11, 9, 3, 4, 5, 12, 13, 6},
	start, _ := flattenChildNodes(h)
	return start
}

func flattenChildNodes(h *DMNode) (*DMNode, *DMNode) {
	var start, end *DMNode
	start = h

	t := h
	for t != nil {
		if t.Child != nil {
			start, end = flattenChildNodes(t.Child)
			t.Child = nil
			n := t.Next
			t.Next = start
			end.Next = n
			t = n
		} else {
			end = t
			t = t.Next
		}
	}

	return h, end
}

//func CreateDoublyLinkedListWithChildNodes(vals []any) *DMNode {
//var h, p, n *DMNode

//vals := []any{1, 2, 3, 4, 5, 6}
//i := 0
//h = &DMNode{Val: vals[i], Prev: p, Next: nil}
//p = h
//
//i++
//n = &DMNode{Val: vals[i], Prev: p, Next: nil}
//p.Next = n
//p = n
//
//i++
//n = &DMNode{Val: vals[i], Prev: p, Next: nil}
//p.Next = n
//p = n
//
//i++
//n = &DMNode{Val: vals[i], Prev: p, Next: nil}
//p.Next = n
//p = n
//
//i++
//n = &DMNode{Val: vals[i], Prev: p, Next: nil}
//p.Next = n
//p = n
//}

// 1, 2, 2, 7, 8, 8, 10, 11, 8, 9, 2, 3, 4, 5, 5, 12, 13, 5, 6
func CreateDoublyLinkedListWithChildNodes(vals []any) *DMNode {
	children, _ := createWithKnownValuesOneListAtATime(0, vals)
	return children
}

// 0, 1, 2, 3, 4, 5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18
// 1, 2, 2, 7, 8, 8, 10, 11, 8, 9,  2,  3,  4,  5,  5, 12, 13,  5,  6
// 1, 2, 2, 7, 8, 9, 2, 3, 4, 5, 5, 12, 13, 5, 6
func createWithKnownValuesOneListAtATime(start any, vals []any) (*DMNode, int) {
	var h, p *DMNode
	var i int
	var last any

	for ; i < len(vals); i++ {
		v := vals[i]
		if v == start {
			i++
			break
		}

		if v == last {
			children, skips := createWithKnownValuesOneListAtATime(v, vals[i+1:])
			i += skips // additional element for the start
			p.Child = children
		} else {
			t := &DMNode{Val: v, Prev: p}
			if p != nil {
				p.Next = t
			}
			if h == nil {
				h = t
			}
			p = t
			last = v
		}
	}

	return h, i
}
