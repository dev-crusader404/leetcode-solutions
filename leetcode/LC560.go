package leetcode

func SubarraySum(nums []int, k int) int {
	resultCount := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1

	sum := 0
	for _, v := range nums {
		sum += v
		if c, ok := sumMap[sum-k]; ok {
			resultCount += c
		}
		sumMap[sum]++
	}
	return resultCount
}
