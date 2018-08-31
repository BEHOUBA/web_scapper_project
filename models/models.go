package models

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	Title   string `json:"title"`
	Price   int    `json:"price"`
	Picture string `json:"picture"`
	Link    string `json:"link"`
	Origin  string `json:"origin"`
}

func formatPriceToInt(priceString string) (price int, err error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	formatedString := reg.ReplaceAllString(priceString, "")
	price, err = strconv.Atoi(formatedString)
	if err != nil {
		return
	}
	return
}

// makeGETRequest set User-Agent header value and make a get request to given url
// make a new html Document with goquery library and then
// return et pointer to goquery.Document struct and an error
func makeGETRequest(url string) (doc *goquery.Document, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if doc == nil {
		return nil, errors.New("Can't create new document from response body")
	}
	if err != nil {
		return
	}
	return
}
