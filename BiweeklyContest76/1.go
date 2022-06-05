package main

func findClosestNumber(nums []int) int {

	dis, res := 1000000, -1000000
	for _, num := range nums {
		if abs(num) < dis {
			dis = abs(num)
			res = num
		} else if abs(num) == dis && num > res {
			res = num
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
