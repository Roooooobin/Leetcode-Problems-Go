package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	hash := make(map[*Node]*Node)
	hash[nil] = nil
	return helper(head, hash)
}

func helper(head *Node, hash map[*Node]*Node) *Node {

	if head == nil {
		return nil
	}
	h := &Node{Val: head.Val}
	hash[head] = h
	if _, ok := hash[head.Random]; ok {
		h.Random = hash[head.Random]
	} else {
		h.Random = helper(head.Random, hash)
	}
	if _, ok := hash[head.Next]; ok {
		h.Next = hash[head.Next]
	} else {
		h.Next = helper(head.Next, hash)
	}
	return h
}

/*
- -回溯+哈希表
*/

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomListBetter(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

//作者：LeetCode-Solution
//链接：https: //leetcode.cn/problems/copy-list-with-random-pointer/solution/fu-zhi-dai-sui-ji-zhi-zhen-de-lian-biao-rblsf/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
