package main

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {

	n := len(s)
	par := make([]int, n)
	for i := range par {
		par[i] = i
	}
	find := func(x int) int {
		r := x
		for par[r] != r {
			r = par[r]
		}
		for x != r {
			tmp := par[x]
			par[x] = r
			x = tmp
		}
		return r
	}
	combine := func(x, y int) {
		x = find(x)
		y = find(y)
		if x != y {
			par[x] = y
		}
	}
	for _, pair := range pairs {
		combine(pair[0], pair[1])
	}
	groups := make([][]int, n)
	for i := range groups {
		groups[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		groups[find(i)] = append(groups[find(i)], i)
	}

	countingSort := func(groups [][]int) []rune {
		res := make([]rune, n)

		for _, group := range groups {
			if len(group) == 0 {
				continue
			}
			a := make([]int, 0)
			for _, idx := range group {
				a = append(a, int(s[idx]-'a'))
			}
			sort.Ints(a)
			for i, idx := range group {
				res[idx] = rune(a[i] + 'a')
			}
		}
		return res
	}
	res := countingSort(groups)
	return string(res)
}

func main() {
	print(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {1, 3}}))
}
