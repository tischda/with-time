package main

import (
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	// Thu Oct 16 17:42:17 CEST 2025
	// Using a fixed time for predictable test outcomes.
	// The timezone is set to a specific location to avoid issues with local timezones.
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Fatalf("could not load location: %v", err)
	}
	testTime := time.Date(2025, 10, 16, 17, 42, 17, 0, loc)

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty format",
			input:    `echo Time is %TIME:%`,
			expected: `echo Time is Thu Oct 16 17:42:17 CEST 2025`,
		},
		{
			name:     "YYYYMMDD-HHmmss format",
			input:    `echo Time is %TIME:YYYYMMDD-HHmmss%`,
			expected: `echo Time is 20251016-174217`,
		},
		{
			name:     "No placeholder",
			input:    `echo "No placeholder here"`,
			expected: `echo "No placeholder here"`,
		},
		{
			name:     "Multiple placeholders",
			input:    `echo "Date: %TIME:YYYY-MM-DD%, Time: %TIME:HH:mm:ss%"`,
			expected: `echo "Date: 2025-10-16, Time: 17:42:17"`,
		},
		{
			name:     "Go layout format",
			input:    `echo "Time is %TIME:Jan 02 2006%"`,
			expected: `echo "Time is Oct 16 2025"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := process(tc.input, testTime)
			if got != tc.expected {
				t.Errorf("process(%q) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}
