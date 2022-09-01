package bstNumNodesInCompleteTree

import (
	bst "challenges/ch27_binary_search_tree"
	"math"
)

type NumNodesInCompleteTreeSolution1DFS struct {
}

func (t NumNodesInCompleteTreeSolution1DFS) findNumNodesInCompleteTree(n *bst.BinaryNode) int {
	if n == nil {
		return 0
	}

	maxLvl := t.findMaxLevel(n, 1)
	numNodes := int(math.Pow(float64(2), float64(maxLvl-1)) - float64(1))
	return t.countNodesAtLevel(n, 1, maxLvl, numNodes)
}

func (t NumNodesInCompleteTreeSolution1DFS) findMaxLevel(n *bst.BinaryNode, lvl int) int {
	maxLvl := lvl

	if n == nil {
		return lvl
	}

	if n.Left != nil {
		newLvl := t.findMaxLevel(n.Left, lvl+1)
		if maxLvl < newLvl {
			maxLvl = newLvl
		}
	}

	return maxLvl
}

func (t NumNodesInCompleteTreeSolution1DFS) countNodesAtLevel(n *bst.BinaryNode, lvl, maxLvl, count int) int {
	if lvl == maxLvl-1 { // only nodes at the 2nd to last level are important
		if n.Left != nil {
			count++
		}

		if n.Right != nil {
			count++
		}
	} else {
		if n.Left != nil {
			count = t.countNodesAtLevel(n.Left, lvl+1, maxLvl, count)
		}

		if n.Right != nil {
			count = t.countNodesAtLevel(n.Right, lvl+1, maxLvl, count)
		}
	}

	return count
}
