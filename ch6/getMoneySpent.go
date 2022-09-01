package ch6

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int32

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int32))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	fmt.Printf("keyboards: %v, drives: %v, b: %v\n", keyboards, drives, b)
	/*
	 * Write your code here.
	 */
	// conditions:
	// both needs to be bought otherwise -1
	// spend max money
	// max elements is 1000
	// budget is 1m

	// make a max heap of keyboards & drives
	kH := &MaxHeap{}
	*kH = append(*kH, keyboards...)
	heap.Init(kH)

	max := int32(-1)

	// iterate over all the keyboards starting from max
	for kH.Len() > 0 {
		maxK := heap.Pop(kH).(int32)
		fmt.Printf("maxKb: %v\n", maxK)
		// pop max kb if maxKb is > b then continue
		if maxK > b {
			fmt.Println("\tcontinue next KB")
			continue
		}

		dH := &MaxHeap{}
		*dH = append(*dH, drives...)
		fmt.Printf("drives: %v\n", len(drives))
		heap.Init(dH)

		// iterate over drives heap
		for dH.Len() > 0 {
			maxD := heap.Pop(dH).(int32)
			s := maxK + maxD
			fmt.Printf("maxH: %v, maxD: %v, s = %v\n", maxK, maxD, s)

			if s > b {
				fmt.Println("\tcontinue next Drive")
				continue
			} else if s > max {
				fmt.Printf("found possible max: old=%v, new=%v\n", max, s)
				max = s
			}
		}

		// fmt.Printf("temp drives Heap=%v, new=%v\n", tmpDh, *dH)
		// dH = &tmpDh
	}

	return max
}
