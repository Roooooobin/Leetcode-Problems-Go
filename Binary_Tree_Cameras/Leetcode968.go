package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCoverGood(root *TreeNode) int {

	res := 0
	var postTraversal func(node *TreeNode) int
	postTraversal = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		lc := postTraversal(node.Left)
		rc := postTraversal(node.Right)
		if lc == 0 || rc == 0 {
			res++
			return 1
		}
		if lc == 1 || rc == 1 {
			return 2
		}
		return 0
	}
	if postTraversal(root) == 0 {
		res++
	}
	return res
}

func minCameraCover(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return 1
	}
	mark := make(map[*TreeNode]int)
	mark[nil] = 2
	res := 0
	var postTraversal func(node *TreeNode)
	postTraversal = func(node *TreeNode) {

		if node == nil {
			return
		}
		postTraversal(node.Left)
		postTraversal(node.Right)
		lc := node.Left
		rc := node.Right
		if mark[lc] == 0 || mark[rc] == 0 {
			mark[node] = 1
			res++
			return
		}
		if mark[lc] == 1 || mark[rc] == 1 {
			mark[node] = 2
		}
		return
	}
	postTraversal(root)
	if mark[root] == 0 {
		res++
	}
	return res
}
