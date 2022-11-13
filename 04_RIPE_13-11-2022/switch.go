package main

import "fmt"

func main() {
	// selezione multiaria
	// switch versione variabile
	num := 0
	// if num == 0
	switch num {
	case 0:
		fmt.Println("Zero")
		num++
	case 1:
		fmt.Println("Uno")
	default:
		fmt.Println("Il numero è maggiore di 1")
	}

	// variabile versione espressioni
	switch {
	case num%2 == 0:
		fmt.Println("Il numero è pari.")
	default:
		fmt.Println("Il numero è dispari.")
	}
}
