package main

func minimizeTheDifference(mat [][]int, target int) int {

	var dfs func(idx, curSum int)
	m, n := len(mat), len(mat[0])
	res := 0x3f3f3f3f
	dp := [71][4903]bool{}
	dfs = func(i, curSum int) {
		if curSum-target > res || dp[i][curSum] {
			return
		}
		dp[i][curSum] = true
		if i == m {
			res = min(res, abs(target-curSum))
			return
		}
		for j := 0; j < n; j++ {
			dfs(i+1, curSum+mat[i][j])
		}
	}
	dfs(0, 0)
	return res
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/solution/liang-chong-fang-fa-jian-dan-yi-dong-ji-7jugp/
