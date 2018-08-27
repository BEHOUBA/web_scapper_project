package models

import (
	"testing"
)

func TestJumiaSearch(t *testing.T) {
	pList, err := JumiaSearch(25, "", "smsung")
	if err != nil {
		t.Fail()
	}
	t.Log(len(pList))
}
