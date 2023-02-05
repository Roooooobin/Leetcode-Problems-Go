package main

import "sort"

func getMaximumConsecutive(coins []int) int {

	res := 1
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > res {
			break
		}
		res += coin
	}
	return res
}

/*
- -贪心
升序前序和如果超不过当前的x, 说明肯定凑不到这个x
*/
