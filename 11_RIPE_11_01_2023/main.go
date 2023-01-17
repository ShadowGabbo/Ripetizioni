package main

import "fmt"

type Persona struct {
	nome string
	anni int
}

func funzione(sl []int) {
	for i := 0; i < len(sl); i++ {
		sl[i] = 5
	}
}

func funzione_mappa(mappa map[string]int) {
	for chiave, _ := range mappa {
		mappa[chiave] = 5
	}
}

func main() {
	mappa := make(map[string]int)
	mappa["a"] = 2
	mappa["c"] = 4
	mappa["d"] = 3

	fmt.Println(mappa)
	for _, valore := range mappa {
		valore += 5
	}
	fmt.Println(mappa)

	/*
		if val, ok := mappa["a"]; ok{

		}
	*/

	var slice []int = []int{1, 2, 3, 4}
	//fmt.Println(slice)
	funzione(slice)
	funzione_mappa(mappa)
	fmt.Println(mappa)
	//fmt.Println(slice)

	// fattoriale 3 = 3*2*1
	fmt.Println(fattoriale(3))

	// fattoriale(3)
	// main : 3 * fattoriale(2)
	// fattoriale(2)
	// main : 3 * 2 * fattoriale(1)
	// fattoriale(1)
	// main : 3 * 2 * 1

	// encapsuling
	persone := []Persona{{"Mario", 23}, {"Simona", 21}, {"Gesu", 2022}}
	fmt.Println(persone)
	//modifica_mario(persone)
	persone[0] = modifica_mario_2(persone[0])
	fmt.Println(persone)
}

func modifica_mario(persone []Persona) {
	persone[0] = Persona{"Cristo Signore", 0}
}

func modifica_mario_2(persona Persona) Persona {
	persona.nome = "Pergiorgio"
	persona.anni = 20
	return persona
}

func fattoriale(num int) int {
	// caso base
	if num == 1 {
		return 1
	}

	// caso ricorsivo, passo ricorsivo
	return num * fattoriale(num-1)
}
