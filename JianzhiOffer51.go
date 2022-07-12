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

func reversePairsBetter(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}

/*
- -归并思想, 树状数组
*/

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/solution/shu-zu-zhong-de-ni-xu-dui-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
