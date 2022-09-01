package ch17_3_cycle_detection

import list "challenges/ch17_linked_lists"

type Problem interface {
	detectCycle(head *list.Node) *list.Node
}
