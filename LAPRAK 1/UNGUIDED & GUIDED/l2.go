package main

import "fmt"

func main() {
	var (
		nama string
		prodi = "S1-IF"
		kelas string
		nim int
	)

	fmt.Println("Masukan nama")
	fmt.Scan(&nama)

	fmt.Println("Masukan kelas")
	fmt.Scan(&kelas)

	fmt.Println("Masukan NIM")
	fmt.Scan(&nim)

	fmt.Println("Perkenalkan saya adalah",nama,"salah satu mahasiswa prodi",prodi,"dari kelas",kelas,"dengan NIM",nim)

}