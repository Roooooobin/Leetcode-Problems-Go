package main

func maximumUniqueSubarray(nums []int) int {

	count := make([]bool, 10007)
	res, cur := 0, 0
	for r, l := 0, 0; r < len(nums); r++ {
		cur += nums[r]
		for l < r && count[nums[r]] {
			count[nums[l]] = false
			cur -= nums[l]
			l++
		}
		res = max(res, cur)
		count[nums[r]] = true
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
