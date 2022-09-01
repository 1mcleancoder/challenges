package ch18_doubly_linked_lists

import (
	"fmt"
)

type DoublyLinkedList interface {
	Print(h *DNode)
	Create() *DNode
	Len(h *DNode) int
}

type DNode struct {
	Val  any
	Prev *DNode
	Next *DNode
}

func Print(h *DNode) int {
	var c int

	var s string
	for h != nil {
		var p, n any
		if h.Prev != nil || h.Next != nil {
			if h.Prev != nil {
				p = h.Prev.Val
			}

			if h.Next != nil {
				n = h.Next.Val
			}
		}
		if n != nil {
			s += fmt.Sprintf("%v (%v, %v) -> ", h.Val, p, n)
		} else {
			s += fmt.Sprintf("%v (%v, %v)", h.Val, p, n)
		}
		h = h.Next

		c++
	}

	fmt.Println(s)

	return c
}

func Create(vals []any) *DNode {
	var h, p *DNode

	for _, v := range vals {
		t := &DNode{Val: v, Prev: p}
		if p != nil {
			p.Next = t
		}
		if h == nil {
			h = t
		}

		p = t
	}

	return h
}
