package main

func splitArray(nums []int, m int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if m == 1 {
		return sum
	}
	n := len(nums)

	helper := func(g int) bool {
		cur := 0
		cnt := 0
		i := 0
		for ; i < n; i++ {
			if nums[i] > g {
				return false
			}
			if cur+nums[i] > g {
				cnt++
				cur = 0
				if cnt == m-1 {
					break
				}
			}
			cur += nums[i]
		}
		cur = 0
		for ; i < n; i++ {
			cur += nums[i]
		}
		if cur <= g {
			return true
		}
		return false
	}
	lo, hi := 0, sum
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if helper(mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

/*
- -二分答案+贪心
*/
