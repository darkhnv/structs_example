package user

import (
	"errors"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
}

func (u *User) GetUserData() error {
	u.FirstName = getUserInput("Please enter your first name: ")
	u.LastName = getUserInput("Please enter your last name: ")
	u.Birthdate = getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	if u.FirstName == "" || u.LastName == "" || u.Birthdate == "" {
		return errors.New("all fields are required")
	}

	return nil
}

func (u *User) PrintUserData() {
	fmt.Printf("Name: %s %s, Birthdate: %s\n", u.FirstName, u.LastName, u.Birthdate)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
