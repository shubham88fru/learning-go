package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPricesJob  struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPricesJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Failed to read file")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err  := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed.")
			fmt.Println(err)
			return
		}
		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPricesJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price*(1+job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPricesJob {
	return &(TaxIncludedPricesJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	})
}