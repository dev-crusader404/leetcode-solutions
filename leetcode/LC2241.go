package leetcode

type ATM struct {
	balance []int
}

func InitATM() ATM {
	return ATM{
		balance: make([]int, 5),
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, count := range banknotesCount {
		this.balance[i] += count
	}
}

func (this *ATM) Withdraw(amount int) []int {
	invalid := []int{-1}
	if amount < 20 {
		return invalid
	}
	original := make([]int, 5)
	copy(original, this.balance)
	n := len(this.balance) - 1
	result := make([]int, 5)
	for i := n; i >= 0; i-- {
		note := getAmount(i)
		for amount >= note && this.balance[i] > 0 {
			amount -= note
			result[i]++
			this.balance[i]--
			if amount == 0 {
				return result
			}
		}
	}
	this.balance = original
	return invalid
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

func RunLC2241() {

}
