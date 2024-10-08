package leetcode

func balancedStringSplit(s string) int {
	var balanceCount, L_Count int
	for _, c := range s {
		if c == 'L' {
			L_Count++
		} else {
			L_Count--
		}
		if L_Count == 0 {
			balanceCount++
		}
	}
	return balanceCount
}
