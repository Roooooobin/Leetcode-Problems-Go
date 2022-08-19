package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}
	maxIdx := findMaxIndex(nums, left, right)
	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = helper(nums, left, maxIdx-1)
	root.Right = helper(nums, maxIdx+1, right)
	return root
}

func findMaxIndex(nums []int, left, right int) int {

	max, maxIdx := nums[left], left
	for i := left; i <= right; i++ {
		if nums[i] > max {
			maxIdx = i
			max = nums[i]
		}
	}
	return maxIdx
}
