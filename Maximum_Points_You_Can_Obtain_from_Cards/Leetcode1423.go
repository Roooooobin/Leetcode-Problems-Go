package main

func maxScore(cardPoints []int, k int) int {

	n := len(cardPoints)
	m := n - k
	prefixSum := make([]int, n+1)
	res := 0x3f3f3f3f
	for i, c := range cardPoints {
		prefixSum[i+1] = prefixSum[i] + c
		if i >= m-1 {
			res = min(res, prefixSum[i+1]-prefixSum[i+1-m])
		}
	}
	return prefixSum[n] - res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
前缀和, 滑动窗口
每次只能拿最左边或最右边, 那么反向考虑, 拿中间的n-k个最小和, 用总和减去即得k个最大和
*/
