package leetcode

//Solution with O(1) space and O(N) time complexity
//Without modifying actual array
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

//Solution with O(1) space and O(N) time complexity
//Modifying actual array
func findDuplicate2(nums []int) int {
	idx := 0
	for idx < len(nums) {
		if nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		} else {
			idx++
		}
	}

	for i, v := range nums {
		if i+1 != v {
			return v
		}
	}
	return nums[0]
}

func RunLC287() {
	n := []int{3, 3, 3, 3, 3}
	println(findDuplicate(n))
	println(findDuplicate2(n))
}
