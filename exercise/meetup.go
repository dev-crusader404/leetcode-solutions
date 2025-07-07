package exercise

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	daysInMonth := getNumberOfDays(year, month)
	switch wSched {
	case First:
		return findNthWeekDayInMonth(firstDayOfMonth, wDay, 1)
	case Second:
		return findNthWeekDayInMonth(firstDayOfMonth, wDay, 2)
	case Third:
		return findNthWeekDayInMonth(firstDayOfMonth, wDay, 3)
	case Fourth:
		return findNthWeekDayInMonth(firstDayOfMonth, wDay, 4)
	case Last:
		return findLastWeekDayInMonth(firstDayOfMonth, wDay, daysInMonth)
	case Teenth:
		return findTeenthDay(firstDayOfMonth, wDay)
	default:
	}
	return 0
}

func getNumberOfDays(year int, month time.Month) int {
	t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}

func findNthWeekDayInMonth(day time.Time, wDay time.Weekday, week int) int {
	firstWeekDay := day.Weekday()
	d := (int(wDay) - int(firstWeekDay) + 7) % 7
	fisrtOccurence := d + 1
	nthOccurence := fisrtOccurence + (week-1)*7
	return nthOccurence
}

func findLastWeekDayInMonth(firstDay time.Time, day time.Weekday, daysInMonth int) int {
	lastDate := time.Date(firstDay.Year(), firstDay.Month(), daysInMonth, 0, 0, 0, 0, time.UTC)
	lastWeekDay := lastDate.Weekday()
	date := (int(lastWeekDay) - int(day) + 7) % 7
	return daysInMonth - date
}

func findTeenthDay(firstDayOfMonth time.Time, wDay time.Weekday) int {
	for i := 13; i < 20; i++ {
		current := time.Date(firstDayOfMonth.Year(), firstDayOfMonth.Month(), i, 0, 0, 0, 0, time.UTC)
		if current.Weekday() == wDay {
			return i
		}
	}
	return 0
}
