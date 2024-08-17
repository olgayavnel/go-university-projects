package main

// alternierendeSpaltensumme checks the array arr to determine if the
// last row contains the alternating column sums, corrects them if necessary,
// and returns the number of incorrect sums.
func alternierendeSpaltensumme(arr *[anzZeilen][anzSpalten]int) int {
	corrections := 0

	// Iterate over each column
	for col := 0; col < anzSpalten; col++ {
		alternatingSum := 0

		// Calculate the alternating sum for this column
		for row := 0; row < anzZeilen-1; row++ { // Exclude the last row
			if row%2 == 0 {
				alternatingSum += arr[row][col]
			} else {
				alternatingSum -= arr[row][col]
			}
		}

		// Compare with the last row's value
		if arr[anzZeilen-1][col] != alternatingSum {
			// If incorrect, update the last row with the correct sum
			arr[anzZeilen-1][col] = alternatingSum
			corrections++
		}
	}

	return corrections
}