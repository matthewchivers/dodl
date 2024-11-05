package dateutils

import "time"

// AddDays adds a number of days to a time.Time object.
// If the number of days is negative, the days are subtracted.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths adds a number of months to a time.Time object.
// If the number of months is negative, the months are subtracted.
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears adds a number of years to a time.Time object.
// If the number of years is negative, the years are subtracted.
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// GetWeekStartDate returns the date of the start of the week containing the given time,
// based on the startDay provided. The startDay should be one of the time.Weekday constants.
func GetWeekStartDate(t time.Time, startDay time.Weekday) time.Time {
	// Calculate the offset based on the start day
	offset := (int(t.Weekday()) - int(startDay) + 7) % 7
	return t.AddDate(0, 0, -offset)
}

// GetDefaultWeekStartDate returns the date of the start of the week containing the given time,
// based on the default start day (Monday).
func GetDefaultWeekStartDate(t time.Time) time.Time {
	return GetWeekStartDate(t, time.Monday)
}

// DaysInYear returns the number of days in the year of the given date.
// e.g. 365 for a non-leap year, 366 for a leap year.
func DaysInYear(t time.Time) int {
	return time.Date(t.Year()+1, 1, 0, 0, 0, 0, 0, t.Location()).YearDay()
}

// DaysInMonth returns the number of days in the month of the given date.
// e.g. 31 for January, 28 for February (non-leap year), 29 for February (leap year).
func DaysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}

// WeekNumber returns the week number of the given date.
// The week number is defined as the ISO 8601 week number.
func WeekNumber(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}
