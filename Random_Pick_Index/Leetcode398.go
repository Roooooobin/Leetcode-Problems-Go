package main

import "math/rand"

type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this Solution) Pick(target int) int {
	cnt := 0
	res := 0
	for i, num := range this.Nums {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				res = i
			}
		}
	}
	return res
}

/*
- -蓄水池抽样法
每次以k/n的概率选择保留该值or not, n为当前第几个
如果是i(i<k)被选中, 计算在k步之前被选中 * k步之后一直不被替换的概率 = 1 * (1 - k/k+1 * 1/k) * k+1/k+2 * ... * N-1/N = k/N
如果是j(j>k)被选中, 被选中的概率k/j且之后一直不被替换的概率 = k/j * (1 - k/j+1 * 1/k) * j+1/j+2 * ... * N-1/N = k/N
*/
