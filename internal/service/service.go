package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// i need to parse pkg.go.doc

// Symbol is representation of common part of every package's symbols
type Symbol struct {
	//Name string
	Signature   string
	Description []string // element is one paragraph of text
	Example     string   // code in text
}

// Constant is type of Symbol
type Constant Symbol

// Var is type of Symbol
type Var Symbol

// Function is type of Symbol
type Function Symbol

// Type is type of Symbol with some Methods and Fields inside
type Type struct {
	Symbol
	Methods []Function
	Fields  []Var
}

// Doc is just list of symbols
type Doc struct {
	PkgName string

	Constants []Constant
	Vars      []Var
	Functions []Function
	Types     []Type

	Dirs []string
}

func NewDoc(pkgName string) (Doc, error) {
	url := "https://pkg.go.dev/" + pkgName

	res, err := http.Get(url)
	if err != nil {
		return Doc{}, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return Doc{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Doc{}, err
	}

	return Doc{
		PkgName:   pkgName,
		Constants: nil,
		Vars:      nil,
		Functions: ParseFunctions(doc), //ParseFunctions(doc)
		Types:     nil, //ParseTypes(doc),
	}, nil
}

func ParseFunctions(doc *goquery.Document) []Function {
	var fs []Function

	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		if v, _ := s.Attr("class"); v == "Documentation-function" {
			// get signature
			signature := ""
			s.Find("div").Each(func(i int, s *goquery.Selection) {
				if v, _ := s.Attr("class"); v == "Documentation-declaration" {
					signature = strings.TrimSpace(s.Text())
				}
			})

			// get description
			var des []string
			for _, e := range strings.Split(s.Find("p").Text(), ".\n") {
				if s := strings.ReplaceAll(e, "\n", " "); s != "" {
					des = append(des, s+".")
				}
			}

			// get example
			example := ""
			s.Find("pre").Each(func(i int, s *goquery.Selection) {
				if v, _ := s.Attr("class"); v == "Documentation-exampleCode" {
					example = s.Text()
				}
			})

			fs = append(fs, Function{
				Signature:   signature,
				Description: des,
				Example:     example,
			})
		}
	})

	return fs
}

func ParseTypes(doc *goquery.Document) []Type {
	var types []Type

	doc.Find("section").Each(func(i int, s *goquery.Selection) {
		if v, _ := s.Attr("class"); v == "Documentation-types" {
			// get signature
			signature := ""
			s.Find("div").Each(func(i int, s *goquery.Selection) {
				if v, _ := s.Attr("class"); v == "Documentation-declaration" {
					signature = strings.TrimSpace(s.Text())
				}
			})

			// get description
			var des []string
			for _, e := range strings.Split(s.Find("p").Text(), ".\n") {
				if s := strings.ReplaceAll(e, "\n", " "); s != "" {
					des = append(des, s+".")
				}
			}

			// get example
			example := ""
			s.Find("pre").Each(func(i int, s *goquery.Selection) {
				if v, _ := s.Attr("class"); v == "Documentation-exampleCode" {
					example = s.Text()
				}
			})

			types = append(types, Type{
				Symbol: Symbol{
					Signature:   signature,
					Description: des,
					Example:     example,
				},
				Methods: nil,
				Fields: nil,
			})
		}
	})

	return types
}
