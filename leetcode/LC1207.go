package leetcode

func uniqueOccurrences(arr []int) bool {
	cMap := make(map[int]int, len(arr))
	for _, v := range arr {
		cMap[v]++
	}
	seen := make(map[int]struct{})
	for _, v := range cMap {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}
