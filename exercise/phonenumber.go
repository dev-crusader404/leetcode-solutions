package exercise

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	if len(phoneNumber) == 0 {
		return "", errors.New("empty input")
	}
	number, err := Normalize(phoneNumber)
	if err != nil {
		return "", err
	}
	return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
	if len(phoneNumber) == 0 {
		return "", errors.New("empty input")
	}
	number, err := Normalize(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	if len(phoneNumber) == 0 {
		return "", errors.New("empty input")
	}
	n, err := Normalize(phoneNumber)
	if err != nil {
		return "", err
	}
	formattedNum := fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:])
	return formattedNum, nil
}

func Normalize(num string) (string, error) {
	num = strings.TrimSpace(num)
	var sb strings.Builder
	for _, c := range num {
		if c >= '0' && c <= '9' {
			sb.WriteRune(c)
		} else if unicode.IsLetter(c) {
			return "", errors.New("error invalid character")
		}
	}
	newNum := sb.String()
	if len(newNum) == 11 {
		if newNum[0] != '1' {
			return "", errors.New("error invalid start")
		}
		newNum = newNum[1:]
	} else if len(newNum) > 11 || len(newNum) < 10 {
		return "", errors.New("error invalid length")
	}
	if newNum[0] < '2' || newNum[3] < '2' {
		return "", errors.New("error invalid character")
	}
	return newNum, nil
}
