package solution

import (
	"strings"
)

// Solution finds the minimum number of water tanks needed for all houses
// Returns -1 if it's not possible to place tanks for all houses
func Solution(S string) int {
	// Check if solution is impossible
	if S == "H" || S == "" || !strings.Contains(S, "-") {
		return -1
	}

	street := []rune(S)
	tanksTotal := 0
	streetLength := len(street)

	for i := 0; i < streetLength; i++ {
		if street[i] == 'H' {
			// Check if house already has water from a tank
			if (i > 0 && street[i-1] == 'T') || (i+1 < streetLength && street[i+1] == 'T') {
				continue
			}

			// Try to place tank to the right first (greedy approach)
			if i+1 < streetLength && street[i+1] == '-' {
				street[i+1] = 'T'
				tanksTotal++
				continue
			}

			// If no space to the right, try to the left
			if i > 0 && street[i-1] == '-' {
				street[i-1] = 'T'
				tanksTotal++
				continue
			}

			// If we get here, no tank could be placed for this house
			return -1
		}
	}

	return tanksTotal
}
