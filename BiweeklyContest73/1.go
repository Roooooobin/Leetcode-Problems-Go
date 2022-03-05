package main

func mostFrequent(nums []int, key int) int {

	hash := make([]int, 1003)
	max := 0
	maxNum := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			hash[nums[i+1]]++
		}
	}
	for x, num := range hash {
		if num > maxNum {
			maxNum = num
			max = x
		}
	}
	return max
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	} else {
//		return y
//	}
//}
