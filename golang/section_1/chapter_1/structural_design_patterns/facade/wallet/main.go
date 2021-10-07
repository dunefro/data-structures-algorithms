package main

import "log"

func main() {
	walletFacade := newWalletFacade("abc", 1234)
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)

	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

}
