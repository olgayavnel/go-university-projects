package main

import "fmt"

// *** Größe der Matrix hier festlegen ***
const anzZeilen = 6 // rows
const anzSpalten = 5 // columns

// Dieses Programm dient lediglich dazu, die Funktion alternierendeSpaltensumme (Alternating Column Sum)
// einmal exemplarisch aufzurufen.
// Wenn Sie Ihre Funktion alternierendeSpaltensumme mit eigenen Daten testen möchten,
// können Sie dazu auch die Datei alternierendeSpaltenSumme_test.go entsprechend ergänzen.
func main() {
	var arr [6][5]int
	arr = [6][5]int{
		{15, 8, 1, 24, 17},
		{16, 14, 7, 5, 23},
		{22, 20, 13, 6, 4},
		{3, 21, 19, 12, 10},
		{9, 2, 25, 18, 11},
		{0, 0, 0, 0, 0},
	}
	fmt.Println("Im Array")
	for _, z := range arr {
		fmt.Println(z)
	}
	fehler := alternierendeSpaltensumme(&arr)
	fmt.Println("waren", fehler, "Spaltensummen fehlerhaft, die nun korrigiert sind:")
	for _, z := range arr {
		fmt.Println(z)
	}
}
