package leetcode

func kClosest2(points [][]int, k int) [][]int {
	if len(points) == 0 || k == 0 {
		return [][]int{}
	}
	low, high := 0, len(points)-1

	for low <= high {
		partitionIdx := partitionData(points, low, high)
		if partitionIdx == k {
			break
		} else if partitionIdx < k {
			low = partitionIdx + 1
		} else {
			high = partitionIdx - 1
		}
	}
	return points[:k]
}

func partitionData(n [][]int, low, high int) int {
	pivot := n[high]
	left := low
	for right := low; right < high; right++ {
		if squareDistance(n[right]) < squareDistance(pivot) {
			n[left], n[right] = n[right], n[left]
			left++
		}
	}
	n[left], n[high] = n[high], n[left]
	return left
}

func squareDistance(n []int) int {
	return n[0]*n[0] + n[1]*n[1]
}
