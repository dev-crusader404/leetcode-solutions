package leetcode

func minDeletions(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	countMap := make(map[int]struct{})
	var numDeletion int

	for _, v := range m {
		isPresent := true
		for isPresent {
			if _, ok := countMap[v]; ok {
				v--
				numDeletion++
			} else {
				isPresent = false
				if v != 0 {
					countMap[v] = struct{}{}
				}
			}
		}
	}
	return numDeletion
}
