package exercise

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	start := `day of Christmas my true love gave to me: `
	gifts := []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
	day := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	verse := fmt.Sprintf("On the %s %s", day[i-1], start)
	for i > 0 {
		verse += gifts[i-1]
		if i > 2 {
			verse += ", "
		} else if i == 2 {
			verse += ", and "
		}
		i--
	}
	verse += "."
	return verse
}

func Song() string {
	var songArr []string
	for i := 1; i <= 12; i++ {
		songArr = append(songArr, Verse(i))
	}
	fullSong := strings.Join(songArr, "\n")
	return fullSong
}
