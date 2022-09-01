package bstNumNodesAtLastLevel

import (
	bst "challenges/ch27_binary_search_tree"
	"math"
)

type NumNodesAtLastLevelSolution2BinarySearch struct {
}

func (n NumNodesAtLastLevelSolution2BinarySearch) findNumNodesAtLastLevel(b *bst.BinaryNode) int {
	if b == nil {
		return 0
	}

	mxLvl := findMaxLevelBinarySearch(b, 1)
	return findNodesAtMaxLevelUsingBinarySearch(b, mxLvl, 1)
}

func findMaxLevelBinarySearch(b *bst.BinaryNode, lvl int) int {
	maxLvl := lvl
	if b.Left != nil {
		newLvl := findMaxLevel(b.Left, lvl+1)
		maxLvl = int(math.Max(float64(newLvl), float64(lvl)))
	}

	return maxLvl

}

func findNodesAtMaxLevelUsingBinarySearch(b *bst.BinaryNode, maxLvl, lvl int) int {
	// calc max elements at last level = 2 ^ (maxLvl -1) - 1 say mx
	maxElements := int(math.Pow(2, float64(maxLvl-1)) - 1)

	min := 0
	max := maxElements

	for min < max {
		mid := (max+min)/2 + 1
		isMidExisting := findElementAtPos(b, maxLvl, 1, 0, maxElements, mid)
		if isMidExisting {
			min = mid
		} else if max-min == 1 {
			max = min // here max is not existing so min must be existing since we already checked min earlier (mid)
			break
		} else {
			// find elements at positions: min (say 0), mid ((min + mx) / 2) + 1) and mx using logN time
			// loop till min < mx
			// if mx is existing then mid = min and calc mid again
			// else mx is mid and calc mid again using min and mid
			max = mid - 1
		}
	}

	return max + 1
}

func findElementAtPos(b *bst.BinaryNode, maxLvl, lvl, left, right, pos int) bool {
	if maxLvl == lvl {
		if b != nil {
			return true
		} else {
			return false
		}
	} else if b == nil {
		return false
	}

	mid := (right+left)/2 + 1
	if pos >= mid {
		return findElementAtPos(b.Right, maxLvl, lvl+1, mid, right, pos)
	} else {
		return findElementAtPos(b.Left, maxLvl, lvl+1, left, mid-1, pos)
	}
}
