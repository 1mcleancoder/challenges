package bstValidate

import bst "challenges/ch27_binary_search_tree"

type BstValidateSolution1BruteForce struct {
}

func (b BstValidateSolution1BruteForce) validateBST(n *bst.BinaryNode) bool {
	if n == nil {
		return true
	}

	max := n.Data
	if n.Left != nil {
		if !validateLeftBSTWithMinAndMaxNums(n.Left, -1, max) {
			return false
		}
	}

	if n.Right != nil {
		if !validateRightBSTWithMinAndMaxNums(n.Right, max, -1) {
			return false
		}
	}

	return true
}

func validateLeftBSTWithMinAndMaxNums(n *bst.BinaryNode, min, max int) bool {
	if n.Left != nil {
		if n.Left.Data < n.Data && (min == -1 || (min != -1 && n.Left.Data > min)) {
			return validateLeftBSTWithMinAndMaxNums(n.Left, min, max)
		} else {
			return false
		}
	}

	if n.Right != nil {
		if n.Right.Data > n.Data && n.Right.Data < max {
			return validateLeftBSTWithMinAndMaxNums(n.Right, n.Data, max)
		} else {
			return false
		}
	}

	return true
}

func validateRightBSTWithMinAndMaxNums(n *bst.BinaryNode, min, max int) bool {
	if n.Left != nil {
		if n.Left.Data < n.Data && n.Left.Data > min {
			return validateRightBSTWithMinAndMaxNums(n.Left, min, n.Data)
		} else {
			return false
		}
	}

	if n.Right != nil {
		if n.Right.Data > n.Data && (max == -1 || (max != -1 && n.Right.Data < max)) {
			return validateRightBSTWithMinAndMaxNums(n.Right, min, max)
		} else {
			return false
		}
	}

	return true
}
