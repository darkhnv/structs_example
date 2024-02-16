package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User     User
}

func (u *User) GetUserData() error {
	u.FirstName = getUserInput("Please enter your first name: ")
	u.LastName = getUserInput("Please enter your last name: ")
	u.Birthdate = getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	if u.FirstName == "" || u.LastName == "" || u.Birthdate == "" {
		return errors.New("all fields are required")
	}

	u.CreatedAt = time.Now()

	return nil
}

func (u *User) PrintUserData() {
	fmt.Printf("Name: %s %s, Birthdate: %s\n", u.FirstName, u.LastName, u.Birthdate)
}

func (a *Admin) PrintAdminData() {
	fmt.Printf("Name: %s %s, Birthdate: %s, Email: %s, Created At: %s\n", a.User.FirstName, a.User.LastName, a.User.Birthdate, a.Email, a.User.CreatedAt.Format(time.RFC3339))
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}

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
