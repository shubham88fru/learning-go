package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

//attach this method to the user struct.
//pass by value
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

//pass by ref.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

//constructor pattern
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == ""  {
		return nil, errors.New("FirstName, last name and birth date are all required")
	}

	return &(User {
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}), nil
}