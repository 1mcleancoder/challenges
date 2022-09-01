package ch23_sorting

type Sort1UsingInbuiltSorting []int

func (s1 Sort1UsingInbuiltSorting) Len() int {
	return len(s1)
}

func (s1 Sort1UsingInbuiltSorting) Swap(i, j int) {
	s1[i], s1[j] = s1[j], s1[i]
}

func (s1 Sort1UsingInbuiltSorting) Less(i, j int) bool {
	return s1[i] < s1[j]
}
