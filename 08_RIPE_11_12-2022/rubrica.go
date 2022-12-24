package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Indirizzo struct {
	via, cap, citta string
	numeroCivico    uint64
}

type Contatto struct {
	cognome, nome, telefono string
	indirizzo               Indirizzo
}

type Rubrica []Contatto

func NuovaRubrica() Rubrica {
	var rubrica Rubrica
	return rubrica
}

func InserisciContatto(r Rubrica, c Contatto) Rubrica {
	for _, contatto := range r {
		if contatto.nome == c.nome && contatto.cognome == c.cognome {
			return r
		}
	}

	r = append(r, c)
	return r
}

func EliminaContatto(r Rubrica, cognome, nome string) Rubrica {
	var nuova_rubrica Rubrica

	for _, c := range r {
		if c.nome == nome && c.cognome == cognome {
			continue
		}
		nuova_rubrica = append(nuova_rubrica, c)
	}

	return nuova_rubrica
}

func StampaContatto(c Contatto) {
	var indirizzo = c.indirizzo
	fmt.Printf("%s %s: %s %v, %s, %s - Tel. %s\n", c.nome, c.cognome, indirizzo.via, indirizzo.numeroCivico, indirizzo.citta, indirizzo.cap, c.telefono)
}

func StampaRubrica(r Rubrica) {
	fmt.Println("Rubrica:")
	for indice, contatto := range r {
		fmt.Printf("[%d] - ", indice+1)
		StampaContatto(contatto)
	}
}

func AggiornaContatto(rubrica Rubrica, cognome, nome string, via string, numero uint64, cap, citta string, telefono string) Rubrica {
	var indirizzo Indirizzo = Indirizzo{via: via, cap: cap, citta: citta, numeroCivico: numero}
	var nuovo_contatto Contatto = Contatto{cognome: cognome, nome: nome, telefono: telefono, indirizzo: indirizzo}

	for i, contatto := range rubrica {
		if contatto.nome == nome && contatto.cognome == cognome {
			rubrica[i] = nuovo_contatto
		}
	}

	return rubrica
}

func CreaContatto(cognome, nome string, via string, numero uint64, cap, citta string, telefono string) Contatto {
	var indirizzo Indirizzo = Indirizzo{via: via, cap: cap, citta: citta, numeroCivico: numero}
	var contatto Contatto = Contatto{cognome: cognome, nome: nome, telefono: telefono, indirizzo: indirizzo}
	return contatto
}

func main() {
	r := NuovaRubrica()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		linea_s := strings.Split(linea, ";")
		if len(linea_s) == 1 {
			StampaRubrica(r)
		} else {
			switch linea_s[0] {
			case "A":
				//cognome := linea_splittata[1]
				numero, _ := strconv.ParseUint(linea_s[4], 10, 64)
				r = AggiornaContatto(r, linea_s[1], linea_s[2], linea_s[3], numero, linea_s[5], linea_s[6], linea_s[7])
			case "I":
				numero, _ := strconv.ParseUint(linea_s[4], 10, 64)
				c := CreaContatto(linea_s[1], linea_s[2], linea_s[3], numero, linea_s[5], linea_s[6], linea_s[7])
				r = InserisciContatto(r, c)
			case "E":
				r = EliminaContatto(r, linea_s[1], linea_s[2])
			default:
				fmt.Println("Operazione non consentita")
			}
		}
	}
}