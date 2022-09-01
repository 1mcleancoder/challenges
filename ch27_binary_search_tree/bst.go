package ch27_binary_search_tree

import "fmt"

type BinaryNode struct {
	Left  *BinaryNode
	Data  int
	Right *BinaryNode
}

func (n *BinaryNode) Insert(val int) {
	if n == nil {
		return
	} else if val <= n.Data {
		if n.Left == nil {
			n.Left = &BinaryNode{Data: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Data: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *BinaryNode) InsertInSpecifiedOrder(vals []int) {
	q := make([]*BinaryNode, 0, len(vals))
	q = append(q, n)
	for i := 1; i < len(vals); {
		p := q[0] // remove element from Q
		q = q[1:] // adjust Q

		p.Left = &BinaryNode{Data: vals[i]}
		q = append(q, p.Left)
		i++

		if i < len(vals) {
			p.Right = &BinaryNode{Data: vals[i]}
			q = append(q, p.Right)
			i++
		}
	}
}

func (n *BinaryNode) preOrder() string {
	if n == nil {
		return ""
	}

	if n.Left != nil {
		return n.Left.preOrder()
	} else if n.Right != nil {
		return n.Right.preOrder()
	} else {
		return fmt.Sprintf("%v", n.Data)
	}
}

func (n *BinaryNode) inOrder() {
	if n == nil {
		return
	}

	if n.Left != nil {
		n.Left.inOrder()
	}

	n.visit()

	if n.Right != nil {
		n.Right.inOrder()
	}
}

func (n *BinaryNode) visit() {
	if n == nil {
		fmt.Printf("")
	} else {
		fmt.Printf("%v", n.Data)
	}
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) Insert(val int) {
	if t == nil {
		t.root = &BinaryNode{Data: val}
	} else {
		t.root.Insert(val)
	}
}
