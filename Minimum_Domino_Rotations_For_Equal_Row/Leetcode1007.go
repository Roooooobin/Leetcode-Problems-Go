package main

func minDominoRotations(tops []int, bottoms []int) int {

	topSetMap := make(map[int]map[int]struct{})
	bottomSetMap := make(map[int]map[int]struct{})
	n := len(tops)
	for i := range tops {
		m, ok := topSetMap[tops[i]]
		if !ok {
			m = make(map[int]struct{})
		}
		m[i] = struct{}{}
		topSetMap[tops[i]] = m
		m2, ok := bottomSetMap[bottoms[i]]
		if !ok {
			m2 = make(map[int]struct{})
		}
		m2[i] = struct{}{}
		bottomSetMap[bottoms[i]] = m2
	}
	res := n + 1
	for num := range topSetMap {
		diff := fully(topSetMap[num], bottomSetMap[num], n)
		if diff != -1 {
			res = min(res, diff)
		}
	}
	if res == n+1 {
		return -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func fully(m1, m2 map[int]struct{}, n int) int {

	cnt1, cnt2 := 0, 0
	for i := 0; i < n; i++ {
		_, ok1 := m1[i]
		_, ok2 := m2[i]
		if !ok1 && !ok2 {
			return -1
		}
		if ok1 {
			cnt1++
		}
		if ok2 {
			cnt2++
		}
	}
	return min(n-cnt1, n-cnt2)
}

func main() {
	minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2})
}

func minDominoRotationsBetter(tops []int, bottoms []int) int {

	a := make([]int, 7)
	b := make([]int, 7)
	tot := make([]int, 7)
	for i := 1; i < 7; i++ {
		tot[i] = len(tops)
	}
	for i := range tops {
		if tops[i] == bottoms[i] {
			tot[tops[i]]--
		} else {
			a[tops[i]]++
			b[bottoms[i]]++
		}
	}
	res := len(tops) + 1
	for i := 1; i < 7; i++ {
		if a[i]+b[i] >= tot[i] {
			res = min(res, min(a[i], b[i]))
		}
	}
	if res == len(tops)+1 {
		return -1
	}
	return res
}

/*
- -贪心
直接记录1-6的个数, 上下加起来数量够, 取6个点数中上下最小值
*/
