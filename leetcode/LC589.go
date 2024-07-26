package leetcode

type NodeN struct {
	Val      int
	Children []*NodeN
}

func preorderN(root *NodeN) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	bfsN(root, &result)
	return result
}

func bfsN(node *NodeN, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	for _, v := range node.Children {
		bfsN(v, result)
	}
}
