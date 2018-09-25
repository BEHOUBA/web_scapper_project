package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	pdastoreciURLFormat = "http://shop.pdastoreci.com/epages/265339.sf/fr_FR/?ViewAction=FacetedSearchProducts&ObjectID=524324&PageSize=30&SearchString=%s&Page=%d"
)

// AllFromPdastoreci search for all article
// related to a specifique query string on shop.pdastoreci.com
func AllFromPdastoreci(query string) (pList []Product, err error) {

	pChan := make(chan []Product)
	for i := 1; i <= 5; i++ {
		page := i
		go func() {
			list, err := PdastoreciSearch(query, page)
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
			fmt.Println("got ", len(pList), "on pdastoreci")
			return
		}
	}
	return
}

// PdastoreciSearch take the query and the category string with page number
// make request to http://shop.pdastoreci.com and return List of product found and error
func PdastoreciSearch(query string, pageCount int) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(pdastoreciURLFormat, url.QueryEscape(query), pageCount)

	doc, err := makeGETRequest(url)
	if err != nil || doc == nil {
		return
	}

	doc.Find(".HotDealList").Find(".HotDeal").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title, _ = s.Find(".ProductName").Attr("title")
		linkPath, _ := s.Find(".ProductName").Attr("href")
		p.Link = "http://shop.pdastoreci.com/epages/265339.sf/fr_FR/" + linkPath
		picturePath, _ := s.Find(".ProductHotDealImage").Attr("src")
		p.Picture = "http://shop.pdastoreci.com" + picturePath
		p.Price, _ = formatPriceToInt(s.Find(".price-value").Text())
		// if err != nil {
		// 	log.Println(err)
		// }
		if p != (Product{}) {
			p.Origin = "PDASTORECI"
			pList = append(pList, p)
		}
	})
	return
}
