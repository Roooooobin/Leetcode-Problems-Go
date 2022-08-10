package main

import (
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {

	if len(s) <= 2 {
		return s
	}
	subs := sort.StringSlice{}
	cnt, l := 0, 0
	for i, ch := range s {
		if ch == '1' {
			cnt++
		} else if cnt--; cnt == 0 {
			subs = append(subs, "1"+makeLargestSpecial(s[l+1:i])+"0")
			l = i + 1
		}
	}
	sort.Sort(sort.Reverse(subs))
	return strings.Join(subs, "")
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/special-binary-string/solution/te-shu-de-er-jin-zhi-xu-lie-by-leetcode-sb7ry/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
