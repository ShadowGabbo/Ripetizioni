package main

import (
	"fmt"
	"os"
)

func isOrdinato(sl []string) bool {
	for i := 0; i < len(sl)-1; i++ {
		if len(sl[i]) < len(sl[i+1]) {
			return false
		}
	}
	return true
}

func ordina(sl []string) []string {
	for i := 0; i < len(sl)-1; i++ {
		if len(sl[i]) < len(sl[i+1]) {
			sl[i], sl[i+1] = sl[i+1], sl[i]
		}
	}
	return sl
}

func main() {
	testo := os.Args[1]
	mappa := make(map[string]int)
	var chiavi []string

	for i := 0; i < len(testo); i++ {
		for j := i + 1; j < len(testo); j++ {
			if testo[i] == testo[j] {
				sottostringa := testo[i : j+1]
				if len(sottostringa) >= 3 {
					mappa[sottostringa]++
				}
			}
		}
	}

	for k := range mappa {
		chiavi = append(chiavi, k)
	}

	//ORDINARE PER LUNGHEZZA
	for {
		flag := isOrdinato(chiavi)
		if flag {
			break
		}
		chiavi = ordina(chiavi)
		//fmt.Printf("ordinato un po di piu: %v\n", chiavi)
	}

	// versione veloce ed elegante
	/*
		sort.Slice(chiavi, func(i, j int) bool {
			return len(chiavi[i]) > len(chiavi[j])
		})
	*/

	for _, k := range chiavi {
		fmt.Printf("%s -> occorrenze: %d\n", k, mappa[k])
	}

}
