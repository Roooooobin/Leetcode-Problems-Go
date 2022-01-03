package main

func findJudge(n int, trust [][]int) int {

	//in := make([]int, n+1)
	//out := make([]int, n+1)
	degree := make([]int, n+1)
	for _, a := range trust {
		degree[a[0]]--
		degree[a[1]]++
	}
	for i := 1; i <= n; i++ {
		if degree[i] == n-1 {
			return i
		}
	}
	return -1
}
