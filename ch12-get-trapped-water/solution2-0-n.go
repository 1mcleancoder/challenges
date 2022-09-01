package ch12_get_trapped_water

import "fmt"

type solution2 struct {
}

// height := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
// sum = 8
func (solution2) getTrappedWater(h []int) int {
	var sum int

	// iterate over the h array with two pointers say maxLeft for left and maxRight for right
	//for maxLeft, maxRight := 0, len(h)-1; maxLeft < maxRight; {
	var maxLeft, maxRight = h[0], h[len(h)-1]

	for pL, pR := 0, len(h)-1; pL < pR; {
		// check if maxLeft <= maxRight
		if maxLeft <= maxRight {
			if h[pL] > maxLeft { // if yes, check h[pL] > maxLeft
				// if yes, then update maxLeft = h[pL] and no need to update sum
				maxLeft = h[pL]
			} else { // if no, then update sum += maxLeft - h[pL]
				sum += maxLeft - h[pL]
			}
			// update pL++
			if maxLeft <= maxRight {
				pL++
			}
		} else {
			if h[pR] > maxRight {
				// if yes, maxRight = h[pR]
				maxRight = h[pR]
			} else { // if no, sum += maxRight - h[pR]
				sum += maxRight - h[pR]
			}

			if maxLeft > maxRight {
				pR-- // update pR--
			}
		}
	}

	fmt.Println(maxLeft, maxRight)
	return sum
}

func minVal(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
