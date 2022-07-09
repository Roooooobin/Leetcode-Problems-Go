package main

func findRedundantConnection(edges [][]int) []int {

	n := len(edges)
	par := make([]int, n+1)
	for i := range par {
		par[i] = i
	}
	find := func(x int) int {
		r := x
		for r != par[r] {
			r = par[r]
		}
		for x != r {
			p := par[x]
			par[x] = r
			x = p
		}
		return r
	}
	union := func(x, y int) bool {
		x = find(x)
		y = find(y)
		if x != y {
			par[x] = y
			return false
		} else {
			return true
		}
	}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if union(x, y) {
			return edge
		}
	}
	return []int{}
}
