package main

func isPossible(nums []int) bool {

	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	endCnt := make(map[int]int)
	for _, num := range nums {
		if cnt[num] == 0 {
			continue
		}
		// 直接加入num-1的序列
		if endCnt[num-1] > 0 {
			endCnt[num]++
			endCnt[num-1]--
			cnt[num]--
			// 新建一个至少为3的序列
		} else if cnt[num+1] > 0 && cnt[num+2] > 0 {
			cnt[num]--
			cnt[num+1]--
			cnt[num+2]--
			endCnt[num+2]++
		} else {
			return false
		}
	}
	return true
}

/*
- -贪心
一个数v, 如果存在以v-1结尾的子序列, 那么加入这个子序列要比再新建一个更优, 以此贪心
*/
