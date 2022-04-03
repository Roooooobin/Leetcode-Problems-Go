package main

type FenwickTreeImpl struct {
	n    int
	tree []int
}

func (ft FenwickTreeImpl) lowBit(x int) int {
	return x & (-x)
}

func (ft FenwickTreeImpl) query(i int) int {

	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= ft.lowBit(i)
	}
	return sum
}

func (ft FenwickTreeImpl) update(i, x int) {
	for i <= ft.n {
		ft.tree[i] += x
		i += ft.lowBit(i)
	}
}

type NumArray struct {
	bit  *FenwickTreeImpl
	nums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1)
	bit := &FenwickTreeImpl{n, tree}
	for i, num := range nums {
		bit.update(i+1, num)
	}
	return NumArray{bit, nums}
}

func (this *NumArray) Update(index int, val int) {
	this.bit.update(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.bit.query(right+1) - this.bit.query(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
