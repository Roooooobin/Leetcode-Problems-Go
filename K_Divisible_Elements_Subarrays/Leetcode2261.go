package main

func countDistinct(nums []int, k int, p int) int {

	n := len(nums)
	res := 0
	mp := make(map[string]bool)
	for i, num := range nums {
		cur := make([]byte, 0)
		cnt := 0
		if num%p == 0 {
			cnt++
		}
		cur = append(cur, byte(num+'0'))
		cur = append(cur, '-')
		key := string(cur)
		if !mp[key] {
			res++
			mp[key] = true
		}
		if cnt > k {
			continue
		}
		for j := i + 1; j < n; j++ {
			if nums[j]%p == 0 {
				cnt++
			}
			if cnt > k {
				break
			}
			cur = append(cur, byte(nums[j]+'0'))
			cur = append(cur, '-')
			key = string(cur)
			if !mp[key] {
				res++
				mp[key] = true
			}
		}
	}
	return res
}
