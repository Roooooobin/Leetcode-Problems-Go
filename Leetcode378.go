package main

func kthSmallest(matrix [][]int, k int) int {
	var n, m = len(matrix), len(matrix[0])
	var l, r = matrix[0][0], matrix[n-1][m-1]
	var mid int
	for l < r {
		mid = (l + r) >> 1
		var cnt int
		var j = n - 1
		for i := 0; i < n; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			cnt += j + 1
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

/*
- -归并排序
还可以用归并排序配合小根堆
func kthSmallest(matrix [][]int, k int) int {
    h := &IHeap{}
    for i := 0; i < len(matrix); i++ {
        heap.Push(h, [3]int{matrix[i][0], i, 0})
    }

    for i := 0; i < k - 1; i++ {
        now := heap.Pop(h).([3]int)
        if now[2] != len(matrix) - 1 {
            heap.Push(h, [3]int{matrix[now[1]][now[2]+1], now[1], now[2]+1})
        }
    }
    return heap.Pop(h).([3]int)[0]
}

type IHeap [][3]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/solution/you-xu-ju-zhen-zhong-di-kxiao-de-yuan-su-by-leetco/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
