package main

import ("fmt")

func main() {
	var numPointer *int
	num := 10
	numPointer = &num
	fmt.Println("The num value is", num)
	fmt.Println("The pointer value is", *numPointer)
	fmt.Println("The pointer is", numPointer)
}