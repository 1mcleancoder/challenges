package ch17_3_cycle_detection

import list "challenges/ch17_linked_lists"

type solution2EfficientFloydsTortoiseHareAlgo struct {
}

func (s2 *solution2EfficientFloydsTortoiseHareAlgo) detectCycle(h *list.Node) *list.Node {
	if h == nil {
		return nil
	}

	// cycle detection
	slow, fast := h, h
	for {
		if slow == nil {
			return nil // no cycles exist
		}
		slow = slow.Next

		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return nil // no cycles exist
		}

		fast = fast.Next.Next // jump two pointers

		if slow == fast {
			break
		}
	}

	fwd, bck := h, fast // starting from head, move one position at a time for both slow and fast pointers say forward and backward pointers
	for fwd != bck {
		fwd = fwd.Next
		bck = bck.Next
	}

	return fwd
}
