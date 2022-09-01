package bstHeightProblem

import bst "challenges/ch27_binary_search_tree"

type HeightProblem interface {
	findMaxHeight(n *bst.BinaryNode) int
}
