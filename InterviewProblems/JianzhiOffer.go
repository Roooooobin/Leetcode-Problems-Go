package main

func subarraySum(nums []int, k int) int {

	hash := make(map[int]int)
	sum := 0
	hash[0] = 1
	res := 0
	for _, num := range nums {
		sum += num
		res += hash[sum-k]
		v := hash[sum]
		hash[sum] = v + 1
	}
	return res
}

func findMaxLength(nums []int) int {

	hash := make(map[int]int)
	diff := 0
	hash[0] = -1
	res := 0
	for i, num := range nums {
		if num == 1 {
			diff++
		} else {
			diff--
		}
		if _, ok := hash[diff]; ok {
			res = max(res, i-hash[diff])
		} else {
			hash[diff] = i
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func countSubstrings(s string) int {

	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	res := n
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}
