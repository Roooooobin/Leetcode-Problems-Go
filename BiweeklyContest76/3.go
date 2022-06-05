package main

type ATM struct {
	count []int64
	value [5]int
}

func Constructor() ATM {

	count := make([]int64, 5)
	return ATM{count, [5]int{20, 50, 100, 200, 500}}
}

func (this *ATM) Deposit(banknotesCount []int) {

	for i := range banknotesCount {
		this.count[i] += int64(banknotesCount[i])
	}
}

func (this *ATM) Withdraw(amount int) []int {

	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		res[i] = int(min(this.count[i], int64(amount/this.value[i])))
		amount -= res[i] * this.value[i]
	}
	if amount > 0 {
		return []int{-1}
	} else {
		for i := 0; i < 5; i++ {
			this.count[i] -= int64(res[i])
		}
		return res
	}
}

func min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
