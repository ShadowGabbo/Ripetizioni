package main

import (
	"fmt"
	"sort"
)

func main() {
	// data una frase mappare le lettere
	// esempio cacca c=3 a=2

	// dichiarazione
	mappa := make(map[string]int)

	// dichiarazione e inizializzazione
	//mappa2 := map[string]int{"a": 1, "b": 2}

	// aggiungere una coppia
	mappa["c"] = 3
	mappa["a"] = 2
	mappa["b"] = 5
	mappa["f"] = 23
	mappa["d"] = 19
	mappa["p"] = 4
	// fmt.Println(mappa)

	// cancellare una coppia data la chiave
	delete(mappa, "b")
	// fmt.Println(mappa)

	// accedere ad un valore data la chiave
	//fmt.Println(mappa["a"])

	// stampare tutte le coppie (loop con range non si puo' usare il for-ternario)
	/*
		for k, v := range mappa {
			fmt.Printf("%s=%d\n", k, v)
		}
	*/

	// sapere se esiste una chiave in una mappa
	// es voglio sapere se la chiave b esiste
	if _, ok := mappa["b"]; ok {
		fmt.Printf("esiste b")
	}

	// stampa della mappa in ordine di chiave (lessicografico)

	// 1) crea una slice di chiavi
	var keys []string
	for k := range mappa {
		fmt.Printf("Chiave: %s - Valore: %d\n", k, mappa[k])
		keys = append(keys, k) // slice di chiavi
	}

	fmt.Println(keys)

	// 2) ordina le chiavi
	sort.Strings(keys)
	fmt.Println(keys)

	// 3) loop su la slice ordinata -> mappa[k] ordinata
	for _, v := range keys {
		fmt.Printf("Chiave: %s - Valore: %d\n", v, mappa[v])
	}
}
