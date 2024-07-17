package leetcode

import "fmt"

func minimumOperations(nums []int) int {
	odd, even := make(map[int]int), make(map[int]int)
	var maxOdd, secOdd, secEven, maxEven, oddEle, evenEle int
	for i, v := range nums {
		if i%2 == 0 {
			odd[v]++
			c := odd[v]
			if c >= maxOdd {
				if v != oddEle {
					secOdd = maxOdd
					oddEle = v
				}
				maxOdd = c
			} else if c >= secOdd {
				secOdd = c
			}
		} else {
			even[v]++
			c := even[v]
			if c >= maxEven {
				if v != evenEle {
					secEven = maxEven
					evenEle = v
				}
				maxEven = c
			} else if c >= secEven {
				secEven = c
			}
		}
	}

	if evenEle != oddEle {
		return len(nums) - maxOdd - maxEven
	}
	return min(len(nums)-maxOdd-secEven, len(nums)-secOdd-maxEven)
}

func RunLC2170() {
	// n := []int{1, 2, 2, 2, 2}
	n := []int{3, 1, 3, 2, 4, 3}
	fmt.Println(minimumOperations(n))
}
