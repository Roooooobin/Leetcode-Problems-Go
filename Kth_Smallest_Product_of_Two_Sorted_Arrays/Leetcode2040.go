package main

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {

	neg1, neg2, pos1, pos2 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	for _, x := range nums1 {
		if x < 0 {
			neg1 = append(neg1, x)
		} else {
			pos1 = append(pos1, x)
		}
	}
	for _, x := range nums2 {
		if x < 0 {
			neg2 = append(neg2, x)
		} else {
			pos2 = append(pos2, x)
		}
	}
	lo, hi := int64(-1e10), int64(1e10)
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		cur := int64(0)
		for i, j := 0, len(pos2)-1; i < len(pos1); i++ {
			for j >= 0 && int64(pos1[i])*int64(pos2[j]) > mi {
				j--
			}
			cur += int64(j + 1)
		}
		for i, j := 0, 0; i < len(neg1); i++ {
			for j < len(pos2) && int64(neg1[i])*int64(pos2[j]) > mi {
				j++
			}
			cur += int64(len(pos2) - j)
		}
		for i, j := 0, 0; i < len(pos1); i++ {
			for j < len(neg2) && int64(pos1[i])*int64(neg2[j]) <= mi {
				j++
			}
			cur += int64(j)
		}
		for i, j := 0, len(neg2)-1; i < len(neg1); i++ {
			for j >= 0 && int64(neg1[i])*int64(neg2[j]) <= mi {
				j--
			}
			cur += int64(len(neg2) - j - 1)
		}
		if cur < k {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return lo
}

/*
二分
求第k个: 可能是利用heap; 也可能是利用二分尝试一个答案, check与k的关系
*/

//https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/solution/yi-ti-san-jie-shuang-zhi-zhen-jie-bu-den-sqsu/
