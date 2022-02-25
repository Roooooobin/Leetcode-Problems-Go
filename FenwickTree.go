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
