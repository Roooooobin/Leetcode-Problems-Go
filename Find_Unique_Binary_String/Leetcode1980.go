package main

import "fmt"

func findDifferentBinaryString(nums []string) string {

	n := len(nums[0])
	set := make(map[int]struct{})
	calc := func(s string) int {
		x := 0
		for i := 0; i < n; i++ {
			x = x*2 + int(s[i]-'0')
		}
		return x
	}
	for _, num := range nums {
		set[calc(num)] = struct{}{}
	}
	for i := 0; i < (1 << n); i++ {
		if _, ok := set[i]; !ok {
			if i == 0 {
				s := ""
				for j := 0; j < n; j++ {
					s += "0"
				}
				return s
			}
			res := ""
			idx := 0
			for i > 0 {
				if i&1 == 1 {
					res = "1" + res
				} else {
					res = "0" + res
				}
				i >>= 1
				idx++
			}
			for j := len(res); j < n; j++ {
				res = "0" + res
			}
			return res
		}
	}
	return ""
}

func main() {
	fmt.Println(findDifferentBinaryString([]string{"11", "10", "00"}))
}
