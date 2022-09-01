package bstRightSideView

import bst "challenges/ch27_binary_search_tree"

type BstRightSideViewSolution2DFS struct {
}

func (b BstRightSideViewSolution2DFS) findRightSideViewValues(n *bst.BinaryNode) []int {
	if n == nil {
		return []int{}
	}

	if n.Left == nil && n.Right == nil {
		return []int{n.Data}
	}

	valsByLvlMap := make(map[int]int)

	visitRightTreeFirst(n, valsByLvlMap, 1)

	vals := make([]int, 0)
	for i := 1; i <= len(valsByLvlMap); i++ {
		vals = append(vals, valsByLvlMap[i])
	}

	return vals
}

func visitRightTreeFirst(n *bst.BinaryNode, valByLvlMap map[int]int, lvl int) {
	if n == nil {
		return
	}

	if _, ok := valByLvlMap[lvl]; !ok { // when there is no element as this level then add it there
		valByLvlMap[lvl] = n.Data
	}

	if n.Right != nil {
		visitRightTreeFirst(n.Right, valByLvlMap, lvl+1)
	}

	if n.Left != nil {
		visitRightTreeFirst(n.Left, valByLvlMap, lvl+1)
	}

	return
}
