package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement3(n int) int {
	if n < 10 {
		return -1
	}
	s := strconv.Itoa(n)
	arr := []byte(s)

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			index := getMinIndex(arr[i:], arr[i-1])
			arr[index+i], arr[i-1] = arr[i-1], arr[index+i]
			toSort := arr[i:]
			sort.Slice(toSort, func(i, j int) bool {
				return toSort[i] < toSort[j]
			})
			break
		}
	}
	if string(arr) == s {
		return -1
	}
	res := string(arr)
	o, _ := strconv.Atoi(res)
	if o > math.MaxInt32 {
		return -1
	}
	return o
}

func getMinIndex(arr []byte, val byte) int {
	minIndex := math.MaxInt
	minimum := byte(minIndex)
	for i, v := range arr {
		if v > val && v < minimum {
			minimum = v
			minIndex = i
		}
	}
	return minIndex
}

func RunLC556() {
	fmt.Println(nextGreaterElement3(2147483486))
}
