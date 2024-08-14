package leetcode

import "fmt"

func sumSubarrayMins(arr []int) int {
	q := []int{}
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	defaultArray(&left, -1)
	defaultArray(&right, n)

	for i := 0; i < n; i++ {
		for len(q) > 0 && arr[q[len(q)-1]] >= arr[i] {
			right[q[len(q)-1]] = i
			q = q[:len(q)-1]
		}
		if len(q) > 0 {
			left[i] = q[len(q)-1]
		}
		q = append(q, i)
	}

	var sum int
	var large int = 1e9 + 7
	for i, v := range arr {
		sum += v * (i - left[i]) * (right[i] - i) % large
	}
	return sum % large
}

func defaultArray(arr *[]int, v int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = v
	}
}

func RunLC907() {
	a := []int{2, 9, 7, 8, 3, 4, 6, 1}
	fmt.Println(sumSubarrayMins(a))
}
