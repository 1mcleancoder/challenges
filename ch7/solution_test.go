package ch7

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	items = []*Item{
		{"Carrot", 30},
		{"Spinach", 5},
		{"Rice", 100},
		{"potato", 45},
	}

	expected = []*Item{
		{"Spinach", 5},
		{"Carrot", 30},
		{"Potato", 45},
		{"Rice", 100},
	}

	pqItemsByExpiry = make(PQItemsByExpiry3, 0, len(items))
)

func TestSolution(t *testing.T) {
	pqItemsByExpiry = append(pqItemsByExpiry, items...)

	heap.Init(&pqItemsByExpiry)

	t.Log("after making heap")
	for _, v := range pqItemsByExpiry {
		t.Log(v)
	}

	for i := 0; pqItemsByExpiry.Len() > 0; i++ {
		actual := heap.Pop(&pqItemsByExpiry)
		assert.Equal(t, expected[i], actual)
	}
}
