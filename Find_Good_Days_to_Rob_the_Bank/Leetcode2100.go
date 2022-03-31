package main

func goodDaysToRobBank(security []int, time int) []int {

	n := len(security)
	res := make([]int, 0)
	if time == 0 {
		for i := range security {
			res = append(res, i)
		}
		return res
	}
	forward := make(map[int]struct{})
	i := 0
	for i < n-1 {
		cnt := 0
		for i < n-1 {
			if security[i+1] <= security[i] {
				cnt++
				if cnt >= time {
					forward[i+1] = struct{}{}
				}
			} else {
				break
			}
			i++
		}
		i++
	}
	i = n - 1
	for i >= 1 {
		cnt := 0
		for i >= 1 {
			if security[i-1] <= security[i] {
				cnt++
				if cnt >= time {
					if _, ok := forward[i-1]; ok {
						res = append(res, i-1)
					}
				}
			} else {
				break
			}
			i--
		}
		i--
	}
	return res
}
