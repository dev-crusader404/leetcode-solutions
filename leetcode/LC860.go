package leetcode

func lemonadeChange(bills []int) bool {
	cash := [3]int{0, 0, 0}
	if bills[0] != 5 {
		return false
	}
	for _, v := range bills {
		bill := getIndex(v)
		cash[bill]++
		bill--
		v -= 5
		for v != 0 && cash[bill] > 0 {
			amt := v / getAmount(bill)
			if amt >= cash[bill] {
				amt = cash[bill]
				cash[bill] = 0
			} else {
				cash[bill] -= amt
			}
			v -= amt * getAmount(bill)
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
