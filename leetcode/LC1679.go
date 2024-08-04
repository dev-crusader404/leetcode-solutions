package leetcode

func maxOperations(nums []int, k int) int {
	if len(nums) < 1 {
		return 0
	}
	m := make(map[int]int)
	count := 0
	for _, v := range nums {
		if _, ok := m[v]; ok && m[v] > 0 {
			m[v]--
			count++
		} else {
			m[k-v]++
		}
	}
	return count
}
