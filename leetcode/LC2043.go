package leetcode

type Bank struct {
	accounts []int64
}

func InitBank(balance []int64) Bank {
	return Bank{
		accounts: balance,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > len(this.accounts) || account2 > len(this.accounts) {
		return false
	}
	isSuccess := this.Withdraw(account1, money)
	if !isSuccess {
		return isSuccess
	}
	return this.Deposit(account2, money)
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > len(this.accounts) {
		return false
	}
	this.accounts[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > len(this.accounts) {
		return false
	}
	if this.accounts[account-1] < money {
		return false
	}
	this.accounts[account-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
