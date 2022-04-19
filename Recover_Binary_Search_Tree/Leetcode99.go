package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var p1, p2, pre *TreeNode
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		if pre != nil && pre.Val > node.Val {
			if p1 == nil {
				p1 = pre
			}
			p2 = node
		}
		pre = node
		inOrder(node.Right)
	}
	inOrder(root)
	tmp := p1.Val
	p1.Val = p2.Val
	p2.Val = tmp
}
