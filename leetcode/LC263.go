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
