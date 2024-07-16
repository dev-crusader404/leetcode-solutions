package leetcode

import "fmt"

func minimumOperations(nums []int) int {
	odd, even := make(map[int]int), make(map[int]int)
	var maxOdd, secOdd, secEven, maxEven, oddEle, evenEle int
	for i, v := range nums {
		if i%2 == 0 {
			odd[v]++
			c, _ := odd[v]
			if c > maxOdd {
				secOdd = maxOdd
				maxOdd = c
				oddEle = v
			}
		} else {
			even[v]++
			c, _ := even[v]
			if c > maxEven {
				secEven = maxEven
				maxEven = c
				evenEle = v
			}
		}
	}

	if evenEle != oddEle {
		return len(nums) - maxOdd - maxEven
	}
	return min(len(nums)-maxOdd-secEven, len(nums)-secOdd-maxEven)
}

func RunLC2170() {
	n := []int{2, 2, 2, 2}
	fmt.Println(minimumOperations(n))
}
