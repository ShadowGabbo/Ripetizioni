package main

import "fmt"

func StampaTris(tris [][]int) {
	for _, i := range tris {
		for _, j := range i {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func main() {
	var n int = 3
	var tris [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		tris[i] = make([]int, n)
	}

	tris[1][1] = 1

	StampaTris(tris)
}
