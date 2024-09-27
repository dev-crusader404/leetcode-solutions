package leetcode

import "fmt"

func addBinary(a string, b string) string {
	var sum, carry byte
	res := make([]byte, max(len(a), len(b))+1)

	for i := 0; i < len(res)-1; i++ {
		var x, y byte
		if i < len(a) {
			x = a[len(a)-1-i] - '0'
		}
		if i < len(b) {
			y = b[len(b)-1-i] - '0'
		}
		sum = x + y + carry
		carry = sum / 2
		sum %= 2
		res[len(res)-1-i] = sum + '0'
	}

	if carry > 0 {
		res[0] = carry + '0'
	} else {
		res = res[1:]
	}
	return string(res)
}

func RunLC67() {
	fmt.Println(addBinary("1", "11"))
}
