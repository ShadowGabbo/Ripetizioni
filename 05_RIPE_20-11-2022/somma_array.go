package main

import "fmt"

func main() {
	//esercizio: somma dei numeri nell'array

	// dichiarazione e inizializzazione
	var arr [4]int = [4]int{1, 2, 3, 4}
	var somma int = 0

	for _, elemento := range arr {
		somma += elemento
	}

	fmt.Printf("La somma è: %d", somma)
}
