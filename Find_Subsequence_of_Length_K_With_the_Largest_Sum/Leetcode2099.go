package main

import "sort"

func maxSubsequence(nums []int, k int) []int {

	orig := make([]int, len(nums))
	copy(orig, nums)
	sort.Ints(nums)
	set := make(map[int]int)
	for _, x := range nums[len(nums)-k:] {
		if _, ok := set[x]; !ok {
			set[x] = 1
		} else {
			set[x]++
		}
	}
	res := make([]int, 0)
	for _, num := range orig {
		if b, ok := set[num]; ok && b > 0 {
			res = append(res, num)
			set[num]--
		}
	}
	return res
}

func main() {
	maxSubsequence([]int{50, -75, -50}, 2)
}
