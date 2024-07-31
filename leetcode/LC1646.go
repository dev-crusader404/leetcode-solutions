package leetcode

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	num := make([]int, n+1)
	num[0], num[1] = 0, 1
	res := 1
	for i := 1; i <= n/2; i++ {
		x, y := 2*i, 2*i+1
		num[x] = num[i]
		res = max(res, num[x])
		if y > n {
			break
		}
		num[y] = num[i] + num[i+1]
		res = max(res, num[y])
	}
	return res
}
