package main

import (
	"math"
	"sort"
)

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

//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//
//	n1, n2 := len(nums1), len(nums2)
//	hp := &maxHeap{nil, nums1, nums2}
//	for i := 0; i < n1 && i < k; i++ {
//		hp.data = append(hp.data, Node{i, 0})
//	}
//	res := make([][]int, 0)
//	for hp.Len() > 0 && len(res) < k {
//		top := heap.Pop(hp).(Node)
//		i, j := top.i, top.j
//		res = append(res, []int{nums1[i], nums2[j]})
//		if j+1 < n2 {
//			heap.Push(hp, Node{i, j + 1})
//		}
//	}
//	return res
//}
//
//type Node struct {
//	i, j int
//}
//type maxHeap struct {
//	data         []Node
//	nums1, nums2 []int
//}
//
//func (h maxHeap) Len() int { return len(h.data) }
//func (h maxHeap) Less(i, j int) bool {
//	return h.nums1[h.data[i].i]+h.nums2[h.data[i].j] < h.nums1[h.data[j].i]+h.nums2[h.data[j].j]
//}
//func (h maxHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
//func (h *maxHeap) Push(v interface{}) { h.data = append(h.data, v.(Node)) }
//func (h *maxHeap) Pop() interface{} {
//	a := h.data
//	n := len(a)
//	v := a[n-1]
//	h.data = a[:n-1]
//	return v
//}

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

// 二叉树中序遍历
func inorderTraversal(root *TreeNode) (res []int) {

	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// 二叉树前序遍历
func preorderTraversal(root *TreeNode) (res []int) {

	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func postorderTraversal(root *TreeNode) (res []int) {

	if root == nil {
		return
	}
	var stack []*TreeNode
	cur := root
	stack = append(stack, cur)
	stack = append(stack, cur)
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 尚未遍历左右子树
		if len(stack) > 0 && cur == stack[len(stack)-1] {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
				stack = append(stack, cur.Left)
			}
		} else {
			res = append(res, cur.Val)
		}
	}
	return
}

func postorderTraversalBetter(root *TreeNode) (res []int) {

	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果没有右子树或者已经访问过了, 则可以访问当前节点
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}

func oddCells(m, n int, indices [][]int) (res int) {

	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	for _, p := range indices {
		for j := range a[p[0]] {
			a[p[0]][j]++
		}
		for _, row := range a {
			row[p[1]]++
		}
	}
	for _, row := range a {
		for _, v := range row {
			res += v % 2
		}
	}
	return
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(q1, q2 *Node) *Node {
	if q1.IsLeaf {
		if q1.Val {
			return &Node{Val: true, IsLeaf: true}
		}
		return q2
	}
	if q2.IsLeaf {
		return intersect(q2, q1)
	}
	o1 := intersect(q1.TopLeft, q2.TopLeft)
	o2 := intersect(q1.TopRight, q2.TopRight)
	o3 := intersect(q1.BottomLeft, q2.BottomLeft)
	o4 := intersect(q1.BottomRight, q2.BottomRight)
	if o1.IsLeaf && o2.IsLeaf && o3.IsLeaf && o4.IsLeaf && o1.Val == o2.Val && o1.Val == o3.Val && o1.Val == o4.Val {
		return &Node{Val: o1.Val, IsLeaf: true}
	}
	return &Node{false, false, o1, o2, o3, o4}
}

func pairSums(nums []int, target int) (res [][]int) {

	sort.Ints(nums)
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		if nums[lo]+nums[hi] == target {
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
		} else if nums[lo]+nums[hi] > target {
			hi--
		} else {
			lo++
		}
	}
	return
}

func search(nums []int, target int) int {

	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		if nums[lo] == target {
			return lo
		}
		mi := lo + (hi-lo)>>1
		if nums[mi] == target {
			hi = mi
		} else if nums[mi] > nums[0] {
			if target < nums[mi] && target >= nums[0] {
				hi = mi - 1
			} else {
				lo = mi + 1
			}
		} else if nums[mi] < nums[0] {
			if target > nums[mi] && target <= nums[n-1] {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		} else {
			lo++
		}
	}
	return -1
}

