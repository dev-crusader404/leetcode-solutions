package leetcode

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	gMap := make(map[int]int)

	for _, v := range nums2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			gMap[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}

	for i, v := range nums1 {
		if x, ok := gMap[v]; ok {
			nums1[i] = x
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for i, v := range nums2 {
		m1[v] = i
	}
	var result []int
	for _, n := range nums1 {
		if index, ok := m1[n]; ok && index != len(nums2)-1 {
			index++
			for index <= (len(nums2)-1) && n > nums2[index] {
				index++
			}
			if index == len(nums2) {
				result = append(result, -1)
			} else {
				result = append(result, nums2[index])
			}
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func RunLC496() {
	a := []int{4, 1, 2}
	b := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(a, b))
	fmt.Println(nextGreaterElement2(a, b))
}
