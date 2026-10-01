package conversion

import (
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	prices := make([]float64, len(strings))
	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, err
		}
		prices[stringIndex] = floatPrice
	}
	return prices, nil
}
