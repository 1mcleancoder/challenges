package ch8

import (
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	items = []*Item{
		{"Carrot", 30},
		{"Spinach", 5},
		{"Rice", 100},
		{"Potato", 45},
	}

	expected = []*Item{
		{"Carrot", 30},
		{"Potato", 45},
		{"Rice", 100},
		{"Spinach", 5},
	}

	pqItemsByName = make(PQItemByName, 0, len(items))
)

func TestSolution(t *testing.T) {
	pqItemsByName = append(pqItemsByName, items...)

	heap.Init(&pqItemsByName)

	fmt.Println("after making heap\n", pqItemsByName)

	for i := 0; pqItemsByName.Len() > 0; i++ {
		actual := heap.Pop(&pqItemsByName)
		assert.Equal(t, expected[i], actual)
	}
}
