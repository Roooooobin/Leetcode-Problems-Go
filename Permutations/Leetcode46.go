package main

func permute(nums []int) [][]int {

	n := len(nums)
	res := make([][]int, 0)
	swap := func(i, j int, a []int) {
		a[i], a[j] = a[j], a[i]
	}
	var backtracking func(idx int, cur []int)
	backtracking = func(idx int, cur []int) {
		if idx == n {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := idx; i < n; i++ {
			swap(idx, i, cur)
			backtracking(idx+1, cur)
			swap(idx, i, cur)
		}
	}
	backtracking(0, nums)
	return res
}
