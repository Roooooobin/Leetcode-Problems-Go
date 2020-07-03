package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r + 1) / 2
	node := &TreeNode{nums[mid], nil, nil}
	node.Left = helper(nums, l, mid-1)
	node.Right = helper(nums, mid+1, r)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}
