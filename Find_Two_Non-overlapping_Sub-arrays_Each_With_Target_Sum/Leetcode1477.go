package main

type Node struct {
	end    int
	length int
}

func minSumOfLengths(arr []int, target int) int {

	mp := make(map[int]int)
	mp[0] = -1
	sum := 0
	nodes := make([]Node, 0)
	INF := 0x3f3f3f3f
	res := INF
	prefixMin := make([]int, 0)
	binarySearch := func(a []Node, target int) int {
		lo, hi := 0, len(a)-1
		if target > a[hi].end {
			return hi
		}
		if target < a[lo].end {
			return -1
		}
		for lo <= hi {
			mi := lo + (hi-lo)>>1
			if a[mi].end == target {
				return mi
			} else if a[mi].end < target {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		}
		return lo
	}
	for i, x := range arr {
		sum += x
		mp[sum] = i
		if v, ok := mp[sum-target]; ok {
			curLen := i - v
			if len(nodes) > 0 {
				idx := binarySearch(nodes, v)
				if idx != -1 {
					res = min(res, prefixMin[idx]+curLen)
				}
			}
			nodes = append(nodes, Node{i, curLen})
			if len(prefixMin) == 0 {
				prefixMin = append(prefixMin, curLen)
			} else {
				prefixMin = append(prefixMin, min(curLen, prefixMin[len(prefixMin)-1]))
			}
		}
	}
	if res == INF {
		return -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	minSumOfLengths([]int{7, 3, 4, 7}, 7)
}

/*
- -二分, 哈希记录窗口和
用哈希记录窗口和,找到区间[x1, y1], [x2, y2]和为target, 用找到区间的左端点x,在前面所有右端点数组中二分,同时记录前缀最小值,加和得到的
即是这次区间所能形成最小的区间对
*/
