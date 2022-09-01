package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {

	res := 0
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {

		if node == nil {
			return 0
		}
		l := helper(node.Left)
		r := helper(node.Right)
		l1, r1 := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			l1 = l + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			r1 = r + 1
		}
		res = max(res, l1+r1)
		return max(l1, r1)
	}
	helper(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
