package main

import (
	"fmt"
	"math"
	"example.com/m/v2/fileops"
)

const inflationRate = 2.5
const fileName = "investment.txt"

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

func main() {
	fmt.Println("Hello World")
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64 = 10

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// fmt.Println("The future value would be", futureValue)
	// fmt.Printf("The future value would be %.2f \n", futureValue)
	// fmt.Println("The future real value would be", futureRealValue)
	fileops.WriteToFile(fileName, investmentAmount, futureValue, futureRealValue)
	fileops.ReadFromFile(fileName)
	communicate()
}
