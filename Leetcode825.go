package main

func numFriendRequests(ages []int) int {

	m := make(map[int]int)
	for _, age := range ages {
		if v, ok := m[age]; !ok {
			m[age] = 1
		} else {
			m[age] = v + 1
		}
	}
	res := 0
	for x := range m {
		for y := range m {
			if doSendRequest(x, y) {
				v := 0
				if x == y {
					v = 1
				}
				res += m[x] * (m[y] - v)
			}
		}
	}
	return res
}

func doSendRequest(x, y int) bool {

	if y <= ((x >> 1) + 7) {
		return false
	}
	return y <= x
}
