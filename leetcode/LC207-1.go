package leetcode

func canFinish2(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	inDegree, list := calculateIndegree(prerequisites, numCourses)
	return canCourseFinish(list, inDegree)
}

func canCourseFinish(list [][]int, inDegree []int) bool {
	stack := findZeroInDegree(inDegree)
	var count int
	for len(stack) > 0 {
		n := len(stack) - 1
		item := stack[n]
		stack = stack[:n]
		count++
		for i := 0; i < len(list[item]); i++ {
			val := list[item][i]
			inDegree[val]--
			if inDegree[val] == 0 {
				stack = append(stack, val)
			}
		}
	}
	if count == len(inDegree) {
		return true
	} else {
		return false
	}
}

func calculateIndegree(prereq [][]int, n int) ([]int, [][]int) {
	inDegree := make([]int, n)
	list := make([][]int, n)
	for _, v := range prereq {
		inDegree[v[0]]++
		list[v[1]] = append(list[v[1]], v[0])
	}
	return inDegree, list
}

func findZeroInDegree(l []int) []int {
	stack := []int{}
	for i, v := range l {
		if v == 0 {
			stack = append(stack, i)
		}
	}
	return stack
}
