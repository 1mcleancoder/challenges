package ch21_queue

type Queue interface {
	Enqueue(v any)
	Dequeue() any
	Peek() any
	IsEmpty() bool
}
