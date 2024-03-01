package main

import (
	"fmt"
	"log"
	"structs/user"
)

func main() {
	// Initialize an empty user
	u := user.User{}

	// Get user data
	err := u.GetUserData()
	if err != nil {
		log.Fatalf("Error getting user data: %v", err)
	}

	// Print user information
	fmt.Println("User Information:")
	u.PrintUserData()

	// Create a new admin
	admin := user.NewAdmin("admin@example.com", "admin1234")

	// Print admin information
	fmt.Println("\nAdmin Information:")
	admin.PrintAdminData()
}
