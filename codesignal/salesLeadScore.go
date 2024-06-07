package codesignal

import (
	"fmt"
	"sort"
)

type Info struct {
	Name     string
	Time     int
	NetValue int
	Score    float32
}

func salesLeadScore(names []string, time []int, netValue []int, isOnVacation []bool) []string {
	var input []Info
	for i := range isOnVacation {
		if !isOnVacation[i] {
			data := NewInfo(names[i], time[i], netValue[i])
			data.Score = calculateScore(netValue[i], time[i])
			input = append(input, data)
		}
	}

	sort.Slice(input, func(i, j int) bool {
		if input[i].Score == input[j].Score {
			if input[i].Time == input[j].Time {
				return input[i].Name < input[j].Name
			} else {
				return input[i].Time > input[j].Time
			}
		} else {
			return input[i].Score > input[j].Score
		}
	})
	var result []string
	for _, v := range input {
		result = append(result, v.Name)
	}
	return result
}

func NewInfo(name string, time, netValue int) Info {
	return Info{
		Name:     name,
		Time:     time,
		NetValue: netValue,
	}
}

func calculateScore(netValue, time int) float32 {
	var val float32
	val = float32(netValue*time) / float32(365)
	return val
}

func RunSalesLeadScore() {
	names := []string{"abcd", "dba", "abc", "abcdd", "abcc", "abd", "abcde", "abcda", "abca", "abcbc", "bba"}
	time := []int{20, 20, 20, 20, 10, 20, 10, 10, 20, 20, 10}
	netValue := []int{1000, 1200, 1000, 1000, 1000, 1200, 1200, 1200, 1000, 1000, 1100}
	isOnVacation := []bool{false, false, false, false, false, true, false, false, true, false, false}
	r := salesLeadScore(names, time, netValue, isOnVacation)
	fmt.Println(r)
}

/*
["dba",
 "abc",
 "abcbc",
 "abcd",
 "abcdd",
 "abcda",
 "abcde",
 "bba",
 "abcc"]

 ["dba",
 "abcd",
 "abc",
 "abcdd",
 "abcbc",
 "abcde",
 "abcda",
 "bba",
 "abcc"]*/
