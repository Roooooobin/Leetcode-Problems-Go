package main

func rearrangeArray(nums []int) []int {

	pots := make([]int, 0)
	negs := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			negs = append(negs, num)
		} else {
			pots = append(pots, num)
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i&1 == 0 {
			res = append(res, pots[i/2])
		} else {
			res = append(res, negs[i/2])
		}
	}
	return res
}
