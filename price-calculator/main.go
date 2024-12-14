package main

import (
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate)
		job.Process()
	}

}