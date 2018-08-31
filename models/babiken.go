package models

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	babikenURLFormat = "https://www.babiken.ci/fr/products?search_api_views_fulltext=%s"
)

// BabikenSearch take the query and the category string with page number
// make request to babiken.ci and return List of product found and error
func BabikenSearch(query string) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(babikenURLFormat, url.QueryEscape(query))

	doc, err := makeGETRequest(url)
	if err != nil || doc == nil {
		return
	}

	doc.Find(".all-products").Find("li").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title = s.Find(".product-ft-title").Text()
		p.Link, _ = s.Find(".field-content").Find("a").Attr("href")
		p.Picture, _ = s.Find(".product-img").Find("img").Attr("src")
		p.Price, _ = formatPriceToInt(s.Find(".product-ft-price").Text())
		// if err != nil {
		// 	log.Println(err)
		// }
		if p != (Product{}) {
			p.Origin = "BABIKEN"
			pList = append(pList, p)
		}
	})
	fmt.Println("got ", len(pList), "on babiken")
	return
}
