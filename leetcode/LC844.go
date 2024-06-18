package leetcode

import (
	"fmt"
	"strings"
)

func backspaceCompares(s string, t string) bool {
	m := len(s) - 1
	n := len(t) - 1
	for {
		if m >= 0 && s[m] == '#' {
			moveBack(s, &m)
		}
		if n >= 0 && t[n] == '#' {
			moveBack(t, &n)
		}
		if m >= 0 && n >= 0 && s[m] == t[n] {
			m--
			n--
		} else {
			break
		}
	}
	return (m == -1 && n == -1)
}

func moveBack(w string, pos *int) {
	delete := 2
	for delete > 0 && (*pos) >= 0 {
		(*pos)--
		delete--
		if (*pos) >= 0 && w[(*pos)] == '#' {
			delete += 2
		}
	}
}

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
	fmt.Println(backspaceCompares("nzp#o#g", "b#nzp#o#g"))
}
