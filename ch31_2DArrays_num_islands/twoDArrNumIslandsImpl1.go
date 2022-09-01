package ch31_2DArrays_num_islands

type TwoDArrNumIslandsImpl1 struct {
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const water, land = 0, 1

type Move struct {
	rowInr int
	colInr int
}

func getMovements() []Move {
	movements := make([]Move, 4, 4)
	movements[UP] = Move{-1, 0}   // up
	movements[RIGHT] = Move{0, 1} // right
	movements[DOWN] = Move{1, 0}  // down
	movements[LEFT] = Move{0, -1} // left

	return movements
}

func (d TwoDArrNumIslandsImpl1) findNumIslands(arr [][]int) int {
	markedPositions := make([][]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		markedPositions[i] = make([]bool, len(arr[i]), len(arr[i]))
	}

	var num int

	for r := range arr {
		for c := range arr[r] {
			v := arr[r][c]
			if v == land && markedPositions[r][c] != true {
				d.markIslandPositions(r, c, arr, markedPositions)
				num++
			}
		}
	}

	return num
}

func (d TwoDArrNumIslandsImpl1) markIslandPositions(r, c int, arr [][]int, markedPositions [][]bool) {

	positions := []TwoD{}
	positions = append(positions, TwoD{r, c})

	for len(positions) > 0 {
		cPos := positions[0]
		positions = positions[1:]

		c++
		markedPositions[cPos.r][cPos.c] = true

		fixedMoves := getMovements()
		for _, m := range fixedMoves {
			if isValid, r, c := d.isLand(cPos.r, cPos.c, m, arr); isValid {
				if !markedPositions[r][c] {
					markedPositions[r][c] = true
					positions = append(positions, TwoD{r, c})
				}
			}
		}
	}
}

func (d TwoDArrNumIslandsImpl1) isLand(r, c int, dir Move, arr [][]int) (bool, int, int) {
	r += dir.rowInr
	if r >= 0 && r < len(arr) {
		c += dir.colInr
		if c >= 0 && c < len(arr[r]) {
			if arr[r][c] == land {
				return true, r, c
			} else {
				return false, -1, -1
			}
		} else {
			return false, -1, -1
		}
	} else {
		return false, -1, -1
	}
}
