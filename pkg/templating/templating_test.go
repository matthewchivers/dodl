package templating

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestRenderTemplate_GeneralPatterns tests rendering general patterns.
func TestRenderTemplate_GeneralPatterns(t *testing.T) {
	testCases := []struct {
		name         string
		template     string
		customFields map[string]interface{}
		testTime     time.Time
		expected     string
	}{
		{
			name:     "Standard date",
			template: "{{ .Now | date \"2006-01\" }}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "2024-10",
		},
		{
			name:     "Standard date - no format specified",
			template: "{{ .Today }}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "2024-10-29 00:00:00 +0000 UTC",
		},
		{
			name:     "Custom Field",
			template: "{{.Author}}",
			customFields: map[string]interface{}{
				"Author": "Alice",
			},
			testTime: time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC),
			expected: "Alice",
		},
		{
			name:     "uppercase custom field",
			template: "{{.Author | upper}}",
			customFields: map[string]interface{}{
				"Author": "Alice",
			},
			testTime: time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC),
			expected: "ALICE",
		},
		{
			name:     "WeekStart - Sunday",
			template: "{{ .WeekStart | date \"2 Jan 2006\" }}",
			testTime: time.Date(2024, time.October, 27, 0, 0, 0, 0, time.UTC),
			expected: "21 Oct 2024",
		},
		{
			name:     "Range from week-start to week-end (start + 6 days)",
			template: "{{ .WeekStart | date \"02 Jan\" }} to {{ addDays .WeekStart 6 | date \"02 Jan\" }}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "28 Oct to 03 Nov",
		},
		{
			name:     "Range from week-start to week-end (start + 6 days) - different months",
			template: "{{ .WeekStart | date \"02 Jan\" }} to {{ addDays .WeekStart 6 | date \"02 Jan\" }}",
			testTime: time.Date(2024, time.April, 30, 0, 0, 0, 0, time.UTC),
			expected: "29 Apr to 05 May",
		},
		{
			name:     "Week Start calculated after pipe",
			template: "{{ .Now | calcWeekStart | date \"02 Jan 2006\" }}",
			testTime: time.Date(2024, time.October, 30, 0, 0, 0, 0, time.UTC),
			expected: "28 Oct 2024",
		},
		{
			name:     "Week number",
			template: "Week {{ .Now | weekNumber }}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "Week 44",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			data := make(map[string]interface{})
			for k, v := range tt.customFields {
				data[k] = v
			}
			result, err := RenderTemplate(tt.template, data, tt.testTime)
			if err != nil {
				t.Fatalf("RenderTemplate failed: %v", err)
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestRenderTemplate_FileNamePatterns tests rendering file name patterns.
func TestRenderTemplate_FileNamePatterns(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		testTime time.Time
		expected string
	}{
		{
			name:     "Standard filename",
			template: "journal {{ .Now | date \"02-01-06\" }}.md",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "journal 29-10-24.md",
		},
		{
			name:     "Filename with month name",
			template: "report_{{ .Now | date \"Jan_2006\" }}.txt",
			testTime: time.Date(2025, time.December, 15, 0, 0, 0, 0, time.UTC),
			expected: "report_Dec_2025.txt",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			data := make(map[string]interface{})
			result, err := RenderTemplate(tt.template, data, tt.testTime)
			if err != nil {
				t.Fatalf("RenderTemplate failed: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestRenderTemplate_FileContentTemplates tests rendering file content templates.
func TestRenderTemplate_FileContentTemplates(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		testTime time.Time
		expected string
	}{
		{
			name:     "Standard content",
			template: "{{ .Now | date \"02 January 2006\" }} - Day {{ .Now.YearDay }}/{{ daysInYear .Now}}",
			testTime: time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
			expected: "29 October 2024 - Day 303/366",
		},
		{
			name:     "Non-leap year",
			template: "{{ .Now | date \"02 January 2006\" }} - Day {{ .Now.YearDay }}/{{ daysInYear .Now}}",
			testTime: time.Date(2023, time.February, 28, 0, 0, 0, 0, time.UTC),
			expected: "28 February 2023 - Day 59/365",
		},
		{
			name:     "Leap day (fixed to 365)",
			template: "{{ .Now | date \"02 January 2006\" }} - Day {{ .Now.YearDay }}/{{ daysInYear .Now}}",
			testTime: time.Date(2020, time.February, 29, 0, 0, 0, 0, time.UTC),
			expected: "29 February 2020 - Day 60/366",
		},
		{
			name:     "Custom format with addDays",
			template: "Date in 5 days: {{ addDays .Now 5 | date \"02-01-2006\" }}",
			testTime: time.Date(2024, time.December, 27, 0, 0, 0, 0, time.UTC),
			expected: "Date in 5 days: 01-01-2025",
		},
		{
			name:     "Custom format with addMonths",
			template: "Date in 3 months: {{ addMonths .Now 3 | date \"02-01-2006\" }}",
			testTime: time.Date(2024, time.December, 27, 0, 0, 0, 0, time.UTC),
			expected: "Date in 3 months: 27-03-2025",
		},
		{
			name:     "Custom format with addYears",
			template: "Date in 2 years: {{ addYears .Now 2 | date \"02-01-2006\" }}",
			testTime: time.Date(2024, time.December, 27, 0, 0, 0, 0, time.UTC),
			expected: "Date in 2 years: 27-12-2026",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			data := make(map[string]interface{})
			result, err := RenderTemplate(tt.template, data, tt.testTime)
			if err != nil {
				t.Fatalf("RenderTemplate failed: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
