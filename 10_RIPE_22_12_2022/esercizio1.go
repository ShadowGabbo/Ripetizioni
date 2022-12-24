package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	epsilon, _ := strconv.ParseFloat(os.Args[1], 64)
	var s, s2 string

	fmt.Scan(&s)

	// prendo solo i numeri
	for i := 0; i < len(s); i++ {
		if !unicode.IsSpace(rune(s[i])) {
			if unicode.IsDigit(rune(s[i])) {
				s2 += string(s[i])
			}
			if string(s[i]) == "." {
				s2 += "."
			}
		}
	}

	fmt.Println(s2)
	var p1, p2, invp1, invp2 string
	s2split := strings.Split(s2, ".")

	// p1 e p2 sono i numeri prima e dopo il punto decimale
	p1 = string(s2split[0])
	p2 = string(s2split[1])

	// reverse della stringa
	for i := len(p1) - 1; i >= 0; i-- {
		invp1 += string(p1[i])
	}
	for i := len(p2) - 1; i >= 0; i-- {
		invp2 += string(p2[i])
	}

	finalS := invp1 + "." + invp2
	fmt.Println(finalS)
	numero_iniziale, _ := strconv.ParseFloat(s2, 64)
	numero_finale, _ := strconv.ParseFloat(finalS, 64)

	// se il valore assoluto della differenza tra il numero iniziale e quello girato e' minore di epsilon allora stampo il numero minore
	if math.Abs(numero_iniziale-numero_finale) > epsilon {
		if numero_finale < numero_iniziale {
			fmt.Println("numero minore:", finalS)
		} else {
			fmt.Println("numero minore:", s2)
		}
	}
}
