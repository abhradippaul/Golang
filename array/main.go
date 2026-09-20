package main

import "fmt"

func main() {
	prices := []float64{10.2, 2.5, 4.8, 7.6}
	newPrices := prices[1:]
	fmt.Println(newPrices)
	fmt.Println(len(newPrices), cap(newPrices))
	updatedPrice := append(prices, 5.9)
	fmt.Println(updatedPrice)
}
