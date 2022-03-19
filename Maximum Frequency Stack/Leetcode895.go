package main

// TLE 31/37
//type node struct {
//	keys map[int]struct{}
//	cnt  int
//}
//
//type FreqStack struct {
//	*list.List
//	nodes   map[int]*list.Element
//	val2Idx map[int][]int
//	counter int
//}
//
//func Constructor() FreqStack {
//	return FreqStack{list.New(), map[int]*list.Element{}, map[int][]int{}, 0}
//}
//
//func (this *FreqStack) Push(val int) {
//	// val已经存在于链表中
//	this.counter++
//	if cur := this.nodes[val]; cur != nil {
//		curNode := cur.Value.(*node)
//		// 如果cnt+1的节点不存在，需要新增一个
//		if nxt := cur.Next(); nxt == nil || nxt.Value.(*node).cnt > curNode.cnt+1 {
//			this.nodes[val] = this.InsertAfter(&node{keys: map[int]struct{}{val: {}}, cnt: curNode.cnt + 1}, cur)
//		} else {
//			nxt.Value.(*node).keys[val] = struct{}{}
//			this.nodes[val] = nxt
//		}
//		delete(curNode.keys, val)
//		if len(curNode.keys) == 0 {
//			this.Remove(cur)
//		}
//		this.val2Idx[val] = append(this.val2Idx[val], this.counter)
//	} else {
//		if this.Front() == nil || this.Front().Value.(*node).cnt > 1 {
//			this.nodes[val] = this.PushFront(&node{keys: map[int]struct{}{val: {}}, cnt: 1})
//		} else {
//			this.Front().Value.(*node).keys[val] = struct{}{}
//			this.nodes[val] = this.Front()
//		}
//		this.val2Idx[val] = []int{this.counter}
//	}
//}
//
//func (this *FreqStack) Pop() int {
//
//	cur := this.Back()
//	curNode := cur.Value.(*node)
//	keys := curNode.keys
//	key := -1
//	nearestToTop := -1
//	// decide which to pop
//	for k := range keys {
//		if this.val2Idx[k][len(this.val2Idx[k])-1] > nearestToTop {
//			nearestToTop = this.val2Idx[k][len(this.val2Idx[k])-1]
//			key = k
//		}
//	}
//	this.val2Idx[key] = this.val2Idx[key][:len(this.val2Idx[key])-1]
//	if curNode.cnt > 1 {
//		if pre := cur.Prev(); pre == nil || pre.Value.(*node).cnt < curNode.cnt-1 {
//			this.nodes[key] = this.InsertBefore(&node{map[int]struct{}{key: {}}, curNode.cnt - 1}, cur)
//		} else {
//			pre.Value.(*node).keys[key] = struct{}{}
//			this.nodes[key] = pre
//		}
//	} else { // key 仅出现一次，将其移出 nodes
//		delete(this.nodes, key)
//	}
//	delete(curNode.keys, key)
//	if len(curNode.keys) == 0 {
//		this.Remove(cur)
//	}
//	return key
//}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

/*
https://leetcode.com/problems/maximum-frequency-stack/solution/
*/

type FreqStack struct {
	freqMap map[int]int
	stacks  map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{map[int]int{}, map[int][]int{}, 0}
}

func (this *FreqStack) Push(val int) {
	freq := this.freqMap[val] + 1
	this.freqMap[val] = freq
	if freq > this.maxFreq {
		this.maxFreq = freq
	}
	v, ok := this.stacks[freq]
	if !ok {
		v = make([]int, 0)
	}
	v = append(v, val)
	this.stacks[freq] = v
}

func (this *FreqStack) Pop() int {
	res := this.stacks[this.maxFreq][len(this.stacks[this.maxFreq])-1]
	this.stacks[this.maxFreq] = this.stacks[this.maxFreq][:len(this.stacks[this.maxFreq])-1]
	this.freqMap[res]--
	if len(this.stacks[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
