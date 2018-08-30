package models

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	sitcomURLFormat = "https://www.sitcom.ci/page/%d/?search_category&s=%s&search_posttype=product"
)

// SitcomSearch take the query and the category string with page number
// make request to babiken.ci and return List of product found and error
func SitcomSearch(pageCount int, query string) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(sitcomURLFormat, pageCount, url.QueryEscape(query))

	doc, err := makeGETRequest(url)
	if err != nil {
		return
	}

	doc.Find("#loop-products").Find("li").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title, _ = s.Find(".item-content").Find("h4").Find("a").Attr("title")
		p.Link, _ = s.Find(".item-content").Find("h4").Find("a").Attr("href")
		p.Picture, _ = s.Find(".product-thumb-hover").Find("img").Attr("src")
		p.Price = s.Find(".item-price").Find("ins").Find(".woocommerce-Price-amount").Text()
		if p.Price == "" {
			p.Price = s.Find(".item-price").Find(".woocommerce-Price-amount").Text()
		}
		if p != (Product{}) {
			p.Origin = "SITCOM"
			pList = append(pList, p)
		}
	})
	fmt.Println("got ", len(pList), "on sitcom")
	return
}
