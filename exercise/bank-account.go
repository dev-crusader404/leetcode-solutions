package exercise

import "sync"

// Define the Account type here.
type Account struct {
	Amount   int64
	isClosed bool
	mu       sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{Amount: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if a.isClosed {
		return 0, false
	}
	return a.Amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed {
		return 0, false
	}
	if amount < 0 && a.Amount < -amount {
		return 0, false
	}
	var mu sync.Mutex
	defer mu.Unlock()
	mu.Lock()
	a.Amount += amount
	return a.Amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed {
		return 0, false
	}
	var mu sync.Mutex
	defer mu.Unlock()
	mu.Lock()
	blnc := a.Amount
	a.Amount = 0
	a.isClosed = true
	return blnc, true
}
