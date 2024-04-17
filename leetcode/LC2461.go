package leetcode

import "fmt"

func maximumSubarraySum(nums []int, k int) int64 {
	uniqueMap := make(map[int]int)
	var left int
	var sum int64 = 0
	for right, val := range nums {
		uniqueMap[val]++

		if right-left+1 == k {
			if len(uniqueMap) == k {
				sum = getMaxSum(uniqueMap, sum)
			}
			uniqueMap[nums[left]]--
			if uniqueMap[nums[left]] == 0 {
				delete(uniqueMap, nums[left])
			}
			left++
		}
	}
	return sum
}

func getMaxSum(m map[int]int, sum int64) int64 {
	var rslt int64
	for i, v := range m {
		rslt += int64(i * v)
	}
	if sum > rslt {
		return sum
	} else {
		return rslt
	}
}

func Run() {
	arr := []int{1, 5, 4, 2, 9, 9, 9}
	fmt.Println(maximumSubarraySum(arr, 3))
}
