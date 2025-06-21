package exercise

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid input")
	}
	width := int(char - 'A')
	if width == 0 {
		return string(char), nil
	}
	result := make([]string, 2*width+1)
	letterA, idx := 'A', 0
	for width >= 0 {
		var sb strings.Builder
		sb.WriteString(strings.Repeat(" ", width))
		start := letterA + rune(idx)
		sb.WriteRune(start)
		if idx > 0 {
			gap := 2*idx - 1
			sb.WriteString(strings.Repeat(" ", gap))
			sb.WriteRune(start)
		}
		sb.WriteString(strings.Repeat(" ", width))
		result[idx] = sb.String()
		result[len(result)-idx-1] = sb.String()
		width--
		idx++
	}

	return strings.Join(result, "\n"), nil
}
