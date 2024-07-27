package leetcode

func numDifferentIntegers(word string) int {
	seen := make(map[string]struct{})
	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			continue
		}
		numStart := i
		for numStart < len(word) {
			if word[numStart] >= 'a' && word[numStart] <= 'z' {
				break
			}
			numStart++
		}
		seen[word[i:numStart]] = struct{}{}
		i = numStart
	}
	return len(seen)
}
