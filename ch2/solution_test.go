package ch2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := createList([]int{9,9,9})
	l2 := createList([]int{9,9,9,9})

	sol := SolutionAddIntNums{}
	testSolution(t, sol, l1, l2)
}

func testSolution(t *testing.T, sol AddIntLists, l1, l2 *ListNode) {
	r := sol.addTwoNumbers(l1, l2)
	assert.NotNil(t, r)

	e := createList([]int{8,9,9,0,1})
	assert.Equal(t, e, r)
}

func createList(nums []int) *ListNode {
	var n *ListNode
	for i:=len(nums)-1; i>=0; i-- {
		t := &ListNode{Val: nums[i], Next: n}
		n = t
	}
	printList(n)
	return n
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Printf("%d ", n.Val)
		n = n.Next
	}

	fmt.Println()
}