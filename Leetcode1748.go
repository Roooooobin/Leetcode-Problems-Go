package main

func sumOfUnique(nums []int) int {

	m := make(map[int]int)
	for _, num := range nums {
		if v, ok := m[num]; ok {
			m[num] = v + 1
		} else {
			m[num] = 1
		}
	}
	res := 0
	for _, num := range nums {
		if m[num] == 1 {
			res += num
		}
	}
	return res
}
