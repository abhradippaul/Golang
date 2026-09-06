package main

import (
	"fmt"
	"time"

	"example.com/m/user"
)

func main() {
	fmt.Println("Custom datatypes")
	var appUser *user.User
	appUser, err := user.CreateUser("Abhradip", "Paul", time.Now())
	if err != nil {
		panic(err)
	}
	appUser.OutputUserDetails()
	appUser.ClearUserDetails()
	appUser.OutputUserDetails()
}
