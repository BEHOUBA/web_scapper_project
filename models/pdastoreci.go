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

// PdastoreciSearch take the query and the category string with page number
// make request to http://shop.pdastoreci.com and return List of product found and error
func PdastoreciSearch(query string, pageCount int) (pList []Product, err error) {

	// construction of url of the request
	url := fmt.Sprintf(pdastoreciURLFormat, url.QueryEscape(query), pageCount)

	doc, err := makeGETRequest(url)
	if err != nil {
		return
	}

	doc.Find(".HotDealList").Find(".HotDeal").Each(func(i int, s *goquery.Selection) {
		p := Product{}
		p.Title, _ = s.Find(".ProductName").Attr("title")
		linkPath, _ := s.Find(".ProductName").Attr("href")
		p.Link = "http://shop.pdastoreci.com/epages/265339.sf/fr_FR/" + linkPath
		picturePath, _ := s.Find(".ProductHotDealImage").Attr("src")
		p.Picture = "http://shop.pdastoreci.com" + picturePath
		p.Price, err = formatPriceToInt(s.Find(".price-value").Text())
		if err != nil {
			log.Println(err)
		}
		if p != (Product{}) {
			p.Origin = "PDASTORECI"
			pList = append(pList, p)
		}
	})
	fmt.Println("got ", len(pList), "on pdastoreci")
	return
}
