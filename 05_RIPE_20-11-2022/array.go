package main

import "fmt"

func main() {
	// dichiarazione di un array
	var array [4]int // array [0,0,0,0]

	// dichiarazione e inizializzazione
	var arr [4]int = [4]int{1, 2, 3, 4}

	//inizializzazione array completa
	array = [4]int{4, 3, 2, 1}

	//inizializzazione parziale
	array[1] = 5

	//for-ternario su array
	for i := 0; i < len(array); i++ {
		elemento := array[i]
		fmt.Printf("%d ", elemento)
	}

	//for-range su array (iterare con indice crescente a passo 1)
	for indice, elemento := range array {
		fmt.Printf("%d %d\n", indice, elemento)
	}
}
