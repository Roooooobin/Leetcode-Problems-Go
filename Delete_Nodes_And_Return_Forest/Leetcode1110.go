package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodesSlow(root *TreeNode, to_delete []int) []*TreeNode {

	set := make(map[int]struct{})
	for _, d := range to_delete {
		set[d] = struct{}{}
	}
	res := make([]*TreeNode, 0)
	if _, ok := set[root.Val]; !ok {
		res = append(res, root)
	}
	var dfs func(node *TreeNode, par *TreeNode, isLeft bool)
	dfs = func(node *TreeNode, par *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		v := node.Val
		lc := node.Left
		rc := node.Right
		if _, ok := set[v]; ok {
			if par != nil {
				if isLeft {
					par.Left = nil
				} else {
					par.Right = nil
				}
			}
			if node.Left != nil {
				res = append(res, node.Left)
			}
			if node.Right != nil {
				res = append(res, node.Right)
			}
			node = nil
		}
		dfs(lc, node, true)
		dfs(rc, node, false)
	}
	dfs(root, nil, true)
	finalRes := make([]*TreeNode, 0)
	for _, node := range res {
		if _, ok := set[node.Val]; !ok {
			finalRes = append(finalRes, node)
		}
	}
	return finalRes
}

// better
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	set := make(map[int]struct{})
	for _, d := range to_delete {
		set[d] = struct{}{}
	}
	res := make([]*TreeNode, 0)
	if _, ok := set[root.Val]; !ok {
		res = append(res, root)
	}
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		if _, ok := set[node.Val]; ok {
			if node.Left != nil {
				res = append(res, node.Left)
			}
			if node.Right != nil {
				res = append(res, node.Right)
			}
			node = nil
		}
		return node
	}
	dfs(root)
	return res
}
