package main

func subarraySum(nums []int, k int) int {

	hash := make(map[int]int)
	sum := 0
	hash[0] = 1
	res := 0
	for _, num := range nums {
		sum += num
		res += hash[sum-k]
		v := hash[sum]
		hash[sum] = v + 1
	}
	return res
}

func findMaxLength(nums []int) int {

	hash := make(map[int]int)
	diff := 0
	hash[0] = -1
	res := 0
	for i, num := range nums {
		if num == 1 {
			diff++
		} else {
			diff--
		}
		if _, ok := hash[diff]; ok {
			res = max(res, i-hash[diff])
		} else {
			hash[diff] = i
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
