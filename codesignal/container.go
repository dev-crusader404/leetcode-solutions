package codesignal

import (
	"fmt"
	"sort"
	"strconv"
)

type Container struct {
	data []int
	size int
}

func NewContainer() *Container {
	return &Container{
		data: []int{},
		size: 0,
	}
}

func (c *Container) Add(x int) string {
	c.data = append(c.data, x)
	c.size++
	return fmt.Sprint(c.size)
}

func (c *Container) Delete(x int) string {
	for i, v := range c.data {
		if v == x {
			c.data = append(c.data[:i], c.data[i+1:]...)
			c.size--
			return "true"
		}
	}
	return "false"
}

func (c *Container) GetMedian() string {
	if c.size == 0 {
		return ""
	} else if c.size == 1 {
		return fmt.Sprint(c.data[0])
	}
	dup := make([]int, c.size)
	copy(dup, c.data)
	sort.Ints(dup)
	mid := (c.size - 1) / 2
	return fmt.Sprintf("%d", dup[mid])
}

func solutionContainer(queries [][]string) []string {
	cntr := NewContainer()
	var result []string
	var value int64
	for _, v := range queries {
		switch v[0] {
		case "ADD":
			value, _ = strconv.ParseInt(v[1], 10, 32)
			result = append(result, cntr.Add(int(value)))
		case "DELETE":
			value, _ = strconv.ParseInt(v[1], 10, 32)
			result = append(result, cntr.Delete(int(value)))
		case "GET_MEDIAN":
			result = append(result, cntr.GetMedian())
		}
	}
	return result
}

func RunContainer() {
	input := [][]string{
		{"GET_MEDIAN"},
		{"ADD", "5"},
		{"ADD", "10"},
		{"ADD", "1"},
		{"GET_MEDIAN"},
		{"ADD", "4"},
		{"GET_MEDIAN"},
		{"DELETE", "1"},
		{"GET_MEDIAN"}}
	res := solutionContainer(input)
	fmt.Println(res)
}
