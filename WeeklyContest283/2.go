package main

import "sort"

func minimalKSum(nums []int, k int) int64 {

	sort.Ints(nums)
	res := int64(0)
	if nums[0] > 1 {
		n := min(k, nums[0]-1)
		res += (int64(1) + int64(n)) * int64(n) / 2
		k -= n
	}
	idx := 0
	for idx < len(nums)-1 && k > 0 {
		if nums[idx] == nums[idx+1] {
			idx++
			continue
		}
		n := min(k, nums[idx+1]-nums[idx]-1)
		res += (int64(nums[idx]+1) + int64(nums[idx]+n)) * int64(n) / 2
		k -= n
		idx++
	}
	if k > 0 {
		res += (int64(nums[len(nums)-1]+1) + int64(nums[len(nums)-1]+k)) * int64(k) / 2
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

func main() {

	// 794
	print(minimalKSum([]int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}, 35))
}
