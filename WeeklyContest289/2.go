package main

func minimumRounds(tasks []int) int {

	hash := make(map[int]int)
	for _, task := range tasks {
		if _, ok := hash[task]; ok {
			hash[task]++
		} else {
			hash[task] = 1
		}
	}
	res := 0
	for _, v := range hash {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			res += v / 3
		} else {
			res += v/3 + 1
		}
	}
	return res
}
