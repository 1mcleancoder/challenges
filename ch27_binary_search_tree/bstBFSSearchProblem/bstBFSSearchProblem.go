package bstBFSSearchProblem

import bst "challenges/ch27_binary_search_tree"

type BfsSearch interface {
	bfs(n *bst.BinaryNode) []int
}
