package bstRightSideView

import bst "challenges/ch27_binary_search_tree"

type BstRightSideViewProblem interface {
	findRightSideViewValues(n *bst.BinaryNode) []int
}
