package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree, list := calIndegree(prerequisites, numCourses)
	return getCompletionPath(indegree, list)
}

func getCompletionPath(indegree []int, list [][]int) []int {
	var result []int
	queue := getZeroIndegree(indegree)

	for len(queue) > 0 {
		n := len(queue) - 1
		curr := queue[n]
		queue = queue[:n]
		result = append(result, curr)
		for i := 0; i < len(list[curr]); i++ {
			element := list[curr][i]
			indegree[element]--
			if indegree[element] == 0 {
				queue = append(queue, element)
			}
		}
	}
	if len(indegree) != len(result) {
		return []int{}
	}
	return result
}

func getZeroIndegree(arr []int) []int {
	stack := []int{}
	for i, v := range arr {
		if v == 0 {
			stack = append(stack, i)
		}
	}
	return stack
}

func calIndegree(prereq [][]int, n int) ([]int, [][]int) {
	inDegree := make([]int, n)
	list := make([][]int, n)
	for _, v := range prereq {
		inDegree[v[0]]++
		list[v[1]] = append(list[v[1]], v[0])
	}
	return inDegree, list
}
