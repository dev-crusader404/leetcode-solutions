package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(arr []any) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return &TreeNode{}
	}
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 0
	for i < (len(arr)-1)/2 {
		node := queue[0]
		queue = queue[1:]
		left := 2*i + 1
		right := 2*i + 2
		node.Left = setChild(node.Left, arr[left])
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		node.Right = setChild(node.Right, arr[right])
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func setChild(node *TreeNode, val any) *TreeNode {
	if val == nil {
		return nil
	}
	node = &TreeNode{Val: val.(int)}
	return node
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}
