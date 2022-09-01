package ch8

import "fmt"

type Item struct {
	name   string
	expiry int
}

func (t *Item) String() string {
	return fmt.Sprintf("[name=%v, expiry=%d]", t.name, t.expiry)
}

type PQItemByName []*Item

func (p PQItemByName) Len() int {
	return len(p)
}

func (p PQItemByName) Less(i, j int) bool {
	return p[i].name < p[j].name
}

func (p PQItemByName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQItemByName) Push(v interface{}) {
	*p = append(*p, v.(*Item))
}

func (p *PQItemByName) Pop() interface{} {
	old := *p
	oldLen := len(old)

	v := old[oldLen-1]
	*p = old[:oldLen-1]

	return v
}
