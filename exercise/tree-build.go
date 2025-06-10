package exercise

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodeMap := make(map[int]*Node)
	for _, rec := range records {
		if _, ok := nodeMap[rec.ID]; ok {
			return &Node{}, errors.New("duplicate node")
		}
		nodeMap[rec.ID] = &Node{ID: rec.ID}
	}

	var root *Node
	for _, rec := range records {
		if rec.ID >= len(records) || rec.Parent > rec.ID {
			return &Node{}, errors.New("invalid input")
		}
		if rec.ID == rec.Parent {
			if rec.ID != 0 || root != nil {
				return &Node{}, errors.New("multiple root node")
			}
			root = nodeMap[rec.ID]
		} else {
			curr, ok := nodeMap[rec.Parent]
			if !ok {
				return &Node{}, errors.New("parent does not exist")
			}
			curr.Children = append(curr.Children, nodeMap[rec.ID])
			sort.Slice(curr.Children, func(i, j int) bool {
				return curr.Children[i].ID < curr.Children[j].ID
			})
		}
	}
	if root == nil {
		return &Node{}, errors.New("no root node")
	}
	return root, nil
}
