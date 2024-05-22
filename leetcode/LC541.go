package leetcode

import "fmt"

func reverseStr(s string, k int) string {
	var left, right int
	b := []byte(s)

	for i := 0; i < len(s); i = i + 2*k {
		left, right = i, min(i+k-1, len(s)-1)

		for left < right && right < len(s) {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}

	return string(b)
}

func RunLC541() {
	s := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	fmt.Println(reverseStr(s, 39))
}
