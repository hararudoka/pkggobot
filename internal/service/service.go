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
	Name        string
	Signature   string
	Description []string // element is one paragraph of text
	Example     string   // code in text
}

// Type is type of Symbol with some Methods and Fields inside
type Type struct {
	Symbol
	Methods []Symbol
	Fields  []Symbol
}

// Doc is just list of symbols
type Doc struct {
	PkgName  string
	Overview string

	Constants []Symbol
	Vars      []Symbol
	Functions []Symbol
	Types     []Type

	Dirs []string
}

func NewDoc(pkgName string) (Doc, error) {
	url := "https://pkg.go.dev/" + pkgName

	page, err := getPage(url)
	if err != nil {
		return Doc{}, err
	}

	return Doc{
		PkgName:   pkgName,
		Overview:  parseOverview(page),
		Constants: nil, // parseConstants(doc)
		Vars:      nil, // parseVars(doc)
		Functions: parseFunctions(page),
		Types:     nil, // parseTypes(doc),
	}, nil
}

// Find finds Symbol by name inside whole Doc and returns it
func (d Doc) Find(name string) Symbol {
	for _, e := range d.Constants {
		if e.Name == name {
			return Symbol(e)
		}
	}
	for _, e := range d.Vars {
		if e.Name == name {
			return Symbol(e)
		}
	}
	for _, e := range d.Functions {
		if e.Name == name {
			return Symbol(e)
		}
	}
	for _, e := range d.Types {
		if e.Name == name {
			return e.Symbol
		}
		for _, v := range e.Methods {
			if v.Name == name {
				return Symbol(v)
			}
		}
	}
	return Symbol{}
}

// getPage returns *goquery.Document form URL
func getPage(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, err
}

// it is not actually Overview - just a description of package
func parseOverview(doc *goquery.Document) string {
	var st string
	if v, ok := doc.Find("head > meta:nth-child(6)").Attr("content"); ok {
		st = v
	} else {
		st = "not found"
	}
	return st
}

// parseConstants returns []Constant from site
func parseConstants(doc *goquery.Document) []Symbol {
	var cs []Symbol

	doc.Find("div.Documentation-constants").Each(func(i int, s *goquery.Selection) {
		// TODO: parseConstants
	})

	return cs
}

// parseVars returns []Var from site
func parseVars(doc *goquery.Document) []Symbol {
	var vs []Symbol

	doc.Find("div.Documentation-variables").Each(func(i int, s *goquery.Selection) {
		// TODO: parseVars
	})

	return vs
}

// parseFunctions returns []Function from site
func parseFunctions(doc *goquery.Document) []Symbol {
	var fs []Symbol

	doc.Find("div.Documentation-function").Each(func(i int, s *goquery.Selection) {
		// get name
		name, _ := s.Find("h4").Attr("id")

		// get signature
		signature := strings.TrimSpace(s.Find("div.Documentation-declaration").Text())

		// get description
		var des []string
		for _, e := range strings.Split(s.Find("p").Text(), ".\n") {
			if s := strings.ReplaceAll(e, "\n", " "); s != "" {
				des = append(des, s+".")
			}
		}

		// get example
		example := s.Find("pre.Documentation-exampleCode").Text()

		fs = append(fs, Symbol{
			Name:        name,
			Signature:   signature,
			Description: des,
			Example:     example,
		})
	})

	return fs
}

// parseTypes returns []Type from site
func parseTypes(doc *goquery.Document) []Type {
	var types []Type

	doc.Find("div.Documentation-type").Each(func(i int, s *goquery.Selection) {
		// get signature
		signature := strings.TrimSpace(s.Find("div.Documentation-type > div.Documentation-declaration > pre:nth-child(1)").Text())

		// get description
		var des []string
		s.Find("p").Each(func(i int, s *goquery.Selection) {
			for _, e := range strings.Split(s.Text(), ".\n") {
				if s := strings.ReplaceAll(e, "\n", " "); s != "" {
					des = append(des, s+".")
				}
			}
		})

		// get example
		example := s.Find("div.Documentation-exampleDetailsBody").Find("pre.Documentation-exampleCode").Text()

		// get methods
		var meths []Symbol
		s.Find("div.Documentation-typeMethod").Each(func(i int, s *goquery.Selection) {
			// get signature
			sig := s.Find("pre").Text()

			// get description
			var desM []string
			s.Find("p").Each(func(i int, s *goquery.Selection) {
				for _, e := range strings.Split(s.Text(), ".\n") {
					if s := strings.ReplaceAll(e, "\n", " "); s != "" {
						desM = append(desM, s+".")
					}
				}
			})

			// TODO: get example
			// ...

			meths = append(meths, Symbol{
				Name:        "name", // TODO name
				Signature:   sig,
				Description: desM,
				Example:     "",
			})
		})

		// get fields
		// not sure need it, it is harder than parsing

		types = append(types, Type{
			Symbol: Symbol{
				Name:        "name", // TODO name
				Signature:   signature,
				Description: des,
				Example:     example,
			},
			Methods: meths,
			Fields:  nil,
		})
	})

	return types
}
