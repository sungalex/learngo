package jobscrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractedJob type
type ExtractedJob struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

var viewjobURL = "https://kr.indeed.com/viewjob?jk="

// GetPageCounts count total pages
func GetPageCounts(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

// GetPage extract all Job Cards
func GetPage(url string, page int, mainC chan<- []ExtractedJob) {
	jobs := []ExtractedJob{}
	c := make(chan ExtractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting:", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		jobs = append(jobs, <-c)
	}

	mainC <- jobs
}

// WriteJobs write Jobs to CSV file
func WriteJobs(jobs []ExtractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Company", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{viewjobURL + job.id, job.title, job.company, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func extractJob(card *goquery.Selection, c chan<- ExtractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	company := cleanString(card.Find(".company").Text())
	location := cleanString(card.Find(".location").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	c <- ExtractedJob{
		id:       id,
		title:    title,
		company:  company,
		location: location,
		salary:   salary,
		summary:  summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with statuscode:", res.StatusCode)
	}
}
