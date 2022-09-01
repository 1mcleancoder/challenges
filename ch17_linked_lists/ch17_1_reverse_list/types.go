package ch17_1_reverse_list

import list "challenges/ch17_linked_lists"

type reverse interface {
	Reverse(n *list.Node) *list.Node
}
