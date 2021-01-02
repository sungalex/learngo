package main

import (
	"fmt"
	"learngo/jobscrapper"
)

var searchWord = "python"
var baseURL = "https://kr.indeed.com/jobs?q=" + searchWord

func main() {
	jobs := []jobscrapper.ExtractedJob{}
	c := make(chan []jobscrapper.ExtractedJob)
	totalPages := jobscrapper.GetPageCounts(baseURL)

	for i := 0; i < totalPages; i++ {
		go jobscrapper.GetPage(baseURL, i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	jobscrapper.WriteJobs(jobs)
	fmt.Println("Done, Extracted", len(jobs))
}

// go routine and channel - final
/*
[Go routines and channels theory]
	- 메인 함수가 종료하면 go routine이 끝나던 아니던 상관없이 프로그램이 종료된다.
	- go routine은 Non-blocking Operation 이지만,
	  channel을 수신하는 operation(<-)은 Blocking Operation 이다.
	  즉, "<-" Operation을 수행하면 Channel에서 message를 수신할 때까지 대기 한다.
	- Channel를 생성할 때는 수신할 메시지의 Type을 지정해야 한다.
	- Channel에서 message를 수신할 때는 한번에 하나의 메시지 만 수신된다.
	  즉, go routine의 수 만큼 channel을 통해 message를 수신하는 "<-" Operation이 있어야 한다.
	- 함수 선언 시 channel type에 "chan<-"을 사용하면 read-only channel이 된다.
*/
/*
import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	c := make(chan requestResult)
	results := map[string]string{}
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
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	status := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
*/

// go routine and channel - 1
/*
import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := []string{"alex", "park", "nico", "flynn"}
	for _, people := range people {
		go isSexy(people, c)
	}
	fmt.Println("Waiting for messages...")
	for i := 0; i < len(people); i++ {
		fmt.Println("Recived this message:", <-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
*/

// for URL Checker
/*
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
*/

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
