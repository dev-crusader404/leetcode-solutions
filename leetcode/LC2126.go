package leetcode

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)

	for _, astr := range asteroids {
		if mass >= astr {
			mass += astr
		} else {
			return false
		}
	}
	return true
}
