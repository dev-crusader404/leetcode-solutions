package leetcode

func asteroidCollision(asteroids []int) []int {
	res := []int{}

	for _, astr := range asteroids {
		if astr > 0 {
			res = append(res, astr)
		} else {
			for len(res) > 0 && -astr > res[len(res)-1] && res[len(res)-1] > 0 {
				res = res[:len(res)-1]
			}
			if len(res) == 0 || res[len(res)-1] < 0 {
				res = append(res, astr)
			} else if res[len(res)-1] == -astr {
				res = res[:len(res)-1]
			}
		}
	}
	return res
}
