package starpackage

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
	Perimetro() float64
}

type Cerchio struct {
	Raggio float64
}

func (c Cerchio) Area() float64 {
	return math.Pi * c.Raggio * c.Raggio
}

func (c Cerchio) Perimetro() float64 {
	return 2 * math.Pi * c.Raggio
}

type Rettangolo struct {
	Larghezza, Altezza float64
}

func (r Rettangolo) Area() float64 {
	return r.Larghezza * r.Altezza
}

func (r Rettangolo) Perimetro() float64 {
	return 2*r.Larghezza + 2*r.Altezza
}

func StampaDettagli(f Forma) {
	switch forma := f.(type) {
	case Cerchio:
		fmt.Println("Cerchio con raggio", forma.Raggio)
	case Rettangolo:
		fmt.Println("Rettangolo con larghezza", forma.Larghezza, "e altezza", forma.Altezza)
	default:
		fmt.Println("Tipo di forma sconosciuto")
	}
	fmt.Println("Area:", f.Area())
	fmt.Println("Perimetro:", f.Perimetro())
}

func IntrospezioneMain() {
	c := Cerchio{Raggio: 5}
	r := Rettangolo{Larghezza: 3, Altezza: 4}

	StampaDettagli(c)
	StampaDettagli(r)
}
