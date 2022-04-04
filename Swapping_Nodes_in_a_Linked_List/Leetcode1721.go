package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {

	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	fast := head
	for i := 0; i < n-k; i++ {
		fast = fast.Next
	}
	p = head
	for i := 0; i < k-1; i++ {
		p = p.Next
	}
	tmp := fast.Val
	fast.Val = p.Val
	p.Val = tmp
	return head
}
