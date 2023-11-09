package main

import (
	"github.com/jrmn-commits/goprogram/helpers"
	"log"
)

func main () {
	log.Println("Hello, World!")

	var myVar helpers.SameType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
