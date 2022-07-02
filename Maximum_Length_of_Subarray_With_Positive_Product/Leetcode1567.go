package main

func getMaxLen(nums []int) int {

	nums = append(nums, 0)
	n := len(nums)
	i := 0
	first, last := -1, -1
	cnt := 0
	res, left := 0, 0
	for i < n {
		if nums[i] < 0 {
			cnt++
			if first == -1 {
				first = i
			}
			last = i
		} else if nums[i] == 0 {
			if cnt&1 == 1 {
				res = max(res, max(i-first-1, last-left))
			} else {
				res = max(res, i-left)
			}
			cnt = 0
			left, first, last = i+1, -1, -1
		}
		i++
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -贪心?逻辑题
如果遇到0(idx:i)就分隔做一次计算,区间[left, i)只用记录第一个负数位置first和最后一个负数位置last和负数的个数cnt
如果cnt是偶数,直接res = max(res, i-left); 如果是奇数, 为了形成正数乘积,只能出现偶数个负数,所以取res = max(res, max(i-first-1, last-left))
*/
