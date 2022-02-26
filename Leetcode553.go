package main

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {

	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	sb := &strings.Builder{}
	sb.Grow(10)
	sb.WriteString(strconv.Itoa(nums[0]))
	sb.WriteByte('/')
	sb.WriteByte('(')
	sb.WriteString(strconv.Itoa(nums[1]))
	for i := 2; i < n; i++ {
		sb.WriteByte('/')
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	sb.WriteByte(')')
	return sb.String()
}
