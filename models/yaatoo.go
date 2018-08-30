package models

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	yaatooURLFormat = "https://www.yaatoo.ci/recherche?controller=search&orderby=position&orderway=desc&search_query=%s&submit_search=&p=%d"
)

// YaatooSearch take the query and the category string with page number
// make request to yaatoo.ci and return List of product found and error
func YaatooSearch(pageCount int, query string) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(yaatooURLFormat, url.QueryEscape(query), pageCount)
	fmt.Println("yaatooURL ", url)
	doc, err := makeGETRequest(url)
	if err != nil {
		return
	}

	doc.Find(".product_list").Find("li").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title, _ = s.Find(".product-meta").Find(".left").Find("h3").Find("a").Attr("title")
		p.Link, _ = s.Find(".product-meta").Find(".left").Find("h3").Find("a").Attr("href")
		p.Picture, _ = s.Find(".product_img_link").Find("img").Attr("src")
		p.Price = s.Find(".product-price").Text()
		if p != (Product{}) {
			p.Origin = "YAATOO"
			pList = append(pList, p)
		}
	})
	fmt.Println("got ", len(pList), "on yaatoo")
	return
}
