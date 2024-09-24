package leetcode

import "fmt"

func addStrings(num1 string, num2 string) string {
	maxLen := max(len(num1), len(num2))
	result := make([]byte, maxLen+1)
	var carry byte

	for i := 0; i < maxLen; i++ {
		var x, y byte
		if i < len(num1) {
			x = num1[len(num1)-1-i] - '0'
		}

		if i < len(num2) {
			y = num2[len(num2)-1-i] - '0'
		}
		sum := x + y + carry
		carry = sum / 10
		result[len(result)-1-i] = sum%10 + '0'
	}
	if carry > 0 {
		result[0] = carry + '0'
	} else {
		result = result[1:]
	}
	return string(result)
}

func RunLC415() {
	fmt.Println(addStrings("11", "123"))
}
