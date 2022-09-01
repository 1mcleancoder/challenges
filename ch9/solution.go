package ch9

type hurdleRaceProblem interface {
	hurdleRace(k int32, heights []int32) int32
}

type HurdleRaceSolution1 struct {
}

func (h HurdleRaceSolution1) hurdleRace(k int32, height []int32) int32 {
	max := int32(0)
	for _, v := range height {
		if v > max {
			max = v
		}
	}

	var jumps int32
	if k >= max {
		jumps = 0
	} else {
		jumps = max - k
	}
	return jumps
}
