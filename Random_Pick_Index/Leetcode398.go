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
