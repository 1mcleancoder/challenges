package ch6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMoneySpent(t *testing.T) {
	keyboards := []int32{3, 1}
	drives := []int32{5, 2, 8}
	b := int32(10)
	moneySpent := getMoneySpent(keyboards, drives, b)
	assert.Equal(t, int32(9), moneySpent)
}
