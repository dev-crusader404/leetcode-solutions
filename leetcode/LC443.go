package leetcode

import "fmt"

func compress(chars []byte) int {
	idx, right := 0, 0
	for right < len(chars) {
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
	return idx
}

func RunLC443() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}
