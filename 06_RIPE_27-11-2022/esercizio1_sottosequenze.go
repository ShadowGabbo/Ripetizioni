package main

import (
	"fmt"
	"os"
)

func palindroma(s string) bool {
	var s_inversa string

	for i := len(s) - 1; i >= 0; i-- {
		s_inversa += string(s[i])
	}

	return s_inversa == s
}

func main() {
	var sottosequenze []string
	parola := os.Args[1]

	for i, _ := range parola {
		for j := i + 1; j <= len(parola); j++ {
			sottosequenza := parola[i:j]
			if len(sottosequenza) > 1 {
				is_palidroma := palindroma(sottosequenza)
				if is_palidroma {
					sottosequenze = append(sottosequenze, sottosequenza)
				}
			}
		}
	}

	fmt.Print(sottosequenze)
}
