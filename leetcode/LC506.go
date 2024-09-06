package leetcode

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	result := make([]string, len(score))
	posScore := make([][2]int, len(score))
	for i := range score {
		posScore[i][0] = score[i]
		posScore[i][1] = i
	}
	sort.Slice(posScore, func(i, j int) bool { return posScore[i][0] > posScore[j][0] })
	for i := range score {
		if i == 0 {
			result[posScore[i][1]] = "Gold Medal"
		} else if i == 1 {
			result[posScore[i][1]] = "Silver Medal"
		} else if i == 2 {
			result[posScore[i][1]] = "Bronze Medal"
		} else {
			result[posScore[i][1]] = strconv.Itoa(i + 1)
		}
	}
	return result
}
