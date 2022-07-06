package main

import "container/list"

/*
- -list.List用法
学习一下list.List的用法
https://leetcode-cn.com/problems/all-oone-data-structure/solution/quan-o1-de-shu-ju-jie-gou-by-leetcode-so-7gdv/
*/

type node struct {
	keys map[string]struct{}
	cnt  int
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{list.New(), map[string]*list.Element{}}
}

func (this *AllOne) Inc(key string) {
	// key已经存在于链表中
	if cur := this.nodes[key]; cur != nil {
		curNode := cur.Value.(node)
		// 如果cnt+1的节点不存在，需要新增一个
		if nxt := cur.Next(); nxt == nil || nxt.Value.(node).cnt > curNode.cnt+1 {
			this.nodes[key] = this.InsertAfter(node{keys: map[string]struct{}{key: {}}, cnt: curNode.cnt + 1}, cur)
		} else {
			nxt.Value.(node).keys[key] = struct{}{}
			this.nodes[key] = nxt
		}
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 {
			this.Remove(cur)
		}
	} else {
		if this.Front() == nil || this.Front().Value.(node).cnt > 1 {
			this.nodes[key] = this.PushFront(node{keys: map[string]struct{}{key: {}}, cnt: 1})
		} else {
			this.Front().Value.(node).keys[key] = struct{}{}
			this.nodes[key] = this.Front()
		}
	}
}

func (this *AllOne) Dec(key string) {

	cur := this.nodes[key]
	curNode := cur.Value.(node)
	if curNode.cnt > 1 {
		if pre := cur.Prev(); pre == nil || pre.Value.(node).cnt < curNode.cnt-1 {
			this.nodes[key] = this.InsertBefore(node{map[string]struct{}{key: {}}, curNode.cnt - 1}, cur)
		} else {
			pre.Value.(node).keys[key] = struct{}{}
			this.nodes[key] = pre
		}
	} else { // key 仅出现一次，将其移出 nodes
		delete(this.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		this.Remove(cur)
	}
}

func (this *AllOne) GetMaxKey() string {

	if b := this.Back(); b != nil {
		for key := range b.Value.(node).keys {
			return key
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if f := this.Front(); f != nil {
		for key := range f.Value.(node).keys {
			return key
		}
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
