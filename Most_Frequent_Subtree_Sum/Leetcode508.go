package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {

	hash := make(map[int]int)
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		sum := left + right + node.Val
		v, _ := hash[sum]
		hash[sum] = v + 1
		return sum
	}
	dfs(root)
	maxm := 0
	for _, v := range hash {
		maxm = max(maxm, v)
	}
	res := make([]int, 0)
	for k, v := range hash {
		if v == maxm {
			res = append(res, k)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
