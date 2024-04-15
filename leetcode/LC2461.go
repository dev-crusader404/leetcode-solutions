package leetcode

func maximumSubarraySum(nums []int, k int) int64 {
	uniqueMap := make(map[int]int)
	var left int
	var sum int64 = 0
	for right, val := range nums {
		if _, ok := uniqueMap[val]; !ok {
			uniqueMap[val]++
		}

		if right-left+1 == k {
			if len(uniqueMap) == k {
				sum = getMaxSum(uniqueMap, sum)
			}
			uniqueMap[nums[left]]--
			if uniqueMap[nums[left]] == 0 {
				delete(uniqueMap, val)
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
