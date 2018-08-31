package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	afrikdiscountURLFormat = "https://afrikdiscount.com/fr/recherche?controller=search&orderby=position&orderway=desc&search_query=%s&submit_search=&p=%d"
)

// AfrikdiscountSearch take the query and the category string with page number
// make request to afrikdiscount.com/fr/ and return List of product found and error
func AfrikdiscountSearch(query string, pageCount int) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(afrikdiscountURLFormat, url.QueryEscape(query), pageCount)

	doc, err := makeGETRequest(url)
	if err != nil {
		return
	}

	doc.Find(".product_list").Find("li").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title = s.Find(".product-name").Text()
		p.Link, _ = s.Find(".product_img_link").Attr("href")
		p.Picture, _ = s.Find(".product_img_link").Find("img").Attr("src")
		p.Price, err = formatPriceToInt(s.Find(".product-price").Text())
		if err != nil {
			log.Println(err)
		}
		if p != (Product{}) {
			p.Origin = "AFRIKDISCOUNT"
			pList = append(pList, p)
		}
	})
	fmt.Println("got ", len(pList), "on afrikdiscount")
	return
}
