package main

import "fmt"

func main() {
	var jam, menit, detik int
	fmt.Scan(&detik)
	jam = detik / 3600
	menit = (detik % 3600) / 60
	detik = detik % 60
	fmt.Println(jam, "jam", menit, "menit", detik, "detik")
}
