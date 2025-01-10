package main

import "fmt"

func main() {
	var totalbelanja, diskon int

	fmt.Print("Total belanja :")
	fmt.Scan(&totalbelanja)
	fmt.Print("Diskon :")
	fmt.Scan(&diskon)
	fmt.Println(diskon <= totalbelanja)
}
