package bstHeightProblem

import bst "challenges/ch27_binary_search_tree"

type MaxHeightSolution1BruteForce struct {
}

func (m MaxHeightSolution1BruteForce) findMaxHeight(n *bst.BinaryNode) int {
	if n == nil {
		return 0
	}

	return m.countMaxHeight(n, 1)
}

func (m MaxHeightSolution1BruteForce) countMaxHeight(n *bst.BinaryNode, count int) int {
	if n == nil {
		return count
	}

	max := count
	if n.Left != nil {
		newCount := m.countMaxHeight(n.Left, count+1)
		if newCount > max {
			max = newCount
		}
	}

	if n.Right != nil {
		newCount := m.countMaxHeight(n.Right, count+1)
		if newCount > max {
			max = newCount
		}
	}

	return max
}
