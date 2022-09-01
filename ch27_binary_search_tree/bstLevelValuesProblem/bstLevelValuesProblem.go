package bstLevelValuesProblem

import bst "challenges/ch27_binary_search_tree"

type LevelValues interface {
	collectLevelValues(n *bst.BinaryNode) [][]int
}
