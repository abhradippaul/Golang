package main

import (
	"fmt"

	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("Price Calculator")
	taxRates := []float64{0, 0.7, .1, .15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}
