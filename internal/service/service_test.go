package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseFunctions(t *testing.T) {
	res, err := http.Get("https://pkg.go.dev/errors")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	fs := ParseFunctions(doc)

	fmt.Println(fs)
}

func TestParseTypes(t *testing.T) {
	res, err := http.Get("https://pkg.go.dev/strings")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	types := ParseTypes(doc)

	fmt.Println(types)
}