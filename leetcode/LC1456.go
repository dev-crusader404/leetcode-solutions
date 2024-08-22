package leetcode

func maxVowels(s string, k int) int {
	var left, vowelCount, count int

	for right, c := range s {
		if isVowel(byte(c)) {
			count++
		}
		if right-left == k {
			if isVowel(s[left]) {
				count--
			}
			left++
		}
		vowelCount = max(vowelCount, count)
	}
	return vowelCount
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}
