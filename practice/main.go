package main

import "fmt"

type stringMap map[string]string

func main() {
	// fmt.Println("Practice")
	// hobbyArray := []string{"Singing", "Dancing", "Gaming"}
	// fmt.Println(hobbyArray)
	// println(hobbyArray[0])
	// newHobbyArray1 := hobbyArray[1:]
	// fmt.Println(newHobbyArray1)
	// newHobbyArray2 := hobbyArray[:2]
	// fmt.Println(newHobbyArray2)

	// map
	websites := stringMap{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedId"] = "https://linkedin.com"
	// delete(websites, "Google")
	for index, value := range websites {
		fmt.Printf("The index is %s and the value is %s", index, value)
		fmt.Println()
	}
}
