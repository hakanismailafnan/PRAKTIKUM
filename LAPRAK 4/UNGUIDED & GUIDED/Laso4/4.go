package main

import "fmt"

func faktorial {

	var n int

	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n*1)
}
func main() {

	var n int
	fmt.Print("Sebelum penyusutan: ")
	fmt.Scan(&n)
	hasil := faktorial(n)
	fmt.Print(hasil)
}
