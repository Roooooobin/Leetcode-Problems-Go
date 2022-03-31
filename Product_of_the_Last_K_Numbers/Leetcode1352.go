package main

type ProductOfNumbers struct {
	list []int
}

func Constructor() ProductOfNumbers {

	return ProductOfNumbers{list: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {

	if num == 0 {
		this.list = []int{1}
		return
	}
	this.list = append(this.list, this.list[len(this.list)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {

	if len(this.list)-1 < k {
		return 0
	}
	return this.list[len(this.list)-1] / this.list[len(this.list)-k-1]
}
