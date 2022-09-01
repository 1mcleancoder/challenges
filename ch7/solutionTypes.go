package ch7

import "fmt"

type Item struct {
	name   string
	expiry int
}

func (t *Item) String() string {
	return fmt.Sprintf("[name=%-7v, expiry=%3d]", t.name, t.expiry)
}
