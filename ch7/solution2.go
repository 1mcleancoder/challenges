package ch7

type PQItemsByExpiry2 []*Item

func (pq PQItemsByExpiry2) Len() int {
	return len(pq)
}

func (pq PQItemsByExpiry2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQItemsByExpiry2) Less(i, j int) bool {
	return pq[i].expiry < pq[j].expiry
}

func (pq *PQItemsByExpiry2) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PQItemsByExpiry2) Pop() interface{} {
	old := *pq
	oldLen := len(old)

	v := old[oldLen-1]

	*pq = old[0 : oldLen-1]
	return v
}
