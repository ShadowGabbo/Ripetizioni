package main

func main() {
	//TO-DO: da vedere come si inizializza in modo dinamico
	const dimensione int

	fmt.Scan("%d", &dimensione)

	var arr [dimensione]int
	fmt.Println(arr)
}