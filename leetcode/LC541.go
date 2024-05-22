package leetcode

import "fmt"

func reverseStr(s string, k int) string {
	var left, right int
	var all bool
	b := []byte(s)
	if k > len(s) {
		right = len(s) - 1
		all = true
	}

	for i := 0; i < len(s); i = i + 2*k {
		left = i
		if !all {
			right = i + k - 1
		}

		for left < right && right < len(s) {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}

	return string(b)
}

func RunLC541() {
	s := "abcd"
	fmt.Println(reverseStr(s, 5))
}
