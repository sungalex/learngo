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

	err2 := dictionary.Update(word, "greetings updated")
	if err2 != nil {
		fmt.Println(err2)
	}
	result, _ := dictionary.Search(word)
	fmt.Println(result)

	dictionary.Delete(word)
	result3, err3 := dictionary.Search(word)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(result3)
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
