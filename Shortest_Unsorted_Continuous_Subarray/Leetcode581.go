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

/*
- -贪心
分成三段, 中间段虽然无序但是满足最小值大于左段最大值, 最大值小于右段最小值
从左到右维护一个最大值max,在进入右段之前，那么遍历到的nums[i]都是小于max的，我们要求的end就是遍历中最后一个小于max元素的位置；
同理，从右到左维护一个最小值min，在进入左段之前，那么遍历到的nums[i]也都是大于min的，要求的begin也就是最后一个大于min元素的位置。

作者：xmblgt
链接：https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/solution/si-lu-qing-xi-ming-liao-kan-bu-dong-bu-cun-zai-de-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
