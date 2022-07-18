package main

type pair struct{ x, y int }

func containVirus(isInfected [][]int) (res int) {

	directions := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	m, n := len(isInfected), len(isInfected[0])
	for {
		var neighbors []map[pair]struct{}
		var firewalls []int
		for i, row := range isInfected {
			for j, infected := range row {
				if infected != 1 {
					continue
				}
				queue := []pair{{i, j}}
				neighbor := map[pair]struct{}{}
				firewall, idx := 0, len(neighbors)+1
				row[j] = -idx
				for len(queue) > 0 {
					p := queue[0]
					queue = queue[1:]
					for _, d := range directions {
						if x, y := p.x+d[0], p.y+d[1]; 0 <= x && x < m && 0 <= y && y < n {
							if isInfected[x][y] == 1 {
								queue = append(queue, pair{x, y})
								isInfected[x][y] = -idx
							} else if isInfected[x][y] == 0 {
								firewall++
								neighbor[pair{x, y}] = struct{}{}
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}

		if len(neighbors) == 0 {
			break
		}

		idx := 0
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}

		res += firewalls[idx]
		for _, row := range isInfected {
			for j, v := range row {
				if v < 0 {
					if v != -idx-1 {
						row[j] = 1
					} else {
						row[j] = 2
					}
				}
			}
		}

		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}

		if len(neighbors) == 1 {
			break
		}
	}
	return
}

/*
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/contain-virus/solution/ge-chi-bing-du-by-leetcode-solution-vn9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
- -BFS
*/
