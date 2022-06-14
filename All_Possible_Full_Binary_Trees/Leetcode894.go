package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {

	res := make([]*TreeNode, 0)
	if n&1 == 0 {
		return res
	}
	if n == 1 {
		res = append(res, &TreeNode{0, nil, nil})
	}
	for i := 1; i < n; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)
		for _, lt := range left {
			for _, rt := range right {
				root := &TreeNode{0, nil, nil}
				root.Left = lt
				root.Right = rt
				res = append(res, root)
			}
		}
	}
	return res
}
