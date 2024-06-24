package leetcode

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	var stack []rune
	for _, v := range s {
		if v == '[' || v == '{' || v == '(' {
			stack = append(stack, getClosure(v))
		} else {
			n := len(stack) - 1
			if n < 0 {
				return false
			}
			pop := stack[n]
			stack = stack[:n]
			if pop != v {
				return false
			}
		}
	}
	return len(stack) == 0
}

func getClosure(a rune) rune {
	var t rune
	switch a {
	case '[':
		t = ']'
	case '(':
		t = ')'
	case '{':
		t = '}'
	default:
		t = a
	}
	return t
}
