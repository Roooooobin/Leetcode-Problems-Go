package main

func divideArray(nums []int) bool {

	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}
	for _, num := range hash {
		if hash[num]%2 == 1 {
			return false
		}
	}
	return true
}
