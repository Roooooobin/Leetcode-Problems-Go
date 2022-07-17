package Champagne_Tower

func champagneTower(poured int, query_row int, query_glass int) float64 {

	// fill all glass 1, 3, 7, 15, 31...2^(row+1)-1
	N := 103
	a := make([][]float64, N)
	for i := 0; i < N; i++ {
		a[i] = make([]float64, N)
	}
	a[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			v := (a[i][j] - 1) / 2.0
			if v > 0 {
				a[i+1][j] += v
				a[i+1][j+1] += v
			}
		}
	}
	return min(1.0, a[query_row][query_glass])
}

/*
- -模拟
模拟往下流的过程
*/

func min(x, y float64) float64 {
	if x < y {
		return x
	} else {
		return y
	}
}
