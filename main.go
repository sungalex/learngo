package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "greetings"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}

// for account package

// import (
// 	"fmt"
// 	"learngo/accounts"
// )

// func main() {
// 	account := accounts.NewAccount("alex")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	account.ChangeOwner("park")
// 	fmt.Println(account)
// }
