package ch17_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	n := Node{}
	list := n.Create([]int{1, 2, 3, 4, 5})
	count := list.Print()
	assert.Equal(t, 5, count)
}
