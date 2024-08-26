package leetcode

func moveZeroes(nums []int) {
	k := 0
	for _, v := range nums {
		if v != 0 {
			nums[k] = v
			k++
		}
	}

	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
