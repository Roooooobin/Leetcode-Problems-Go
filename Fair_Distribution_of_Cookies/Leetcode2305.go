package main

func distributeCookies(cookies []int, k int) int {

	sums := make([]int, k)
	res := 0x3f3f3f3f
	var backtracking func(idx int, curSums []int)
	backtracking = func(idx int, curSums []int) {
		if idx == len(cookies) {
			maxm := curSums[0]
			for _, cur := range curSums {
				maxm = max(maxm, cur)
			}
			res = min(res, maxm)
			return
		}
		for i := range curSums {
			curSums[i] += cookies[idx]
			if curSums[i] < res {
				backtracking(idx+1, curSums)
			}
			curSums[i] -= cookies[idx]
		}
	}
	backtracking(0, sums)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
