package ch17_3_cycle_detection

import list "challenges/ch17_linked_lists"

type solution1_bf struct {
}

func (s1 *solution1_bf) detectCycle(h *list.Node) *list.Node {
	if h == nil {
		return h
	}

	var cycle *list.Node
	var nodeByValue = make(map[any]*list.Node)

	for h != nil {
		if n, ok := nodeByValue[h.Val]; ok {
			cycle = n
			break
		}

		nodeByValue[h.Val] = h

		h = h.Next
	}

	return cycle
}
