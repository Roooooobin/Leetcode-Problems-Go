package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {

	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res += abs(left) + abs(right)
		return node.Val + left + right - 1
	}
	dfs(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

// https://leetcode.cn/problems/distribute-coins-in-binary-tree/solution/zai-er-cha-shu-zhong-fen-pei-ying-bi-by-leetcode/
/*
- -递归, 树的后序遍历
每个节点当做中转站, 要中转的值是|1-val|, 后序遍历, 左右往上累加
*/
