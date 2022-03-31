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
