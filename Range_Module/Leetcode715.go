package main

import "github.com/emirpasic/gods/trees/redblacktree"

type RangeModule struct {
	*redblacktree.Tree
}

func Constructor() RangeModule {
	return RangeModule{redblacktree.NewWithIntComparator()}
}

func (this RangeModule) AddRange(left, right int) {
	if node, ok := this.Floor(left); ok {
		r := node.Value.(int)
		if r >= right {
			return
		}
		if r >= left {
			left = node.Key.(int)
			this.Remove(left)
		}
	}
	for node, ok := this.Ceiling(left); ok && node.Key.(int) <= right; node, ok = this.Ceiling(left) {
		right = max(right, node.Value.(int))
		this.Remove(node.Key)
	}
	this.Put(left, right)
}

func (this RangeModule) QueryRange(left, right int) bool {
	node, ok := this.Floor(left)
	return ok && node.Value.(int) >= right
}

func (this RangeModule) RemoveRange(left, right int) {
	if node, ok := this.Floor(left); ok {
		l, r := node.Key.(int), node.Value.(int)
		if r >= right {
			if l == left {
				this.Remove(l)
			} else {
				node.Value = left
			}
			if right != r {
				this.Put(right, r)
			}
			return
		}
		if r > left {
			node.Value = left
		}
	}
	for node, ok := this.Ceiling(left); ok && node.Key.(int) < right; node, ok = this.Ceiling(left) {
		r := node.Value.(int)
		this.Remove(node.Key)
		if r > right {
			this.Put(right, r)
			break
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/range-module/solution/range-mo-kuai-by-leetcode-solution-4utf/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
