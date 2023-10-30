package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var whatToSay string = "Goodbye World!"
	var i int = 7

	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, otherThingSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, otherThingSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}