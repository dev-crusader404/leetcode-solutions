package leetcode

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	for _, k := range t {
		c, ok := m[k]
		if !ok || c == 0 {
			return byte(k)
		}
		m[k]--
	}
	return byte(t[0])
}