func findString(a []string, s string) int {

	lo, hi := 0, len(a)-1
	for lo <= hi {
		for lo <= hi && a[lo] == "" {
			lo++
		}
		for lo <= hi && a[hi] == "" {
			hi--
		}
		mi := lo + (hi-lo)>>1
		for mi <= hi && a[mi] == "" {
			mi++
		}
		if a[mi] == s {
			return mi
		}
		if a[mi] < s {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return -1
}

/*
https://leetcode.cn/problems/sparse-array-search-lcci
- -二分变式
*/

func searchMatrix(a [][]int, target int) bool {

	if len(a) == 0 {
		return false
	}
	m, n := len(a), len(a[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if a[i][j] == target {
			return true
		}
		if a[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func smallestDifference(a []int, b []int) int {

	sort.Ints(a)
	sort.Ints(b)
	i, j := 0, 0
	res := math.MaxInt32
	for i < len(a) && j < len(b) {
		res = min(res, abs(a[i]-b[j]))
		if a[i] == b[j] {
			return 0
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
https://leetcode.cn/problems/smallest-difference-lcci/
- -排序+双指针
*/

func maxAliveYear(birth []int, death []int) int {

	a := make([]int, 102)
	for _, y := range birth {
		a[y-1900]++
	}
	for _, y := range death {
		a[y-1900+1]--
	}
	maxN := 0
	maxY := -1
	ps := 0
	for i := 0; i < len(a); i++ {
		ps += a[i]
		if ps > maxN {
			maxY = i + 1900
			maxN = ps
		}
	}
	return maxY
}

func findSwapValues(a1 []int, a2 []int) []int {

	sum1 := 0
	sum2 := 0
	for _, x := range a1 {
		sum1 += x
	}
	mp := make(map[int]bool)
	for _, x := range a2 {
		sum2 += x
		mp[x] = true
	}
	diff := sum2 - sum1
	if diff&1 == 1 {
		return []int{}
	}
	diff >>= 1
	for _, x := range a1 {
		if mp[x+diff] {
			return []int{x, x + diff}
		}
	}
	return []int{}
}

func minSubsequence(nums []int) []int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	tot := 0
	for _, num := range nums {
		tot += num
	}
	for i, sum := 0, 0; ; i++ {
		sum += nums[i]
		if sum > tot-sum {
			return nums[:i+1]
		}
	}
}

func missingTwo(nums []int) []int {
	xorSum := 0
	n := len(nums) + 2
	for _, num := range nums {
		xorSum ^= num
	}
	for i := 1; i <= n; i++ {
		xorSum ^= i
	}
	kBit := xorSum & -xorSum
	// use kBit divides into two groups
	group1, group2 := 0, 0
	for _, num := range nums {
		if num&kBit > 0 {
			group1 ^= num
		} else {
			group2 ^= num
		}
	}
	for i := 1; i <= n; i++ {
		if i&kBit > 0 {
			group1 ^= i
		} else {
			group2 ^= i
		}
	}
	return []int{group1, group2}
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/missing-two-lcci/solution/xiao-shi-de-liang-ge-shu-zi-by-leetcode-zuwq3/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func getKthMagicNumber(k int) int {

	dp := make([]int, k)
	idx3, idx5, idx7 := 0, 0, 0
	dp[0] = 1
	for i := 1; i < k; i++ {
		dp[i] = min(dp[idx3]*3, min(dp[idx5]*5, dp[idx7]*7))
		if dp[i] == dp[idx3]*3 {
			idx3++
		}
		if dp[i] == dp[idx5]*5 {
			idx5++
		}
		if dp[i] == dp[idx7]*7 {
			idx7++
		}
	}
	return dp[k-1]
}
