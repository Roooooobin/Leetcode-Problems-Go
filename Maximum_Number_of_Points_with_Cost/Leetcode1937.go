package main

//https://leetcode.cn/problems/maximum-number-of-points-with-cost/solution/dp-you-hua-ji-qiao-chai-xiang-qian-hou-z-5vvc/
func maxPoints(points [][]int) int64 {

	m := len(points[0])
	dp := make([][2]int, m)
	suffixMax := make([]int, m)
	res := 0
	for i, row := range points {
		if i == 0 {
			for j, v := range row {
				res = max(res, v)
				dp[j][0] = v + j
				dp[j][1] = v - j
			}
		} else {
			preMax := -0x3f3f3f3f
			for j, v := range row {
				preMax = max(preMax, dp[j][0])
				cur := max(v-j+preMax, v+j+suffixMax[j])
				res = max(res, cur)
				dp[j][0] = cur + j
				dp[j][1] = cur - j
			}
		}
		suffixMax[m-1] = dp[m-1][1]
		for j := m - 2; j >= 0; j-- {
			suffixMax[j] = max(suffixMax[j+1], dp[j][1])
		}
	}
	return int64(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -DP
算完每一层之后,用前缀最小值和后缀最小值降低复杂度
*/
