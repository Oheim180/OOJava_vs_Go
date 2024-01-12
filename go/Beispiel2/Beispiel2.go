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

// Buerger ist eine Struktur, die das Verhalten eines Buerger darstellt.
type Buerger struct {
	Person
}

// Sammlung von Methoden fuer Personen
type person interface {
	sprechen() string
}

func (p Person) sprechen() string {
	return "Ich bin Person " + p.name
}
func (m Mitarbeiter) sprechen() string {
	return "Ich bin Mitarbeiter " + m.name
}
func (k Kunde) sprechen() string {
	return "Ich bin Kunde " + k.name
}
func main() {
	p := Person{"Max"}
	m := Mitarbeiter{Person: Person{"Bernd"}, mitarbeiterNr: 123}
	k := Kunde{Person: Person{"Herbert"}, kundenNr: 456}
	b := Buerger{Person: Person{"Manfred"}}

	fmt.Printf(p.sprechen() + "\n")
	fmt.Printf(m.sprechen() + "\n")
	fmt.Printf(k.sprechen() + "\n")
	fmt.Printf(b.sprechen())
}

// Beliebiger Subtyp von Person moeglich
func sprechen(x person) string {
	return x.sprechen()
}
