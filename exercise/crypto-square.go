package exercise

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	pt = normalizeInput(pt)
	length := len(pt)
	if length <= 1 {
		return pt
	}
	squaredown := math.Sqrt(float64(length))
	var row, column int
	row = int(squaredown)
	if squaredown == math.Floor(squaredown) {
		column = row
	} else {
		column = row + 1
	}
	row = int(math.Round(squaredown))
	arr := make([]byte, row*column)
	rIdx, cIdx := 0, 0
	for i := 0; i < length; i++ {
		if cIdx >= column {
			rIdx++
			cIdx = 0
		}
		arr[(row*cIdx)+rIdx] = pt[i]
		cIdx++
	}
	fillZeroByte(arr)
	resultArr := getResultArray(arr, row, column)
	return strings.Join(resultArr, " ")
}

func getResultArray(arr []byte, row, column int) []string {
	var result []string
	for i := 0; i < column; i++ {
		start := i * row
		end := start + row
		word := arr[start:end]
		result = append(result, string(word))
	}
	return result
}

func fillZeroByte(arr []byte) {
	for i := range arr {
		if arr[i] == 0x00 {
			arr[i] = ' '
		}
	}
}

func normalizeInput(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	return re.ReplaceAllString(s, "")
}
