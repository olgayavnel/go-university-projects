package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// *** Anzahl der Orte hier festlegen ***
const numberOfCities = 4

// Das Programm demonstriert die Verwendung der Funktion rundreise,
// indem es eine zufällige Entfernungsmatrix erzeugt, eine Rundreise
// berechnet und diese ausgibt.
func main() {
	// Anzahl der Einträge oberhalb der Diagonalen einer symmetrischen Matrix;
	// hier also die (mindestens) benötigte Anzahl von unterschiedlichen Entfernungswerten,
	// die per Zufall erzeugt werden müssen
	const anzEntf int = numberOfCities * (numberOfCities - 1) / 2

	const skalEnt int = 20  // gewünschte Skalierung der Entfernungen (mindestens 1)
	const mindEntf int = 11 // gewünschte Mindestentfernung zwischen zwei Orten

	// Anzahl der max. Stellen eines Entfernungswerts ermitteln (für saubere Ausgabe der Matrix)
	stellen := int(math.Ceil(math.Log10(float64(skalEnt*anzEntf + mindEntf))))
	// entsprechenden Format-String erstellen
	format := "%" + strconv.Itoa(stellen) + "d"

	// Matrix mit zufälligen Entfernungen (gemäß obigen Vorgaben) generieren
	var entf [numberOfCities][numberOfCities]int
	zufall := rand.Perm(skalEnt * anzEntf) // benötigte Mindestanzahl anzEntf wird um Faktor skalEnt skaliert
	zNr := 0
	for zeile := 0; zeile < numberOfCities; zeile++ {
		for spalte := zeile + 1; spalte < numberOfCities; spalte++ {
			abstand := zufall[zNr] + mindEntf
			zNr = zNr + 1
			entf[zeile][spalte] = abstand
			entf[spalte][zeile] = abstand
		}
	}

	// saubere Ausgabe der Entfernungsmatrix (inkl. Zeilen-Spalten-Beschriftung) erzeugen,
	// indem jeder Eintrag entsprechend der (max.) Stellen-Anzahl passend eingerückt wird
	fmt.Println("Entfernungsmatrix der Orte:")

	fmt.Printf(" %"+strconv.Itoa(stellen)+"s", "")
	for i := range numberOfCities {
		fmt.Printf(" "+format, i+1)
	}
	fmt.Println()
	for zeile := 0; zeile < numberOfCities; zeile++ {
		fmt.Printf(format+" [", zeile+1)
		for spalte := 0; spalte < numberOfCities; spalte++ {
			fmt.Printf(format, entf[zeile][spalte])
			if spalte == numberOfCities-1 {
				fmt.Println("]")
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()

	// Rundreise berechnen
	route, gesamtstrecke := calculateRoundTrip(entf)

	// Ergebnisse ausgeben
	fmt.Println("Route der Rundreise:", route)
	fmt.Println("Gesamtstrecke:      ", gesamtstrecke)
}
