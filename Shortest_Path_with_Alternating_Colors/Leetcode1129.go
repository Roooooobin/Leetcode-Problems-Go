package main

type pair struct {
	x, color int
}

func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {

	graph := make([][]pair, n)
	for _, e := range redEdges {
		graph[e[0]] = append(graph[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		graph[e[0]] = append(graph[e[0]], pair{e[1], 1})
	}
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	seen := make([][2]bool, n)
	seen[0] = [2]bool{true, true}
	q := []pair{{0, 0}, {0, 1}}
	for level := 0; len(q) > 0; level++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x := p.x
			if res[x] < 0 {
				res[x] = level
			}
			for _, to := range graph[x] {
				if to.color != p.color && !seen[to.x][to.color] {
					seen[to.x][to.color] = true
					q = append(q, to)
				}
			}
		}
	}
	return res
}
