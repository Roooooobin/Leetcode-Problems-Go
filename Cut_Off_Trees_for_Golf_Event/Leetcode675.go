package main

import "sort"

func cutOffTree(forest [][]int) int {
	// sum of the shortest distance between consequent nodes
	m, n := len(forest), len(forest[0])
	// x*n+y, height
	nodes := make([][]int, 0)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if forest[x][y] > 1 {
				nodes = append(nodes, []int{x*n + y, forest[x][y]})
			}
		}
	}
	// sort by height asc
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i][1] < nodes[j][1]
	})
	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	bfs := func(start, end int) int {
		queue := make([][]int, 0)
		queue = append(queue, []int{start, 0})
		seen := make(map[int]struct{})
		seen[start] = struct{}{}
		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]
			x, y := front[0]/n, front[0]%n
			if x*n+y == end {
				return front[1]
			}
			for _, direction := range directions {
				nX, nY := x+direction[0], y+direction[1]
				nextCord := nX*n + nY
				if nX >= 0 && nX < m && nY >= 0 && nY < n && forest[nX][nY] > 0 {
					if _, ok := seen[nextCord]; ok {
						continue
					}
					seen[nextCord] = struct{}{}
					queue = append(queue, []int{nextCord, front[1] + 1})
				}
			}
		}
		return -1
	}
	preCord := 0
	res := 0
	for _, node := range nodes {
		dist := bfs(preCord, node[0])
		if dist == -1 {
			return -1
		}
		res += dist
		forest[preCord/n][preCord%n] = 1
		preCord = node[0]
	}
	return res
}
