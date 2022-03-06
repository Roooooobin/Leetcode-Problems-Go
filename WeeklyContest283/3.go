package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {

	nodes := make(map[int][]int)
	inDegree := make([]int, 100003)
	for _, description := range descriptions {
		node1 := description[0]
		if _, ok := nodes[node1]; !ok {
			nodes[node1] = make([]int, 2)
		}
		if description[2] == 1 {
			nodes[node1][0] = description[1]
		} else {
			nodes[node1][1] = description[1]
		}
		inDegree[description[1]]++
	}
	queue := make([]*TreeNode, 0)
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, &TreeNode{Val: node})
		}
	}
	head := queue[0]
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if _, ok := nodes[front.Val]; !ok {
			continue
		}
		if nodes[front.Val][0] != 0 {
			left := &TreeNode{Val: nodes[front.Val][0]}
			front.Left = left
			queue = append(queue, left)
		}
		if nodes[front.Val][1] != 0 {
			right := &TreeNode{Val: nodes[front.Val][1]}
			front.Right = right
			queue = append(queue, right)
		}
	}
	return head
}

func main() {

	createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}})
}
