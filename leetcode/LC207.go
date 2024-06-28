package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	list := buildAdjacency(prerequisites, numCourses)

	for i := 0; i < numCourses; i++ {
		queue := []int{}
		m := make(map[int]struct{})
		for j := 0; j < len(list[i]); j++ {
			adjList := list[i][j]
			queue = append(queue, adjList)
		}

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]
			m[item] = struct{}{}
			if item == i {
				return false
			}
			for k := 0; k < len(list[item]); k++ {
				val := list[item][k]
				if _, ok := m[val]; !ok {
					queue = append(queue, list[item][k])
				}
			}
		}
	}

	return true
}

func buildAdjacency(arr [][]int, n int) [][]int {
	list := make([][]int, n)
	for _, v := range arr {
		list[v[1]] = append(list[v[1]], v[0])
	}
	return list
}

func RunLC207() {
	// [[0,10],[3,18],[5,5],[6,11],[11,14],[13,1],[15,1],[17,4]]
	// a := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	// a := [][]int{{1, 0}}
	// a := [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}
	a := [][]int{{1, 0}, {2, 1}, {0, 2}, {2, 3}}
	canFinish(4, a)
	canFinish2(4, a)
}
