package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	wordArray := strings.Split(sentence, " ")
	for i, w := range wordArray {
		if w[0] != '$' {
			continue
		}
		value, err := strconv.Atoi(w[1:])
		if err != nil {
			continue
		}
		discountPrice := float64(value*(100-discount)) / float64(100)
		wordArray[i] = "$" + fmt.Sprintf("%.2f", discountPrice)
	}
	return strings.Join(wordArray, " ")
}
