package main

//https://leetcode.cn/problems/naming-a-company/solution/by-endlesscheng-ruz8/

func distinctNames(ideas []string) (res int64) {

	group := [26]map[string]bool{}
	for i := range group {
		group[i] = make(map[string]bool)
	}
	for _, s := range ideas {
		group[s[0]-'a'][s[1:]] = true
	}
	for i, a := range group {
		for j := 0; j < i; j++ {
			b := group[j]
			cnt := 0
			for s := range a {
				if b[s] {
					cnt++
				}
			}
			res += int64(len(a)-cnt) * int64(len(b)-cnt)
		}
	}
	return res * 2
}

/*
- -分组讨论
首字母分组, 当前组有m个,另一个组里n个,如果有后缀相同的cnt个,那么res+=(m-cnt)*(n-cnt)
*/
