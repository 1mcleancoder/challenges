package bstRightSideView

import bst "challenges/ch27_binary_search_tree"

type BstRightSideViewSolution3DFSWithArray struct {
}

func (b BstRightSideViewSolution3DFSWithArray) findRightSideViewValues(n *bst.BinaryNode) []int {
	if n == nil {
		return []int{}
	}

	if n.Left == nil && n.Right == nil {
		return []int{n.Data}
	}

	return visitRightTreeFirstWithLevel(n, make([]int, 0), 1)
}

func visitRightTreeFirstWithLevel(n *bst.BinaryNode, vals []int, lvl int) []int {
	if n == nil {
		return vals
	}

	if lvl > len(vals) { // when there is no element as this level then add it there
		vals = append(vals, n.Data)
	}

	if n.Right != nil {
		vals = visitRightTreeFirstWithLevel(n.Right, vals, lvl+1)
	}

	if n.Left != nil {
		vals = visitRightTreeFirstWithLevel(n.Left, vals, lvl+1)
	}

	return vals
}
