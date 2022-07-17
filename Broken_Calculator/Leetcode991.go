package main

func brokenCalc(startValue int, target int) int {

	res := 0
	for target > startValue {
		if target%2 == 1 {
			target = (target + 1) / 2
			res += 2
		} else {
			target /= 2
			res++
		}
	}
	return res + startValue - target
}

/*
- -贪心
反向从target开始, 总是优先做除2的操作, 如果是奇数先做加法再除法
*/
// BFS TLE
//seen := make(map[int]struct{})
//queue := make([][]int, 0)
//queue = append(queue, []int{startValue, 0})
//for len(queue) > 0 {
//	cur := queue[0]
//	if cur[0] == target {
//		return cur[1]
//	}
//	queue = queue[1:]
//	if cur[0] > 2*target {
//		continue
//	}
//	nxt := cur[0] * 2
//	if _, ok := seen[nxt]; !ok {
//		seen[nxt] = struct{}{}
//		queue = append(queue, []int{nxt, cur[1] + 1})
//	}
//	nxt = cur[0] - 1
//	if _, ok := seen[nxt]; !ok {
//		seen[nxt] = struct{}{}
//		queue = append(queue, []int{nxt, cur[1] + 1})
//	}
//}
//return -1
