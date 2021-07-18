package util

import "time"

var Test = false

func IsToday(year int, month int, day int) bool {
	var Today time.Time
	if Test {
		return true
	}
	Today = time.Now()

	var TodayYear, TodayMonth, TodayDay = Today.Date()

	if TodayYear != year ||
		TodayMonth != time.Month(month) ||
		TodayDay != day {
		return false
	}
	return true
}
