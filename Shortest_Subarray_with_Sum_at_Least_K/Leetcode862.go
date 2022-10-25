package main

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSumArr := make([]int, n+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
	}
	res := n + 1
	var q []int
	for i, curSum := range preSumArr {
		for len(q) > 0 && curSum-preSumArr[q[0]] >= k {
			res = min(res, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && preSumArr[q[len(q)-1]] >= curSum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if res < n+1 {
		return res
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
