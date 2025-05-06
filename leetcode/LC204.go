package leetcode

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] == false {
			count++
			for j := 2; i*j < n; j++ {
				isPrime[i*j] = true
			}
		}
	}
	return count
}
