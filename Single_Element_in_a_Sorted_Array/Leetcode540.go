package main

func singleNonDuplicate(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if mi == 0 {
			if nums[mi] != nums[mi+1] {
				return nums[mi]
			}
			lo = mi + 1
		} else if mi == len(nums)-1 {
			if nums[mi] != nums[mi-1] {
				return nums[mi]
			}
			hi = mi - 1
		} else {
			if mi&1 == 1 {
				if nums[mi] != nums[mi-1] {
					hi = mi - 1
				} else {
					lo = mi + 1
				}
			} else {
				if nums[mi] != nums[mi+1] {
					hi = mi - 1
				} else {
					lo = mi + 1
				}
			}
		}
	}
	return nums[lo]
}

//better
func singleNonDuplicateBetter(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

//作者：LeetCode-Solution
//链接：https: //leetcode.cn/problems/skFtm2/solution/pai-xu-shu-zu-zhong-zhi-chu-xian-yi-ci-d-jk8w/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
- -二分
根据当前下标是奇数偶数来判断二分往左还是往右
*/
