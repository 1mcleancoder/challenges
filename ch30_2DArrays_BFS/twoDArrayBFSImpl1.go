package ch30_2DArrays_BFS

type TwoDArrayBFSImpl1 struct {
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Move struct {
	rowInr int
	colInr int
}

func (d TwoDArrayBFSImpl1) twoDArrayTraverseBFS(arr [][]int, start TwoD) []int {
	fixedMoves := getMovements()

	positions := []TwoD{}
	positions = append(positions, start)

	visited := make([][]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		visited[i] = make([]bool, len(arr[i]), len(arr[i]))
	}

	bfs := make([]int, 0, len(arr)*len(arr[0]))

	for len(positions) > 0 {
		cPos := positions[0]
		positions = positions[1:]

		bfs = append(bfs, arr[cPos.r][cPos.c])
		visited[cPos.r][cPos.c] = true

		for _, m := range fixedMoves {
			if isValid, r, c := d.isValid(cPos.r, cPos.c, m, visited); isValid {
				visited[r][c] = true
				positions = append(positions, TwoD{r, c})
			}
		}
	}

	return bfs
}

func getMovements() []Move {
	movements := make([]Move, 4, 4)
	movements[UP] = Move{-1, 0}   // up
	movements[RIGHT] = Move{0, 1} // right
	movements[DOWN] = Move{1, 0}  // down
	movements[LEFT] = Move{0, -1} // left

	return movements
}

func (d TwoDArrayBFSImpl1) traverseDFS(r, c int, arr [][]int, visited [][]bool, movements []Move, dfs []int) []int {
	dfs = append(dfs, arr[r][c])
	visited[r][c] = true

	for i := range movements {
		if isValid, newRow, newCol := d.isValid(r, c, movements[i], visited); isValid {
			dfs = d.traverseDFS(newRow, newCol, arr, visited, movements, dfs)
			break
		}
	}

	return dfs
}

func (d TwoDArrayBFSImpl1) isValid(r, c int, dir Move, visited [][]bool) (bool, int, int) {
	r += dir.rowInr
	if r >= 0 && r < len(visited) {
		c += dir.colInr
		if c >= 0 && c < len(visited[r]) {
			if !visited[r][c] {
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
