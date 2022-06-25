package main

func findMaximumXOR(nums []int) (x int) {

	const highBit = 30
	for k := highBit; k >= 0; k-- {
		seen := make(map[int]bool)
		for _, num := range nums {
			seen[num>>k] = true
		}
		xNext := x*2 + 1
		flag := false
		for _, num := range nums {
			if seen[num>>k^xNext] {
				flag = true
				break
			}
		}
		if flag {
			x = xNext
		} else {
			x = xNext - 1
		}
	}
	return
}

//https://leetcode.cn/problems/ms70jA/solution/zui-da-de-yi-huo-by-leetcode-solution-hr7m/

/*
字典树记录各个二进制
*/
