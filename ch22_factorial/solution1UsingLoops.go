package ch22_factorial

type solution1UsingLoops struct {
}

func (s1 *solution1UsingLoops) calcFactorial(v int) int {
	if v <= 2 {
		return v
	}

	// declare fact to store the result
	fact := 1
	// start a loop from i=1 <= v
	for i := 2; i <= v; i++ {
		// multiply fact with i and store the result in fact
		fact *= i
	}
	// end loop

	return fact
}
