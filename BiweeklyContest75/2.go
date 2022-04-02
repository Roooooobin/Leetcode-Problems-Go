package main

func triangularSum(nums []int) int {

	t := nums
	for len(t) != 1 {
		newT := make([]int, len(t)-1)
		for i := 0; i < len(t)-1; i++ {
			newT[i] = (t[i] + t[i+1]) % 10
		}
		t = newT
	}
	return t[0]
}
