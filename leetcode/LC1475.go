package leetcode

func finalPrices(prices []int) []int {
	stk := []int{}
	res := make([]int, len(prices))
	copy(res, prices)
	for i := len(prices) - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] > prices[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			res[i] -= stk[len(stk)-1]
		}
		stk = append(stk, prices[i])
	}
	return res
}
