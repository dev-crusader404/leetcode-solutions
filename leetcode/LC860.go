package leetcode

import "fmt"

// Solved with O(1) space and O(N) time
func lemonadeChange(bills []int) bool {
	var five, ten int

	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			five--
			ten++
		} else if ten > 0 {
			ten--
			five--
		} else {
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func lemonadeChange2(bills []int) bool {
	cash := [3]int{0, 0, 0}
	if bills[0] != 5 {
		return false
	}
	for _, v := range bills {
		bill := getIndex(v)
		cash[bill]++
		bill--
		v -= 5
		for bill >= 0 && v != 0 {
			amt := v / getBalance(bill)
			if amt >= cash[bill] {
				amt = cash[bill]
				cash[bill] = 0
			} else {
				cash[bill] -= amt
			}
			v -= amt * getBalance(bill)
			bill--
		}
		if v != 0 {
			return false
		}
	}
	return true
}

func getIndex(x int) int {
	switch x {
	case 5:
		return 0
	case 10:
		return 1
	case 20:
		return 2
	}
	return 0
}

func getBalance(x int) int {
	if x == 2 {
		return 20
	}
	return 5*x + 5
}

func RunLC860() {
	n := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	fmt.Println(lemonadeChange(n))
}
