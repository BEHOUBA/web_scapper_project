package test

import (
	"testing"

	"github.com/behouba/webScrapperApp/models"
	_ "github.com/behouba/webScrapperApp/routers"
)

var randomSearchValues = []string{"samsung", "samsung galaxy", "iphone", "Xbox", "LG "}

// func TestJumiaSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.JumiaSearch(1, "", s)
// 		if err != nil {
// 			t.Fail()
// 		}
// 		t.Log("found ", len(pList), " items")
// 	}
// }

func TestAfrimarket(t *testing.T) {
	for _, s := range randomSearchValues {
		pList, err := models.AfrimarketSearch(1, "", s)
		if err != nil {
			t.Errorf("%s", err)
		}
		t.Log("found ", len(pList))
	}
}
