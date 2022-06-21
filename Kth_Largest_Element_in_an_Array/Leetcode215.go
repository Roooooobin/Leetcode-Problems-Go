package main

func findKthLargest(nums []int, k int) int {

	n := len(nums)
	lo, hi := 0, n-1
	k = n - k
	partition := func(l, r int) int {
		i, j := l, r
		nums[i], nums[i+(j-i)>>1] = nums[i+(j-i)>>1], nums[i]
		pivot := nums[i]
		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}
			for i < j && nums[i] <= pivot {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = pivot
		return i
	}
	for true {
		idx := partition(lo, hi)
		if idx == k {
			return nums[idx]
		} else if idx < k {
			lo = idx + 1
		} else {
			hi = idx - 1
		}
	}
	return 0
}
