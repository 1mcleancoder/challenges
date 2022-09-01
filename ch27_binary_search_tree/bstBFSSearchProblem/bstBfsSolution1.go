package bstBFSSearchProblem

import "challenges/ch27_binary_search_tree"

type BstBFSSearchSolution1 struct {
}

func (b BstBFSSearchSolution1) bfs(n *ch27_binary_search_tree.BinaryNode) []int {
	var q = make([]*ch27_binary_search_tree.BinaryNode, 0)
	q = append(q, n)

	allData := make([]int, 0)

	for len(q) > 0 {
		v := q[0]

		allData = append(allData, v.Data)

		q = q[1:] // remove the first element

		if v.Left != nil {
			q = append(q, v.Left)
		}
		if v.Right != nil {
			q = append(q, v.Right)
		}
	}

	return allData
}
