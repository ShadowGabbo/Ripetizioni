package main

import "fmt"

func get_n() int {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return n
}

func main() {
	// funzione che da standard input ritorna n
	n := get_n()

	for i := 0; i < n*2; i++ { //righe
		for j := 0; j < n*2-1; j++ { //colonne
			if (i == 0 && j < n) || (i == n*2-1 && j >= n-1) {
				fmt.Print("*")
			} else if j == n-1 {
				fmt.Print("*")
			} else if (i == j && i < n-1) || (i == j+1 && i > n) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
