package main

/*
https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/solution/zi-dian-xu-de-di-kxiao-shu-zi-by-leetcod-bfy0/
*/

func findKthNumber(n int, k int) int {

	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		// 往下一个兄弟节点去
		if steps <= k {
			k -= steps
			cur++
			// 往子树下去
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

// 当前cur之下有多少个数
func getSteps(cur, n int) int {

	first, last := cur, cur
	steps := 0
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return steps
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
