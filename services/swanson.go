package services

import (
	"encoding/json"
	"net/http"
	"time"
)

func FetchSwansonQuote() chan string {
	var quotes []string
	quoteChan := make(chan string)

	go func() {
		// To simulate longer load times
		time.Sleep(time.Second * 3)

		res, _ := http.Get("https://ron-swanson-quotes.herokuapp.com/v2/quotes")
		decoder := json.NewDecoder(res.Body)

		decoder.Decode(&quotes)

		quoteChan <- quotes[0]
	}()

	return quoteChan
}
