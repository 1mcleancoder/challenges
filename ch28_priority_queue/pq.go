package ch28_priority_queue

type PriorityQueue interface {
	Less(a, b int) bool
	Swap(a, b int)
	Len() int
	Peek() int
	Push(v int)
	Pop() int
	IsEmpty() bool
}
