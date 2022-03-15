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
		Constants: nil, // ParseConstants(doc)
		Vars:      nil, // ParseVars(doc)
		Functions: ParseFunctions(doc),
		Types:     nil, //ParseTypes(doc),
	}, nil
}

// ParseConstants returns []Constant from site
func ParseConstants(doc *goquery.Document) []Constant {
	var cs []Constant

	doc.Find("div.Documentation-constants").Each(func(i int, s *goquery.Selection) {
		// TODO: ParseConstants
	})

	return cs
}

// ParseVars returns []Var from site
func ParseVars(doc *goquery.Document) []Var {
	var vs []Var

	doc.Find("div.Documentation-variables").Each(func(i int, s *goquery.Selection) {
		// TODO: ParseVars
	})

	return vs
}

// ParseFunctions returns []Function from site
func ParseFunctions(doc *goquery.Document) []Function {
	var fs []Function

	doc.Find("div.Documentation-function").Each(func(i int, s *goquery.Selection) {
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

		fs = append(fs, Function{
			Signature:   signature,
			Description: des,
			Example:     example,
		})
	})

	return fs
}

// ParseTypes returns []Type from site
func ParseTypes(doc *goquery.Document) []Type {
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
		var meths []Function
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

			meths = append(meths, Function{
				Signature:   sig,
				Description: desM,
				Example:     "",
			})
		})

		// get fields
		// not sure need it, it is harder than parsing

		types = append(types, Type{
			Symbol: Symbol{
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
