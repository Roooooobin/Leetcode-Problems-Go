package main

func minOperations(nums []int, x int) int {

	n := len(nums)
	prefixSum := 0
	prefixSum2IdxMap := make(map[int]int)
	for i, num := range nums {
		prefixSum += num
		prefixSum2IdxMap[prefixSum] = i + 1
	}
	res := n + 1
	if v, ok := prefixSum2IdxMap[x]; ok {
		res = v
	}
	postSum := 0
	for i := n - 1; i >= 0; i-- {
		postSum += nums[i]
		if postSum > x {
			break
		}
		if postSum == x {
			res = min(res, n-i)
		}
		if v, ok := prefixSum2IdxMap[x-postSum]; ok {
			res = min(res, v+n-i)
		}
	}
	if res == n+1 {
		return -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
