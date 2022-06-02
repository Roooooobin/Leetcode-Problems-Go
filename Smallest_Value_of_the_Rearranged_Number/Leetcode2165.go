package main

import (
	"fmt"
	"sort"
)

func smallestNumber(num int64) int64 {

	if num == 0 {
		return 0
	}
	if num < 0 {
		s := make([]int, 0)
		num = -num
		for num > 0 {
			s = append(s, int(num%10))
			num /= 10
		}
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
		res := int64(0)
		for _, i := range s {
			res = res*10 + int64(i)
		}
		return -res
	} else {
		s := make([]int, 0)
		for num > 0 {
			s = append(s, int(num%10))
			num /= 10
		}
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		if s[0] == 0 {
			i := 0
			for i < len(s) {
				if s[i] != 0 {
					s[0], s[i] = s[i], s[0]
					break
				}
				i++
			}
			if i == len(s) {
				return num
			}
		}
		res := int64(0)
		for _, i := range s {
			res = res*10 + int64(i)
		}
		return res
	}
}

func main() {
	fmt.Println(smallestNumber(-30100))
	// byte to int
	s := "123"
	c := s[0]
	fmt.Println(c)
	fmt.Println(int(c - '0'))
	fmt.Println(string(2 + '0'))
}
