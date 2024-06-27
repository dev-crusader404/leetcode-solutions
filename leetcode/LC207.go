package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	list := buildAdjacency(prerequisites, numCourses)
	m := make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		soFar := make(map[int]struct{})
		soFar[i] = struct{}{}
		isCycle := cycleInCourse(list, i, m, soFar)
		if isCycle {
			return false
		}
	}

	return true
}

func cycleInCourse(list [][]int, curr int, seen, soFar map[int]struct{}) bool {
	if _, ok := soFar[curr]; ok && len(soFar) > 1 {
		return true
	} else {
		soFar[curr] = struct{}{}
	}
	if _, ok := seen[curr]; ok {
		return false
	}
	seen[curr] = struct{}{}
	for _, k := range list[curr] {
		isCycle := cycleInCourse(list, k, seen, soFar)
		if isCycle {
			return true
		}
	}
	return false
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
	a := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	canFinish(19, a)
}
