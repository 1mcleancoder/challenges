package ch17_1_reverse_list

import (
	list "challenges/ch17_linked_lists"
)

type Solution2ReverseList struct {
}

func (s2 *Solution2ReverseList) Reverse(listNodes *list.Node) *list.Node {
	// 1 -> 2 -> 3 -> nil
	// 3 -> 2 -> 1 -> nil

	// steps:
	// 1 -> nil, prev = 1, curr = 2
	// 		next = curr.Next
	//		curr.Next = prev => 2 -> 1 -> nil
	//		prev = curr => prev = 2
	//		curr = next => curr = 3
	// 2 -> 1 -> nil
	// 3 -> 2 -> 1 -> nil
	// return 3's pointer

	// create two pointers say prev (prev) and cNode (temp)
	// here the idea is that we loop the nodes such that
	// at every node we create a new node say n and set n.Next = prev (initially nil)
	// the say prev = n
	// store the newly created node's pointer in prev such that the next node can point to it.
	// new Node with 1 say newNode
	// In the end, n will have the reversed's list's head node and we can return the same.

	if listNodes == nil || listNodes.Next == nil {
		return listNodes
	}

	var prev *list.Node = nil

	cNode := listNodes
	for cNode != nil {
		next := cNode.Next
		cNode.Next = prev
		prev = cNode
		cNode = next
	}

	return prev
}
