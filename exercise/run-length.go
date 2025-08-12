package exercise

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) < 2 {
		return input
	}
	var sb strings.Builder
	var count int
	for i := 0; i < len(input); i++ {
		count++
		if i == len(input)-1 || (i < len(input)-1 && input[i] != input[i+1]) {
			if count > 1 {
				sb.WriteString(strconv.Itoa(count))
			}
			sb.WriteByte(input[i])
			count = 0
		}
	}
	return sb.String()
}

func RunLengthDecode(input string) string {
	if len(input) < 2 {
		return input
	}
	var left, num int
	var digitFound bool
	var sb strings.Builder
	for i, c := range input {
		if unicode.IsDigit(c) {
			if !digitFound {
				left = i
				digitFound = true
			}
			continue
		}
		if digitFound {
			num, _ = strconv.Atoi(input[left:i])
			for num > 0 {
				sb.WriteRune(c)
				num--
			}
		} else {
			sb.WriteRune(c)
		}
		digitFound = false
	}
	return sb.String()
}
