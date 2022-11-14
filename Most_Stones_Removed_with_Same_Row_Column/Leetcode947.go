package main

func removeStones(stones [][]int) int {

	par, rank := make(map[int]int), make(map[int]int)
	var find func(int) int
	find = func(x int) int {
		if _, has := par[x]; !has {
			par[x], rank[x] = x, 1
		}
		if par[x] != x {
			par[x] = find(par[x])
		}
		return par[x]
	}
	combine := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		par[fy] = fx
	}

	const MAX = 10001
	for _, stone := range stones {
		combine(stone[0], stone[1]+MAX)
	}
	res := len(stones)
	for x, fx := range par {
		if x == fx {
			res--
		}
	}
	return res
}

/*
- -并查集
同列或同行, 在一个集合
*/
