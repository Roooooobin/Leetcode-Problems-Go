package main

func networkBecomesIdle(edges [][]int, patience []int) int {

	res := 0
	n := len(patience)
	// 建图
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	dist := make([]int, n)
	queue := [][]int{{0, 0}}
	dist[0] = -1
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, node := range graph[front[0]] {
			// 还没访问过，更新距离
			if dist[node] == 0 {
				dist[node] = front[1] + 1
				queue = append(queue, []int{node, front[1] + 1})
			}
		}
	}
	for i := 1; i < n; i++ {
		d := dist[i] * 2
		res = max(res, (d-1)/patience[i]*patience[i]+d+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
