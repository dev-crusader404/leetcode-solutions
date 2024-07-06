package leetcode

type ATM struct {
	balance  []int
	withdraw []int
}

var invalid []int = []int{-1}

func InitATM() ATM {
	return ATM{
		balance:  make([]int, 5),
		withdraw: make([]int, 5),
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, count := range banknotesCount {
		this.balance[i] += count
	}
}

func (this *ATM) Withdraw(amount int) []int {
	if amount < 20 {
		return invalid
	}
	n := len(this.balance) - 1

	for i := n; i >= 0; i-- {
		note := getAmount(i)
		this.withdraw[i] = amount / note

		if this.balance[i] < this.withdraw[i] {
			this.withdraw[i] = this.balance[i]
		}
		amount -= this.withdraw[i] * note
	}

	if amount != 0 {
		return invalid
	}

	for i := range this.withdraw {
		this.balance[i] -= this.withdraw[i]
	}
	return this.withdraw
}

func getAmount(i int) int {
	switch i {
	case 4:
		return 500
	case 3:
		return 200
	case 2:
		return 100
	case 1:
		return 50
	case 0:
		return 20
	}
	return 0
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
