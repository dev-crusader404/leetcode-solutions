package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return bst(0, len(nums)-1, nums)
}

func bst(low, high int, nums []int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	var root TreeNode
	root.Val = nums[mid]
	root.Left = bst(low, mid-1, nums)
	root.Right = bst(mid+1, high, nums)
	return &root
}
