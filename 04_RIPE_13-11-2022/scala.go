package main

import "fmt"

func StampaScala(gradini int) {
	for i := 0; i < gradini; i++ {
		StampaGradino(i)
	}
}

func StampaGradino(gradino int) {
	if gradino == 0 {
		fmt.Print("***\n  *\n")
		return
	}

	for j := 0; j < gradino*2; j++ {
		fmt.Print(" ")
	}

	fmt.Print("***\n")

	for k := 0; k < gradino*2+2; k++ {
		fmt.Print(" ")
	}
	fmt.Print("*\n")
}

func main() {
	var n int = 0
	fmt.Print("Inserisci il numero dei gradini: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Print("Dimensione non sufficiente")
		return
	}

	StampaScala(n)
}
