package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ { //righe
		for j := 0; j < n; j++ { //colonne
			if i%3 == 0 {
				fmt.Print("* ")
			} else if i%3 == 1 {
				fmt.Print("+ ")
			} else {
				fmt.Print("o ")
			}
		}
		fmt.Println()
	}
}
