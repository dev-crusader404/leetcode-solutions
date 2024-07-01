package leetcode

// T: O(N) 		S: O(N)
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return -1
	}
	n := len(cost) - 1
	result := make([]int, len(cost))
	result[n] = cost[n]
	result[n-1] = cost[n-1]
	for i := n - 2; i >= 0; i-- {
		result[i] = cost[i] + min(result[i+1], result[i+2])
	}

	return min(result[0], result[1])
}

// T: O(N)		S: O(1)
func minCostClimbingStairs2(cost []int) int {
	if len(cost) < 2 {
		return -1
	}
	topOne := cost[0]
	topTwo := cost[1]
	for i := 2; i < len(cost); i++ {
		curr := cost[i] + min(topOne, topTwo)
		topOne = topTwo
		topTwo = curr
	}

	return min(topOne, topTwo)
}
