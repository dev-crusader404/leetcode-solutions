package leetcode

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	var sb strings.Builder
	s = strings.TrimSpace(s)
	arr := strings.Fields(s)

	for i := len(arr) - 1; i >= 0; i-- {
		sb.WriteString(arr[i])
		if i > 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func reverseWords2(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.TrimSpace(s)
	var sb strings.Builder
	for i, v := range s {
		if v == ' ' && i+1 < len(s) && s[i+1] == ' ' {
			continue
		}
		sb.WriteRune(v)
	}
	arr := strings.Split(sb.String(), " ")

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return strings.Join(arr, " ")
}

func RunLC151() {
	s := "example   good   a"
	fmt.Println(reverseWords(s))
}
