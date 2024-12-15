package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings [] string) ([] float64, error) {
	floats := make([] float64, 0)
	for _, line := range strings {
		floatPrice, err  := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed.")
			fmt.Println(err)
			return nil, errors.New("converting price to float failed")
		}
		floats = append(floats, floatPrice)
	}

	return floats, nil
}