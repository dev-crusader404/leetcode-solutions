package leetcode

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]byte, len(num1)+len(num2))
	m, n := len(num1)-1, len(num2)-1

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			product := (num1[i]-'0')*(num2[j]-'0') + result[i+j+1]
			result[i+j+1] = (product % 10)
			result[i+j] += (product / 10)
		}
	}
	if result[0] == byte(0) {
		result = result[1:]
	}
	for i := range result {
		result[i] += '0'
	}
	return string(result)
}

func RunLC43() {
	fmt.Println(multiply("123", "456"))
}
