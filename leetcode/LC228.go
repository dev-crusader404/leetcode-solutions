package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
    var result []string
    var temp string
    if len(nums) == 0 { return result }
    start := nums[0]
    for i, n := range nums {
        if i+1 < len(nums) {
            if n+1 == nums[i+1] {
                continue
            }
        }
        if start == n {
            temp = fmt.Sprintf("%d", n)
        } else {
            temp = fmt.Sprintf("%d->%d", start, n)
        }
        if i+1 < len(nums) {
            start = nums[i+1]
        }
        result = append(result, temp)
    }
    return result
}