package main

import "log"

func main() {
// SWITCH STATEMENTS
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("I'm a cat!")
	case "dog":
		log.Println("I'm a dog!")
	case "fish":
		log.Println("I'm a fish!")
	default:
		log.Println("I don't know what that is!")
	}
	// INT IF STATEMENTS
	//	myNum := 10
	//	isTrue := false

	//	if myNum > 99 && !isTrue {
	//		log.Println("myNum is greater than 99 and isTrue is not true")
	//	} else if myNum < 100 && isTrue {
	//		log.Println("myNum is less than 100 and isTrue is true")
	//	} else if myNum == 100 || isTrue {
	//		log.Println("myNum is 100 or isTrue is true")
	//	}

	// VAR IF STATEMENTS
	//	cat := "cat"

	//	if cat == "cat" {
	//		log.Println("I'm a cat")
	//	} else {
	//		log.Println("I'm not a cat")
	//	}

	// BOOL IF STATEMENTS
	//	var isTrue bool

	//	isTrue = false

	//	if isTrue == true{
	//		log.Println("is True is", isTrue)
	//	} else {
	//		log.Println("is True is", isTrue)
	//	}
}
