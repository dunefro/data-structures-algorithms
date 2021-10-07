package main

import "fmt"

type account struct {
	name string
}

func newAccount(name string) *account {
	acc := &account{
		name: name,
	}
	return acc
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}
