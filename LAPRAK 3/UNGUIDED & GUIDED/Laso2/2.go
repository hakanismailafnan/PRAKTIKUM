package main

import "fmt"

func main() {
	var bmi, tinggibadan, beratbadan float64
	
	fmt.Print("Masukkan nilai (kg): ")
	fmt.Scanln(&bmi)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggibadan)
	beratbadan = bmi * (tinggibadan * tinggibadan)
	fmt.Printf("Berat badan anda: %.f", beratbadan)
}
