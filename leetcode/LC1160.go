package leetcode

func countCharacters(words []string, chars string) int {
	charArray := make(map[rune]int)
	for _, v := range chars {
		charArray[v]++
	}
	var total, breakIdx int
	for _, word := range words {
		found := true
		breakIdx = len(word)
		for i, v := range word {
			c, ok := charArray[v]
			if !ok || c <= 0 {
				found = false
				breakIdx = i
				break
			}
			charArray[v]--
		}

		if found {
			total += len(word)
		}

		for i, v := range word {
			if i >= breakIdx {
				break
			}
			charArray[v]++
		}
	}
	return total
}
