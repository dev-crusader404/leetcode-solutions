package leetcode

import "fmt"

func fizzBuzz(n int) []string {
	l := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			l = append(l, "FizzBuzz")
		} else if i%3 == 0 {
			l = append(l, "Fizz")
		} else if i%5 == 0 {
			l = append(l, "Buzz")
		} else {
			l = append(l, fmt.Sprintf("%v", i))
		}
	}
	return l
}
