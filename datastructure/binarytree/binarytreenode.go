package binarytree

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Data       int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func (t *TreeNode) Insert(key int) {
	if key < t.Data {
		if t.LeftChild == nil {
			t.LeftChild = &TreeNode{Data: key}
		} else {
			t.LeftChild.Insert(key)
		}
	} else if key > t.Data {
		if t.RightChild == nil {
			t.RightChild = &TreeNode{Data: key}
		} else {
			t.RightChild.Insert(key)
		}
	}
}

func (n *TreeNode) InOrderTransversal() {
	if n.LeftChild != nil {
		n.LeftChild.InOrderTransversal()
	}
	fmt.Printf("%d ", n.Data)

	if n.RightChild != nil {
		n.RightChild.InOrderTransversal()
	}
}

func (n *TreeNode) PreOrderTransversal() {
	fmt.Printf("%d ", n.Data)

	if n.LeftChild != nil {
		n.LeftChild.PreOrderTransversal()
	}

	if n.RightChild != nil {
		n.RightChild.PreOrderTransversal()
	}
}

func (n *TreeNode) PostOrderTransversal() {
	if n.LeftChild != nil {
		n.LeftChild.PostOrderTransversal()
	}

	if n.RightChild != nil {
		n.RightChild.PostOrderTransversal()
	}
	fmt.Printf("%d ", n.Data)
}

func (t *TreeNode) Delete(key int) *TreeNode {
	if t == nil {
		return nil
	}

	if t.Data > key {
		t.LeftChild = t.LeftChild.Delete(key)
	} else if t.Data < key {
		t.RightChild = t.RightChild.Delete(key)
	} else {
		if t.RightChild == nil {
			return t.LeftChild
		} else if t.LeftChild == nil {
			return t.RightChild
		} else {
			t.Data = t.RightChild.Min()
			t.RightChild = t.RightChild.Delete(t.Data)
		}
	}
	return t
}

func (t *TreeNode) Search(key int) bool {

	if t.Data == key {
		return true
	} else if key < t.Data {
		return t.LeftChild.Search(key)
	} else {
		return t.RightChild.Search(key)
	}
}

func (n *TreeNode) Min() int {
	if n.LeftChild != nil {
		return n.LeftChild.Min()
	}

	return n.Data
}

func (n *TreeNode) Max() int {
	if n.RightChild != nil {
		return n.RightChild.Max()
	}

	return n.Data
}

func (n *TreeNode) binaryTreePaths() []string {
	result, arr := []string{}, []string{}
	transversePath(n, &arr, &result)
	return result
}

func transversePath(root *TreeNode, arr, result *[]string) {
	*arr = append(*arr, fmt.Sprintf("%d", root.Data))
	if root.LeftChild == nil && root.RightChild == nil {
		*result = append(*result, strings.Join(*arr, "->"))
		(*arr) = (*arr)[:len(*arr)-1]
		return
	}
	if root.LeftChild != nil {
		transversePath(root.LeftChild, arr, result)
	}

	if root.RightChild != nil {
		transversePath(root.RightChild, arr, result)
	}
	(*arr) = (*arr)[:len(*arr)-1]
}
