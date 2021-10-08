package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}
type account struct {
	customerName   string
	accountDetails *AccountDetails
}

func (account *account) setDetails(id string, accountType string) {
	account.accountDetails = &AccountDetails{
		id:          id,
		accountType: accountType,
	}
}

func (account *account) getId() string {
	return account.accountDetails.id
}
func (account *account) getAccountType() string {
	return account.accountDetails.accountType
}

func main() {
	var account *account = &account{
		customerName: "John Smith",
	}
	account.setDetails("4532", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private class hidden", string(jsonAccount))
	fmt.Println("Account Customer Name", account.customerName)
	fmt.Println("Account id", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
