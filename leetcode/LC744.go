package leetcode

import "fmt"

// Time complexity - O(N)
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

// Time complexity - O(logN)
func nextGreater2(letters []byte, target byte) byte {
	last := letters[len(letters)-1]
	if target == 'z' || last <= target {
		return letters[0]
	}

	low, high := 0, len(letters)-1

	for low < high {
		mid := (low + high) / 2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return letters[low]
}

func NextGreaterLetter() {
	list := []byte{'c', 'f', 'j'}
	fmt.Println(string(nextGreatestLetter(list, 'c')))
	fmt.Println(string(nextGreater2(list, 'c')))
}
