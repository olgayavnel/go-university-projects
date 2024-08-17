package main

import (
	"math/rand"
	"testing"
)

func TestAlternierendeSummeWerte(t *testing.T) {
	tests := [...]struct {
		a [6][5]int
		b [6][5]int
		f int
	}{
		{[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {27, 65, 13, 31, -1}},
			[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {27, -5, 13, 31, -1}},
			1},
		{[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {65, -5, 13, 65, -1}},
			[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {27, -5, 13, 31, -1}},
			2},
		{[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {65, 65, 65, 31, -1}},
			[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {27, -5, 13, 31, -1}},
			3},
		{[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {65, -5, 65, 65, 65}},
			[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {27, -5, 13, 31, -1}},
			4},
		{[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {65, 65, 65, 65, 65}},
			[6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10}, {9, 2, 25, 18, 11}, {27, -5, 13, 31, -1}},
			5},
	}
	for _, tt := range tests {
		a := tt.a
		got := alternierendeSpaltensumme(&tt.a)
		if got != tt.f || tt.a != tt.b {
			t.Errorf("\nalternierendeSpaltensumme(%v)\n= %v, want %v Fehler\nwant %v\ngot  %v\n", a, got, tt.f, tt.b, tt.a)
		}
	}
}

func TestAlternierendeSummeEigenschaften(t *testing.T) {
	var tt [anzZeilen][anzSpalten]int
	for nr := 1; nr <= 100; nr++ {
		for i := 0; i < anzZeilen; i++ {
			for j := 0; j < anzSpalten; j++ {
				tt[i][j] = rand.Intn(1000) - 500
			}
		}
		a := tt
		got := alternierendeSpaltensumme(&tt)
		if got < 0 || got > anzSpalten {
			t.Errorf("\nalternierendeSpaltensumme(%v)\n= %v Fehler\n aber es kann nur 0 bis %v Fehler geben", a, got, anzSpalten)
		} else if got == 0 && a != tt {
			t.Errorf("\nalternierendeSpaltensumme(%v)\n= %v Fehler\n aber Array differiert nach Korrektur: %v\n", a, got, tt)
		} else if got != 0 && a == tt {
			t.Errorf("\nalternierendeSpaltensumme(%v)\n= %v Fehler, aber Array identisch nach Korrektur.\n", a, got)
		} else {
			b := tt
			got := alternierendeSpaltensumme(&tt)
			if got != 0 {
				t.Errorf("%v wurde zu\n%v korrigiert, aber\nalternierendeSpaltensumme(%v)\n= %v Fehler\n ", a, tt, b, got)
			}
		}
	}
}
