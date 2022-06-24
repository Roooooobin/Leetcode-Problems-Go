package main

import (
	"math"
)

//func minimumTime(time []int, totalTrips int) int64 {
//
//	n := len(time)
//	if n == 1 {
//		return int64(time[0] * totalTrips)
//	}
//	idx := make([]int, n)
//	for i, _ := range idx {
//		idx[i] = 1
//	}
//	res := int64(0)
//	sum := 0
//	for true {
//		min_ := int64(math.MaxInt64)
//		for i := 0; i < n; i++ {
//			min_ = min(min_, int64(idx[i]*time[i]))
//		}
//		for i := 0; i < n; i++ {
//			if int64(idx[i]*time[i]) == min_ {
//				idx[i]++
//				sum++
//			}
//		}
//		if sum >= totalTrips {
//			res = min_
//			break
//		}
//	}
//	return res
//}
//
//func min(x, y int64) int64 {
//	if x < y {
//		return x
//	} else {
//		return y
//	}
//}

func minimumTime(time []int, totalTrips int) int64 {

	lo, hi := int64(1), int64(math.MaxInt64)
	total := int64(totalTrips)
	for lo <= hi {
		total = int64(totalTrips)
		mid := lo + (hi-lo)/2
		for _, t := range time {
			total -= mid / int64(t)
			if total <= 0 {
				hi = mid - 1
				break
			}
		}
		if total > 0 {
			lo = mid + 1
		}
	}
	return lo
}

//func main() {
//
//	res := minimumTime([]int{39, 82, 69, 37, 78, 14, 93, 36, 66, 61, 13, 58, 57, 12, 70, 14, 67, 75, 91, 1, 34, 68, 73, 50, 13, 40, 81, 21, 79, 12, 35, 18, 71, 43, 5, 50, 37, 16, 15, 6, 61, 7, 87, 43, 27, 62, 95, 45, 82, 100, 15, 74, 33, 95, 38, 88, 91, 47, 22, 82, 51, 19, 10, 24, 87, 38, 5, 91, 10, 36, 56, 86, 48, 92, 10, 26, 63, 2, 50, 88, 9, 83, 20, 42, 59, 55, 8, 15, 48, 25}, 4187)
//	//res := minimumTime([]int{2, 3}, 3) //[]int{39, 82, 69, 37, 78, 14, 93, 36, 66, 61, 13, 58, 57, 12, 70, 14, 67, 75, 91, 1, 34, 68, 73, 50, 13, 40, 81, 21, 79, 12, 35, 18, 71, 43, 5, 50, 37, 16, 15, 6, 61, 7, 87, 43, 27, 62, 95, 45, 82, 100, 15, 74, 33, 95, 38, 88, 91, 47, 22, 82, 51, 19, 10, 24, 87, 38, 5, 91, 10, 36, 56, 86, 48, 92, 10, 26, 63, 2, 50, 88, 9, 83, 20, 42, 59, 55, 8, 15, 48, 25}, 4187)
//	fmt.Println(res)
//}
