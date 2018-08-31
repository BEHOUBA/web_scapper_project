package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AfJSONData struct {
	ResultsList []Results `json:"results"`
}
type Results struct {
	HitsList []Hits `json:"hits"`
}

type Hits struct {
	Title   string `json:"name"`
	Link    string `json:"url"`
	Picture string `json:"image_url"`
	Price   struct {
		XOF struct {
			Default       int    `json:"default"`
			PriceFormated string `json:"default_formated"`
		} `json:"XOF"`
	} `json:"price"`
}

const (
	afIndexName      = "magento2_cote_ivoire_local_products"
	headersFormat    = `{"requests":[{"indexName":"%s","params":"query=%s&hitsPerPage=20&maxValuesPerFacet=10&page=%d%s"}]}`
	afURLEnd         = "&facets=%5B%22price.XOF.default%22%2C%22color%22%2C%22size%22%2C%22our_brand%22%5D&tagFilters=&numericFilters=%5B%22visibility_search%3D1%22%5D"
	afrimarketAPIKEY = "Mjc0Mzk5OGI5YmY0NGM2MDRlNTNhNTIzOWI5NTY4NGNhMzNjNDI5NjFkZmY1NTc1YWFiZWI0OGEzNmUxYjk4MHRhZ0ZpbHRlcnM9"
	afrimarketURL    = "https://fk542t44ex-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20vanilla%20JavaScript%20(lite)%203.24.9%3Binstantsearch.js%202.4.1%3BMagento2%20integration%20(1.6.0)%3BJS%20Helper%202.23.2&x-algolia-application-id=FK542T44EX&x-algolia-api-key="
	apiURL           = afrimarketURL + afrimarketAPIKEY
)

// AllFromAfrimarket search for all article
// related to a specifique query string on afrimarket.ci
func AllFromAfrimarket(query string) (pList []Product, err error) {

	pChan := make(chan []Product)
	for i := 1; i <= 5; i++ {
		page := i
		go func() {
			list, err := AfrimarketSearch(page, "", query)
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
			fmt.Println("got ", len(pList), "on afrimarket")
			return
		}
	}
	return
}

// AfrimarketSearch take the query and the category string with page number
// make request to afrimarket.ci and return List of product found and error
func AfrimarketSearch(pageCount int, category, query string) (pList []Product, err error) {

	body, err := makePOSTReqToAF(query, category, pageCount)
	if err != nil {
		return
	}
	jsonData := AfJSONData{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return
	}
	for _, p := range jsonData.ResultsList[0].HitsList {
		product := Product{
			Title:   p.Title,
			Link:    p.Link,
			Picture: p.Picture,
		}

		product.Price, _ = formatPriceToInt(p.Price.XOF.PriceFormated)
		// if err != nil {
		// 	log.Println(err)
		// }
		if product.Price > 90000000 {
			continue
		}
		if product != (Product{}) {
			product.Origin = "AFRIMARKET"
			pList = append(pList, product)
		}
	}
	return
}

func makePOSTReqToAF(query, category string, pageCount int) (body []byte, err error) {
	requestData := fmt.Sprintf(headersFormat, afIndexName, url.QueryEscape(query), pageCount, afURLEnd)
	jsonStr := []byte(requestData)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://afrimarket.ci")
	req.Header.Set("Referer", "https://afrimarket.ci/")
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
