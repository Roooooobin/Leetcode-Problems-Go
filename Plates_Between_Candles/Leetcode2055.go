package main

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {

	n := len(s)
	q := len(queries)
	candlePos := make([][]int, 0)
	idx := 0
	for idx < n && s[idx] == '*' {
		idx++
	}
	res := make([]int, q)
	if idx == n {
		return res
	}
	for idx < n {
		for idx < n && s[idx] == '|' {
			idx++
		}
		if idx == n {
			break
		}
		l := idx
		for idx < n && s[idx] == '*' {
			idx++
		}
		if idx < n {
			candlePos = append(candlePos, []int{l, idx - 1})
		}
	}
	if len(candlePos) == 0 {
		return res
	}
	// 3 4, 6 8, 10 10
	left := make([]int, n)
	right := make([]int, n)
	for j := 0; j < candlePos[0][0]; j++ {
		left[j] = -1
		right[j] = 0
	}
	for i, p := range candlePos {
		if i == 0 {
			for j := p[0]; j <= p[1]; j++ {
				left[j] = i - 1
				right[j] = i + 1
			}
		} else {
			for j := candlePos[i-1][1] + 1; j < p[0]; j++ {
				left[j] = i - 1
				right[j] = i
			}
			for j := p[0]; j <= p[1]; j++ {
				left[j] = i - 1
				right[j] = i + 1
			}
		}
	}
	for i := candlePos[len(candlePos)-1][1] + 1; i < n; i++ {
		left[i] = len(candlePos) - 1
		right[i] = len(candlePos)
	}
	prefixSum := make([]int, len(candlePos))
	for i := range candlePos {
		if i == 0 {
			prefixSum[i] = candlePos[i][1] - candlePos[i][0] + 1
		} else {
			prefixSum[i] = prefixSum[i-1] + candlePos[i][1] - candlePos[i][0] + 1
		}
	}
	//for i, v := range prefixSum {
	//	fmt.Printf("%d %d\n", i, v)
	//}
	for i, query := range queries {
		l := left[query[1]]
		r := right[query[0]]
		print(l, r)
		rv := 0
		if r == 0 {
			rv = 0
		} else {
			rv = prefixSum[r-1]
		}
		if l >= 0 && r >= 0 && l >= r {
			res[i] = prefixSum[l] - rv
		}
	}
	for i, v := range res {
		fmt.Printf("%d %d\n", i, v)
	}

	return res
}

func main() {

	s := "**|**|***|*|"
	//s := "******|*|"
	queries := [][]int{{1, 10}, {2, 5}}
	platesBetweenCandles(s, queries)
}

/*
"***|**|*****|**||**|*"
[[1,17],[4,5],[14,17],[5,11],[15,16]]
*/
