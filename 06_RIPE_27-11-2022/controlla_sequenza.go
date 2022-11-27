package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var valido bool = true
	var args []string = make([]string, len(os.Args[1:]))
	copy(args, os.Args[1:])

	for posizione, elemento := range args {
		// controllo se è un numero
		numero, err := strconv.Atoi(elemento)
		if err != nil {
			fmt.Printf("Valore in posizione %d non valido.", posizione)
			valido = false
			break
		}

		if posizione > 0 {
			// controllo la regola 2
			precedente, _ := strconv.Atoi(args[posizione-1])
			if posizione%2 == 1 && numero >= precedente {
				fmt.Printf("Valore in posizione %d non valido.", posizione)
				valido = false
				break
			}

			// controllo la regola 3
			if posizione%2 == 0 && numero <= precedente {
				fmt.Printf("Valore in posizione %d non valido.", posizione)
				valido = false
				break
			}
		}
	}

	if valido {
		fmt.Printf("Sequenza valida")
	}
}
