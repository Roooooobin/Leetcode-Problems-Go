package main

func findMin(a []int) int {

	n := len(a)
	lo, hi := 0, n-1
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if a[mi] > a[hi] {
			lo = mi + 1
		} else if a[mi] < a[hi] {
			hi = mi
		} else {
			hi--
		}
	}
	return a[lo]
}

/*
- -二分
与最右的元素比较, 注意如果相等那么只能排除hi这一个元素, hi--之后继续在[lo, hi]之间找
*/
