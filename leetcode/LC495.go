package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	var total int

	poisonUntil := timeSeries[0] + duration - 1
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] > poisonUntil {
			total += duration
		} else {
			total += timeSeries[i] - timeSeries[i-1]
		}
		poisonUntil = timeSeries[i] + duration - 1
	}
	return total + duration
}
