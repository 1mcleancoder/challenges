package ch29_2DArrays_DFS

type TwoDArrayImpl1 struct {
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

func (d TwoDArrayImpl1) twoDArrayTraverse(arr [][]int) []int {
	visited := make([][]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		visited[i] = make([]bool, len(arr[i]), len(arr[i]))
	}

	movements := make([]Move, 4, 4)
	movements[UP] = Move{-1, 0}   // up
	movements[RIGHT] = Move{0, 1} // right
	movements[DOWN] = Move{1, 0}  // down
	movements[LEFT] = Move{0, -1} // left

	return d.traverseDFS(0, 0, arr, visited, movements, make([]int, 0))
}

func (d TwoDArrayImpl1) traverseDFS(r, c int, arr [][]int, visited [][]bool, movements []Move, dfs []int) []int {
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

func (d TwoDArrayImpl1) isValid(r, c int, dir Move, visited [][]bool) (bool, int, int) {
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
