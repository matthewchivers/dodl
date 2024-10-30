package templating

import (
	"time"

	"text/template"
)

func addCustomFuncs(funcMap template.FuncMap) template.FuncMap {
	funcMap["addDays"] = addDays
	funcMap["addMonths"] = addMonths
	funcMap["addYears"] = addYears
	funcMap["WeekStart"] = WeekStart
	funcMap["daysInYear"] = daysInYear
	funcMap["daysInMonth"] = daysInMonth
	return funcMap
}

// addDays adds a number of days to a time.Time object.
func addDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// addMonths adds a number of months to a time.Time object.
func addMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// addYears adds a number of years to a time.Time object.
func addYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// WeekStart returns the date of the Monday of the week containing the given date.
func WeekStart(t time.Time) time.Time {
	offset := (int(t.Weekday()) + 6) % 7 // Adjust so that Monday is 0
	return t.AddDate(0, 0, -offset)
}

// daysInYear returns the number of days in the year of the given date.
func daysInYear(t time.Time) int {
	return time.Date(t.Year()+1, 1, 0, 0, 0, 0, 0, t.Location()).YearDay()
}

// daysInMonth returns the number of days in the month of the given date.
func daysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}
