package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	sitcomURLFormat = "https://www.sitcom.ci/page/%d/?search_category&s=%s&search_posttype=product"
)

// AllFromSitcom search for all article
// related to a specifique query string on sitcom.ci
func AllFromSitcom(query string) (pList []Product, err error) {

	pChan := make(chan []Product)
	for i := 1; i <= 5; i++ {
		page := i
		go func() {
			list, err := SitcomSearch(page, query)
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
			fmt.Println("got ", len(pList), "on sitcom")
			return
		}
	}
	return
}

// SitcomSearch take the query and the category string with page number
// make request to sitcom.ci and return List of product found and error
func SitcomSearch(pageCount int, query string) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(sitcomURLFormat, pageCount, url.QueryEscape(query))

	doc, err := makeGETRequest(url)
	if err != nil || doc == nil {
		return
	}

	doc.Find("#loop-products").Find("li").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		var pString string
		p.Title, _ = s.Find(".item-content").Find("h4").Find("a").Attr("title")
		p.Link, _ = s.Find(".item-content").Find("h4").Find("a").Attr("href")
		p.Picture, _ = s.Find(".product-thumb-hover").Find("img").Attr("src")
		pString = s.Find(".item-price").Find("ins").Find(".woocommerce-Price-amount").Text()
		if pString == "" {
			pString = s.Find(".item-price").Find(".woocommerce-Price-amount").Text()
		}
		p.Price, err = formatPriceToInt(pString)
		if err != nil {
			log.Println(err)
		}
		if p != (Product{}) {
			p.Origin = "SITCOM"
			pList = append(pList, p)
		}
	})
	return
}
