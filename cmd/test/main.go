package main

import (
	"fmt"

	"github.com/hararudoka/pkggobot/internal/service"
)

func main () { // TODO: remove this file
	doc, err := service.NewDoc("strings")
	if err != nil {
		panic(err)
	}

	fmt.Println(doc.Types[0].Signature)
}