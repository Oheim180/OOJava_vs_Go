package main

import "fmt"

// Person ist eine Struktur, die das Verhalten einer Person darstellt.
type Person struct {
	name string
}

// Mitarbeiter ist eine Struktur, die das Verhalten eines Mitarbeiters darstellt.
type Mitarbeiter struct {
	Person
	mitarbeiterNr int
}

// Kunde ist eine Struktur, die das Verhalten eines Kunden darstellt.
type Kunde struct {
	Person
	kundenNr int
}

//Sammlung von Methoden für Personen
type person interface {
	//sprechen() string
}

func checkTyp(any interface{}) string {
	switch v := any.(type) {
	case Person:
		return "Der Typ ist Person und hat den Namen: " + v.name
	case Mitarbeiter:
		return "Der Typ ist Mitarbeiter und hat den Namen: " + v.name
	case Kunde:
		return "Der Typ ist Kunde und hat den Namen: " + v.name
	default:
		return "Der Typ ist weder Person noch Mitarbeiter oder Kunde"
	}
}
func main() {
	p := Person{"Max"}
	m := Mitarbeiter{Person: Person{"Bernd"}, mitarbeiterNr: 123}
	k := Kunde{Person: Person{"Herbert"}, kundenNr: 456}

	fmt.Printf(checkTyp(p) + "\n")
	fmt.Printf(checkTyp(m) + "\n")
	fmt.Printf(checkTyp(k))
}
