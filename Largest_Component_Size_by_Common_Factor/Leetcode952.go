package main

func largestComponentSize(nums []int) (res int) {

	m := 0
	for _, num := range nums {
		m = max(m, num)
	}
	par := make([]int, m+1)
	for i := range par {
		par[i] = i
	}
	find := func(x int) int {
		r := x
		for r != par[r] {
			r = par[r]
		}
		for x != r {
			tmp := par[x]
			par[x] = r
			x = tmp
		}
		return r
	}
	union := func(x, y int) {
		x = find(x)
		y = find(y)
		if x != y {
			par[x] = y
		}
	}
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				union(num, i)
				union(num, num/i)
			}
		}
	}
	cnt := make([]int, m+1)
	for _, num := range nums {
		r := find(num)
		cnt[r]++
		res = max(res, cnt[r])
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
