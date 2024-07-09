package leetcode

func findTheDifference2(s string, t string) byte {
	var s1, s2 int
	for _, v := range s {
		s1 += int(v)
	}

	for _, k := range t {
		s2 += int(k)
	}
	return byte(s2 - s1)
}

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
