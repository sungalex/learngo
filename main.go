package main

import (
	"fmt"

	"learngo/accounts"
)

func main() {
	account := accounts.NewAccount("alex")
	fmt.Println(account)
}
