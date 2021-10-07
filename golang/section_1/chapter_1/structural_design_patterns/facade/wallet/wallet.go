package main

import "fmt"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	w := &wallet{
		balance: 0,
	}
	fmt.Println("Wallet has been created")
	return w
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Money credited to the account")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Insufficient Balance in the account")
	}
	w.balance -= amount
	fmt.Println("Money debited from the account")
	return nil
}
