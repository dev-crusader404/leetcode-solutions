package exercise

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	key int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return shift{key: 3}
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift{key: distance}
}

func (c shift) Encode(input string) string {
	if c.key == 0 {
		return input
	}
	input = strings.ToLower(input)
	return Shifter(input, c.key)
}

func (c shift) Decode(input string) string {
	if c.key == 0 {
		return input
	}
	input = strings.ToLower(input)
	return Shifter(input, -c.key)
}

func Shifter(input string, key int) string {
	var result []rune
	key = (key%26 + 26) % 26
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		shiftedRune := 'a' + (r+rune(key)-'a')%26
		result = append(result, shiftedRune)
	}
	return string(result)
}

func NewVigenere(key string) Cipher {
	if !validate(key) {
		return nil
	}
	return vigenere{key: key}
}

func normalize(input, key string) string {
	var newKey string = key
	for len(input) >= len(newKey) {
		newKey += key
	}
	return newKey
}

func (v vigenere) Encode(input string) string {
	var result []rune
	input = strings.ToLower(input)
	v.key = normalize(input, v.key)
	var idx int
	for i, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}
		shiftedRune := 'a' + (input[i]-'a'+v.key[idx]-'a')%26
		result = append(result, rune(shiftedRune))
		idx++
	}
	return string(result)
}

func (v vigenere) Decode(input string) string {
	var result []rune
	input = strings.ToLower(input)
	v.key = normalize(input, v.key)
	var idx int
	for i, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}
		effective := ('a' - v.key[i] + 26) % 26
		shiftedRune := 'a' + (input[i]+effective-'a')%26
		result = append(result, rune(shiftedRune))
		idx++
	}
	return string(result)
}

func validate(s string) bool {
	var aCount int
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		} else if c == 'a' {
			aCount++
		}
	}
	return !(aCount == len(s))
}
