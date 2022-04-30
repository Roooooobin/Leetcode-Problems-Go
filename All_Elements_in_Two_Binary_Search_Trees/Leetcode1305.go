package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	inOrder := make([]int, 0)
	var inOrderTraversal func(root *TreeNode)
	inOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraversal(root.Left)
		inOrder = append(inOrder, root.Val)
		inOrderTraversal(root.Right)
	}
	inOrderTraversal(root1)
	inOrderTraversal(root2)
	sort.Ints(inOrder)
	return inOrder
}
