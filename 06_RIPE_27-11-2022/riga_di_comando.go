package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// somma dei numeri passati da riga di comando
	var somma int

	// slice di stringa os.Args
	for _, comando := range os.Args[1:] {
		addendo, _ := strconv.Atoi(comando)
		somma += addendo
	}
	fmt.Printf("La somma è : %d", somma)
}
