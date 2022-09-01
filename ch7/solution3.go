package ch7

type PQItemsByExpiry3 []*Item

func (p PQItemsByExpiry3) Len() int {
	return len(p)
}

func (p PQItemsByExpiry3) Less(i, j int) bool {
	//return p[i].expiry < p[j].expiry
	return p[i].name < p[j].name
}

func (p PQItemsByExpiry3) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQItemsByExpiry3) Push(v interface{}) {
	*p = append(*p, v.(*Item))
}

func (p *PQItemsByExpiry3) Pop() interface{} {
	old := *p
	oldLen := len(old)

	v := old[oldLen-1]

	*p = old[:oldLen-1]
	return v
}
