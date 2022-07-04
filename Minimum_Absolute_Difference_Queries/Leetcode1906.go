package main

func minDifference(nums []int, queries [][]int) []int {

	bucket := make([][]int, 101)
	for i := range bucket {
		bucket[i] = make([]int, 0)
	}
	for i, num := range nums {
		bucket[num] = append(bucket[num], i)
	}
	// 用来判断区间是否重叠, 返回为负表示没有找到, 返回插入位置, 为正表示找到, 用l, r分别返回之后如果不同为负且相等,则表示重叠
	binarySearch := func(a []int, t int) int {
		lo, hi := 0, len(a)-1
		if t < a[lo] {
			return -hi - 23
		}
		if t > a[hi] {
			return -hi - 10
		}
		for lo <= hi {
			mi := lo + (hi-lo)>>1
			if a[mi] < t {
				lo = mi + 1
			} else if a[mi] == t {
				return mi
			} else {
				hi = mi - 1
			}
		}
		return -lo
	}
	res := make([]int, len(queries))
	INF := 0x3f3f3f3f
	for i, q := range queries {
		diff, pre := INF, -1
		l, r := q[0], q[1]
		for b := 1; b < 101; b++ {
			if len(bucket[b]) == 0 {
				continue
			}
			idx1 := binarySearch(bucket[b], l)
			idx2 := binarySearch(bucket[b], r)
			if idx1 < 0 && idx2 < 0 && idx1 == idx2 {
				continue
			}
			if pre != -1 {
				diff = min(diff, b-pre)
			}
			pre = b
		}
		if diff == INF {
			diff = -1
		}
		res[i] = diff
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
	//minDifference([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}})
	minDifference([]int{4, 2, 15, 1, 14, 6, 2, 9, 4, 5}, [][]int{{0, 2}})
}

/*
- -桶排序
用[1,100]的桶记录当前值的索引数组, 对于每个查询[l,r], 判断是否与其索引数组有重叠部分,如果有表明这个查询中包含了该值,枚举最小diff
*/
