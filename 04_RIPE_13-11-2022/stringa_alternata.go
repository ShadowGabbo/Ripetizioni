package main

import "fmt"

func StringheAlternate(s1, s2 string) (risultato string) {
	var len_piu_lunga int
	if len(s1) >= len(s2) {
		len_piu_lunga = len(s1)
	} else {
		len_piu_lunga = len(s2)
	}

	// cane (6)
	// i = 0 (c) ... i = 3 (e)
	var lettera1, lettera2 string
	for i := 0; i < len_piu_lunga; i++ {
		if i < len(s1) {
			lettera1 = string(s1[i])
		} else {
			lettera1 = "-"
		}

		if i < len(s2) {
			lettera2 = string(s2[i])
		} else {
			lettera2 = "-"
		}

		risultato += lettera1 + lettera2
	}
	return
}

func main() {
	var parola1, parola2 string

	fmt.Scan(&parola1)
	fmt.Scan(&parola2)

	fmt.Print(StringheAlternate(parola1, parola2))
}
