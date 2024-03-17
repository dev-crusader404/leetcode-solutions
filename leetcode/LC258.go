package leetcode

func AddDigits(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num != 0 {
		rem := num % 10
		num = num / 10
		sum += rem
	}
	return AddDigits(sum)
}

func AddDigits2(num int) int {
	if num < 10 {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
