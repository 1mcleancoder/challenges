package ch31_2DArrays_num_islands

type NumIslands interface {
	findNumIslands(arr [][]int) int
}

type TwoD struct {
	r, c int
}
