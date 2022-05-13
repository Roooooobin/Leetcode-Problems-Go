package main

// O(nlogn)
//func findUnsortedSubarray(nums []int) int {
//
//	n := len(nums)
//	a := make([]int, n)
//	for i, num := range nums {
//		a[i] = num
//	}
//	sort.Ints(a)
//	l, r := 0, n-1
//	for ; l < n; l++ {
//		if a[l] != nums[l] {
//			break
//		}
//	}
//	if l == n-1 {
//		return 0
//	}
//	for ; r >= l; r-- {
//		if a[r] != nums[r] {
//			break
//		}
//	}
//	return r - l + 1
//}

// O(n)
func findUnsortedSubarray(nums []int) int {

	n := len(nums)
	l, r := 0, -1
	min, max := nums[n-1], nums[0]
	for i := 0; i < n; i++ {
		if nums[i] < max {
			r = i
		} else {
			max = nums[i]
		}
		if nums[n-i-1] > min {
			l = n - i - 1
		} else {
			min = nums[n-i-1]
		}
	}
	return r - l + 1
}
