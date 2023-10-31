package main

import "log"

type myStruct struct {
	firstName string
}

func (m *myStruct) printFirstName() string {
	return m.firstName
}

func main() {
	var myVar myStruct
	myVar.firstName = "John"

	myVar2 := myStruct{
		firstName: "Mary",
	}

	log.Println("myVar:", myVar.printFirstName(), "myVar2:", myVar2.printFirstName())
}
