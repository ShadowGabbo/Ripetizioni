package main

import "fmt"

func main() {
	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ { //righe
		for j := 0; j < n; j++ { //colonne
			if j%2 == 1 {
				fmt.Print("+ ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
