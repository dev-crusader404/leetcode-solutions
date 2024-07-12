package leetcode

import "fmt"

func numberOfSpecialChars(word string) int {
	var count int
	m := make(map[rune]int)

	for i, v := range word {
		if _, ok := m[v]; ok && v < 97 {
			continue
		}
		m[v] = i
	}

	for i, v := range word {
		if v < 97 {
			if index, ok := m[v+32]; ok {
				if index < i {
					count++
					delete(m, v+32)
				}
			}
		}
	}
	return count
}

func RunLC3121() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
}
