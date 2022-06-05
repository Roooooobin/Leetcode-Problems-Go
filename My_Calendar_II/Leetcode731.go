package main

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarTwo struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{redblacktree.NewWithIntComparator()}
}

func (this *MyCalendarTwo) add(x, v int) {
	if val, ok := this.Get(x); ok {
		v += val.(int)
	}
	this.Put(x, v)
}

func (this *MyCalendarTwo) Book(start int, end int) bool {

	this.add(start, 1)
	this.add(end, -1)
	cur := 0
	for it := this.Iterator(); it.Next(); {
		cur += it.Value().(int)
		if cur > 2 {
			this.add(start, -1)
			this.add(end, 1)
			if v, ok := this.Get(start); ok && v == 0 {
				this.Remove(start)
			}
			if v, ok := this.Get(end); ok && v == 0 {
				this.Remove(end)
			}
			return false
		}
	}
	return true
}

func main() {
	ins := Constructor()
	ins.Book(37, 50)
	ins.Book(33, 50)
	ins.Book(4, 17)
	ins.Book(35, 48)
	ins.Book(8, 25)
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
