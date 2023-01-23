package main

func maxValue(n int, index int, maxSum int) int {

	calcSum := func(x, length int) int {
		if length == 0 {
			return 0
		}
		if x < length {
			return x*(x+1)/2 + length - x
		} else {
			return (2*x - length + 1) * length / 2
		}
	}

	pass := func(x int) bool {
		return x+calcSum(x-1, index)+calcSum(x-1, n-index-1) <= maxSum
	}

	lo, hi := 1, maxSum
	for lo <= hi {
		mi := lo + (hi-lo)/2
		if pass(mi) {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return hi
}

//func main() {
//	maxValue(3, 2, 18)
//}

/*
- -二分答案+贪心
贪心:两边的数组最小和(等差数列)之和要小于等于maxSum-mid
二分查找该值
*/
