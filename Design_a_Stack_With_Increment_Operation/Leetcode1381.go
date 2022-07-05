package main

type CustomStack struct {
	stack []int
	incre []int
	top   int
}

func Constructor(maxSize int) CustomStack {
	stack := make([]int, maxSize)
	incre := make([]int, maxSize)
	return CustomStack{stack, incre, -1}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack)-1 > this.top {
		this.top++
		this.stack[this.top] = x
	}
}

func (this *CustomStack) Pop() int {
	if this.top == -1 {
		return -1
	}
	res := this.stack[this.top] + this.incre[this.top]
	if this.top != 0 {
		this.incre[this.top-1] += this.incre[this.top]
	}
	this.incre[this.top] = 0
	this.top--
	return res
}

func (this *CustomStack) Increment(k int, val int) {

	limit := min(k-1, this.top)
	if limit >= 0 {
		this.incre[limit] += val
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
