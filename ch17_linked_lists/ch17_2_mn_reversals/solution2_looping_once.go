package ch17_2_mn_reversals

import (
	list "challenges/ch17_linked_lists"
)

type solution2_looping_once struct {
}

// 1, 2, 3, 4, 5, 6, 7			3, 5
// 1, 2, 5, 4, 3, 6, 7

// start = 2, end = 3
// 3 -> null
// 4 -> 3 -> null
// revList = 5 -> 4 -> 3 -> null
// start.Next = revList => 1, 2, 5, 4, 3, null
// currNode = 6
// end.Next = currNode = 6
// result: 1 -> 2 -> 5 -> 4 -> 3 -> 6 -> 7

// {[]int{1, 2, 3, 4, 5, 6, 7}, 3, 5, "1,2,5,4,3,6,7"},
func (s2 solution2_looping_once) Reverse(head *list.Node, m, n int) *list.Node {
	if head == nil || m == n {
		return head
	}

	if m <= 1 && n <= 1 {
		return head
	}

	var start, revListEnd *list.Node
	var pos int = 1

	currNode := head
	for pos < m {
		start = currNode

		currNode = currNode.Next
		pos++
	}

	revListEnd = currNode.Next

	var revListSoFar, next *list.Node

	for currNode != nil && pos <= n {
		next = currNode.Next
		currNode.Next = revListSoFar
		revListSoFar = currNode

		currNode = next

		pos++
	}

	if revListEnd != nil {
		revListEnd.Next = next
	}

	if start != nil {
		// attaching the reversed list to the start node
		start.Next = revListSoFar

		return head // head of the list is unchanged
	} else {
		return revListSoFar // new list is the head now
	}
}
