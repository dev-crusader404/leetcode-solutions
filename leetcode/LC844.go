package leetcode

import (
	"fmt"
	"strings"
)

func backspaceCompare(s string, t string) bool {
	a := process(s)
	b := process(t)
	return a == b
}

func process(s string) string {
	var t []string
	for _, v := range s {
		if v != '#' {
			t = append(t, fmt.Sprintf("%s", v))

		} else if len(t) != 0 {
			t = t[:len(t)-1]
		}
	}
	return strings.Join(t, "")
}

func RunLC844() {
	fmt.Println(backspaceCompare("a#c", "b"))
}
