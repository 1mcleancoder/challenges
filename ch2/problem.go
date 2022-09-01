package ch2

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

type AddIntLists interface {
	addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode
}
