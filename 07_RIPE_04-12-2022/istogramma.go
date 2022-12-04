package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func LeggiTesto() (righe string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		righe += riga + " "
	}
	return
}

func Occorrenze(s string) map[rune]int {
	mappa := make(map[rune]int)
	for _, lettera := range s {
		if unicode.IsLetter(lettera) {
			mappa[lettera]++
		}
	}
	return mappa
}

func stampa_asterischi(num int) (s string) {
	for i := 0; i < num; i++ {
		s += "*"
	}
	return
}

func StampaIstogramma(occorrenze map[rune]int) {
	// slice di chiavi
	chiavi := []rune{}
	for k, _ := range occorrenze {
		chiavi = append(chiavi, k)
	}

	// ordiniamo la slice di rune
	sort.Slice(chiavi, func(i, j int) bool {
		return chiavi[i] < chiavi[j]
	})

	// stampa istogramma
	fmt.Print("Istogramma: \n")
	for _, v := range chiavi {
		fmt.Print(string(v)+": ", stampa_asterischi(occorrenze[v]), "\n")
	}
}

func main() {
	testo := LeggiTesto()
	mappa := Occorrenze(testo)
	StampaIstogramma(mappa)
}
