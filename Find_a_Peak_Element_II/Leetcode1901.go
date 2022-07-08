package main

func findPeakGrid(mat [][]int) []int {

	n := len(mat[0])
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		row := getColMax(mat, mi)
		if mi == 0 {
			if mat[row][mi] > mat[row][mi+1] {
				return []int{row, mi}
			} else {
				lo = mi + 1
			}
		} else if mi == n-1 {
			if mat[row][mi] > mat[row][mi-1] {
				return []int{row, mi}
			} else {
				hi = mi - 1
			}
		} else {
			if mat[row][mi] > mat[row][mi+1] && mat[row][mi] > mat[row][mi-1] {
				return []int{row, mi}
			} else if mat[row][mi] < mat[row][mi+1] {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		}
	}
	return []int{}
}

func getColMax(nums [][]int, col int) int {
	res := 0
	resRow := 0
	for i, a := range nums {
		if res < a[col] {
			res = a[col]
			resRow = i
		}
	}
	return resRow
}

/*
- -二分hard
找到列的最大值, 往左右二分
*/
