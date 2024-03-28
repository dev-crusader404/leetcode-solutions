package datastructure

import "fmt"

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Data       int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &TreeNode{Data: key}
	} else {
		t.Root.Insert(key)
	}
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

func (t *Tree) Search(key int) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.Search(key)
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

func (t *Tree) InOrderTransversal() {
	if t.Root != nil {
		t.Root.InOrderTransversal()
	}
}

func (n *TreeNode) InOrderTransversal() {
	if n.LeftChild != nil {
		n.LeftChild.InOrderTransversal()
	}
	fmt.Println(n.Data)

	if n.RightChild != nil {
		n.RightChild.InOrderTransversal()
	}
}

func (t *Tree) PreOrderTransversal() {
	if t.Root != nil {
		t.Root.PreOrderTransversal()
	}
}

func (n *TreeNode) PreOrderTransversal() {
	fmt.Println(n.Data)

	if n.LeftChild != nil {
		n.LeftChild.PreOrderTransversal()
	}

	if n.RightChild != nil {
		n.RightChild.PreOrderTransversal()
	}
}

func (t *Tree) PostOrderTransversal() {
	if t.Root != nil {
		t.Root.PostOrderTransversal()
	}
}

func (n *TreeNode) PostOrderTransversal() {
	if n.LeftChild != nil {
		n.LeftChild.PostOrderTransversal()
	}

	if n.RightChild != nil {
		n.RightChild.PostOrderTransversal()
	}
	fmt.Println(n.Data)
}

func BST() {
	t := &Tree{}
	t.Insert(35)
	t.Insert(165)
	t.Insert(47)
	t.Insert(243)
	t.Insert(65)
	t.Insert(146)
	t.Insert(10)
	t.Insert(6)
	t.Insert(15)

	fmt.Println(t.Search(10))
	fmt.Println("\nInOrder Transversal: ")
	t.InOrderTransversal()
	fmt.Println("\nPreOrder Transversal: ")
	t.PreOrderTransversal()
	fmt.Println("\nPostOrder Transversal: ")
	t.PostOrderTransversal()
}
