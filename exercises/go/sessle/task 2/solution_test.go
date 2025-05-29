package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		street        string
		expectedTanks int
		description   string
	}{
		{"-H-HH--", 2, "Example from problem statement"},
		{"H", -1, "Only one house, no empty plots (impossible)"},
		{"----H----", 1, "One house with many empty plots"},
		{"HHH", -1, "All houses, no empty plots (impossible)"},
		{"H-H-H", 3, "Alternating houses and empty plots"},
		{"-H-H-HH-", 3, "Complex case with multiple houses"},
		{"--H-H--H--", 3, "Multiple houses with multiple empty plots"},
		{"", -1, "Empty string"},
		{"-----", 0, "Only empty plots, no houses"},
		{"H--H", 2, "Two houses with space between"},
		{"HH-", 1, "Two houses with one empty plot at the end"},
		{"-HH", 1, "Two houses with one empty plot at the beginning"},
	}

	for _, tc := range testCases {
		result := Solution(tc.street)
		if result != tc.expectedTanks {
			t.Errorf("For street %q, expected %d tanks, but got %d. %s",
				tc.street, tc.expectedTanks, result, tc.description)
		}
	}
}
