package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithIndex struct {
	Node  *TreeNode
	Index int
}

func widthOfBinaryTree(root *TreeNode) int {

	queue := make([]TreeNodeWithIndex, 0)
	queue = append(queue, TreeNodeWithIndex{Index: 1, Node: root})
	res := 1
	nLevel := 1
	for len(queue) > 0 {
		nLevel = len(queue)
		nums := make([]int, 0)
		for i := 0; i < nLevel; i++ {
			front := queue[0].Node
			idxFront := queue[0].Index
			queue = queue[1:]
			if front.Left != nil {
				nums = append(nums, idxFront*2)
				queue = append(queue, TreeNodeWithIndex{Node: front.Left, Index: idxFront * 2})
			}
			if front.Right != nil {
				nums = append(nums, idxFront*2+1)
				queue = append(queue, TreeNodeWithIndex{Node: front.Right, Index: idxFront*2 + 1})
			}
		}
		sort.Ints(nums)
		if len(nums) <= 1 {
			continue
		}
		res = max(res, nums[len(nums)-1]-nums[0]+1)
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
