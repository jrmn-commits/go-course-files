package main

import (
	"log"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	birthDate   time.Time
}

func main() {
	user := &User{
		firstName:   "John",
		lastName:    "Doe",
		phoneNumber: "123-456-7890",
		age:         25,
		birthDate:   time.Now(),
	}

	log.Println(user.firstName, user.lastName, user.phoneNumber, "Birth Date:", user.birthDate)
}
