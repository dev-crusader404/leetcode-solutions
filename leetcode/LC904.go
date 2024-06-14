package leetcode

import "fmt"

func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}
	fruitMap := make(map[int]int)
	var left, maxLen int
	k := 2
	for right, val := range fruits {
		fruitMap[val]++
		for len(fruitMap) > k {
			fruitMap[fruits[left]]--
			if fruitMap[fruits[left]] == 0 {
				delete(fruitMap, fruits[left])
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func RunLC904() {
	a := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(a))
}
