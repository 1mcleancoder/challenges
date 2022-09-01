package bstValidate

import (
	bst "challenges/ch27_binary_search_tree"
	"math"
)

type BstValidateSolution2Efficient2 struct {
}

func (b BstValidateSolution2Efficient2) validateBST(n *bst.BinaryNode) bool {
	return validate1(n, math.MinInt, math.MaxInt)
}

func validate1(n *bst.BinaryNode, grThan, lessThan int) bool {
	if grThan < n.Data && n.Data < lessThan {
		if n.Left != nil && !validate(n.Left, grThan, n.Data) {
			return false
		}

		if n.Right != nil && !validate(n.Right, n.Data, lessThan) {
			return false
		}
	} else {
		return false
	}

	return true
}
