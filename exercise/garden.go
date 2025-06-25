package exercise

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	children   []string
	presentMap map[string]struct{}
	diagram    []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) == 0 || len(children) == 0 {
		return nil, errors.New("invalid input")
	}
	g := &Garden{presentMap: map[string]struct{}{}}
	rows, err := ParseInput(diagram, len(children))
	if err != nil {
		return nil, err
	}
	g.diagram = rows
	for _, name := range children {
		if _, ok := g.presentMap[name]; ok {
			return &Garden{}, errors.New("duplicate name")
		}
		g.presentMap[name] = struct{}{}
		g.children = append(g.children, name)
	}
	if len(g.children) > 1 {
		sort.Slice(g.children, func(i int, j int) bool { return g.children[i] < g.children[j] })
	}
	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	var plantsArray []string
	if _, ok := g.presentMap[child]; !ok {
		return plantsArray, false
	}
	plantEncode := map[byte]string{'G': "grass", 'C': "clover", 'R': "radishes", 'V': "violets"}
	index := 0
	for i := 0; i < len(g.children); i++ {
		if g.children[i] == child {
			index = i
			break
		}
	}
	for _, code := range g.diagram {
		seed1 := code[2*index]
		seed2 := code[2*index+1]
		plantsArray = append(plantsArray, plantEncode[seed1], plantEncode[seed2])
	}
	return plantsArray, true
}

func ParseInput(s string, size int) ([]string, error) {
	var result []string
	arr := strings.Split(s, "\n")
	if arr[0] != "" || arr[len(arr)-1] == "" {
		return []string{}, errors.New("invalid diagram format")
	}
	arr = arr[1:]
	for i, row := range arr {
		if row == "" || len(row) != (size*2) || len(row)%2 != 0 {
			return []string{}, errors.New("invalid length")
		}
		if row != strings.ToUpper(row) {
			return []string{}, errors.New("invalid code")
		}
		if i+1 < len(arr) && len(row) != len(arr[i+1]) {
			return []string{}, errors.New("invalid code length")
		}
		result = append(result, row)
	}
	return result, nil
}
