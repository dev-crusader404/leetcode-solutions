package leetcode

func largestAltitude(gain []int) int {
	prefixSum := 0
	maxGain := 0
	for i := 1; i <= len(gain); i++ {
		prefixSum += gain[i-1]
		maxGain = max(maxGain, prefixSum)
	}
	return maxGain
}
