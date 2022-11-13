package main

import (
	"fmt"
	"unicode"
)

func main() {
	var stringa string = "cioè"

	count := 0
	// for ternario sulle stringe stampa i byte (ok per US-ASCII usare il for che itera sui caratteri)
	for i := 0; i < len(stringa); i++ {
		fmt.Printf("%d\n", stringa[i])
		count++
	}
	fmt.Printf("La parola è formata da %d lettere.\n", count)

	// for range io itero sui caratteri (rune)
	for indice, lettera := range stringa {
		fmt.Printf("Indice: %d, Lettera: %c.\n", indice, lettera)
	}

	// Stampa il numero di lettere in una frase US-ASCII
	stringa2 := "Ciao \n come stai"

	// versione for ternario
	count2 := 0
	for i := 0; i < len(stringa2); i++ {
		lettera := rune(stringa2[i])
		if unicode.IsSpace(lettera) != true {
			count2++
		}
	}
	fmt.Printf("Versione for ternario - Frase formata da %d lettere.\n", count2)

	count2 = 0
	// versione for range
	for _, lettera := range stringa2 {
		if unicode.IsSpace(lettera) != true {
			count2++
		}
	}
	fmt.Printf("Versione for range - Frase formata da %d lettere.\n", count2)
}
