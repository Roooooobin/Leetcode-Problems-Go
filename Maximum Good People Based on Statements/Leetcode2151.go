package main

import (
	"fmt"
	"math/bits"
)

func maximumGood(statements [][]int) int {

	n := len(statements)
	m := 1 << n
	res := 0
	for mask := 1; mask < m; mask++ {
		goodSet := make(map[int]struct{})
		for j := 1; j <= n; j++ {
			if 1&(mask>>(j-1)) != 0 {
				goodSet[j] = struct{}{}
			}
		}
		flag := true
		for goodPerson1 := range goodSet {
			for j := 1; j <= n; j++ {
				if _, ok := goodSet[j]; ok && statements[goodPerson1-1][j-1] == 0 {
					flag = false
					break
				}
				if _, ok := goodSet[j]; !ok && statements[goodPerson1-1][j-1] == 1 {
					flag = false
					break
				}
			}
			if !flag {
				break
			}
		}
		if flag {
			if res < len(goodSet) {
				res = len(goodSet)
			}
		}
	}
	return res
}

func maximumGood2(statements [][]int) int {
	n := len(statements)
	ans := 0
	for mask := 1; mask < (1 << n); mask++ {
		check := func() bool {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if i == j {
						continue
					}
					if statements[i][j] == 0 {
						if ((mask & (1 << i)) > 0) && ((mask & (1 << j)) > 0) {
							return false
						}
					} else if statements[i][j] == 1 {
						if ((mask & (1 << i)) > 0) && ((mask & (1 << j)) == 0) {
							return false
						}
					}
				}
			}
			return true
		}

		if check() {
			ans = max(ans, bits.OnesCount(uint(mask)))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/maximum-good-people-based-on-statements/solution/ji-yu-chen-shu-tong-ji-zui-duo-hao-ren-s-lfn9/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {

	fmt.Println(maximumGood([][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}))
}
