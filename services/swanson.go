package services

import (
	"encoding/json"
	"net/http"
)

func FetchSwansonQuote() chan string {
	var quotes []string
	quoteChan := make(chan string)

	go func() {
		res, _ := http.Get("https://ron-swanson-quotes.herokuapp.com/v2/quotes")
		decoder := json.NewDecoder(res.Body)

		decoder.Decode(&quotes)

		quoteChan <- quotes[0]
	}()

	return quoteChan
}
