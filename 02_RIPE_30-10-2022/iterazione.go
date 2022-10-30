package main

func main() {
	/* stampare i numeri da 1 a 100
	//for inzio, condizione di fine, step
	for (var i int = 1; i <= 100; i++){
		ogni volta che i<= 100 è vero entro qua
	}
	*/

	/* esempio 1
	// casting: conversine di tipo
	// da stringa a intero atoi, da intero a stringa itoa (strconv)
	for i := 1; i <= 100; i++ {
		// fmt.Println("il numero è:",i)
		fmt.Println("il numero è: " + strconv.Itoa(i))
	}
	*/

	// NB: da evitare
	// break e continue
	// break interrompe il flusso
	// continue passa alla prossima iterazione

	/* esempio 2 (selezione e iterazione)
	// stampa i numeri pari da 1 a 100
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("il numero è: " + strconv.Itoa(i))
		}
	}


	for i := 1; i <= 100; i++ {
		fmt.Println("il numero è: " + strconv.Itoa(i))
		if i%2 == 1 {
			break
		}
	}
	*/

	/*
		var n int

		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&n)

		var i int
		for i = 1; i <= n;{
			fmt.Print(i, " ")
			i++
		}

		fmt.Println()
	*/

}
