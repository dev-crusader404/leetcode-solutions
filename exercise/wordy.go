package exercise

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	quesArr, isErr := cleanup(question)
	if !isErr {
		return 0, isErr
	}
	return evaluate(quesArr)
}

func cleanup(question string) ([]string, bool) {
	if len(question) == 0 || !strings.HasPrefix(question, "What is ") || !strings.HasSuffix(question, "?") {
		return []string{}, false
	}
	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSuffix(question, "?")
	quesArr := strings.Fields(question)
	if len(quesArr) == 0 {
		return []string{}, false
	}
	return quesArr, true
}

func evaluate(arr []string) (int, bool) {
	num1, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, false
	}
	for i := 1; i < len(arr); i += 2 {
		oprn := arr[i]
		if oprn == "multiplied" || oprn == "divided" {
			i++
		}
		if i+1 >= len(arr) {
			return 0, false
		}
		nextNum, err := strconv.Atoi(arr[i+1])
		if err != nil {
			return 0, false
		}
		switch oprn {
		case "plus":
			num1 += nextNum
		case "minus":
			num1 -= nextNum
		case "multiplied":
			num1 *= nextNum
		case "divided":
			num1 /= nextNum
		default:
			return 0, false
		}
	}
	return num1, true
}
