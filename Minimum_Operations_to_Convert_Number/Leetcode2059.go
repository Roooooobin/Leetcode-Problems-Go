package main

func minimumOperations(nums []int, start int, goal int) int {

	seen := make(map[int]struct{})
	seen[start] = struct{}{}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{start, 0})
	valid := func(x int) bool {
		if x >= 0 && x <= 1000 {
			if _, ok := seen[x]; !ok {
				return true
			}
		}
		return false
	}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, step := front[0], front[1]
		if x == goal {
			return step
		}
		for _, num := range nums {
			nx1 := x + num
			if nx1 == goal {
				return step + 1
			}
			if valid(nx1) {
				queue = append(queue, [2]int{nx1, step + 1})
				seen[nx1] = struct{}{}
			}
			nx2 := x - num
			if nx2 == goal {
				return step + 1
			}
			if valid(nx2) {
				queue = append(queue, [2]int{nx2, step + 1})
				seen[nx2] = struct{}{}
			}
			nx3 := x ^ num
			if nx3 == goal {
				return step + 1
			}
			if valid(nx3) {
				queue = append(queue, [2]int{nx3, step + 1})
				seen[nx3] = struct{}{}
			}
		}
	}
	return -1
}
