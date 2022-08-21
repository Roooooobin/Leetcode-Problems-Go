package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calDepth(node *TreeNode) int {

	height := 0
	if node.Left != nil {
		height = calDepth(node.Left) + 1
	}
	if node.Right != nil {
		height = max(height, calDepth(node.Right)+1)
	}
	return height
}

func printTree(root *TreeNode) [][]string {

	height := calDepth(root)
	m := height + 1
	n := 1<<m - 1
	res := make([][]string, m)
	for i := range res {
		res[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		res[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
