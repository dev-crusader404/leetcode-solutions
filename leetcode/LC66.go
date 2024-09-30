package leetcode

// Quick solution
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}

func plusOne2(digits []int) []int {
	digits[len(digits)-1] += 1
	var carry int
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}
	if carry == 0 {
		return digits
	}
	res := make([]int, 0)
	res = append(res, carry)
	res = append(res, digits[0:]...)
	return res
}
