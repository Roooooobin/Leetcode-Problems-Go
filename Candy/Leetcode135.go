package main

import "sort"

func candy(ratings []int) int {

	n := len(ratings)
	nodes := make([][]int, 0)
	for i, r := range ratings {
		nodes = append(nodes, []int{r, i})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i][0] < nodes[j][0]
	})
	res := 0
	placed := make([]int, n)
	for _, node := range nodes {
		r, idx := node[0], node[1]
		place := 1 // 至少要放1
		// 如果比左右大, 得比左右至少大1
		if idx > 0 {
			if r > ratings[idx-1] {
				place = max(place, placed[idx-1]+1)
			}
		}
		if idx < n-1 {
			if r > ratings[idx+1] {
				place = max(place, placed[idx+1]+1)
			}
		}
		placed[idx] = place
		res += place
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func candyBetter(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/candy/solution/fen-fa-tang-guo-by-leetcode-solution-f01p/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
- -贪心
左右两次遍历, 先从左往右, 如果比左边的rating高则比左边的大1, 然后再从右往左, 如果高则大1, 取两者的最大值
*/
