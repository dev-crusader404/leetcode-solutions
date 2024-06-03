package codesignal

import "fmt"

func solution1(sequence []int) bool {
	if len(sequence) == 1 {
		return true
	}

	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			if i == 0 {
				return isIncreasing(sequence[i+1:])
			}
			if (i + 1) == len(sequence)-1 {
				return true
			}

			return check(sequence[i-1], sequence[i+1:]) || check(sequence[i], sequence[i+2:])
		}
	}

	return true
}

func check(x int, a []int) bool {
	if x >= a[0] {
		return false
	}
	return isIncreasing(a)
}

func isIncreasing(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func RunInc() {
	a := []int{1, 2, 3, 4, 3, 6}
	fmt.Println(solution1(a))
}
