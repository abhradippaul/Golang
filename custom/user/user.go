package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	TimeStamp time.Time
	secret    string
}

func CreateUser(firstName string, lastName string, timeStamp time.Time) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("First Name and Last Name is required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		TimeStamp: timeStamp,
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println("Firstname:", u.FirstName)
	fmt.Println("Lirstname:", u.LastName)
	fmt.Println("Timestamp:", u.TimeStamp)
	fmt.Println("Secret:", u.secret)
}

func (u *User) ClearUserDetails() {
	u.FirstName = ""
	u.LastName = ""
}
