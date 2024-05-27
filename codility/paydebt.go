package codility

import (
	"fmt"
	"sort"
)

type Combo struct {
	debts    float64
	interest float64
}

func solution(s int, debts []int, interests []int) float64 {
	Arr := make([]Combo, 0)
	for i := range debts {
		Arr = append(Arr, Combo{
			debts:    float64(debts[i]),
			interest: float64(interests[i]),
		})
	}
	sort.SliceStable(Arr, func(i, j int) bool {
		return Arr[i].interest > Arr[j].interest
	})
	var totalDebts float64
	totalDebts = calculateDebt(&Arr, 0)
	var totals float64
	var month int

	for totalDebts != 0 {
		spend := 0.1 * float64(s)

		if month > 0 {
			totalDebts = calculateDebt(&Arr, month)
		}
		for i := range Arr {

			if Arr[i].debts == 0 {
				continue
			} else if spend == 0 {
				break
			} else if float64(Arr[i].debts) > spend {
				Arr[i].debts -= spend
				totalDebts -= spend
				totals += spend
				break
			} else {
				totalDebts -= Arr[i].debts
				spend -= Arr[i].debts
				totals += Arr[i].debts
				Arr[i].debts = 0
			}
		}
		month++
	}
	return totals
}

func calculateDebt(arr *[]Combo, c int) float64 {
	var t float64

	for i := 0; i < len(*arr); i++ {
		if c > 0 {
			(*arr)[i].debts += (*arr)[i].debts * ((*arr)[i].interest / 100)
		}
		t += (*arr)[i].debts
	}
	return t
}

func RunTest() {
	s := 50
	debt := []int{2, 2, 5}
	interest := []int{200, 100, 150}
	fmt.Println(solution(s, debt, interest))
}
