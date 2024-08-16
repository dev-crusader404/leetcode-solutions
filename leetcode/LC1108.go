package leetcode

import "strings"

func defangIPaddr(address string) string {
	var sb strings.Builder

	for _, v := range address {
		if v == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteRune(v)
		}
	}
	return sb.String()
}