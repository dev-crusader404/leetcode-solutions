package leetcode

import "strconv"

func calPoints(operations []string) int {
	var result int
	stack := []int{}
	for _, v := range operations {
		n := len(stack)
		switch v {
		case "C":
			stack = stack[:n-1]
		case "D":
			stack = append(stack, stack[n-1]*2)
		case "+":
			stack = append(stack, stack[n-1]+stack[n-2])
		default:
			d, _ := strconv.Atoi(v)
			stack = append(stack, d)
		}
	}
	for _, val := range stack {
		result += val
	}
	return result
}
