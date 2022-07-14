package main

import (
	"fmt"
	"sort"
)

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

func reversePairs(nums []int) int {

	n := len(nums)
	a := make([]int, 0, 2*n)
	for _, num := range nums {
		a = append(a, num, num*2)
	}
	sort.Ints(a)
	k := 1
	rankOrigMap := map[int]int{a[0]: k}
	for i := 1; i < 2*n; i++ {
		if a[i] != a[i-1] {
			k++
			rankOrigMap[a[i]] = k
		}
	}

	ft := FenwickTree{
		n:    k,
		tree: make([]int, k+1),
	}
	res := 0
	for i := 0; i < n; i++ {
		q := ft.query(rankOrigMap[nums[i]*2])
		res += i - q
		ft.update(rankOrigMap[nums[i]], 1)
	}
	return res
}

func main() {
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
}

/*
- -树状数组
逆序对, 也可以用归并排序的思想做
*/
