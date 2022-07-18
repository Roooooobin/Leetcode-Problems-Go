package main

func goodDaysToRobBank(security []int, time int) []int {

	n := len(security)
	res := make([]int, 0)
	if time == 0 {
		for i := range security {
			res = append(res, i)
		}
		return res
	}
	forward := make(map[int]struct{})
	i := 0
	for i < n-1 {
		cnt := 0
		for i < n-1 {
			if security[i+1] <= security[i] {
				cnt++
				if cnt >= time {
					forward[i+1] = struct{}{}
				}
			} else {
				break
			}
			i++
		}
		i++
	}
	i = n - 1
	for i >= 1 {
		cnt := 0
		for i >= 1 {
			if security[i-1] <= security[i] {
				cnt++
				if cnt >= time {
					if _, ok := forward[i-1]; ok {
						res = append(res, i-1)
					}
				}
			} else {
				break
			}
			i--
		}
		i--
	}
	return res
}

/*
- -DP
我的做法是前后两次遍历, 复杂度有点高
*/

func goodDaysToRobBankBetter(security []int, time int) (ans []int) {
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
		if security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/find-good-days-to-rob-the-bank/solution/gua-he-da-jie-yin-xing-de-ri-zi-by-leetc-z6r1/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
