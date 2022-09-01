package ch17_2_mn_reversals

import list "challenges/ch17_linked_lists"

type Reverse interface {
	Reverse(headNode *list.Node, m, n int) *list.Node
}
