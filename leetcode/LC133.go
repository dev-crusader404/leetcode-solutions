package leetcode

// Definition for a Node133.
type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return node
	}

	cloneResult := make(map[*Node133]*Node133)
	copyGraph(node, cloneResult)
	return cloneResult[node]
}

func copyGraph(node *Node133, result map[*Node133]*Node133) {
	if node == nil {
		return
	}
	newNode := new(Node133)
	newNode.Val = node.Val
	result[node] = newNode
	for _, neigh := range node.Neighbors {
		if _, ok := result[neigh]; !ok {
			copyGraph(neigh, result)
		}
		newNode.Neighbors = append(newNode.Neighbors, result[neigh])
	}
}
