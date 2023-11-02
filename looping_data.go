package main

import "log"

func main() {
	type User struct {
		firstName string
		lastName  string
		age       int
		email     string
	}

	var users []User
	users = append(users, User{"John", "Doe", 30, "john@gmail.com"})
	users = append(users, User{"Jane", "Doe", 20, "jane@gmail.com"})
	users = append(users, User{"John", "Doe", 30, "john@gmail.com"})
	users = append(users, User{"Jane", "Doe", 20, "jane@gmail.com"})

	for _, l := range users {
		log.Println(l.firstName, l.lastName, l.age, l.email)
    }
}