package ch18_doubly_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := Create([]any{1, 2, 3, 4, 5})
	count := Print(list)
	assert.Equal(t, 5, count)
}
