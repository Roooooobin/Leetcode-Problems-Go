package main

// backtracking
//func maximumBobPoints(numArrows int, aliceArrows []int) []int {
//
//	n := len(aliceArrows)
//	res := make([]int, n)
//	maxSum := 0
//	var backtracking func(idx int, cur []int, curSum, left int)
//	backtracking = func(idx int, cur []int, curSum, left int) {
//		if left < 0 {
//			return
//		}
//		if curSum > maxSum {
//			res = cur
//			t := numArrows
//			for i := 1; i < n; i++ {
//				t -= res[i]
//			}
//			res[0] = t
//			maxSum = curSum
//		}
//		if idx == n || left == 0 {
//			return
//		}
//		for i := idx; i < n; i++ {
//			curSum += i
//			nxt := make([]int, n)
//			for j := range cur {
//				nxt[j] = cur[j]
//			}
//			nxt[i] = aliceArrows[i] + 1
//			backtracking(i+1, nxt, curSum, left-nxt[i])
//			curSum -= i
//		}
//	}
//	cur := make([]int, n)
//	backtracking(0, cur, 0, numArrows)
//	return res
//}

// bit
func maximumBobPoints(numArrows int, aliceArrows []int) []int {

	n := len(aliceArrows)
	maxMask := (1 << n) - 1
	res := make([]int, n)
	maxSum := 0
	for mask := maxMask; mask >= 1; mask-- {
		sum := 0
		t := numArrows
		arr := make([]int, n)
		for i := 1; i < n; i++ {
			if (1<<i)&mask != 0 {
				sum += i
				t -= aliceArrows[i] + 1
				arr[i] = aliceArrows[i] + 1
			}
		}
		if t >= 0 {
			arr[0] = t
			if maxSum < sum {
				maxSum = sum
				res = arr
			}
		}
	}
	return res
}
