package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		var ano = target - num
		if j, ok := hash[ano]; ok {
			var ans = []int{j, i}
			return ans
		} else {
			hash[num] = i
		}
	}
	return []int{0, 0}
}
