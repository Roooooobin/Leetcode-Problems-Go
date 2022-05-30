package main

func findDifference(nums1 []int, nums2 []int) [][]int {

	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}
	res1 := make([]int, 0)
	res2 := make([]int, 0)
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			res1 = append(res1, k)
		}
	}
	for k := range set2 {
		if _, ok := set1[k]; !ok {
			res2 = append(res2, k)
		}
	}
	return [][]int{res1, res2}
}
