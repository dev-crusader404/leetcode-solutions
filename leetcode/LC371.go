package leetcode

import "fmt"

func getSum(a int, b int) int {
	sum := a ^ b
	carry := a & b

	for carry != 0 {
		carry = carry << 1
		temp := sum
		sum = sum ^ carry
		carry = carry & temp
	}
	return sum
}

func RunLC371() {
	fmt.Println(getSum(121, 99))
}
