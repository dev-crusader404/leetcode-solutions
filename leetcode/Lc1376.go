package leetcode

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	if n < 2 {
		return 0
	}
	adList := buildAdjList(manager)
	fmt.Println(adList)
	return findInformTime(headID, 0, adList, informTime)
}

func findInformTime(start, maxTime int, list [][]int, infoTime []int) int {
	if len(list[start]) == 0 {
		return maxTime
	}
	total := maxTime + infoTime[start]
	for _, v := range list[start] {
		maxTime = max(maxTime, findInformTime(v, total, list, infoTime))
	}
	return maxTime
}

func buildAdjList(m []int) [][]int {
	n := len(m)
	list := make([][]int, n)

	for i, v := range m {
		if v != -1 {
			list[v] = append(list[v], i)
		}
	}

	return list
}

func RunLC1376() {
	m := []int{2, 2, 4, 6, -1, 4, 4, 5}
	inf := []int{0, 0, 4, 0, 7, 3, 6, 0}
	id := 4
	time := numOfMinutes(len(m), id, m, inf)
	fmt.Println(time)
}
