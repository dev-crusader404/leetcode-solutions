package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, c := range candies {
		maxCandy = max(maxCandy, c)
	}
	res := make([]bool, len(candies))
	for i, c := range candies {
		if c+extraCandies >= maxCandy {
			res[i] = true
		}
	}
	return res
}
