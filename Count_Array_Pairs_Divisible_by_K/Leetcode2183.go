package main

func countPairs(nums []int, K int) int64 {

	gcdMap := make(map[int]int, 0)
	res := int64(0)
	for _, num := range nums {
		g := gcd(num, K)
		for k, v := range gcdMap {
			if int64(g)*int64(k)%int64(K) == 0 {
				res += int64(v)
			}
		}
		v, ok := gcdMap[g]
		if !ok {
			gcdMap[g] = 1
		} else {
			gcdMap[g] = v + 1
		}
	}
	return res
}

func gcd(x, y int) int {

	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
