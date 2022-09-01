package bstLevelValuesProblem

import "challenges/ch27_binary_search_tree"

type BSTLevelValuesSolution1BruteForce struct {
}

func (b BSTLevelValuesSolution1BruteForce) collectLevelValues(n *ch27_binary_search_tree.BinaryNode) [][]int {
	type BNodeWithLevel struct {
		node *ch27_binary_search_tree.BinaryNode
		lvl  int
	}

	valsByLvl := make(map[int][]int)

	q := make([]*BNodeWithLevel, 0)
	q = append(q, &BNodeWithLevel{n, 1})

	for len(q) > 0 {
		v := q[0]
		if vals, ok := valsByLvl[v.lvl]; ok {
			vals = append(vals, v.node.Data)
			valsByLvl[v.lvl] = vals
		} else {
			vals = make([]int, 0)
			vals = append(vals, v.node.Data)
			valsByLvl[v.lvl] = vals
		}

		q = q[1:] // ignore the first element

		if v.node.Left != nil {
			q = append(q, &BNodeWithLevel{v.node.Left, v.lvl + 1})
		}

		if v.node.Right != nil {
			q = append(q, &BNodeWithLevel{v.node.Right, v.lvl + 1})
		}
	}

	lvlVals := make([][]int, 0, len(valsByLvl))
	for i := 1; i <= len(valsByLvl); i++ {
		v := valsByLvl[i]
		lvlVals = append(lvlVals, v)
	}

	return lvlVals
}
