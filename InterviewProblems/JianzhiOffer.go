package main

import "container/heap"

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

func asteroidCollision(asteroids []int) []int {

	stack := make([]int, 0)
	for _, asteroid := range asteroids {
		if asteroid < 0 {
			flag := true // should add this asteroid
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if top < 0 {
					break
				}
				if top+asteroid < 0 {
					stack = stack[:len(stack)-1]
				} else if top+asteroid == 0 {
					stack = stack[:len(stack)-1]
					flag = false
					break
				} else {
					flag = false
					break
				}
			}
			if flag {
				stack = append(stack, asteroid)
			}
		} else {
			stack = append(stack, asteroid)
		}
	}
	return stack
}

func dailyTemperatures(temperatures []int) []int {

	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if t > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}

type MovingAverage struct {
	window []int
	sum    float64
	size   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {

	window := make([]int, 0)
	return MovingAverage{window, 0, size}
}

func (this *MovingAverage) Next(val int) float64 {

	if len(this.window) < this.size {
		this.window = append(this.window, val)
		this.sum += float64(val)
		return this.sum / float64(len(this.window))
	}
	this.sum -= float64(this.window[0])
	this.sum += float64(val)
	this.window = this.window[1:]
	this.window = append(this.window, val)
	return this.sum / float64(len(this.window))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxPathSum(root *TreeNode) int {

	res := -0x3f3f3f3f
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))
		res = max(res, left+right+node.Val)
		return node.Val + max(left, right)
	}
	dfs(root)
	return res
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}
	left := inorderSuccessor(root.Left, p)
	if left == nil {
		return root
	}
	return left
}

//https://leetcode.cn/problems/7WqeDu/solution/zhi-he-xia-biao-zhi-chai-du-zai-gei-ding-94ei/
/*
- -桶排序
按照值划分桶,每个桶只存一个元素,如果当前值x对应的桶已经有说明找到了,如果没有找上一个和下一个桶并且判断值差是否满足. 桶中的值只保留k范围内
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	getID := func(x int) int {
		if x >= 0 {
			return x / (t + 1)
		}
		return (x+1)/(t+1) - 1
	}
	mp := make(map[int]int)
	for i, x := range nums {
		id := getID(x)
		if _, ok := mp[id]; ok {
			return true
		}
		if y, ok := mp[id-1]; ok && abs(x-y) <= t {
			return true
		}
		if y, ok := mp[id+1]; ok && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k]))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	n1, n2 := len(nums1), len(nums2)
	hp := &maxHeap{nil, nums1, nums2}
	for i := 0; i < n1 && i < k; i++ {
		hp.data = append(hp.data, Node{i, 0})
	}
	res := make([][]int, 0)
	for hp.Len() > 0 && len(res) < k {
		top := heap.Pop(hp).(Node)
		i, j := top.i, top.j
		res = append(res, []int{nums1[i], nums2[j]})
		if j+1 < n2 {
			heap.Push(hp, Node{i, j + 1})
		}
	}
	return res
}

type Node struct {
	i, j int
}
type maxHeap struct {
	data         []Node
	nums1, nums2 []int
}

func (h maxHeap) Len() int { return len(h.data) }
func (h maxHeap) Less(i, j int) bool {
	return h.nums1[h.data[i].i]+h.nums2[h.data[i].j] < h.nums1[h.data[j].i]+h.nums2[h.data[j].j]
}
func (h maxHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *maxHeap) Push(v interface{}) { h.data = append(h.data, v.(Node)) }
func (h *maxHeap) Pop() interface{} {
	a := h.data
	n := len(a)
	v := a[n-1]
	h.data = a[:n-1]
	return v
}

func searchInsert(nums []int, target int) int {

	lo, hi := 0, len(nums)-1
	if target > nums[hi] {
		return len(nums)
	}
	if target < nums[lo] {
		return lo
	}
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] >= target {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return lo
}

func generateParenthesis(n int) []string {

	res := make([]string, 0)
	var backtracking func(cur string, left, right int)
	backtracking = func(cur string, left, right int) {
		if left+right == 2*n {
			res = append(res, cur)
		}
		if left < n {
			backtracking(cur+"(", left+1, right)
		}
		if left > right {
			backtracking(cur+")", left, right+1)
		}
	}
	backtracking("", 0, 0)
	return res
}

func coinChange(coins []int, amount int) int {

	INF := 0x3f3f3f3f
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == INF {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
