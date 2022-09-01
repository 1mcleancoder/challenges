package ch23_sorting

type bubbleSort1 struct {
}

func (s1 bubbleSort1) sort(b []int) {
	for i := 0; i < len(b)-1; i++ {
		max := len(b) - 1 - i // optimization
		for j := 0; j < max; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
}
