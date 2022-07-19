package main

import "fmt"

func bestRotation(nums []int) int {

	n := len(nums)
	diff := make([]int, n+1)
	for i, num := range nums {
		if num == 0 {
			diff[0]++
			diff[n]--
		} else if i < num {
			diff[i+1]++
			diff[n-num+i+1]--
		} else {
			diff[0]++
			diff[i-num+1]--
			diff[i+1]++
			diff[n]--
		}
	}
	prefix := make([]int, n+1)
	max := 0
	resIdx := 0
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + diff[i]
		if prefix[i+1] > max {
			max = prefix[i+1]
			resIdx = i
		}
	}
	//for i := 1; i <= n; i++ {
	//	fmt.Println(prefix[i])
	//}
	return resIdx
}

func main() {
	fmt.Println(bestRotation([]int{2, 3, 1, 4, 0}))
}

/*
- -差分数组的巧用
*/
