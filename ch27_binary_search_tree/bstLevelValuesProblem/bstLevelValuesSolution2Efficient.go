package bstLevelValuesProblem

import "challenges/ch27_binary_search_tree"

type BSTLevelValuesSolution2Efficient struct {
}

func (b BSTLevelValuesSolution2Efficient) collectLevelValues(n *ch27_binary_search_tree.BinaryNode) [][]int {
	q := make([]*ch27_binary_search_tree.BinaryNode, 0)
	q = append(q, n)

	lvlVals := make([][]int, 0)

	for len(q) > 0 {
		count, length := 0, len(q)
		vals := make([]int, 0, length)
		for count < length {
			node := q[0]
			q = q[1:] // remove element from Q

			vals = append(vals, node.Data)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			count++
		}

		lvlVals = append(lvlVals, vals)
	}

	return lvlVals
}
