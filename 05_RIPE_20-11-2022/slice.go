package main

import "fmt"

func main() {
	// dichiaro slice
	var s []int

	// inizilizzo a valore di default n numeri
	s = make([]int, 5)
	fmt.Printf("La capacità è: %d, la lunghezza è: %d\n", cap(s), len(s))

	// dichiarazione e inizializzazione completa
	b := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%T", b)

	//inizializzazione personalizzata
	a := [...]int{1: 2, 3: 9}
	fmt.Println(a)

	//aggiungo un elemento alla slice
	s = append(s, 4)
	fmt.Printf("La capacità è: %d, la lunghezza è: %d\n", cap(s), len(s))
	fmt.Printf("Array: %v\n", s)

	//for-ternario sulle slide
	for i := 0; i < len(s); i++ {
		elemento := s[i]
		fmt.Printf("%d ", elemento)
	}
	fmt.Println()

	//for-range su array (iterare con indice crescente a passo 1)
	for indice, elemento := range s {
		fmt.Printf("%d %d\n", indice, elemento)
	}

	//slicing
	fmt.Println(b)

	//slice da x a y (escluso)
	fmt.Println(b[2:5])

	//slice da inizio a x(escluso)
	fmt.Println(b[:5])

	//slice da x a fine
	fmt.Println(b[2:])

	//
	fmt.Println("INIZIO")
	c := b // occhio collegati dal puntantore
	fmt.Println("b:", b)
	// b = append(b, 10) crea una nuova slice aggiungendo in coda il numero
	c[2] = 12
	fmt.Println("b:", b)
}
