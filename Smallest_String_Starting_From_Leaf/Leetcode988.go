package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {

	res := "~"
	var dfs func(node *TreeNode, s []byte)
	dfs = func(node *TreeNode, s []byte) {
		if node == nil {
			return
		}
		s = append(s, byte(node.Val+'a'))
		if node.Left == nil && node.Right == nil {
			cur := reverse(s)
			if cur < res {
				res = cur
			}
			reverse(s)
			return
		}
		dfs(node.Left, s)
		dfs(node.Right, s)
		s = s[:len(s)-1]
	}
	dfs(root, []byte{})
	return res
}

func reverse(s []byte) string {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return string(s)
}
