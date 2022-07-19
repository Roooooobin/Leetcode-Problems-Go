package main

func maxSubArray(nums []int) int {

	sum := 0
	res := nums[0]
	for _, num := range nums {
		if sum+num < num {
			sum = 0
		}
		sum += num
		res = max(res, sum)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
