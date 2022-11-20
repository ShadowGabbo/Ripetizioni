package main

import "fmt"

func main() {
	// data un numero di elementi e gli elementi, sommali.
	// 4
	// 0 1 2 3
	// La somma è : 6

	var numeri []int
	var volte int

	fmt.Println("Input: ")
	fmt.Scan(&volte)
	numeri = make([]int, volte)

	for i := 0; i < volte; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(numeri)

	somma := 0
	for _, num := range numeri {
		somma += num
	}
	fmt.Printf("La somma è : %d", somma)
}
