package main

import (
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {

	bytes := []byte(strconv.Itoa(n))
	if len(bytes) == 1 {
		return -1
	}
	var i int
	for i = len(bytes) - 1; i >= 1; i-- {
		if bytes[i] > bytes[i-1] {
			break
		}
	}
	if i == 0 {
		return -1
	}
	var j int
	for j = len(bytes) - 1; j >= i; j-- {
		if bytes[j] > bytes[i-1] {
			break
		}
	}
	bytes[i-1], bytes[j] = bytes[j], bytes[i-1]
	// from i to j, reverse
	j = len(bytes) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	s := string(bytes)
	res, _ := strconv.Atoi(s)
	if res > math.MaxInt32 {
		return -1
	}
	return res
}

/*
- -双指针
从后往前找到第一个降序对(i-1, i),然后从[i, len-1]中找到第一个j,a[j]>a[i-1]交换,然后reverse a[i:]
*/
