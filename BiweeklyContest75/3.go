package main

func numberOfWays(s string) int64 {

	n := len(s)
	num1 := int64(0)
	num10 := int64(0)
	num101 := int64(0)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			num10 += num1
		} else if s[i] == '1' {
			num1++
			num101 += num10
		}
	}
	num0 := int64(0)
	num01 := int64(0)
	num010 := int64(0)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			num01 += num0
		} else if s[i] == '0' {
			num0++
			num010 += num01
		}
	}
	return num101 + num010
}
