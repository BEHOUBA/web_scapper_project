package models

import (
	"log"
	"regexp"
	"strconv"
)

type Product struct {
	Title   string `json:"title"`
	Price   int    `json:"price"`
	Picture string `json:"picture"`
	Link    string `json:"link"`
	Origin  string `json:"origin"`
}

func formatPriceToInt(priceString string) (price int, err error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	formatedString := reg.ReplaceAllString(priceString, "")
	price, err = strconv.Atoi(formatedString)
	if err != nil {
		return
	}
	return
}
