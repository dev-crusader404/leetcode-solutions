package leetcode

func maxProduct1464(nums []int) int {
	var n1, n2 int
	for _, n := range nums {
		if n >= n1 {
			n2 = n1
			n1 = n
		} else if n >= n2 {
			n2 = n
		}
	}
	return (n1 - 1) * (n2 - 1)
}
