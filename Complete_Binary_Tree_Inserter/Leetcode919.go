package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{root}
}

func (this *CBTInserter) Insert(val int) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, this.root)
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front.Left == nil {
			front.Left = &TreeNode{val, nil, nil}
			return front.Val
		} else {
			queue = append(queue, front.Left)
		}
		if front.Right == nil {
			front.Right = &TreeNode{val, nil, nil}
			return front.Val
		} else {
			queue = append(queue, front.Right)
		}
	}
	return 0
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
