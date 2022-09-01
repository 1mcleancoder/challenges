package ch17_1_reverse_list

import (
	list "challenges/ch17_linked_lists"
)

type Solution1ReverseList struct {
}

func (s1 *Solution1ReverseList) Reverse(listNodes *list.Node) *list.Node {
	// 1 -> 2 -> 3 -> nil
	// 3 -> 2 -> 1 -> nil

	// steps:
	// 1 -> nil
	// 2 -> 1 -> nil
	// 3 -> 2 -> 1 -> nil
	// return 3's pointer

	// create two pointers say p (prev) and t (temp)
	// here the idea is that we loop the nodes such that
	// at every node we create a new node say n and set n.Next = p (initially nil)
	// the say p = n
	// store the newly created node's pointer in prev such that the next node can point to it.
	// new Node with 1 say newNode
	// In the end, n will have the reversed's list's head node and we can return the same.

	if listNodes == nil || listNodes.Next == nil {
		return listNodes
	}

	var p *list.Node = nil

	t := listNodes
	for t != nil {
		n := &list.Node{t.Val, p}
		p = n
		t = t.Next
	}

	return p
}
