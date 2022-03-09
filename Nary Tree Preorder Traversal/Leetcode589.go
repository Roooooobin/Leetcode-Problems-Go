package Nary_Tree_Preorder_Traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (res []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return
}
