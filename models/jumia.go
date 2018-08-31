package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"
	jumiaURLFormat = "https://www.jumia.ci/%s/?page=%d&q=%s"
)

// AllFromJumia search for all article
// related to a specifique query string on jumia.ci
func AllFromJumia(query string) (pList []Product, err error) {

	pChan := make(chan []Product)
	for i := 1; i <= 5; i++ {
		page := i
		go func() {
			list, err := JumiaSearch(page, "", query)
			pChan <- list
			if err != nil {
				log.Println(err)
				return
			}
			return
		}()
	}

	for j := 1; j <= 5; j++ {
		pList = append(pList, <-pChan...)
		if j == 5 {
			close(pChan)
			fmt.Println("got ", len(pList), "on jumia")
			return
		}
	}
	return
}

// JumiaSearch take the query and the category string with page number
// make request to jumia.ci and return List of product found and error
func JumiaSearch(pageCount int, category, query string) (pList []Product, err error) {
	if category == "" {
		category = "catalog"
	}
	// construction of url of the request
	url := fmt.Sprintf(jumiaURLFormat, category, pageCount, url.QueryEscape(query))

	doc, err := makeGETRequest(url)
	if err != nil || doc == nil {
		return
	}

	doc.Find(".products").Find(".sku").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title = s.Find(".title").Text()
		p.Link, _ = s.Find(".link").Attr("href")
		p.Picture, _ = s.Find(".image").Attr("data-src")
		p.Price, _ = formatPriceToInt(s.Find(".price").First().Text())
		// if err != nil {
		// 	log.Println(err)
		// }
		if p != (Product{}) {
			p.Origin = "JUMIA"
			pList = append(pList, p)
		}
	})
	return
}
