package leetcode

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countMap := make(map[rune]int)
	for _, v := range s {
		countMap[v]++
	}
	for _, v := range t {
		if _, ok := countMap[v]; ok {
			countMap[v]--
			if countMap[v] == 0 {
				delete(countMap, v)
			}
		} else {
			return false
		}
	}
	return len(countMap) == 0
}
