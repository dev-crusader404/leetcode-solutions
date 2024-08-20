package leetcode

import "fmt"

func compress(chars []byte) int {
	idx := 0
	for right := range chars {
		left := right
		for right < len(chars) && chars[left] == chars[right] {
			right++
		}
		chars[idx] = chars[left]
		idx++
		var countChar string
		if right-left > 1 {
			countChar = fmt.Sprintf("%d", right-left)
		}
		for _, v := range countChar {
			chars[idx] = byte(v)
			idx++
		}
	}
	fmt.Println(chars)
	return idx
}
