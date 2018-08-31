package test

import (
	"testing"

	"github.com/behouba/webScrapperApp/models"
	_ "github.com/behouba/webScrapperApp/routers"
)

var randomSearchValues = []string{"samsung", "samsung galaxy", "iphone", "Xbox", "LG ", "lenovo ideapad"}

// func TestJumiaSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.JumiaSearch(1, "", s)
// 		if err != nil {
// 			t.Fail()
// 		}
// 		t.Log("found ", len(pList), " items")
// 	}
// }

// func TestAfrimarket(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.AfrimarketSearch(1, "", s)
// 		if err != nil {
// 			t.Errorf("%s", err)
// 		}
// 		t.Log("found ", len(pList))
// 	}
// }

// func TestYaatooSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.YaatooSearch(1, s)
// 		if err != nil {
// 			t.Errorf("%s", err)
// 		}
// 		t.Log("found ", len(pList))
// 	}
// }

// func TestBabikenSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.BabikenSearch(s)
// 		if err != nil {
// 			t.Errorf("%s", err)
// 		}
// 		t.Log("found ", len(pList))
// 	}
// }

// func TestSitcomSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.SitcomSearch(1, s)
// 		if err != nil {
// 			t.Errorf("%s", err)
// 		}
// 		t.Log("found ", len(pList))
// 	}
// }

// func TestAfrikdiscountSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.AfrikdiscountSearch(s, 1)
// 		if err != nil {
// 			t.Errorf("%s", err)
// 		}
// 		t.Log("found ", len(pList))
// 		t.Log("first ", pList[0])
// 	}
// }

// func TestPdastoreciSearch(t *testing.T) {
// 	for _, s := range randomSearchValues {
// 		pList, err := models.PdastoreciSearch(s, 1)
// 		if err != nil {
// 			t.Errorf("%s", err)
// 		}
// 		t.Log("found ", len(pList))
// 		t.Log("first ", pList[0])
// 	}
// }

func TestSitcomAll(t *testing.T) {
	for _, s := range randomSearchValues {
		pList, err := models.AllFromSitcom(s)
		if err != nil {
			t.Errorf("%s", err)
		}
		t.Log("found ", len(pList))
		// t.Log("first ", pList[0])
	}
}
