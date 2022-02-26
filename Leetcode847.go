package main

func shortestPathLength(graph [][]int) int {

	n := len(graph)
	queue := make([][3]int, 0)
	seen := make([][]bool, n)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		queue = append(queue, [3]int{i, 1 << i, 0})
	}
	res := 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		curNode, mask, dist := front[0], front[1], front[2]
		if mask == (1<<n)-1 {
			return dist
		}
		for _, v := range graph[curNode] {
			curMask := mask | (1 << v)
			if !seen[v][curMask] {
				queue = append(queue, [3]int{v, curMask, dist + 1})
				seen[v][curMask] = true
			}
		}
	}
	return res
}
