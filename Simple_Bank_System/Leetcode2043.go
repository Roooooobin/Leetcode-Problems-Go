package main

type Bank struct {
	accountBalance map[int]int64
}

func Constructor(balance []int64) Bank {
	accountBalance := make(map[int]int64)
	for i, b := range balance {
		accountBalance[i+1] = b
	}
	return Bank{accountBalance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {

	if _, ok := this.accountBalance[account1]; !ok {
		return false
	}
	if _, ok := this.accountBalance[account2]; !ok {
		return false
	}
	if this.accountBalance[account1] < money {
		return false
	}
	this.accountBalance[account1] -= money
	this.accountBalance[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {

	if _, ok := this.accountBalance[account]; !ok {
		return false
	}
	this.accountBalance[account] = this.accountBalance[account] + money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {

	if v, ok := this.accountBalance[account]; !ok || v < money {
		return false
	}
	this.accountBalance[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
