package main

import (
	"fmt"
)

func main() {
	// dichiarare variabili
	var s string
	var x int
	/*
		var b bool
		var f float64
		var c rune
	*/

	// inizializzazione variabili
	s = "ciao"
	x = 4
	/*
		b = true
		f = 1.4
		c = 'c'
	*/

	// oppure insieme
	var z int = 10
	// var somma int = z + x //14

	// inferenza di tipo
	//y := "pippo"

	// tipo di una variabile
	// fmt.Println(reflect.TypeOf(y))

	// esempi
	s2 := "simo"

	// output
	fmt.Println(s + s2) //concatenazione
	fmt.Println(s, s2)  // dopo , aggiunge uno spazio e va a capo
	fmt.Print(s, s2)    // dopo , non aggiuge uno spazio e non va a capo
	// operazioni con varibili
	fmt.Printf("\nLa somma tra %d + %d è: %d.\n", x, z, x+z)

	// standard input
	var numero1 int
	fmt.Scan(&numero1)
	fmt.Printf("Il numero inserito è: %d.\n", numero1)
	// redirezione dell'input go run <nomefile>.go < <nomefileinput>.txt
}
