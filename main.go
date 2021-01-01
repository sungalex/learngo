package main

import (
	"fmt"

	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("alex")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	account.ChangeOwner("park")
	fmt.Println(account)
}
