package main

func numSubmatrixSumTarget(a [][]int, target int) int {

	m, n := len(a), len(a[0])
	ps := make([][]int, m+1)
	for i := range ps {
		ps[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ps[i][j] = ps[i-1][j] + ps[i][j-1] - ps[i-1][j-1] + a[i-1][j-1]
		}
	}
	res := 0
	for b := 1; b <= m; b++ {
		for t := 1; t <= b; t++ {
			cur := 0
			hash := make(map[int]int)
			for j := 1; j <= n; j++ {
				cur = ps[b][j] - ps[t-1][j]
				if cur == target {
					res++
				}
				if v, ok := hash[cur-target]; ok {
					res += v
				}
				v := hash[cur]
				hash[cur] = v + 1
			}
		}
	}
	return res
}

/*
- -二维前缀和+哈希优化
*/
