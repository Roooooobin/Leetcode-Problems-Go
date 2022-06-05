package main

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarThree struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.NewWithIntComparator()}
}

func (this *MyCalendarThree) add(x, v int) {

	if val, ok := this.Get(x); ok {
		v += val.(int)
	}
	this.Put(x, v)
}

func (this *MyCalendarThree) Book(start int, end int) int {

	this.add(start, 1)
	this.add(end, -1)

	res := 0
	cur := 0
	for it := this.Iterator(); it.Next(); {
		cur += it.Value().(int)
		if cur > res {
			res = cur
		}
	}
	return res
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
