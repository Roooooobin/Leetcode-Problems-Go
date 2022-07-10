package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	sb := &strings.Builder{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("nil,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	a := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if a[0] == "nil" {
			a = a[1:]
			return nil
		}
		val, _ := strconv.Atoi(a[0])
		a = a[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

/*
- -前序遍历
前序遍历用,隔开同时记录nil,反序列化时先处理左子树再处理右子树, 每次拿出第一个元素
*/
