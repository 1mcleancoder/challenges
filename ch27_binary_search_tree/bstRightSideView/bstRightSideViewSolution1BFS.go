package bstRightSideView

import bst "challenges/ch27_binary_search_tree"

type BstRightSideViewSolution1BFS struct {
}

func (b BstRightSideViewSolution1BFS) findRightSideViewValues(n *bst.BinaryNode) []int {
	if n == nil {
		return []int{}
	}

	if n.Left == nil && n.Right == nil {
		return []int{n.Data}
	}

	rsv := make([]int, 0)

	q := make([]*bst.BinaryNode, 0)
	q = append(q, n)

	for len(q) > 0 {
		count, lvlLen := 0, len(q)

		for count < lvlLen {
			// remove element from Q
			node := q[0]
			q = q[1:]

			if count == (lvlLen - 1) { // right side view element is the last element
				rsv = append(rsv, node.Data)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
			
			count++
		}
	}

	return rsv
}
