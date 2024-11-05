package templating

import (
	"text/template"

	"github.com/matthewchivers/dodl/pkg/dateutils"
)

func addCustomFuncs(funcMap template.FuncMap) template.FuncMap {
	funcMap["addDays"] = dateutils.AddDays
	funcMap["addMonths"] = dateutils.AddMonths
	funcMap["addYears"] = dateutils.AddYears
	funcMap["calcWeekStart"] = dateutils.GetDefaultWeekStartDate
	funcMap["daysInYear"] = dateutils.DaysInYear
	funcMap["daysInMonth"] = dateutils.DaysInMonth
	funcMap["weekNumber"] = dateutils.WeekNumber
	return funcMap
}
