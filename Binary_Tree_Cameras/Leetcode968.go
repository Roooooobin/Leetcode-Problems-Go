package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCoverBetter(root *TreeNode) int {

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

/*
- -后序遍历
后序遍历延迟安排摄像头
分情况讨论, 如果左儿子右儿子其中一个为0, 表示必须安排一个了, 如果有某一个为1, 表示能被辐射到, 但是本身自己没有摄像头
最后还要单独判断root
*/
