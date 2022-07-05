package main

func longestConsecutive(nums []int) int {

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	res := 1
	for x := range set {
		if set[x-1] {
			continue
		}
		y := x + 1
		for set[y] {
			y++
		}
		res = max(res, y-x)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -哈希集合
先将所有数放入set, 然后遍历set找到连续序列, 每次只从开头开始, 即如果x-1存在于集合中则跳过
*/
