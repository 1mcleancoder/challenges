package ch30_2DArrays_BFS

type TwoDArrayBFS interface {
	twoDArrayTraverseBFS(arr [][]int, start TwoD) []int
}

type TwoD struct {
	r, c int
}
