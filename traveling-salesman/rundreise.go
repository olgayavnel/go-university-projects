package main

import "math"

// TODO: write a function to calculate the route and the total distance

// distanceMatrix := [3][3]int{
//     {0, 10, 20}, // Distances from City 1 to City 1, City 2, City 3
//     {10, 0, 15}, // Distances from City 2 to City 1, City 2, City 3
//     {20, 15, 0}, // Distances from City 3 to City 1, City 2, City 3
// }
//          City 1  City 2  City 3
// 	City 1       0      10      20
// 	City 2      10       0      15
// 	City 3      20      15       0

func calculateRoundTrip(distanceMatrix [numberOfCities][numberOfCities]int) ([numberOfCities + 1]int, int) {
	var visitedCities [numberOfCities]bool // array of boolean values with length numberOfCities, all initialized to false
	var route [numberOfCities+1]int
	totalDistance := 0

	currentCity := 0 // start at the fist city
	route[0] = currentCity + 1 // add the starting city to the route (1-indexed)
	visitedCities[currentCity] = true

	// main loop
	for i := 1; i < numberOfCities; i++ {
		nearestCity := -1
		shortestDistance := math.MaxInt

		// find the nearest unvisited city
		for j := 0; j < numberOfCities; j++ {
			if !visitedCities[j] && distanceMatrix[currentCity][j] < shortestDistance {
				shortestDistance = distanceMatrix[currentCity][j]
				nearestCity = j
			}
		}

		route[i] = nearestCity+1 // add the nearest city to the route (1-indexed)
		visitedCities[nearestCity] = true
		totalDistance += shortestDistance
		currentCity = nearestCity
	}

	// return to the starting city
	route[numberOfCities] = 1
	totalDistance += distanceMatrix[currentCity][0]
	
	return route, totalDistance // e.g. [1 4 3 2 1], 282
}

