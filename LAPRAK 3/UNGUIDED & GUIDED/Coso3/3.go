package main

import "fmt"

func main() {
	var beratbadan, tinggibadan, bmi float64
	fmt.Print("Masukan berat badan (Kg): ")
	fmt.Scan(&beratbadan)
	fmt.Print("Masukan tinggi badan (m): ")
	fmt.Scan(&tinggibadan)
	bmi = beratbadan / (tinggibadan * tinggibadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}
