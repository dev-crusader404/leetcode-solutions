package leetcode

func maxProfit(prices []int) int {
	minm := prices[0]
	var maxProfit int
	if len(prices) < 2 {
		return maxProfit
	}
	maxProfit = max(prices[1]-prices[0], 0)
	for i := 1; i < len(prices); i++ {
		if prices[i]-minm > maxProfit {
			maxProfit = prices[i] - minm
		}
		minm = min(minm, prices[i])
	}
	return maxProfit
}
