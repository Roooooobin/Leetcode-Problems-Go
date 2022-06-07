package main

func maxTrailingZeros(grid [][]int) int {

	calcTwo := func(x int) int {
		for i := 9; i >= 1; i-- {
			if x%(1<<i) == 0 {
				return i
			}
		}
		return 0
	}
	calcFive := func(x int) int {
		five := 625
		for i := 4; i >= 1; i-- {
			if x%five == 0 {
				return i
			}
			five /= 5
		}
		return 0
	}

	m, n := len(grid), len(grid[0])
	two := make([][]int, m)
	for i := range two {
		two[i] = make([]int, n)
	}
	five := make([][]int, m)
	for i := range five {
		five[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			two[i][j] = calcTwo(grid[i][j])
			five[i][j] = calcFive(grid[i][j])
		}
	}
	preTwoRow := make([][]int, m)
	for i := range preTwoRow {
		preTwoRow[i] = make([]int, n+1)
	}
	preFiveRow := make([][]int, m)
	for i := range preFiveRow {
		preFiveRow[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preTwoRow[i][j+1] = preTwoRow[i][j] + two[i][j]
			preFiveRow[i][j+1] = preFiveRow[i][j] + five[i][j]
		}
	}
	res := 0
	for j := 0; j < n; j++ {
		cnt2, cnt5 := 0, 0
		for i := range grid {
			cnt2 += two[i][j]
			cnt5 += five[i][j]
			res = max(res, max(min(cnt2+preTwoRow[i][j], cnt5+preFiveRow[i][j]), min(cnt2+preTwoRow[i][n]-preTwoRow[i][j+1], cnt5+preFiveRow[i][n]-preFiveRow[i][j+1])))
		}
		cnt2, cnt5 = 0, 0
		for i := m - 1; i >= 0; i-- {
			cnt2 += two[i][j]
			cnt5 += five[i][j]
			res = max(res, max(min(cnt2+preTwoRow[i][j], cnt5+preFiveRow[i][j]), min(cnt2+preTwoRow[i][n]-preTwoRow[i][j+1], cnt5+preFiveRow[i][n]-preFiveRow[i][j+1])))
		}
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
