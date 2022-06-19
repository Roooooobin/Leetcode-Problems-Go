package main

func subarraySum(nums []int, k int) int {

	hash := make(map[int]int)
	sum := 0
	hash[0] = 1
	res := 0
	for _, num := range nums {
		sum += num
		res += hash[sum-k]
		v := hash[sum]
		hash[sum] = v + 1
	}
	return res
}

func findMaxLength(nums []int) int {

	hash := make(map[int]int)
	diff := 0
	hash[0] = -1
	res := 0
	for i, num := range nums {
		if num == 1 {
			diff++
		} else {
			diff--
		}
		if _, ok := hash[diff]; ok {
			res = max(res, i-hash[diff])
		} else {
			hash[diff] = i
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func countSubstrings(s string) int {

	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	res := n
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
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
