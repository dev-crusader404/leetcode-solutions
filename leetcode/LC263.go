package leetcode

func IsUgly(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else if n%2 == 0 {
		return IsUgly(n / 2)
	} else if n%3 == 0 {
		return IsUgly(n / 3)
	} else if n%5 == 0 {
		return IsUgly(n / 5)
	} else {
		return false
	}
}

func IsUgly2(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	resultChan := make(chan int, 3)
	pf := []int{2, 3, 5}
	for _, v := range pf {
		go func(n, factor int) {
			for ; n%factor == 0; n /= factor {
			}
			resultChan <- n
		}(n, v)
	}

	go func() {
		defer close(resultChan)
	}()

	num := 1
	for v := range resultChan {
		num *= v
	}
	return num == n*n
}
