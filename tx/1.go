package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{1, &ListNode{0, &ListNode{1, &ListNode{1, nil}}}}
	b := &ListNode{0, &ListNode{1, &ListNode{0, &ListNode{1, nil}}}}
	c := xorList(a, b)
	fmt.Println(c)
}

func xorList(a *ListNode, b *ListNode) *ListNode {
	// write code here
	// reverse two lists first
	a = reverseList(a)
	c := &ListNode{0, nil}
	p := c
	for a != nil || b != nil {
		var av, bv int
		if a != nil {
			av = a.Val
			a = a.Next
		}
		if b != nil {
			bv = b.Val
			b = b.Next
		}
		p.Next = &ListNode{av ^ bv, nil}
		p = p.Next
	}
	c = c.Next
	p = reverseList(c)
	for p != nil {
		if p.Val == 0 {
			p = p.Next
		} else {
			break
		}
	}
	if p == nil {
		p = &ListNode{0, nil}
	}
	return p
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}
