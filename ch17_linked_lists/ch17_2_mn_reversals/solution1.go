package ch17_2_mn_reversals

import (
	list "challenges/ch17_linked_lists"
)

type solution1 struct {
}

// {[]int{1, 2, 3, 4, 5}, 2, 4, "1,4,3,2,5"},
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 and 2, 5
// 1 -> 5 -> 4 -> 3 -> 2 -> 6

// 1 -> 2 -> 3 -> 4 -> 5 -> 6 and 2, 5
// 1 -> 5 -> 3 -> 4 -> 5 -> 6 and 2, 5
// 1 -> 5 -> 4 -> 3 -> 5 -> 6 and 2, 5
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 and 2, 5
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 and 2, 5
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 and 2, 5
/*
	iterate to the m position, note down previous position,
	keep iterating to the n position, note down the to position
	prevToFromNode = 1
	toNode = 5
	lastNodes = toNode.Next() -- 6

	pNode = prevToFromNode (1)
	while pNode.next != lastNode
		fwdNode = pNode.Next (2)
		pNode.Next = toNode (5)
		pNode = prevNode.Next

*/
func (s1 solution1) Reverse(h *list.Node, m, n int) *list.Node {
	if m <= 1 && n <= 1 {
		return h
	}

	// iterate the list and save the beforeM node and after n node
	// reverse the list till n
	bM, bN := findBeforeMAndNNodes(h, m, n)
	//fmt.Printf("bM=%v, bN=%v", bM, bN)

	var lastNode *list.Node
	if bN != nil {
		lastNode = bN.Next
	}

	if m <= 1 {
		return reverseBetweenNodes(bM, lastNode)
	} else {
		revList := reverseBetweenNodes(bM.Next, lastNode)
		bM.Next = revList
		return bM
	}
}

func findBeforeMAndNNodes(h *list.Node, m, n int) (*list.Node, *list.Node) {
	var bM, aN *list.Node
	var c int

	t := h
	pN := h
	for t != nil {
		c++

		if m == c && bM == nil {
			bM = pN
		}

		if n == c {
			aN = t
			break
		}

		t = t.Next
	}

	return bM, aN
}

// 1 -> 2 -> 3 -> 4 -> 5 , 2, 4
// 2 -> 3 -> 4 -> 5
// 1 -> 5 -> 4 -> 3 -> 2 -> 1 -> 6
func reverseBetweenNodes(m, last *list.Node) *list.Node {
	var c *list.Node

	c = m
	prev := last

	for c != nil && c != last {
		t := c.Next
		c.Next = prev
		prev = c
		c = t
	}

	return prev
}
