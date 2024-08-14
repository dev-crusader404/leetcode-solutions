package leetcode

import "fmt"

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

func finalPrices2(prices []int) []int {
	stk := make([]int, 0, len(prices))
	for i, v := range prices {
		for len(stk) > 0 && v <= prices[stk[len(stk)-1]] {
			prices[stk[len(stk)-1]] -= v
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return prices
}

func RunLC1475() {
	a := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices2(a))
}
