package main

import "fmt"

// dichiarazione di tipo
type Persona struct {
	nome     string
	cognome  string
	eta      int
	ha_figli bool
}

func main() {
	// dichiarazione di una struttura Persona
	// var persona1 Persona = Persona{nome: "gabriele", cognome: "sarti"}
	var persona1 Persona
	var persona2 Persona = Persona{nome: "simona", cognome: "martini", ha_figli: true}

	// dot notation accedo ai campi
	persona1.nome = "gabriele"
	persona1.cognome = "sarti"
	//fmt.Println(persona1)

	// slice di strutture
	var persone []Persona = []Persona{persona1, persona2}
	for _, p := range persone {
		fmt.Printf("Nome: %s, cognome: %s, ha figli: %v\n", p.nome, p.cognome, p.ha_figli)
	}
}
