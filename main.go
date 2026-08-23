package main

import (
	"fmt"
	// "reflect"
	// "strconv"
)

func main() {
	// var first_name, last_name string = "Abhradip", "Paul"
	// var (
	// 	age       int     = 24
	// 	net_worth float32 = 2.66
	// )
	// is_adult := true
	// fmt.Print("Hello", "\n")
	// fmt.Println("World")
	// fmt.Printf("The name of the user is %v %s \n", first_name, last_name)
	// fmt.Printf("Age of the user is %d \n", age)
	// fmt.Println("Is the user adult", is_adult)
	// fmt.Println("The networth of the user is", net_worth)
	// var (
	// 	random_string string
	// 	random_number int
	// )
	// fmt.Print("Enter random string and random number: ")
	// count, err := fmt.Scanf("%s %d", &random_string, &random_number)
	// fmt.Println("Count:", count)
	// fmt.Println("Err:", err)
	// fmt.Printf("Here is the random string: %s %d \n", random_string, random_number)
	// fmt.Printf("Data type of the variable %s is %T \n", random_string, random_string)
	// fmt.Printf("Type: %v \n", reflect.TypeOf(8.8))
	// fmt.Println(int(net_worth), reflect.TypeOf(strconv.Itoa(age)))

	// Array
	// marks := [3]int{10, 20, 30}
	// fmt.Println(marks)
	// var grades [5]int = [5]int{10, 20, 30, 40, 50}
	// fmt.Println(grades)
	// fruits := [...]string{"apples", "oranges"}
	// for index, element := range fruits {
	// 	fmt.Println(index, "=>", element)
	// }
	// fmt.Println(len(fruits))
	fmt.Println("Sum of two number is", addNumbers(1, 1))
}

func addNumbers(a int, b int) int {
	return a + b
}
