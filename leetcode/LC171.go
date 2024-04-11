package leetcode

import "strings"

func TitleToNumber(columnTitle string) int {
	if len(columnTitle) == 0 {
		return 0
	}
	columnTitle = strings.ToUpper(columnTitle)

	sum := 0
	for _, v := range columnTitle {
		i := int(v) - int('A') + 1
		sum = sum*26 + i
	}
	return sum
}
