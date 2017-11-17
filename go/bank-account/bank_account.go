package account

import "sync"


type Account struct{
	isLocked bool
	balance int64
	sync.Mutex
}


func (a *Account) Close() (int64, bool){
	a.Lock()
	defer a.Unlock()

	if a.isLocked {
		return 0, false
	}

	a.isLocked = true
	return a.balance, true
}

func (a Account) Balance() (int64, bool){
	return a.balance, !a.isLocked
}

func (a *Account) Deposit(amount int64) (int64, bool){
	a.Lock()
	defer a.Unlock()

	newAmount := a.balance + amount
	if !a.isLocked && newAmount >= 0{
		a.balance = newAmount
	}

	return a.balance, !a.isLocked && newAmount >= 0
}

func Open(initialDeposit int64) *Account{
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		isLocked : false,
		balance : initialDeposit,
	}
}