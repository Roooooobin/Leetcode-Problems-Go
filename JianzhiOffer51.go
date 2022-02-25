package main

import "sort"

func reversePairs(nums []int) int {

	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}

	ft := FenwickTree{
		n:    n,
		tree: make([]int, n+1),
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		res += ft.query(nums[i] - 1)
		ft.update(nums[i], 1)
	}
	return res
}

type FenwickTree struct {
	n    int
	tree []int
}

func (ft FenwickTree) lowBit(x int) int {
	return x & (-x)
}

func (ft FenwickTree) query(i int) int {

	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= ft.lowBit(i)
	}
	return sum
}

func (ft FenwickTree) update(i, x int) {
	for i <= ft.n {
		ft.tree[i] += x
		i += ft.lowBit(i)
	}
}
