// user.go

package user

import (
	"errors"
	"fmt"
	"time"
)

// User represents a user's information
type User struct {
	FirstName string    // First name of the user
	LastName  string    // Last name of the user
	Birthdate string    // Birthdate of the user (MM/DD/YYYY)
	CreatedAt time.Time // Time when the user was created
}

// Admin represents an administrator's information
type Admin struct {
	Email    string // Email of the admin
	Password string // Password of the admin
	User     User   // Embedded User struct for admin's basic information
}

// GetUserData prompts the user to input their information
func (u *User) GetUserData() error {
	u.FirstName = getUserInput("Please enter your first name: ")
	u.LastName = getUserInput("Please enter your last name: ")
	u.Birthdate = getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	// Check if all fields are provided
	if u.FirstName == "" || u.LastName == "" || u.Birthdate == "" {
		return errors.New("all fields are required")
	}

	// Record the time when the user data was created
	u.CreatedAt = time.Now()

	return nil
}

// PrintUserData prints the user's information
func (u *User) PrintUserData() {
	fmt.Printf("Name: %s %s, Birthdate: %s\n", u.FirstName, u.LastName, u.Birthdate)
}

// PrintAdminData prints the admin's information
func (a *Admin) PrintAdminData() {
	fmt.Printf("Name: %s %s, Birthdate: %s, Email: %s, Created At: %s\n", a.User.FirstName, a.User.LastName, a.User.Birthdate, a.Email, a.User.CreatedAt.Format(time.RFC3339))
}

// getUserInput prompts the user for input and returns the value
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}

// NewAdmin creates a new admin with default values
func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			Birthdate: "01/01/2000",
		},
	}
}
