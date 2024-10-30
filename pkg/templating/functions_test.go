package templating

import (
	"testing"
	"time"
)

// TestAddDays tests the addDays function.
func TestAddDays(t *testing.T) {
	testCases := []struct {
		name     string
		start    time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "Add 1 day",
			start:    time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			days:     1,
			expected: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Subtract 1 day",
			start:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			days:     -1,
			expected: time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Add 30 days",
			start:    time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			days:     30,
			expected: time.Date(2023, time.January, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Leap year February end",
			start:    time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC),
			days:     1,
			expected: time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := addDays(tc.start, tc.days)
			if !result.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestAddMonths tests the addMonths function.
func TestAddMonths(t *testing.T) {
	testCases := []struct {
		name     string
		start    time.Time
		months   int
		expected time.Time
	}{
		{
			name:     "Add 1 month",
			start:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			months:   1,
			expected: time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Subtract 1 month",
			start:    time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			months:   -1,
			expected: time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "End of month handling",
			start:    time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC),
			months:   1,
			expected: time.Date(2024, time.March, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Cross year boundary",
			start:    time.Date(2023, time.November, 30, 0, 0, 0, 0, time.UTC),
			months:   3,
			expected: time.Date(2024, time.March, 01, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := addMonths(tc.start, tc.months)
			if !result.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestAddYears tests the addYears function.
func TestAddYears(t *testing.T) {
	testCases := []struct {
		name     string
		start    time.Time
		years    int
		expected time.Time
	}{
		{
			name:     "Add 1 year",
			start:    time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC),
			years:    1,
			expected: time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Subtract 1 year",
			start:    time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			years:    -1,
			expected: time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Leap year handling",
			start:    time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			years:    1,
			expected: time.Date(2025, time.March, 01, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Century leap year",
			start:    time.Date(2000, time.February, 29, 0, 0, 0, 0, time.UTC),
			years:    100,
			expected: time.Date(2100, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := addYears(tc.start, tc.years)
			if !result.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestWeekStart tests the WeekStart function.
func TestWeekStart(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Regular Monday",
			date:     time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Middle of the week",
			date:     time.Date(2024, time.October, 30, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Sunday before Monday",
			date:     time.Date(2024, time.November, 3, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "End of year",
			date:     time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, time.December, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := WeekStart(tc.date)
			if !result.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestDaysInYear tests the daysInYear function.
func TestDaysInYear(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "Non-leap year",
			date:     time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "Leap year",
			date:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "Century non-leap year",
			date:     time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "Century leap year",
			date:     time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := daysInYear(tc.date)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

// TestDaysInMonth tests the daysInMonth function.
func TestDaysInMonth(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "Standard month",
			date:     time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "February in leap year",
			date:     time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC),
			expected: 29,
		},
		{
			name:     "February in non-leap year",
			date:     time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC),
			expected: 28,
		},
		{
			name:     "April (30 days)",
			date:     time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := daysInMonth(tc.date)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
