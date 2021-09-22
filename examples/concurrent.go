package examples

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Concurrent struct{}
type ChuckNorris struct {
	Quote      string   `json:"value"`
	Url        string   `json:"url"`
	Id         string   `json:"id"`
	Categories []string `json:"categories"`
}

func (Concurrent) Run() {
	startTime := time.Now()
	defer printExecutionTime(startTime)
	getQuotesConcurrently(100)
}

const baseUrl = "https://api.chucknorris.io/jokes/random"

func getQuote() (quote *ChuckNorris, err error) {
	response, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func printExecutionTime(t time.Time) {
	fmt.Println("Execution time: ", time.Since(t))
}

func getQuotesSequentially(numOfQuotes int) {
	quotesMap := make(map[int]*ChuckNorris, numOfQuotes)

	for i := 0; i < numOfQuotes; i++ {
		fmt.Println(i)
		quote, err := getQuote()
		if err != nil {
			continue
		}
		quotesMap[i] = quote
		fmt.Printf("New Chuck Norris quote: %v\n", quote.Quote)
	}
}

func getQuotesConcurrently(numOfQuotes int) {
	//quotesMap := make(map[int]*ChuckNorris, numOfQuotes)
	wg := sync.WaitGroup{}
	var quotesMap sync.Map

	for i := 0; i < numOfQuotes; i++ {
		wg.Add(1)
		go func(idx int) {
			quote, err := getQuote()
			if err != nil {
				panic(err)
			}
			//quotesMap[idx] = quote
			quotesMap.Store(idx, quote)

			//fmt.Printf("New Chuck Norris quote: %v\n", quote.Quote)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
