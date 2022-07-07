package main

func minFallingPathSum(matrix [][]int) int {

	INF := 0x3f3f3f3f
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = INF
	}
	preMin, preSecondMin := INF, INF
	for j := 0; j < n; j++ {
		dp[j] = matrix[0][j]
		if preMin == INF || dp[j] < preMin {
			preSecondMin = preMin
			preMin = dp[j]
		} else if preSecondMin == INF || dp[j] < preSecondMin {
			preSecondMin = dp[j]
		}
	}
	for i := 1; i < m; i++ {
		cur := make([]int, n)
		for j := 0; j < n; j++ {
			if dp[j] == preMin {
				cur[j] = preSecondMin + matrix[i][j]
			} else {
				cur[j] = preMin + matrix[i][j]
			}
		}
		dp = cur
		preMin, preSecondMin = INF, INF
		for j := 0; j < n; j++ {
			if preMin == INF || dp[j] < preMin {
				preSecondMin = preMin
				preMin = dp[j]
			} else if preSecondMin == INF || dp[j] < preSecondMin {
				preSecondMin = dp[j]
			}
		}
	}
	return preMin
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
- -DP
每一行结束后记录这一行的最小值和第二最小值,在下一行计算时,如果上一行该数等于最小值(因为不能用同一列)那么就用第二最小值累加,否则用最小值累加
一维滚动数组优化空间复杂度
*/

func main() {
	minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
