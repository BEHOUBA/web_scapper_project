package models

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	Title   string `json:"title"`
	Price   string `json:"price"`
	Picture string `json:"picture"`
	Link    string `json:"link"`
}

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"
	jumia     = "https://www.jumia.ci/catalog"
)

// jumiaSearch take the query and the category string with page number
// make request to jumia.ci and return List of product found and error
func JumiaSearch(pageCount int, category, query string) (pList []Product, err error) {
	if category == "" {
		category = "catalog"
	}
	// construction of url of the request
	url := fmt.Sprintf("%s/%s/?q=%s&page=%d", jumia, category, query, pageCount)

	doc, err := makeGETRequest(url)
	if err != nil {
		return
	}

	doc.Find(".products").Find(".sku").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title = s.Find(".title").Text()
		p.Link, _ = s.Find(".link").Attr("href")
		p.Picture, _ = s.Find(".image").Attr("data-src")
		p.Price = s.Find(".price").First().Text()
		if p != (Product{}) {
			pList = append(pList, p)
		}
	})

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
	if err != nil {
		return
	}
	return
}
