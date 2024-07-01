package leetcode

import "math"

// Implementing soltion using Bellman Ford Algorithm
func networkDelayTime3(times [][]int, n int, k int) int {
	distance := make([]int, n)
	for i := range distance {
		if i != k-1 {
			distance[i] = math.MaxInt
		}
	}

	for i := 0; i < n-1; i++ {
		updateCount := 0
		for _, v := range times {
			if distance[v[0]-1] != math.MaxInt && distance[v[0]-1]+v[2] < distance[v[1]-1] {
				distance[v[1]-1] = distance[v[0]-1] + v[2]
				updateCount++
			}
		}
		if updateCount == 0 {
			break
		}
	}

	var signalTime int
	for _, v := range distance {
		if v == math.MaxInt {
			return -1
		}
		signalTime = max(signalTime, v)
	}
	return signalTime
}
