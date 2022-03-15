package main

import (
	"container/list"
	"fmt"
)

type nodeLRU struct {
	*list.List
	nodesLRU map[int]*list.Element
}

type nodeLFU struct {
	lru *nodeLRU
	cnt int
}

type LFUCache struct {
	capacity int
	*list.List
	nodesLFU map[int]*list.Element
	k2v      map[int]int
	total    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity, list.New(), map[int]*list.Element{}, map[int]int{}, 0}
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.k2v[key]; !ok {
		return -1
	}
	this.Update(key)
	return this.k2v[key]
}

func (this *LFUCache) Update(key int) {
	cur := this.nodesLFU[key]
	curNode := cur.Value.(*nodeLFU)
	if nxt := cur.Next(); nxt == nil || nxt.Value.(*nodeLFU).cnt > curNode.cnt+1 {
		newL := list.New()
		newL.PushFront(key)
		this.nodesLFU[key] = this.InsertAfter(&nodeLFU{&nodeLRU{newL,
			map[int]*list.Element{key: newL.Front()}}, curNode.cnt + 1}, cur)
	} else {
		lru := nxt.Value.(*nodeLFU).lru
		lru.PushFront(key)
		lru.nodesLRU[key] = lru.List.Front()
		nxt.Value.(*nodeLFU).lru = lru
		this.nodesLFU[key] = nxt
	}
	curNode.lru.Remove(curNode.lru.nodesLRU[key])
	delete(curNode.lru.nodesLRU, key)
	if len(curNode.lru.nodesLRU) == 0 {
		this.Remove(cur)
	}
}

func (this *LFUCache) Put(key int, value int) {
	// 如果是新key
	if this.capacity == 0 {
		return
	}
	if _, ok := this.k2v[key]; !ok {
		// 超过容量，需要T掉LFU的头里的LRU的尾
		if this.total >= this.capacity {
			f := this.Front()
			fNode := f.Value.(*nodeLFU)
			lru := fNode.lru
			backKey := lru.List.Back().Value.(int)
			lru.Remove(lru.Back())
			delete(lru.nodesLRU, backKey)
			delete(this.k2v, backKey)
			if len(fNode.lru.nodesLRU) == 0 {
				this.Remove(f)
			}
		} else {
			this.total++
		}

		if this.Front() == nil {
			newL := list.New()
			newL.PushFront(key)
			this.nodesLFU[key] = this.PushFront(&nodeLFU{&nodeLRU{newL,
				map[int]*list.Element{key: newL.Front()}}, 1})
		} else {
			front := this.Front()
			frontNode := front.Value.(*nodeLFU)
			if frontNode.cnt > 1 {
				newL := list.New()
				newL.PushFront(key)
				this.nodesLFU[key] = this.PushFront(&nodeLFU{&nodeLRU{newL,
					map[int]*list.Element{key: newL.Front()}}, 1})
			} else {
				lru := front.Value.(*nodeLFU).lru
				lru.PushFront(key)
				lru.nodesLRU[key] = lru.List.Front()
				front.Value.(*nodeLFU).lru = lru
				this.nodesLFU[key] = front
			}
		}

	} else {
		this.Update(key)
	}
	this.k2v[key] = value
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	obj.Put(4, 4)
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
