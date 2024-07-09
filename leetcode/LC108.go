package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var head *TreeNode
	if len(nums) == 0 {
		return head
	}
	head = &TreeNode{Val: nums[0]}
	curr := head
	for _, v := range nums {
		bst(curr, v)
	}
	return head
}

func bst(node *TreeNode, val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			bst(node.Left, val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			bst(node.Right, val)
		}
	}
}

func RunLC108() {

}
