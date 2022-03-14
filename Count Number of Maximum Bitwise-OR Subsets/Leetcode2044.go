package main

import "fmt"

func countMaxOrSubsets(nums []int) int {

	maximum := 0
	maxCount := 0
	var dfs func(a []int, idx, cur int)
	dfs = func(a []int, idx, cur int) {
		if cur == maximum {
			maxCount++
		} else if cur > maximum {
			maximum = cur
			maxCount = 1
		}
		if idx >= len(a) {
			return
		}
		for i := idx; i < len(a); i++ {
			dfs(a, i+1, cur|a[i])
		}
	}
	dfs(nums, 0, 0)
	return maxCount
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))
}
