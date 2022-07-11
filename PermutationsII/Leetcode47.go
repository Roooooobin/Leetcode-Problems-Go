package main

import "sort"

func permuteUnique(nums []int) (res [][]int) {

	sort.Ints(nums)
	n := len(nums)
	cur := make([]int, 0)
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			res = append(res, append([]int(nil), cur...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			cur = append(cur, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/permutations-ii/solution/quan-pai-lie-ii-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
/*
- -回溯
去重挺难想
*/
