package main

func findKDistantIndices(nums []int, key int, k int) []int {

	res := make([]int, 0)
	n := len(nums)
	for i := range nums {
		for j := max(0, i-k); j <= min(n-1, i+k); j++ {
			if nums[j] == key {
				res = append(res, i)
				break
			}
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

//func min(x, y int) int {
//	if x < y {
//		return x
//	} else {
//		return y
//	}
//}
