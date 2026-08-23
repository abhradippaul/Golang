package main

import "fmt"

func main() {
	var first_name, last_name string = "Abhradip", "Paul"
	var (
		age       int     = 24
		net_worth float32 = 2.66
	)
	is_adult := true
	fmt.Print("Hello", "\n")
	fmt.Println("World")
	fmt.Printf("The name of the user is %v %s \n", first_name, last_name)
	fmt.Printf("Age of the user is %d \n", age)
	fmt.Println("Is the user adult", is_adult)
	fmt.Println("The networth of the user is", net_worth)
}
