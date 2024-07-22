package leetcode

import "fmt"

func maxDivScore(nums []int, divisors []int) int {
	maxCount, divisor := 0, divisors[0]

	for _, i := range divisors {
		count := 0
		for _, j := range nums {
			if j%i == 0 {
				count++
			}
		}
		if count >= maxCount {
			if count == maxCount {
				divisor = min(i, divisor)
			} else {
				divisor = i
			}
			maxCount = count
		}
	}
	return divisor
}

func RunLC2644() {
	a := []int{73, 13, 20, 6}
	b := []int{56, 75, 83, 26, 24, 53, 56, 61}
	fmt.Println(maxDivScore(a, b))
}
