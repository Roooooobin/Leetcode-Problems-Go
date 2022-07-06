package main

import "container/list"

type LRUCache struct {
	l        *list.List
	hash     map[int]*list.Element
	keyValue map[int]int
	capacity int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	hash := make(map[int]*list.Element)
	m := make(map[int]int)
	return LRUCache{l, hash, m, capacity}
}

func (this *LRUCache) Get(key int) int {

	if _, ok := this.keyValue[key]; !ok {
		return -1
	}
	this.l.Remove(this.hash[key])
	this.l.PushFront(key)
	this.hash[key] = this.l.Front()
	return this.keyValue[key]
}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.keyValue[key]; ok {
		this.keyValue[key] = value
		this.l.Remove(this.hash[key])
		this.l.PushFront(key)
		this.hash[key] = this.l.Front()
		return
	}
	if this.l.Len() >= this.capacity {
		de := this.l.Back().Value.(int)
		this.l.Remove(this.l.Back())
		delete(this.hash, de)
		delete(this.keyValue, de)
	}
	this.l.PushFront(key)
	this.hash[key] = this.l.Front()
	this.keyValue[key] = value
}

/*
- -LRU, 链表+哈希表
Put也要更新频率
*/
