package leetcode

func numberOfLines(widths []int, s string) []int {
	line, pixel := 1, 0

	for _, v := range s {
		wdt := widths[v-'a']
		if pixel+wdt > 100 {
			pixel = wdt
			line++
		} else {
			pixel += wdt
		}
	}
	return []int{line, pixel}
}
