package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) (res int) {

	var preOrder func(node *TreeNode, path int)
	preOrder = func(node *TreeNode, path int) {
		if node == nil {
			return
		}
		path = path ^ (1 << node.Val)
		if node.Left == nil && node.Right == nil {
			if (path)&(path-1) == 0 {
				res++
			}
		}
		preOrder(node.Left, path)
		preOrder(node.Right, path)
	}
	preOrder(root, 0)
	return
}

/*
- -前序遍历, 利用位运算判断
最多只有一个元素的频率是奇数, 用位1<<node.Val记录频率
*/
