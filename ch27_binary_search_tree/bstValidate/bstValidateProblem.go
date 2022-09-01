package bstValidate

import bst "challenges/ch27_binary_search_tree"

type BstValidate interface {
	validateBST(node *bst.BinaryNode) bool
}
