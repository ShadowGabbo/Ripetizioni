package main

import "fmt"

/*
Scrivere un programma che, dopo aver richiesto all'utente di inserire da
standard input un numero intero n, stampi a video la corrispondente
tabellina (moltiplicando n per i numeri naturali da 1 a 10)
come mostrato nell'Esempio d'esecuzione.
*/

func main() {
	var num int = 0

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	var i int
	for i = 1; i <= 10; i++ {
		var prodotto int
		prodotto = i * num
		//fmt.Println(i, "x", num, "=", prodotto)
		fmt.Printf("%d x %d = %d\n", i, num, prodotto) // (i*num)
	}
}
