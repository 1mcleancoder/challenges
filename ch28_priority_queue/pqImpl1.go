package ch28_priority_queue

type PriorityQueueImpl1 struct {
	arr []int
}

func (pq *PriorityQueueImpl1) Less(a, b int) bool {
	if a > pq.Len()-1 {
		return false
	}
	if b > pq.Len()-1 {
		return false
	}

	if pq.arr[a] < pq.arr[b] {
		return true
	} else {
		return false
	}
}

func (pq *PriorityQueueImpl1) Swap(a, b int) {
	pq.arr[a], pq.arr[b] = pq.arr[b], pq.arr[a]
}

func (pq *PriorityQueueImpl1) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueueImpl1) Peek() int {
	if pq.IsEmpty() {
		return -1
	}

	return pq.arr[0]
}

func (pq *PriorityQueueImpl1) Push(v int) {
	if pq.IsEmpty() {
		pq.arr = append(pq.arr, v)
		return
	}

	pq.arr = append(pq.arr, v)
	pq.percolateUp()
}

func (pq *PriorityQueueImpl1) Pop() int {
	if pq.IsEmpty() {
		return -1
	}

	// delete the top value
	v := pq.arr[0]
	// add the last value at the top position
	pq.arr[0] = pq.arr[pq.Len()-1]
	pq.arr = pq.arr[0 : pq.Len()-1]

	pq.percolateDown()
	return v
}

func (pq *PriorityQueueImpl1) IsEmpty() bool {
	if pq.Len() == 0 {
		return true
	}

	return false
}

func (pq *PriorityQueueImpl1) percolateUp() {
	c := pq.Len() - 1
	p := pq.parentIndex(c)

	for c != 0 && pq.Less(p, c) {
		//swap value and parent
		pq.Swap(c, p)
		c = p
		p = pq.parentIndex(c)
	}
}

func (pq *PriorityQueueImpl1) parentIndex(idx int) int {
	return (idx - 1) / 2
}

func (pq *PriorityQueueImpl1) leftChildIndex(idx int) int {
	return (2 * idx) + 1
}

func (pq *PriorityQueueImpl1) rightChildIndex(idx int) int {
	return (2 * idx) + 2
}

func (pq *PriorityQueueImpl1) percolateDown() {
	if pq.IsEmpty() {
		return
	}

	p := 0

	for {
		left := pq.leftChildIndex(p)
		if left > pq.Len()-1 {
			break // it's the last element
		}

		var grChildIdx int

		right := pq.rightChildIndex(p)
		if right < pq.Len() {
			if pq.Less(left, right) {
				grChildIdx = right
			} else {
				grChildIdx = left
			}
		} else {
			grChildIdx = left
		}

		if pq.Less(p, grChildIdx) {
			pq.Swap(p, grChildIdx) //swap value and parent
			p = grChildIdx
		} else {
			break
		}
	}
}
