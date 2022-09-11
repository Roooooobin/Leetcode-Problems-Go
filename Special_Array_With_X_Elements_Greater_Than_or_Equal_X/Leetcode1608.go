package main

func specialArray(nums []int) int {

	n := len(nums)
	for i := 0; i <= n; i++ {
		cnt := 0
		for _, num := range nums {
			if num >= i {
				cnt++
			}
		}
		if cnt == i {
			return i
		}
	}
	return -1
}
