package exercise

import (
	"fmt"
	"math"
)

var colorMap = map[string]int{
	"black": 0, "brown": 1, "red": 2,
	"orange": 3, "yellow": 4, "green": 5, "blue": 6,
	"violet": 7, "grey": 8, "white": 9,
}

func Label(colors []string) string {
	if len(colors) == 0 {
		return ""
	}
	var total int
	var result string
	for i := 0; i < len(colors)-1; i++ {
		total = total * 10
		if i < 2 {
			total += colorMap[colors[i]]
		}
	}
	suffix := `ohms`
	cal := colorMap[colors[len(colors)-1]]
	if total != 0 {
		calculatePower(&total, &cal)
	}
	if cal < 3 {
		m := int(math.Pow10(cal))
		total *= m
		result = fmt.Sprintf("%d %s", total, suffix)
	} else if cal < 6 {
		m := cal - 3
		total *= int(math.Pow10(m))
		result = fmt.Sprintf("%d kilo%s", total, suffix)
	} else if cal < 9 {
		m := cal - 6
		total *= int(math.Pow10(m))
		result = fmt.Sprintf("%d mega%s", total, suffix)
	} else if cal == 9 {
		result = fmt.Sprintf("%d giga%s", total, suffix)
	}
	return result
}

func calculatePower(total, power *int) {
	for (*total)%10 == 0 {
		(*power)++
		*total /= 10
	}
}
