package leetcode

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	last := letters[len(letters)-1]
	if target == 'z' || last < target {
		return letters[0]
	}

	for _, v := range letters {
		if v > target {
			return v
		}
	}
	return letters[0]
}

func NextGreaterLetter() {
	list := []byte{'c', 'f', 'j'}
	fmt.Println(string(nextGreatestLetter(list, 'c')))
}
