package templating

import (
	"testing"
	"time"
)

func prepareTestTimes(testTime time.Time) map[string]interface{} {
	data := map[string]interface{}{
		"Today":          testTime,
		"WeekCommencing": weekCommencing(testTime),
	}
	return data
}

func TestRenderTemplate_DirectoryPatterns(t *testing.T) {
	testCases := []struct {
		name         string
		template     string
		customFields map[string]interface{}
		testTime     time.Time
		expected     string
	}{
		{
			name:     "Standard date",
			template: "{{ .Today | date \"2006/01\" }}/wc {{ .WeekCommencing | date \"02-01-06\" }}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "2024/10/wc 28-10-24",
		},
		{
			name:     "Custom Field",
			template: "{{ .Today | date \"2006/01\" }}/{{.Author}}/wc {{ .WeekCommencing | date \"02-01-06\" }}",
			customFields: map[string]interface{}{
				"Author": "Alice",
			},
			testTime: time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC),
			expected: "2023/11/Alice/wc 13-11-23",
		},
		{
			name:     "Edge case - Sunday",
			template: "{{ .Today | date \"2006/01\" }}/wc {{ .WeekCommencing | date \"2 Jan 2006\" }}",
			testTime: time.Date(2024, time.January, 7, 0, 0, 0, 0, time.UTC), // Sunday
			expected: "2024/01/wc 1 Jan 2024",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := prepareTestTimes(tc.testTime)
			for k, v := range tc.customFields {
				data[k] = v
			}
			result, err := RenderTemplate(tc.template, data)
			if err != nil {
				t.Fatalf("RenderTemplate failed: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

func TestRenderTemplate_FileNamePatterns(t *testing.T) {
	testCases := []struct {
		name         string
		template     string
		customFields map[string]interface{}
		testTime     time.Time
		expected     string
	}{
		{
			name:     "Standard filename",
			template: "journal {{ .Today | date \"02-01-06\" }}.md",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "journal 29-10-24.md",
		},
		{
			name:     "Filename with month name",
			template: "report_{{ .Today | date \"Jan_2006\" }}.txt",
			testTime: time.Date(2025, time.December, 15, 0, 0, 0, 0, time.UTC),
			expected: "report_Dec_2025.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := prepareTestTimes(tc.testTime)
			result, err := RenderTemplate(tc.template, data)
			if err != nil {
				t.Fatalf("RenderTemplate failed: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

func TestRenderTemplate_FileContentTemplates(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		testTime time.Time
		expected string
	}{
		{
			name:     "Standard content",
			template: "{{ .Today | date \"02 January 2006\" }} - Day {{ .Today.YearDay }}/{{ daysInYear .Today}}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "29 October 2024 - Day 303/366",
		},
		{
			name:     "Non-leap year",
			template: "{{ .Today | date \"02 January 2006\" }} - Day {{ .Today.YearDay }}/{{ daysInYear .Today}}",
			testTime: time.Date(2023, time.February, 28, 0, 0, 0, 0, time.UTC),
			expected: "28 February 2023 - Day 59/365",
		},
		{
			name:     "Leap day (fixed to 365)",
			template: "{{ .Today | date \"02 January 2006\" }} - Day {{ .Today.YearDay }}/{{ daysInYear .Today}}",
			testTime: time.Date(2020, time.February, 29, 0, 0, 0, 0, time.UTC),
			expected: "29 February 2020 - Day 60/366",
		},
		{
			name:     "Custom format with addDays",
			template: "Date in 5 days: {{ addDays .Today 5 | date \"02-01-2006\" }}",
			testTime: time.Date(2024, time.December, 27, 0, 0, 0, 0, time.UTC),
			expected: "Date in 5 days: 01-01-2025",
		},
		{
			name:     "Custom format with addMonths",
			template: "Date in 3 months: {{ addMonths .Today 3 | date \"02-01-2006\" }}",
			testTime: time.Date(2024, time.December, 27, 0, 0, 0, 0, time.UTC),
			expected: "Date in 3 months: 27-03-2025",
		},
		{
			name:     "Custom format with addYears",
			template: "Date in 2 years: {{ addYears .Today 2 | date \"02-01-2006\" }}",
			testTime: time.Date(2024, time.December, 27, 0, 0, 0, 0, time.UTC),
			expected: "Date in 2 years: 27-12-2026",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := prepareTestTimes(tc.testTime)
			result, err := RenderTemplate(tc.template, data)
			if err != nil {
				t.Fatalf("RenderTemplate failed: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
