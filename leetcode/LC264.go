package leetcode

import "fmt"

func nthUglyNumber(n int) int {
	l, j, k := 2, 3, 5
	var x, y, z int
	ugly := make([]int, n)
	ugly[0] = 1
	for i := 1; i < n; i++ {
		m := min(min(l, j), k)
		ugly[i] = m
		if m == l {
			x++
			l = 2 * ugly[x]
		}
		if m == j {
			y++
			j = 3 * ugly[y]
		}
		if m == k {
			z++
			k = 5 * ugly[z]
		}
	}
	return ugly[n-1]
}

func RunLC264() {
	fmt.Println(nthUglyNumber(10))
}
