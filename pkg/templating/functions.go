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
	funcMap["weekNumber"] = weekNumber
	return funcMap
}

// addDays adds a number of days to a time.Time object.
// If the number of days is negative, the days are subtracted.
func addDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// addMonths adds a number of months to a time.Time object.
// If the number of months is negative, the months are subtracted.
func addMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// addYears adds a number of years to a time.Time object.
// If the number of years is negative, the years are subtracted.
func addYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// WeekStart returns the date of the Monday of the week containing the given t (time.Time).
func WeekStart(t time.Time) time.Time {
	offset := (int(t.Weekday()) + 6) % 7 // Adjust so that Monday is 0
	return t.AddDate(0, 0, -offset)
}

// daysInYear returns the number of days in the year of the given date.
// e.g. 365 for a non-leap year, 366 for a leap year.
func daysInYear(t time.Time) int {
	return time.Date(t.Year()+1, 1, 0, 0, 0, 0, 0, t.Location()).YearDay()
}

// daysInMonth returns the number of days in the month of the given date.
// e.g. 31 for January, 28 for February (non-leap year), 29 for February (leap year).
func daysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}

// weekNumber returns the week number of the given date.
// The week number is defined as the ISO 8601 week number.
func weekNumber(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}
