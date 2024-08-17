package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestcalculateRoundTripEigenschaften(t *testing.T) {
	for i := 1; i <= 10000; i++ {
		// benötigte Mindestanzahl von Werten (Einträge oberhalb der Hauptdiagonalen) wird um Faktor 100 skaliert
		zufall := rand.Perm(100 * (numberOfCities * (numberOfCities - 1) / 2))
		zNr := 0
		var entf [numberOfCities][numberOfCities]int
		for zeile := 0; zeile < numberOfCities; zeile++ {
			for spalte := zeile + 1; spalte < numberOfCities; spalte++ {
				abstand := zufall[zNr] + 1
				zNr = zNr + 1
				entf[zeile][spalte] = abstand
				entf[spalte][zeile] = abstand
			}
		}
		orte, gesamtstrecke := calculateRoundTrip(entf)
		if orte[0] != 1 {
			t.Errorf("calculateRoundTrip(%v)\n= %v, %v\n aber der erste Ort muss 1 sein\n", entf, orte, gesamtstrecke)
		} else if orte[numberOfCities+1-1] != 1 {
			t.Errorf("calculateRoundTrip(%v)\n= %v, %v\n aber der letzte Ort muss 1 sein\n", entf, orte, gesamtstrecke)
		} else {
			sortiert := orte
			slices.Sort(sortiert[:])
			ok := true
			for i := 2; i < numberOfCities+1; i++ {
				if sortiert[i] != i {
					ok = false
				}
			}
			if !ok {
				t.Errorf("calculateRoundTrip(%v)\n= %v, %v\n aber die calculateRoundTrip muss bis auf den ersten alle Orte genau einmal enthalten.\n", entf, orte, gesamtstrecke)
			}
		}
	}
}

func TestcalculateRoundTripWerte(t *testing.T) {

	if numberOfCities != 4 {
		t.Errorf("Dieser Test ist nur durchführbar, wenn numberOfCities == 4 gilt.")
		return
	}

	tests := [...]struct {
		e [numberOfCities][numberOfCities]int
		o [numberOfCities + 1]int
		g int
	}{
		{[numberOfCities][numberOfCities]int{
			{0, 1, 2, 3},
			{1, 0, 4, 5},
			{2, 4, 0, 6},
			{3, 5, 6, 0},
		},
			[numberOfCities + 1]int{1, 2, 3, 4, 1},
			14},
		{[numberOfCities][numberOfCities]int{
			{0, 2, 5, 4},
			{2, 0, 6, 1},
			{5, 6, 0, 3},
			{4, 1, 3, 0},
		},
			[numberOfCities + 1]int{1, 2, 4, 3, 1},
			11},
		{[numberOfCities][numberOfCities]int{
			{0, 3, 1, 6},
			{3, 0, 4, 5},
			{1, 4, 0, 7},
			{6, 3, 7, 0},
		},
			[numberOfCities + 1]int{1, 3, 2, 4, 1},
			16},
		{[numberOfCities][numberOfCities]int{
			{0, 5, 1, 3},
			{5, 0, 8, 4},
			{1, 8, 0, 7},
			{3, 4, 7, 0},
		},
			[numberOfCities + 1]int{1, 3, 4, 2, 1},
			17},
		{[numberOfCities][numberOfCities]int{
			{0, 7, 8, 3},
			{7, 0, 5, 9},
			{8, 5, 0, 4},
			{3, 9, 4, 0},
		},
			[numberOfCities + 1]int{1, 4, 3, 2, 1},
			19},
		{[numberOfCities][numberOfCities]int{
			{0, 7, 5, 1},
			{7, 0, 4, 2},
			{5, 4, 0, 6},
			{1, 2, 6, 0},
		},
			[numberOfCities + 1]int{1, 4, 2, 3, 1},
			12},
	}
	for _, tt := range tests {
		o, g := calculateRoundTrip(tt.e)
		if o != tt.o || g != tt.g {
			t.Errorf("calculateRoundTrip(%v)\n=    %v, %v\nwant %v, %v\n", tt.e, o, g, tt.o, tt.g)
		}
	}
}
