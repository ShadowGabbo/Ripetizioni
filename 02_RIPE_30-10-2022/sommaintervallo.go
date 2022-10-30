package main

import "fmt"

/*
Scrivere un programma che legga da standard input due numeri
interi a e b e stampi a video la somma dei numeri
pari compresi tra a e b (estremi esclusi).
*/

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	var i int
	var somma int
	for i = a + 1; i < b; i++ {
		if i%2 == 0 {
			somma += i
		}
	}
	fmt.Printf("Somma = %d", somma)
}
