package ch12_get_trapped_water

type solution3 struct {
}

// height := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
// sum = 8
func (solution3) getTrappedWater(h []int) int {
	var sum int

	// iterate over the h array with two pointers say maxLeft for left and maxRight for right
	//for maxLeft, maxRight := 0, len(h)-1; maxLeft < maxRight; {
	var maxLeft, maxRight = h[0], h[len(h)-1]

	for pL, pR := 0, len(h)-1; pL < pR; {
		if h[pL] <= h[pR] {
			if maxLeft <= h[pL] {
				maxLeft = h[pL]
			} else {
				sum += maxLeft - h[pL]
			}

			pL++
		} else {
			if maxRight <= h[pR] {
				maxRight = h[pR]
			} else {
				sum += maxRight - h[pR]
			}

			pR--
		}
	}

	//fmt.Println(maxLeft, maxRight)
	return sum
}
