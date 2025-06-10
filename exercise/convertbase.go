package exercise

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var result []int
	var decimalValue int
	var err error
	if inputBase < 2 {
		return result, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return result, errors.New("output base must be >= 2")
	}
	if inputBase == outputBase {
		return inputDigits, nil
	}

	decimalValue, err = convertToDecimal(inputBase, inputDigits)
	if err != nil {
		return result, err
	}
	list := convertToOutputBase(decimalValue, outputBase)
	if len(list) == 0 {
		return append(result, 0), nil
	}
	return list, nil
}

func convertToDecimal(base int, inputDigits []int) (int, error) {
	var total int
	size := len(inputDigits)
	for i, ele := range inputDigits {
		if ele >= base || ele < 0 {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		total += ele * int(math.Pow(float64(base), float64(size-i-1)))
	}
	return total, nil
}

func convertToOutputBase(total, opBase int) []int {
	temp := make([]int, 0)
	i := 0
	for total != 0 {
		temp = append(temp, total%opBase)
		total /= opBase
		i++
	}
	var result []int
	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}
	return result
}
