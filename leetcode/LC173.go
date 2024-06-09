package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Node []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	itr := BSTIterator{}

	for root != nil {
		itr.Node = append(itr.Node, root)
		root = root.Left
	}
	return itr
}

func (this *BSTIterator) Next() int {
	n := len(this.Node)
	current := this.Node[n-1]
	pop := current.Val
	this.Node = this.Node[:n-1]
	current = current.Right
	for current != nil {
		this.Node = append(this.Node, current)
		current = current.Left
	}
	return pop
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Node) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
