package ch17_linked_lists

type LinkedList interface {
	Print(n Node) int
	Create(vals []int) Node
	Length() int
}
