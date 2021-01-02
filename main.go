package main

// for URL Checker & Go Routines
import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var results = map[string]string{}
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

// for mydict package
/*
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
*/

// for account package
/*
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
*/
