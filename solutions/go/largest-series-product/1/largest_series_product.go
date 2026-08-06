package largestseriesproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return 0, errors.New("span must be smaller than or equal to digits length")
	}
	for _, ch := range digits {
		if !unicode.IsDigit(ch) {
			return 0, errors.New("digits input must only contain digits")
		}
	}
	if span == 0 {
		return 1, nil
	}
	var maxProduct int64 = 0
	for i := 0; i <= len(digits)-span; i++ {
		var currentProduct int64 = 1
		for j := i; j < i+span; j++ {
			digit := int64(digits[j] - '0')
			currentProduct *= digit
		}
		if currentProduct > maxProduct {
			maxProduct = currentProduct
		}
	}
	return maxProduct, nil
}