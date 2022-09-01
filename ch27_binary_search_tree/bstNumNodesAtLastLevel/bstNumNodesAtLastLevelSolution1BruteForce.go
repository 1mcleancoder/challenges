package bstNumNodesAtLastLevel

import (
	bst "challenges/ch27_binary_search_tree"
	"math"
)

type NumNodesAtLastLevelSolution1 struct {
}

func (n NumNodesAtLastLevelSolution1) findNumNodesAtLastLevel(b *bst.BinaryNode) int {
	if b == nil {
		return 0
	}

	mxLvl := findMaxLevel(b, 1)
	return findNodesAtMaxLevel(b, mxLvl, 1)
}

func findMaxLevel(b *bst.BinaryNode, lvl int) int {
	maxLvl := lvl
	if b.Left != nil {
		newLvl := findMaxLevel(b.Left, lvl+1)
		maxLvl = int(math.Max(float64(newLvl), float64(lvl)))
	}

	return maxLvl
}

func findNodesAtMaxLevel(b *bst.BinaryNode, maxLvl, lvl int) int {
	numNodes := 0
	lvlBeforeMax := maxLvl - 1
	if lvl < lvlBeforeMax {
		if b.Left != nil {
			numNodes += findNodesAtMaxLevel(b.Left, maxLvl, lvl+1)
		}

		if b.Right != nil {
			numNodes += findNodesAtMaxLevel(b.Right, maxLvl, lvl+1)
		}
	} else {
		if b.Left != nil {
			numNodes++
		}

		if b.Right != nil {
			numNodes++
		}
	}

	return numNodes
}
