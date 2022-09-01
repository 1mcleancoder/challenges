package ch17_linked_lists

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) Print() int {
	var c int
	t := n
	for t != nil {
		fmt.Printf("%v ", t.Val)
		t = t.Next

		c++
	}

	fmt.Println()

	return c
}

func (n *Node) String() string {
	if n == nil || n.Val == 0 {
		return ""
	}
	var vals = make([]string, 0)

	t := n
	for t != nil {
		vals = append(vals, strconv.Itoa(t.Val))
		t = t.Next
	}

	return strings.Join(vals, ",")
}

func (n *Node) Create(vals []int) *Node {
	if len(vals) == 0 {
		return &Node{}
	}

	h := &Node{Val: vals[0]}
	if len(vals) == 1 {
		return h
	}

	t := h
	for i := 1; i < len(vals); i++ {
		n := &Node{Val: vals[i]}
		t.Next = n
		t = n
	}

	return h
}

func (n *Node) CreateListWithCycleWhenSimilarValues(vals []int) *Node {
	if len(vals) == 0 {
		return &Node{}
	}

	var h, prev *Node

	var nodeByValMap = make(map[any]*Node)

	for _, v := range vals {
		if node, ok := nodeByValMap[v]; ok {
			prev.Next = node // cycle made
			break
		}

		t := &Node{Val: v}
		nodeByValMap[v] = t
		if h == nil {
			h = t
		} else {
			prev.Next = t
		}

		prev = t
	}

	return h
}

func (n *Node) Length() int {
	var c int

	t := n
	for t != nil {
		c++

		t = t.Next
	}

	return c
}
