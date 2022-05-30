package main

func minDeletion(nums []int) int {

	n := len(nums)
	if n <= 1 {
		return n
	}
	res := 0
	l, r := 0, 1
	for r < n {
		for r < n && nums[l] == nums[r] {
			res++
			r++
		}
		if r == n {
			res++
			break
		}
		l = r + 1
		r += 2
	}
	if l == n-1 {
		res++
	}
	return res
}
