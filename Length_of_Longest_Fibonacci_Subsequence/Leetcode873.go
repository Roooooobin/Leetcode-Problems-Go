package main

func lenLongestFibSubseq(arr []int) int {

	n := len(arr)
	dp := make(map[int]int)
	mp := make(map[int]int)
	for i, x := range arr {
		mp[x] = i
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] >= arr[j] {
				continue
			}
			if idx, ok := mp[arr[i]-arr[j]]; ok {
				dp[j*n+i] = max(dp[j*n+i], dp[idx*n+j]+1)
				res = max(dp[j*n+i]+2, res)
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -DP
dp[j][i] = dp[idx][j]+1 idx是arr[i] - arr[j],形成idx, j ,i的fibonacci数列
*/
