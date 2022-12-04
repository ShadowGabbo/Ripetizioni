package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1]
	_, err := strconv.Atoi(s)

	if err == nil {
		sottoseq := Sottostringhe(s)
		fmt.Println(sottoseq)
	}
}

func Sottostringhe(s string) []string {

	var ss []string

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] <= s[j-1] {
				break
			} else {
				ss = append(ss, string(s[i:j+1]))
			}

		}
	}
	return ss
}
