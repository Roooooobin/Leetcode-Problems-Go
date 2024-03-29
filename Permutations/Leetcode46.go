package main

func permute(nums []int) [][]int {

	n := len(nums)
	res := make([][]int, 0)
	var backtracking func(idx int, cur []int)
	backtracking = func(idx int, cur []int) {
		if idx == n {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := idx; i < n; i++ {
			cur[idx], cur[i] = cur[i], cur[idx]
			backtracking(idx+1, cur)
			cur[idx], cur[i] = cur[i], cur[idx]
		}
	}
	backtracking(0, nums)
	return res
}

/*
- -回溯
0-n-1每个位置都试着与后面交换
最后记录结果要用res = append(res, append([]int{}, cur...)), 不然因为浅拷贝, 后续会影响
*/
