package solution

import (
	"fmt"
	"strings"
)

// MinimumWaterTanks calculates the minimum number of water tanks needed
// to provide water to all houses in a street.
// Returns the count of tanks and the resulting street arrangement.
// Returns -1 if it's not possible to place tanks for all houses.
//
// Parameters:
//   - street: A string representing the street layout
//     'H' represents a house
//     '-' represents an empty plot where a tank can be placed
//
// Returns:
//   - int: The minimum number of tanks needed (-1 if impossible)
//   - string: The street layout with tanks placed ('T' for tank)
func MinimumWaterTanks(street string) (int, string) {
	// Convert to uppercase for consistency
	street = strings.ToUpper(street)

	// Check if there are any empty plots to place tanks
	if !strings.Contains(street, "-") {
		return -1, ""
	}

	// Convert the street string to a rune array for modification
	layout := []rune(street)
	streetLength := len(street)
	tankCount := 0

	// Process each position in the street
	for i := 0; i < streetLength; i++ {
		// Only process house positions
		if street[i] == 'H' {
			// Skip if this house already has a water source
			// (a tank either to its left or right)
			if (i > 0 && layout[i-1] == 'T') || (i+1 < streetLength && layout[i+1] == 'T') {
				continue
			}

			// Try to place a tank for this house
			tankPlaced := false

			// First strategy: Place tank to the right of the house if possible
			// This is optimal as it can potentially serve the next house too
			if i+1 < streetLength && street[i+1] == '-' {
				layout[i+1] = 'T'
				tankCount++
				tankPlaced = true
			}

			// Second strategy: If can't place to the right, try placing to the left
			if !tankPlaced && i > 0 && street[i-1] == '-' {
				layout[i-1] = 'T'
				tankCount++
				tankPlaced = true
			}

			// If no tank could be placed for this house, it's impossible to solve
			if !tankPlaced {
				return -1, ""
			}
		}
	}

	return tankCount, string(layout)
}

// Solution calculates the minimum number of water tanks needed
// to provide water to all houses in a street.
func Solution(S string) int {
	if notPossibleSolution(S) {
		return -1
	}

	street := []rune(S)
	tanksTotal := 0
	streetLength := len(street)
	for i := 0; i < streetLength; i++ {
		if HouseHere(street[i]) {
			if (i > 0 && TankHere(street[i-1])) || (i+1 < streetLength && TankHere(street[i+1])) {
				continue
			}

			if i+1 < streetLength && EmptySpot(street[i+1]) {
				street[i+1] = 'T'
				tanksTotal++
				continue
			}

			if i-1 >= 0 && EmptySpot(street[i-1]) {
				street[i-1] = 'T'
				tanksTotal++
				continue
			}

			// If we can't place a tank for a house, it's impossible
			return -1
		}
	}
	return tanksTotal
}

func HouseHere(spot rune) bool {
	return spot == 'H'
}

func TankHere(spot rune) bool {
	return spot == 'T'
}

func EmptySpot(spot rune) bool {
	return spot == '-'
}

func notPossibleSolution(S string) bool {
	return S == "H" || S == "" || !strings.Contains(S, "-")
}

func main() {
	// Example from the problem statement
	exampleStreet := "-H-HH--"
	tankCount, resultLayout := MinimumWaterTanks(exampleStreet)

	fmt.Println("--- Example from problem statement ---")
	fmt.Println("Street layout:     ", exampleStreet)
	fmt.Println("Minimum tanks:     ", tankCount)
	fmt.Println("Resulting layout:  ", resultLayout)
	fmt.Println()

	// Test cases to verify the algorithm
	testCases := []struct {
		street      string
		description string
	}{
		{"H", "Only one house, no empty plots (impossible)"},
		{"----H----", "One house with many empty plots"},
		{"HHH", "All houses, no empty plots (impossible)"},
		{"H-H-H", "Alternating houses and empty plots"},
		{"-H-H-HH-", "Complex case with multiple houses"},
		{"--H-H--H--", "Multiple houses with multiple empty plots"},
	}

	fmt.Println("--- Additional Test Cases ---")
	for _, tc := range testCases {
		count, result := MinimumWaterTanks(tc.street)
		status := "Success"
		if count == -1 {
			status = "Impossible"
		}

		fmt.Printf("Street: %-12s | Tanks: %-2d | Result: %-12s | Status: %-10s | %s\n",
			tc.street, count, result, status, tc.description)
	}
}
