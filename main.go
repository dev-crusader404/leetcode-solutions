package main

import (
	lc "github.com/dev-crusader/leetcode-solutions/leetcode"
)

func main() {
	// fmt.Println(lc.ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

	// fmt.Println(lc.AddDigits(54613475))
	// fmt.Println(lc.AddDigits2(54613475))
	// fmt.Println(lc.IsUgly2(30))

	// lc.RunLC67()
	lc.RunLC2()
	lc.RunLC2487()
	lc.RunLC143()
	lc.RunRotatedSearch()
	lc.MaxProduct([]int{4, 5, 0, -7, -5, 2})
	lc.FindMinInRotatedArray()
	lc.CheckSubarraySum([]int{23, 2, 3, 4, 7, 3}, 6)
	lc.SubarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)
	// lc.CharacterReplacement("AABABBA", 1)
}
