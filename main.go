package main

import (
	"fmt"
	"log"
	"structs/user"
)

func main() {
	u := user.User{}

	err := u.GetUserData()
	if err != nil {
		log.Fatalf("Error getting user data: %v", err)
	}

	fmt.Println("User Information:")
	u.PrintUserData()
}
