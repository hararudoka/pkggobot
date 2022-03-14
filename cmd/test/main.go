package main

import (
	"fmt"

	"github.com/hararudoka/pkggobot/internal/service"
)

func main () {
	doc, err := service.NewDoc("strings")
	if err != nil {
		panic(err)
	}

	fmt.Println(doc.Types[0].Signature)
}