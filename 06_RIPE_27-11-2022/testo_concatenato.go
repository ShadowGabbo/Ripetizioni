package main

import (
	"bufio"
	"fmt"
	"os"
)

func LeggiTesto() string {
	var testo string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if verifica := scanner.Scan(); verifica {
			linea := scanner.Text()
			testo += linea + "\n"
		} else {
			break
		}
	}

	return testo
}

func main() {
	fmt.Println("Inserisci testo (termina con CTRL+D):")
	testo := LeggiTesto()
	fmt.Println(testo)
}
