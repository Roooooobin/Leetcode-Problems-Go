package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {

	inIdx, largerIdx := -1, -1
	res := 0
	for i, num := range nums {
		if num >= left && num <= right {
			inIdx = i
			res += i - largerIdx
		} else if num < left {
			res += max(0, inIdx-largerIdx)
		} else {
			largerIdx = i
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

func numSubarrayBoundedMaxBetter(nums []int, left int, right int) int {

	count := func(a []int, target int) int {
		res, cur := 0, 0
		for _, x := range a {
			if x <= target {
				cur++
			} else {
				cur = 0
			}
			res += cur
		}
		return res
	}
	return count(nums, right) - count(nums, left-1)
}

/*
- -将范围拆分为单点计算
count(right) - count(left-1)
*/
