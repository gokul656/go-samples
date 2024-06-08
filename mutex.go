package main

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

type Account struct {
	sync.Mutex
	balance decimal.Decimal
}

func (a *Account) deposit(amount decimal.Decimal) {
	a.Lock()
	defer a.Unlock()

	a.balance = a.balance.Add(amount)
}

func (a *Account) withdraw(amount decimal.Decimal) {
	a.Lock()
	defer a.Unlock()

	a.balance = a.balance.Sub(amount)
}

func (a *Account) getBalance() decimal.Decimal {
	return a.balance
}

func mutexExample() {

	wg := &sync.WaitGroup{}
	account := &Account{
		balance: decimal.New(10, 2),
	}

	for range 1000 {
		// deposit
		wg.Add(1)
		go func(acc *Account) {
			defer wg.Done()
			acc.deposit(decimal.New(10, 2))
		}(account)

		// withdraw
		wg.Add(1)
		go func(acc *Account) {
			defer wg.Done()
			acc.withdraw(decimal.New(1, 2))
		}(account)
	}

	wg.Wait()
	fmt.Println(account.getBalance())
}
