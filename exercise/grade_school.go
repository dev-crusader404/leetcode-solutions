package exercise

import (
	"fmt"
	"sort"
)

// Define the Grade and School types here.
type School struct {
	Class        []Grade
	EnrolledList map[string]struct{}
}

type Grade struct {
	Level    int
	Students []string
}

func NewSchool() *School {
	return &School{Class: make([]Grade, 9), EnrolledList: map[string]struct{}{}}
}

func (s *School) Add(student string, grade int) {
	if _, ok := s.EnrolledList[student]; ok {
		fmt.Println("Already enrolled: " + student)
		return
	}
	s.Class[grade-1].Level = grade
	if s.Class[grade-1].Students == nil {
		s.Class[grade-1].Students = []string{}
	}
	s.Class[grade-1].Students = append(s.Class[grade-1].Students, student)
	s.EnrolledList[student] = struct{}{}
	if len(s.Class[grade-1].Students) > 1 {
		sort.Strings(s.Class[grade-1].Students)
	}
}

func (s *School) Grade(level int) []string {
	var list []string
	l := s.Class[level-1]
	list = append(list, l.Students...)
	return list
}

func (s *School) Enrollment() []Grade {
	grade := []Grade{}
	for _, level := range s.Class {
		if len(level.Students) > 0 {
			g := Grade{Level: level.Level, Students: level.Students}
			grade = append(grade, g)
		}
	}
	return grade
}
