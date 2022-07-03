package main

func wiggleMaxLength(nums []int) int {

	pos, neg := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			pos = neg + 1
		} else if nums[i+1] < nums[i] {
			neg = pos + 1
		}
	}
	return max(pos, neg) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -贪心,DP
上升,下降交替
*/
