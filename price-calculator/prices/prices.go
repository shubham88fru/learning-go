package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPricesJob  struct {
	IOManager filemanager.FileManager
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPricesJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("error in filemenager")
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Converting price to float failed.")
		fmt.Println(err)
		return
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

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {
	return &(TaxIncludedPricesJob{
		IOManager: fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	})
}