package leetcode

func areAlmostEqual(s1 string, s2 string) bool {
	var count int
	idx := make([]int, 0)
	if s1 == s2 {
		return true
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			idx = append(idx, i)
			count++
		}
		if count > 2 {
			return false
		}
	}
	return count == 0 || (count == 2 && s1[idx[0]] == s2[idx[1]] && s1[idx[1]] == s2[idx[0]])
}
