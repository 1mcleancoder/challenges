package ch7

import "fmt"

type PQItemByExpiry []*Item

func (p PQItemByExpiry) Len() int {
	return len(p)
}

func (p PQItemByExpiry) Less(i, j int) bool {
	return p[i].expiry < p[j].expiry
}

func (p PQItemByExpiry) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQItemByExpiry) Push(v interface{}) {
	*p = append(*p, v.(*Item))
}

func (p *PQItemByExpiry) Pop() interface{} {
	old := *p
	oldLen := len(old)

	v := old[oldLen-1]

	*p = old[:oldLen-1]
	return v
}

func (p *PQItemByExpiry) String() string {
	s := ""
	for i := 0; len(*p) > 0; i++ {
		s += fmt.Sprintf("%v ", *(*p)[i])
	}

	return s
}
