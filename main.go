package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func stockHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Path[len("/stock/"):]
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)
	interval := "1d"

	url, err := url.Parse("https://query1.finance.yahoo.com/v7/finance/download/" + symbol)
	if err != nil {
		log.Fatal(err)
	}

	q := url.Query()
	q.Set("period1", strconv.FormatInt(yearAgo.Unix(), 10))
	q.Set("period2", strconv.FormatInt(now.Unix(), 10))
	q.Set("interval", interval)
	q.Set("events", "history")
	q.Set("includeAdjustedClose", "true")
	url.RawQuery = q.Encode()

	var client http.Client
	resp, err := client.Get(url.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "hi"+string(bodyBytes))

}

func main() {
	fmt.Println("Listen http://localhost:8080")
	http.HandleFunc("/stock/", stockHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
