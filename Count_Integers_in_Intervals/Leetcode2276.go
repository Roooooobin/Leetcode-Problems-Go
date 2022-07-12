package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

type CountIntervals struct {
	*redblacktree.Tree
	cnt int
}

func Constructor() CountIntervals {
	return CountIntervals{redblacktree.NewWithIntComparator(), 0}
}

func (t *CountIntervals) Add(left, right int) {
	for node, _ := t.Ceiling(left); node != nil && node.Value.(int) <= right; node, _ = t.Ceiling(left) {
		l, r := node.Value.(int), node.Key.(int)
		if l < left {
			left = l
		}
		if r > right {
			right = r
		}
		t.cnt -= r - l + 1
		t.Remove(r)
	}
	t.cnt += right - left + 1
	t.Put(right, left)
}

func (t *CountIntervals) Count() int {
	return t.cnt
}

/*
- -红黑树
*/

//作者：endlesscheng
//链接：https://leetcode.cn/problems/count-integers-in-intervals/solution/by-endlesscheng-clk2/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
//type CountIntervals struct {
//	*redblacktree.Tree
//}
//
//func Constructor() CountIntervals {
//	return CountIntervals{redblacktree.NewWithIntComparator()}
//}
//
//func (this *CountIntervals) Add(left int, right int) {
//	if v, ok := this.Get(left); ok {
//		if v == -1 {
//			this.Put(left, 0)
//		} else {
//			this.Put(left, v.(int)+1)
//		}
//	} else {
//		this.Put(left, 1)
//	}
//	if v, ok := this.Get(right); ok {
//		if v == 1 {
//			this.Put(right, 0)
//		} else {
//			this.Put(right, v.(int)-1)
//		}
//	} else {
//		this.Put(right, -1)
//	}
//}
//
//func (this *CountIntervals) Count() int {
//
//	res := 0
//	it := this.Iterator()
//	if !it.Next() {
//		return 0
//	}
//	preVal := it.Key().(int)
//	signSum := it.Value().(int)
//	if it.Value().(int) == 0 {
//		res++
//		signSum = 0
//	}
//	for it.Next() {
//		curVal, curSign := it.Key().(int), it.Value().(int)
//		if signSum == 0 {
//			if curSign == 0 {
//				res++
//				continue
//			}
//			preVal = curVal
//			signSum = 1
//			continue
//		}
//		signSum += curSign
//		if signSum == 0 {
//			res += curVal - preVal + 1
//		}
//	}
//	return res
//}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */

func main() {
	obj := Constructor()
	//fmt.Println(obj.Count())
	//obj.Add(8, 43)
	//fmt.Println(obj.Count())
	//obj.Add(13, 16)
	//fmt.Println(obj.Count())
	//obj.Add(26, 33)
	//fmt.Println(obj.Count())
	//
	//obj.Add(28, 36)
	//fmt.Println(obj.Count())
	//
	//obj.Add(29, 37)
	obj.Add(5, 10)
	obj.Add(5, 10)
	//obj.Add(6, 46)
	fmt.Println(obj.Count())
}
