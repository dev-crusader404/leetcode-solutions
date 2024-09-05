package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	var result [][]int
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	seen := make(map[int]struct{})
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}
	var missing []int
	for _, v := range nums1 {
		_, s := seen[v]
		if _, ok := m2[v]; !ok && !s {
			missing = append(missing, v)
			seen[v] = struct{}{}
		}
		m1[v] = struct{}{}
	}
	seen = make(map[int]struct{})
	result = append(result, missing)
	missing = nil
	for _, v := range nums2 {
		_, s := seen[v]
		if _, ok := m1[v]; !ok && !s {
			missing = append(missing, v)
			seen[v] = struct{}{}
		}
	}
	result = append(result, missing)
	return result
}
