package main

import "fmt"

func input() (int, int) {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	return n1, n2
}

func somma(n1 int, n2 int) int {
	return n1 + n2
}

func sottrazione(n1 int, n2 int) int {
	return n1 - n2
}

func stampa(somma int) {
	fmt.Printf("La somma è : %d", somma)
}

func main() {
	// Programma che dati due numeri da standard input restituisce la somma con una funzione

	// funzione che prende numeri da standard input e li restituisce
	x, y := input()

	// funzione che restituisce la somma di due numeri
	somma := somma(x, y)

	// funzione che stampa la somma
	stampa(somma)
}
