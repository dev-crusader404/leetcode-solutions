package leetcode

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]struct{})
	for _, r := range jewels {
		m[r] = struct{}{}
	}
	var count int
	for _, v := range stones {
		if _, ok := m[v]; ok {
			count++
		}
	}
	return count
}
